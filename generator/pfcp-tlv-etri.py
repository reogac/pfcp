# Copyright (C) 2019-2023 by Sukchan Lee <acetcom@gmail.com>

# This file is part of Open5GS.

# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

from docx import Document
import re, os, sys, string, io
import datetime
import getopt
import getpass
from collections import Counter

version = "0.1.0"

msg_list = {}
type_list = {}
group_list = {}

verbosity = 0
filename = ""
outdir = './'
cachedir = './cache/'

ndnTlv ="github.com/usnistgov/ndn-dpdk/ndn/tlv"

FAIL = '\033[91m'
INFO = '\033[93m'
ENDC = '\033[0m'

def d_print(string):
    if verbosity > 0:
        sys.stdout.write(string)

def d_info(string):
    sys.stdout.write(INFO + string + ENDC + "\n")

def d_error(string):
    sys.stderr.write(FAIL + string + ENDC + "\n")
    sys.exit(0)

def write_file(f, string):
    f.write(string)
    d_print(string)

def v_upper(v):
    return re.sub('5GS', 'FIVEGS', re.sub('3GPP', '', re.sub('\'', '_', re.sub('/', '_', re.sub('-', '_', re.sub(' ', '_', v)))).upper()))

def v_lower(v):
    return re.sub('5gs', 'fivegs', re.sub('3gpp', '', re.sub('\'', '_', re.sub('/', '_', re.sub('-', '_', re.sub(' ', '_', v)))).lower()))

def v_cap(v):
    return re.sub('5GS', 'fivegs', re.sub('3GPP', '', re.sub('\'', '', re.sub('/', '', re.sub('-', '', re.sub(' ', '', v))))))

def get_types(ieSize):
    m = "Bytes"
    t = "[]byte"
    if ieSize == 1:
        m = "Uint8"
        t = "uint8"
    elif ieSize == 2:
        m = "Uint16"
        t = "uint16"
    elif ieSize == 3:
        m = "3Bytes"
        t = "uint32"
    elif ieSize == 4:
        m = "Uint32"
        t = "uint32"
    return m, t

# merge_unknown writes the line that adds what a decoded grouped IE skipped to
# the list of the IE or message holding it, so that the message's list holds
# every IE type skipped anywhere in it
def merge_unknown(f, isGroup, value):
    if isGroup:
        f.write("t.Unknown = append(t.Unknown, %s.Unknown...)\n" % value)

def write_unmarshal_multi_ie(f, ieName, customType, ieSize, isGroup=False, ieKey=None):
    m, t = get_types(ieSize) 
    if customType:
        f.write("var v %s\n" % customType)
        f.write("if err = de.UnmarshalValue(&v); err == nil {\n")
        f.write("t." + ieName + " = append(t." + ieName + ", v)\n")
        merge_unknown(f, isGroup, "v")
        f.write("}\n")
    elif t == "[]byte":
        f.write("t." + ieName + " = append(t." + ieName + ", de.Value)\n")
    else:
        f.write("var v " + t+ "\n")
        f.write("if err = " + unmarshal_fn(ieKey, m) + "&v); err == nil {\n")
        f.write("t." + ieName + " = append(t." + ieName + ", v)\n")
        f.write("}\n")

# mandatory_on_n4 is whether a row is M on N4: a row is mandatory on the
# interfaces its applicability columns mark, and the N4 column is the one this
# library is for
def mandatory_on_n4(ie):
    return ie["presence"] == "M" and ie.get("n4") == "X" and not ie.get("collapsed")

# write_mandatory_checks ends a decoder: a mandatory IE that did not arrive is a
# MissingIE naming its type (TS 29.244 7.6.6). checks are (row, seen index or
# None for a list, field name, tag)
def write_mandatory_checks(f, checks):
    for (ie, idx, ieName, ieTag) in checks:
        if idx is None:
            f.write("if len(t.%s) == 0 {\n" % ieName)
        else:
            f.write("if !seen[%d] {\n" % idx)
        f.write("return &pfcp.MissingIE{Type: %s}\n}\n" % ieTag)

# write_first_only guards the decoding of a single IE: an IE repeated where
# repetition is not specified is handled once, at its first occurrence, and its
# later ones ignored (TS 29.244 7.6.10). seen is an array on the decoder's
# stack, so the guard allocates nothing
def write_first_only(f, idx):
    f.write("if seen[%d] {\ncontinue\n}\nseen[%d] = true\n" % (idx, idx))

def write_unmarshal_ie(f, ieName, customType, ieSize, presence, isGroup=False, ieKey=None):
    m, t = get_types(ieSize) 
    if customType:
        if presence == "M":
            f.write("if err = de.UnmarshalValue(&t.%s); err == nil {\n" % ieName)
            merge_unknown(f, isGroup, "t." + ieName)
            f.write("}\n")
        else:
            f.write("var v %s\n" % customType)
            f.write("if err = de.UnmarshalValue(&v); err == nil {\n")
            f.write("t." + ieName + " = &v\n")
            merge_unknown(f, isGroup, "v")
            f.write("}\n")
    elif t == "[]byte":
        f.write("t.%s = de.Value\n" % ieName)
    else:
        if presence == "M":
            f.write("err = " + unmarshal_fn(ieKey, m) + "&t.%s)\n" % ieName)
        else:
            f.write("var v %s\n" % t)
            f.write("if err = " + unmarshal_fn(ieKey, m) + "&v); err == nil {\n")
            f.write("t." + ieName + " = &v\n")
            f.write("}\n")

def write_field_multi_ie(f, ieName, ieTag, customType, isGroup, ieSize, v="ie"):
    m, t = get_types(ieSize) 
    if customType:
        if isGroup:
            ieMarshal = "%s.Field()" % v
        else:
            ieMarshal = "tlv.TLVFrom(%s, tlv.Bytes(%s.Bytes()))" % (ieTag, v)
    elif t == "[]byte":
        ieMarshal = "tlv.TLVFrom(%s, tlv.Bytes(%s))" % (ieTag, v)
    else: 
        ieMarshal = "tlv.TLVFrom(%s, tlv.Bytes(pfcp.Marshal%s(%s)))" % (ieTag, m, v)

    #f.write("if len(t." + ieName + ") >0 {\n")
    f.write("for _, %s := range t.%s {\n" % (v, ieName))
    f.write("fields = append(fields, %s)\n" % ieMarshal)
    f.write("}\n")
    #f.write("}\n")

def write_field_ie(f, ieName, ieTag, customType, isGroup, ieSize, presence):
    m, t = get_types(ieSize) 
    if customType:
        if isGroup:
            ieMarshal = "t.%s.Field()" % ieName
        else:
            ieMarshal = "tlv.TLVFrom(%s, tlv.Bytes(t.%s.Bytes()))" % (ieTag, ieName)
        if presence != "M":
            f.write("if t.%s != nil {\n" % ieName)
            f.write("fields = append(fields, %s)\n" % ieMarshal)
            f.write("}\n")
        else:
            f.write("fields = append(fields, %s)\n" % ieMarshal)
    elif t == "[]byte":
        ieMarshal = "t.%s" % ieName
        if presence != "M":
            f.write("if len(t.%s) > 0 {\n" % ieName)
            f.write("fields = append(fields, tlv.TLVFrom(%s, tlv.Bytes(%s)))\n"\
                    % (ieTag, ieMarshal))
            f.write("}\n")
        else:
            f.write("fields = append(fields, tlv.TLVFrom(%s, tlv.Bytes(%s)))\n"\
                    % (ieTag, ieMarshal))
    else: 
        if presence == "M":
            ieMarshal = "pfcp.Marshal%s(t.%s)" % (m, ieName)
            f.write("fields = append(fields, tlv.TLVFrom(%s,\
                    tlv.Bytes(%s)))\n" % (ieTag, ieMarshal))
        else:
            ieMarshal = "pfcp.Marshal%s(*t.%s)" % (m, ieName)
            f.write("if t.%s !=nil {\n" % ieName)
            f.write("fields = append(fields, tlv.TLVFrom(%s,\
                    tlv.Bytes(%s)))\n" % (ieTag, ieMarshal))
            f.write("}\n")

def write_custom_type(ieType, ieSize):
    type_file = ieDir + ieType + "IE.go"
    if os.path.isfile(type_file):
        return

    f = open(type_file, 'w')
    f.write("package pfcp\n")
    f.write("type %s struct {\n" % ieType)
    if ieSize == 0:
        f.write("bytes []byte\n")
    elif ieSize == 1:
        f.write("Value uint8\n")
    elif ieSize == 2:
        f.write("Value uint16\n")
    elif ieSize == 4:
        f.write("Value uint32\n")
    else:
        f.write("Data [3]byte\n")
    f.write("}\n")

    f.write("func (ie *%s) Bytes() []byte {\n" % ieType)
    if ieSize == 0:
        f.write("return ie.bytes\n")
    elif ieSize == 1:
        f.write("return []byte{ie.Value}\n")
    elif ieSize == 2:
        f.write("return marshalUint16(ie.Value)\n")
    elif ieSize == 4:
        f.write("return marshalUint32(ie.Value)\n")
    else:
        f.write("return ie.Data[:]\n")
    f.write("}\n")

    f.write("func (ie *%s) UnmarshalBinary(wire []byte) (err error) {\n" % ieType)
    if ieSize == 0:
        f.write("ie.bytes=make([]byte,len(wire))\n copy(ie.bytes, wire)\nreturn nil\n")
    elif ieSize == 1:
        f.write("return pfcp.UnmarshalUint8(wire,&ie.Value)\n")
    elif ieSize == 2:
        f.write("return pfcp.UnmarshalUint16(wire,&ie.Value)\n")
    elif ieSize == 4:
        f.write("return pfcp.UnmarshalUint32(wire, &ie.Value)\n")
    else:
        f.write("if len(wire) < 3 {\nreturn tlv.ErrIncomplete\n} else if\
                len(wire) >3 {\nreturn tlv.ErrTail\n} else {\ncopy(ie.Data[:],\
                wire)\nreturn nil\n}\n")
    f.write("}\n")
    f.close()
    return
""

def output_header_to_file(f):
    now = datetime.datetime.now()
    f.write("package pfcp\n\n")

modulePath = "github.com/reogac/pfcp"

# write_go_file writes a generated file of package pkg: the body, behind the
# imports it uses. The module's own packages come first and its tlv package
# after a blank line, as gofmt leaves the library's files.
def write_go_file(path, pkg, body):
    own = []
    if "pfcp." in body:
        own.append(modulePath)
    if pkg != "ie" and "ie." in body:
        own.append(modulePath + "/ie")
    f = open(path, 'w')
    f.write("package %s\n\nimport (\n" % pkg)
    for imp in own:
        f.write("\"%s\"\n" % imp)
    if own:
        f.write("\n")
    f.write("\"%s/tlv\"\n)\n" % modulePath)
    f.write(body)
    f.close()

def usage():
    print("Python generating TLV build/parser for PFCP v%s" % (version))
    print("Usage: python pfcp-tlv.py [options]")
    print("Available options:")
    print("-d        Enable script debug")
    print("-f [file] Input file to parse")
    print("-o [dir]  Output files to given directory")
    print("-c [dir]  Cache files to given directory")
    print("-h        Print this help and return")

# n4_column is the index of a table's N4 applicability column, found by its
# header cell, or None for a table with none
def n4_column(table, header_rows):
    for row in table.rows[:header_rows]:
        for i, c in enumerate(row.cells):
            if c.text.strip() == "N4":
                return i
    return None

def get_cells(cells, n4idx=None):
    note = cells[0].text
    if note.find('NOTE') != -1:
        return None
    comment = cells[2].text.encode('ascii', 'ignore').decode('utf-8')
    comment = re.sub('\n|\"|\'|\\\\', '', comment);
    #print(comment)
    ie_type = re.sub('\s*$', '', re.sub('\'\s*\n*\s*\(NOTE.*\)*', '', cells[-1].text))

    #if ie_type.find('Usage Report') != -1:
    if ie_type == 'Usage Report':
        if comment.find('Report Type') != -1:
            ie_type = "Usage Report Session Report Request"
        elif comment.find('Query URR') != -1:
            ie_type = "Usage Report Session Modification Response"
        elif comment.find('provisioned ') != -1:
            ie_type = "Usage Report Session Deletion Response"
        else:
             assert False, "Unknown IE type : [Usage Report]"

    if ie_type == 'Update BAR':
        if comment.find('7.5.4.11-1') != -1:
            ie_type = "Update BAR Session Modification Request"
        elif comment.find('7.5.9.2-1') != -1:
            ie_type = "Update BAR PFCP Session Report Response"
        else:
             assert False, "Unknown IE type : [Update BAR]"

    if ie_type.find('PFD Contents') != -1:
        ie_type = 'PFD contents'
    elif ie_type.find('UE IP address') != -1:
        ie_type = 'UE IP Address'
    elif ie_type.find('SxSMReq-Flags') != -1:
        ie_type = 'PFCPSMReq-Flags'
    elif ie_type.find('PFCPSRRsp-Flags2') != -1:
        ie_type = 'PFCPSRRsp-Flags'
    elif ie_type.find('IPv4 Configuration Parameters (IP4CP)') != -1:
        ie_type = 'IP4CP'
    elif ie_type.find('QoS Information') != -1:
        ie_type = 'QoS Information in GTP-U Path QoS Report'
    elif ie_type.find('IP Multicast Addressing Info') != -1:
        ie_type = 'IP Multicast Addressing Info within PFCP Session Establishment Request'
    elif ie_type.find('Redundant Transmission Detection Parameters') != -1:
        ie_type = 'Redundant Transmission Parameters'
    elif ie_type.find('MAC Addressed Detected') != -1:
        ie_type = 'MAC Addresses Detected'
    elif ie_type.find('Join IP Multicast Information') != -1:
        ie_type = 'Join IP Multicast Information IE within Usage Report'
    elif ie_type.find('Leave IP Multicast Information') != -1:
        ie_type = 'Leave IP Multicast Information IE within Usage Report'
    elif ie_type.find('PFCP Session Retention Information') != -1:
        ie_type = 'PFCP Session Retention Information within PFCP Association Setup Request'
    elif ie_type.find('GTP-U Path QoS Report') != -1:
        ie_type = 'GTP-U Path QoS Report PFCP Node Report Request'
    elif ie_type.find('Provide RDS configuration information') != -1:
        ie_type = 'Provide RDS Configuration Information'
    elif ie_type.find('RDS configuration information') != -1:
        ie_type = 'RDS Configuration Information'
    elif ie_type.find('Group Id') != -1:
        ie_type = 'Group ID'
    elif ie_type.find('Group-Id') != -1:
        ie_type = 'Group ID'
    elif ie_type.find('L2TP session Indications') != -1:
        ie_type = 'L2TP Session Indications'
    elif ie_type.find('L2TP User Authentication') != -1:
        ie_type = 'L2TP User Authentication IE'
    elif ie_type.find('IP Address and Port Number Replacement') != -1:
        ie_type = 'IP Address and Port number Replacement'
    elif ie_type.find('User Plane Node Management Information Container') != -1:
        ie_type = 'Bridge Management Information Container'
    elif ie_type.find('TSC Management Information') != -1:
        ie_type = 'TSC Management Information IE within PFCP Session Modification Request'
    elif ie_type.find('Query Packet Rate Status') != -1:
        ie_type = 'Query Packet Rate Status IE within PFCP Session Modification Request'
    if ie_type not in type_list.keys():
        assert False, "Unknown IE type : [" \
                + cells[-1].text + "]" + "(" + ie_type + ")"
    presence = cells[1].text
    ie_value = re.sub('\s*\n*\s*\([^\)]*\)*', '', cells[0].text)
    ie_value = re.sub('\n', '', ie_value)
    if ie_value[len(ie_value)-1] == ' ':
        ie_value = ie_value[:len(ie_value)-1]

    tlv_more = "0"  # PFCP has no tlv_more
    if ie_type == 'Create PDR' or ie_type == 'Created PDR' or ie_type == 'Update PDR' or ie_type == "Remove PDR":
        tlv_more = "15"
    if ie_type == 'Create FAR' or ie_type == 'Update FAR' or ie_type == "Remove FAR":
        tlv_more = "15"
    if ie_type == 'Create URR' or ie_type == 'Update URR' or ie_type == "Remove URR":
        tlv_more = "15"
    if ie_type == 'Create QER' or ie_type == 'Update QER' or ie_type == "Remove QER":
        tlv_more = "3"
    if ie_type == 'User Plane IP Resource Information':
        tlv_more = "3"
    if ie_type == 'SDF Filter':
        tlv_more = "7"
    if (ie_type == 'Framed-Route' or
        ie_type == 'Framed-IPv6-Route'):
        tlv_more = "7"
    if (ie_type == 'Usage Report Session Report Request' or
        ie_type == 'Usage Report Session Deletion Response' or
        ie_type == 'Usage Report Session Modification Response'):
        tlv_more = "7"
    if ie_type == 'URR ID' and comment.find('Several IEs within the same IE type may be present') != -1:
        tlv_more = "7"

    if int(tlv_more) > int(type_list[ie_type]["max_tlv_more"]):
        type_list[ie_type]["max_tlv_more"] = tlv_more
        write_file(f, "type_list[\"" + ie_type + "\"][\"max_tlv_more\"] = \"" + tlv_more + "\"\n")

    # whether the row applies on N4: "X" or "-", "" for a table with no column.
    # A row listing flags marks each in its own paragraph, so the cell is X when
    # it marks N4 at all
    n4 = ""
    if n4idx is not None and n4idx < len(cells):
        n4 = "X" if "X" in cells[n4idx].text else "-"
    return { "ie_type" : ie_type, "ie_value" : ie_value, "presence" : presence, "tlv_more" : tlv_more, "comment" : comment, "n4" : n4 }

def write_cells_to_file(name, cells):
    write_file(f, name + ".append({ \"ie_type\" : \"" + cells["ie_type"] + \
        "\", \"ie_value\" : \"" + cells["ie_value"] + \
        "\", \"presence\" : \"" + cells["presence"] + \
        "\", \"tlv_more\" : \"" + cells["tlv_more"] + \
        "\", \"comment\" : \"" + cells["comment"] + \
        "\", \"n4\" : \"" + cells.get("n4", "") + "\"})\n")

try:
    opts, args = getopt.getopt(sys.argv[1:], "df:ho:c:", ["debug", "file", "help", "output", "cache"])
except getopt.GetoptError as err:
    # print help information and exit:
    usage()
    sys.exit(2)

for o, a in opts:
    if o in ("-d", "--debug"):
        verbosity = 1
    if o in ("-f", "--file"):
        filename = a
    if o in ("-o", "--output"):
        outdir = a
        if outdir.rfind('/') != len(outdir):
            outdir += '/'
    if o in ("-c", "--cache"):
        cachedir = a
        if cachedir.rfind('/') != len(cachedir):
            cachedir += '/'
    if o in ("-h", "--help"):
        usage()
        sys.exit(2)

ieDir = outdir + 'ie/'
msgDir = outdir + 'message/'
os.makedirs(ieDir, exist_ok=True)
os.makedirs(msgDir, exist_ok=True)

if os.path.isfile(filename) and os.access(filename, os.R_OK):
    file = open(filename, 'r')
else:
    d_error("Cannot find file : " + filename)

d_info("[Message List]")
cachefile = cachedir + 'tlv-msg-list.py'
if os.path.isfile(cachefile) and os.access(cachefile, os.R_OK):
    exec(open(cachefile).read())
    print("Read from " + cachefile)
else:
    document = Document(filename)
    f = open(cachefile, 'w')

    msg_table = ""
    for i, table in enumerate(document.tables):
        try:
            cell = table.rows[0].cells[0]
        except:
            continue;
        else:
            if cell.text.find('Message Type value') != -1:
                msg_table = table
                d_print("Table Index = %d\n" % i)

    for row in msg_table.rows[2:-3]:
        key = row.cells[1].text
        type = row.cells[0].text
        if type.isdigit() is False:
            continue
        if key.find('Reserved') != -1:
            continue
        key = re.sub('\s*\n*\s*\([^\)]*\)*', '', key)
        key = re.sub('\n', '', key)
        msg_list[key] = { "type": type }
        write_file(f, "msg_list[\"" + key + "\"] = { \"type\" : \"" + type + "\" }\n")
    f.close()

d_info("[IE Type List]")
cachefile = cachedir + 'tlv-type-list.py'
if os.path.isfile(cachefile) and os.access(cachefile, os.R_OK):
    exec(open(cachefile).read())
    print("Read from " + cachefile)
else:
    document = Document(filename)
    f = open(cachefile, 'w')

    ie_table = ""
    for i, table in enumerate(document.tables):
        try:
            cell = table.rows[0].cells[0]
        except:
            pass
        else:
            if cell.text.find('IE Type value') != -1:
                ie_table = table
                d_print("Table Index = %d\n" % i)

    for row in ie_table.rows[1:]:
        key = row.cells[1].text
        type = row.cells[0].text
        if type.isdigit() is False:
            continue
        if key.find('Reserved') != -1:
            continue
        key = re.sub('\(', '', key)
        key = re.sub('\)', '', key)
        key = re.sub('\s*$', '', key)

        # the 3rd column: Fixed Length, Variable Length or Extendable; the
        # 4th: the octets the IE had when first defined (TS 29.244 8.1.2)
        kind = re.sub(r'\s*/.*$', '', row.cells[2].text.strip().split('\n')[0]).strip()
        fixed = row.cells[3].text.strip()
        type_list[key] = { "type": type , "max_tlv_more" : "0", "kind" : kind, "fixed" : fixed }
        write_file(f, "type_list[\"" + key + "\"] = { \"type\" : \"" + type)
        write_file(f, "\", \"max_tlv_more\" : \"0\", \"kind\" : \"" + kind + "\", \"fixed\" : \"" + fixed + "\" }\n")
    f.close()

d_info("[Group IE List]")
cachefile = cachedir + 'tlv-group-list.py'
if os.path.isfile(cachefile) and os.access(cachefile, os.R_OK):
    exec(open(cachefile).read())
    print("Read from " + cachefile)
else:
    document = Document(filename)
    f = open(cachefile, 'w')

    for i, table in enumerate(document.tables):
        try:
            cell = table.rows[0].cells[0]
        except:
            pass
        else:
            if cell.text.find('Octet') != -1 and \
               table.rows[0].cells[1].text.find('Outer Header to be created') == -1:

                num = 0;
                if len(table.rows[0].cells) > 2 and table.rows[0].cells[2].text.find('IE Type') != -1:
                    num = 2
                elif len(table.rows[0].cells) > 3 and table.rows[0].cells[3].text.find('IE Type') != -1:
                    num = 3
                elif len(table.rows[0].cells) > 4 and table.rows[0].cells[4].text.find('IE Type') != -1:
                    num = 4

                if num == 0:
                    continue;

                row = table.rows[0];

                d_print("Table Index = %d[%d:%s]\n" % (i, num, row.cells[num].text))

                if len(re.findall('\d+', row.cells[num].text)) == 0:
                    continue;
                ie_type = re.findall('\d+', row.cells[num].text)[-1]
                ie_name = re.sub('\s*IE Type.*', '', row.cells[num].text)

                d_print("TYPE:%s NAME:%s\n" % (ie_type, ie_name))

                # SKIP Access Forwarding Action Information
                if (int(ie_type) == 78):
                    ie_name =  "Usage Report Session Modification Response"
                elif (int(ie_type) == 79):
                    ie_name =  "Usage Report Session Deletion Response"
                elif (int(ie_type) == 80):
                    ie_name =  "Usage Report Session Report Request"
                elif (int(ie_type) == 86):
                    ie_name =  "Update BAR Session Modification Request"
                elif (int(ie_type) == 12):
                    ie_name =  "Update BAR PFCP Session Report Response"
                elif (int(ie_type) == 183):
                    ie_name =  "PFCP Session Retention Information within PFCP Association Setup Request"
                elif (int(ie_type) == 188):
                    ie_name =  "IP Multicast Addressing Info within PFCP Session Establishment Request"
                elif (int(ie_type) == 189):
                    ie_name =  "Join IP Multicast Information IE within Usage Report"
                elif (int(ie_type) == 190):
                    ie_name =  "Leave IP Multicast Information IE within Usage Report"
                elif (int(ie_type) == 199):
                    ie_name =  "TSC Management Information IE within PFCP Session Modification Request"
                elif (int(ie_type) == 200):
                    ie_name =  "TSC Management Information IE within PFCP Session Modification Response"
                elif (int(ie_type) == 201):
                    ie_name =  "TSC Management Information IE within PFCP Session Report Request"
                elif (int(ie_type) == 239):
                    ie_name =  "GTP-U Path QoS Report PFCP Node Report Request"
                elif (int(ie_type) == 240):
                    ie_name =  "QoS Information in GTP-U Path QoS Report"
                elif (int(ie_type) == 255):
                    ie_name =  "Redundant Transmission Parameters"
                elif (int(ie_type) == 263):
                    ie_name =  "Query Packet Rate Status IE within PFCP Session Modification Request"
                elif (int(ie_type) == 264):
                    ie_name =  "Packet Rate Status Report IE within PFCP Session Modification Response"

                if ie_name.find('Non-3GPP Access Forwarding Action Information') != -1:
                    ie_idx = str(int(ie_type)+100)
                    group_list[ie_name] = { "index" : ie_idx, "type" : ie_type, "ies" : ies }
                    write_file(f, "group_list[\"" + ie_name + "\"] = { \"index\" : \"" + ie_idx + "\", \"type\" : \"" + ie_type + "\", \"ies\" : ies }\n")
                    continue

                if ie_name not in group_list.keys():
                    ies = []
                    write_file(f, "ies = []\n")
                    n4idx = n4_column(table, 4)
                    for row in table.rows[4:]:
                        cells = get_cells(row.cells, n4idx)
                        if cells is None:
                            continue

                        ies.append(cells)
                        write_cells_to_file("ies", cells)

                    ie_idx = str(int(ie_type)+100)
                    group_list[ie_name] = { "index" : ie_idx, "type" : ie_type, "ies" : ies }
                    write_file(f, "group_list[\"" + ie_name + "\"] = { \"index\" : \"" + ie_idx + "\", \"type\" : \"" + ie_type + "\", \"ies\" : ies }\n")
    f.close()

# every message is parsed from its own table, found by its caption ("Table
# 7.5.2.1-1: Information Elements in a PFCP Session Establishment Request")
# rather than by its index in the docx, which a new release moves. Two messages
# are parsed with no IEs: the PFCP Version Not Supported Response, which has no
# table, and the PFCP Session Deletion Request, whose table is skipped as it
# always was
no_ies = ("PFCP Session Deletion Request", "PFCP Version Not Supported Response")
for key in msg_list.keys():
    msg_list[key]["table"] = True

def caption_key(name):
    name = re.sub(r'^(a|an) ', '', name.strip())
    return re.sub(r'^PFCP ', '', name).lower()

# message_table is the docx table whose caption names message key; exactly one
# must, so a release that renames or splits a table stops the parse
def message_table(document, key):
    found, caption = [], ""
    tables = list(document.tables)
    i = 0
    for el in document.element.body.iterchildren():
        tag = el.tag.split('}')[1]
        if tag == 'p':
            text = ''.join(x.text or '' for x in el.iter() if x.tag.endswith('}t')).strip()
            if re.match(r'Table \d', text):
                caption = text
        elif tag == 'tbl':
            m = re.search(r'Information Elements in (.*)$', caption)
            if m and caption_key(m.group(1)) == caption_key(key):
                found.append(i)
            caption = ""
            i += 1
    assert len(found) == 1, "%s: %d tables captioned for it" % (key, len(found))
    return tables[found[0]]

for key in msg_list.keys():
    if msg_list[key]["table"]:
        d_info("[" + key + "]")
        cachefile = cachedir + "tlv-msg-" + msg_list[key]["type"] + ".py"
        if os.path.isfile(cachefile) and os.access(cachefile, os.R_OK):
            exec(open(cachefile).read())
            print("Read from " + cachefile)
        else:
            document = Document(filename)
            f = open(cachefile, 'w')

            table = None if key in no_ies else message_table(document, key)
#            if key.find('Association') != -1:
#                start_i = 1
#            elif key.find('Heartbeat') != -1:
#                start_i = 1
#            else:
#                start_i = 2
            start_i = 2

            ies = []
            write_file(f, "ies = []\n")
            if key not in no_ies:
                n4idx = n4_column(table, start_i)
                for row in table.rows[start_i:]:
                    cells = get_cells(row.cells, n4idx)
                    if cells is None:
                        continue

                    item = {
                        "ie_type" : cells["ie_type"],
                        "ie_value" : cells["ie_value"],
                        "presence" : cells["presence"],
                        "tlv_more" : cells["tlv_more"],
                        "comment" : cells["comment"],
                        "n4" : cells["n4"]
                        }
                    ies.append(item)
                    write_cells_to_file("ies", item)

            msg_list[key]["ies"] = ies
            write_file(f, "msg_list[key][\"ies\"] = ies\n")
            f.close()

#set size for some IEs
type_list["Cause"]["size"] = 1                              # Type 19
type_list["Source Interface"]["size"] = 1                   # Type 20
type_list["Gate Status"]["size"] = 1                        # Type 25
type_list["QER Correlation ID"]["size"] = 4                 # Type 28
type_list["Precedence"]["size"] = 4                         # Type 29
type_list["Time Threshold"]["size"] = 4                     # Type 32
type_list["Reporting Triggers"]["size"] = 3                 # Type 37
type_list["Report Type"]["size"] = 1                        # Type 39
type_list["Offending IE"]["size"] = 2                       # Type 40
type_list["Destination Interface"]["size"] = 1              # Type 42
type_list["Apply Action"]["size"] = 2                       # Type 44
type_list["PFCPSMReq-Flags"]["size"] = 1                    # Type 49
type_list["PFCPSRRsp-Flags"]["size"] = 1                    # Type 50
type_list["PDR ID"]["size"] = 2                             # Type 56
type_list["Measurement Method"]["size"] = 1                 # Type 62
type_list["Measurement Information"]["size"] = 4
type_list["Downlink Data Notification Delay"]["size"] = 1
type_list["Suggested Buffering Packets Count"]["size"] = 1
type_list["Usage Report Trigger"]["size"] = 3               # Type 63
type_list["Measurement Period"]["size"] = 4                 # Type 64
type_list["Duration Measurement"]["size"] = 4               # Type 67
type_list["Time of First Packet"]["size"] = 4               # Type 69
type_list["Time of Last Packet"]["size"] = 4                # Type 70
type_list["Quota Holding Time"]["size"] = 4                 # Type 71
type_list["Time Quota"]["size"] = 4                         # Type 74
type_list["Start Time"]["size"] = 4                         # Type 75
type_list["End Time"]["size"] = 4                           # Type 76
type_list["URR ID"]["size"] = 4                             # Type 81
type_list["BAR ID"]["size"] = 1                             # Type 88
type_list["CP Function Features"]["size"] = 1               # Type 89
type_list["Recovery Time Stamp"]["size"] = 4                # Type 96
type_list["UR-SEQN"]["size"] = 4                            # Type 104
type_list["FAR ID"]["size"] = 4                             # Type 108
type_list["QER ID"]["size"] = 4                             # Type 109
type_list["PDN Type"]["size"] = 1                           # Type 113
type_list["RQI"]["size"] = 1                                # Type 123
type_list["QFI"]["size"] = 1                                # Type 124
type_list["Event Quota"]["size"] = 4                        # Type 148
type_list["Event Threshold"]["size"] = 4                    # Type 149
type_list["Averaging Window"]["size"] = 4                   # Type 157
type_list["Paging Policy Indicator"]["size"] = 1            # Type 158
type_list["PFCPSRReq-Flags"]["size"] = 1                    # Type 161
type_list["PFCPAUReq-Flags"]["size"] = 1                    # Type 162
type_list["Quota Validity Time"]["size"] = 4                # Type 181
type_list["PFCPSEReq-Flags"]["size"] = 1                    # Type 186
type_list["Data Status"]["size"] = 1                        # Type 260


# every hard-coded size is checked against the IE type table: a Fixed Length IE
# has exactly its octets, an Extendable one at least the octets it was first
# defined with. A size the table contradicts stops the generator
for name, t in type_list.items():
    if "size" not in t:
        continue
    fixed = int(t["fixed"])
    if t["kind"].startswith("Fixed"):
        assert t["size"] == fixed, "%s: size %d, a Fixed Length IE of %d octets" % (name, t["size"], fixed)
    else:
        assert t["kind"].startswith("Extendable") and t["size"] >= fixed, \
            "%s: size %d, a %s IE first defined with %d octets" % (name, t["size"], t["kind"], fixed)

# unmarshal_fn is the call that decodes a sized IE, up to its first argument
# after the value: an Extendable one through pfcp.UnmarshalExtendable, which
# takes any length from the first-defined one up and ignores octets beyond
# those this library knows, as the legacy receiver of TS 29.244 8.1.2 must
def unmarshal_fn(ieKey, m):
    t = type_list.get(ieKey, {})
    if ieKey and t.get("kind", "").startswith("Extendable") and "size" in t:
        return "pfcp.UnmarshalExtendable(de.Value, %d, %d, " % (int(t["fixed"]), t["size"])
    return "pfcp.Unmarshal%s(de.Value, " % m

#TLV tags definitions
f = open(ieDir + 'tags.go', 'w')
f.write("package ie\n\n")

f.write("/* PFCP message type */\n")

tmp = [(k, v["type"]) for k, v in msg_list.items()]
sorted_msg_list = sorted(tmp, key=lambda tup: int(tup[1]))
f.write("const (\n")
for (k, v) in sorted_msg_list:
    f.write("MSG_" + v_upper(k) + "_TYPE uint32 =" + v + "\n")
f.write(")\n")

tmp = [(k, v["type"]) for k, v in type_list.items()]
sorted_type_list = sorted(tmp, key=lambda tup: int(tup[1]))

f.write("\n/* Group/element type */\n")
f.write("const (\n")
for (k, v) in sorted_type_list:
    f.write("  " + v_upper(k) + "_TYPE uint32 =" + v + "\n")
f.write(")\n")

f.close()

for k, v in group_list.items():
    if v_lower(k) == "ethernet_packet_filter":
        v["index"] = "1"
    if v_lower(k) == "redundant_transmission_parameters":
        v["index"] = "2"
    if v_lower(k) == "ip_multicast_addressing_info_within_pfcp_session_establishment_request":
        v["index"] = "3"
    if v_lower(k) == "pdi":
        v["index"] = "4"
    if v_lower(k) == "transport_delay_reporting":
        v["index"] = "5"
    if v_lower(k) == "create_pdr":
        v["index"] = "6"
    if v_lower(k) == "forwarding_parameters":
        v["index"] = "7"
    if v_lower(k) == "duplicating_parameters":
        v["index"] = "8"
    if v_lower(k) == "redundant_transmission_forwarding_parameters":
        v["index"] = "9"
    if v_lower(k) == "mbs_multicast_parameters":
        v["index"] = "10"
    if v_lower(k) == "add_mbs_unicast_parameters":
        v["index"] = "11"
    if v_lower(k) == "create_far":
        v["index"] = "12"
    if v_lower(k) == "update_forwarding_parameters":
        v["index"] = "13"
    if v_lower(k) == "update_duplicating_parameters":
        v["index"] = "14"
    if v_lower(k) == "update_far":
        v["index"] = "15"
    if v_lower(k) == "pfd_context":
        v["index"] = "16"
    if v_lower(k) == "application_id_s_pfds":
        v["index"] = "17"
    if v_lower(k) == "ethernet_traffic_information":
        v["index"] = "18"
    if v_lower(k) == "_access_forwarding_action_information":
        v["index"] = "19"
    if v_lower(k) == "non__access_forwarding_action_information":
        v["index"] = "20"
    if v_lower(k) == "update__access_forwarding_action_information":
        v["index"] = "21"
    if v_lower(k) == "update_non__access_forwarding_action_information":
        v["index"] = "22"
    if v_lower(k) == "access_availability_report":
        v["index"] = "23"
    if v_lower(k) == "qos_monitoring_report":
        v["index"] = "24"
    if v_lower(k) == "mptcp_parameters":
        v["index"] = "25"
    if v_lower(k) == "atsss_ll_parameters":
        v["index"] = "26"
    if v_lower(k) == "pmf_parameters":
        v["index"] = "27"
    if v_lower(k) == "join_ip_multicast_information_ie_within_usage_report":
        v["index"] = "28"
    if v_lower(k) == "leave_ip_multicast_information_ie_within_usage_report":
        v["index"] = "29"

tmp = [(k, v["index"]) for k, v in group_list.items()]
sorted_group_list = sorted(tmp, key=lambda tup: int(tup[1]), reverse=False)



# repetition: whether a row may occur more than once, read from the row's own
# text. TS 29.244 says so in one of these phrasings, and a row saying none is a
# single IE; a repeated IE where repetition is not specified is a later
# occurrence to ignore (7.6.10). This replaces Open5GS's tlv_more, a list of the
# IEs Open5GS itself needed as arrays, which left most repeated IEs single
repeat_phrases = re.compile(
    r"[Ss]everal IEs? (with|within) the same (IE )?type"
    r"|[Mm]ore than one IE with this type"
    r"|[Tt]here may be several instances"
    r"|[Tt]wo IEs with the same IE type"
    r"|list of all the")

# repeat_overrides are rows that repeat although their text uses none of the
# phrases, each with the wording that says so
repeat_overrides = {
    # "shall be included per PDR ID reported in the message. When multiple PDR ID
    # IEs are present in the message, ..."
    ("Downlink Data Report", "Downlink Data Service Information"),
}

# repeats is whether row ie of message or grouped IE k is a list
def repeats(k, ie):
    return bool(ie.get("collapsed")) or (k, ie["ie_value"]) in repeat_overrides \
        or bool(repeat_phrases.search(ie["comment"]))

# field_decl writes one field of a struct: a list, a value, or a pointer. It is
# the one place a field's Go type is decided, so that the type of a repeated IE
# can change later in one edit
def field_decl(f, ieName, ieType, ieKey, multi, value):
    if multi:
        f.write("%s []%s //%s\n" % (ieName, ieType, ieKey))
    elif value:
        f.write("%s %s //%s\n" % (ieName, ieType, ieKey))
    else:
        f.write("%s *%s //%s\n" % (ieName, ieType, ieKey))

# unknown_field is the last field of every message and grouped IE: the types of
# the IEs its decoder skipped, as unknown or not defined for it, its grouped
# IEs' included. TS 29.244 7.6.9 has the receiver discard such an IE, log an
# error and process the message, so decoding never fails on one; the caller
# logs what this holds. Not in the JSON form, which carries what was decoded
def unknown_field(f):
    f.write("Unknown []uint32 `json:\"-\"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)\n")

# name_overrides correct a row's name where the specification's spelling would
# name the field differently from every other row of its kind
name_overrides = {
    # every other Network Instance row is spelled with a capital I
    ("Update Forwarding Parameters", "Network instance"): "Network Instance",
}
for (k, v) in group_list.items():
    for ie in v["ies"]:
        if (k, ie["ie_value"]) in name_overrides:
            ie["ie_value"] = name_overrides[(k, ie["ie_value"])]

# presence overrides, each a row whose P column cannot be taken as it is.
# MBS Session N4mb Control Information is M on N4mb alone: its Sxa, Sxb, Sxc
# and N4 columns are "-", and a value field sent it on every N4 request
presence_overrides = {
    ("PFCP Session Establishment Request", "MBS Session N4mb Control Information"): "O",
}
for (k, v) in sorted_msg_list:
    for ie in msg_list[k]["ies"]:
        if (k, ie["ie_value"]) in presence_overrides:
            ie["presence"] = presence_overrides[(k, ie["ie_value"])]

#fix attributes having same TLV tag in a message
#replace them with a single list of same type
for (k, v) in sorted_msg_list:
    msgType = v_cap(k)
    msgTag = "MSG_" + v_upper(k)+ "_TYPE"
    counters = Counter([ie["ie_type"] for ie in msg_list[k]["ies"]])
    dupIes = [k for k, c in counters.items() if c>1]
    repIes = []
    for ie in msg_list[k]["ies"]:
        if ie["ie_type"] in dupIes:
            newIe = ie
            newIe["ie_value"] = ie["ie_type"] + " List"
            newIe["tlv_more"] = 5 #set to a non-zero number
            newIe["collapsed"] = True
            type_list[ie["ie_type"]]["max_tlv_more"] = 5
            repIes.append(newIe)
            break

    newIes = [ie for ie in msg_list[k]["ies"] if ie["ie_type"] not in dupIes]
    newIes.extend(repIes)
    msg_list[k]["ies"] = newIes



#write list of types to file if it has not been created
field_ies_file = 'field-ies.py'
if not os.path.isfile(field_ies_file):
    f = open(field_ies_file, 'w')
    f.write("custom_types = [\n")
    f.write(",\n".join([ "#\""+k+"\"" for (k, v) in sorted_type_list if k not in\
        group_list]))
    f.write("\n]\n")
    f.close()

#load the edited types file (uncomment types for generating its data structure
exec(open(field_ies_file).read())

#write data structure template for custom type, not-overwrinting
for t in custom_types:
    ieType = v_cap(t)
    ieSize = 0
    if "size" in type_list[t]:
        ieSize = type_list[t]["size"]

    write_custom_type(ieType, ieSize)


for (k, v) in sorted_type_list:
    if k in custom_types:
        print("%s\n" % k)

#set size for attribute in grouped IEs
for (k, v) in sorted_group_list:
    for ie in group_list[k]["ies"]:
        ieKey = ie["ie_type"]
        ieSize = 0
        if "size" in type_list[ieKey]:
            ieSize = type_list[ieKey]["size"]

        if ieSize == 0: 
           ie["goType"] = "[]byte"
           ie["size"] = 0 
        elif ieSize == 1:
           ie["size"] = 1 
           ie["goType"] = "uint8"
        elif ieSize == 2:
           ie["size"] = 2 
           ie["goType"] = "uint16"
        elif ieSize == 3:
           ie["size"] = 3 
           ie["goType"] = "uint32"
        elif ieSize == 4:
           ie["size"] = 4 
           ie["goType"] = "uint32"
        else:
            assert False, "Unknown size = %d for key = %s" % (ieSize, ieKey)
 
# pointer_fields are mandatory fields of grouped IEs generated as pointers, as
# optional ones are. None: a mandatory IE is a value, and a decoder reports one
# that is missing (write_mandatory_checks) rather than leaving a nil to find.
# Create PDR's PDI and Create QER's Gate Status were made pointers by hand and
# are values again
pointer_fields = set()

# groupPresence is the presence a field of grouped IE g is generated with
def groupPresence(g, ie, customType):
    if customType and ie["presence"] == "M" and (g, ie["ie_value"]) in pointer_fields:
        return "C"
    return ie["presence"]

tmp = []
# Two rows of one IE type in a grouped IE: the Local F-TEID and the Local F-TEID
# for Redundant Transmission of Created PDR, Created Traffic Endpoint and Updated
# PDR, both of type F-TEID. TS 29.244 tells them apart by neither an order nor
# any marker, and Updated PDR's table lists them the other way round. So this is
# an assumption, stated once: the first occurrence on the wire is the Local
# F-TEID and a second the one for redundant transmission, and they are encoded
# in that order whatever the table's. Any other pair of one type stops the
# generator, so that a new release cannot add one unseen
redundant_suffix = " for Redundant Transmission"

def paired_rows(k):
    pairs = {}
    rows = group_list[k]["ies"]
    for ie in rows:
        same = [r for r in rows if r["ie_type"] == ie["ie_type"]]
        if len(same) < 2:
            continue
        first = [r for r in same if not r["ie_value"].endswith(redundant_suffix)]
        second = [r for r in same if r["ie_value"].endswith(redundant_suffix)]
        assert len(same) == 2 and len(first) == 1 and len(second) == 1, \
            "%s: rows %s share IE type %s" % (k, [r["ie_value"] for r in same], ie["ie_type"])
        pairs[ie["ie_type"]] = (first[0], second[0])
    return pairs

# encoding_order is the rows in the table's order, but for the second of a
# pair, moved to just after the first
def encoding_order(k, pairs):
    out = []
    for ie in group_list[k]["ies"]:
        if ie["ie_type"] in pairs:
            first, second = pairs[ie["ie_type"]]
            if ie is first:
                out.extend([first, second])
            continue
        out.append(ie)
    return out

def custom_type_of(ieKey):
    if ieKey in group_list or ieKey in custom_types:
        return v_cap(ieKey)
    return None

#write group IEs data structures
for (k, v) in sorted_group_list:
    groupName = v_cap(k)
    groupType = groupName
    rows = group_list[k]["ies"]
    pairs = paired_rows(k)

    f = io.StringIO()

    f.write("\n/* \'" + k + "\' data structure */\n")
    f.write("type " + groupType + " struct {\n")
    for ie in rows:
        ieKey = ie["ie_type"]
        customType = custom_type_of(ieKey)
        ieType = customType if customType else ie["goType"]
        field_decl(f, v_cap(ie["ie_value"]), ieType, ieKey, repeats(k, ie),
                   groupPresence(k, ie, customType) == "M" or ieType == "[]byte")
    unknown_field(f)
    f.write("}\n")
    f.write("\n")

    #fielder
    f.write("func (t *" + groupType + ") Field() tlv.Field {\n")
    f.write("fields := []tlv.Field{}\n")
    for ie in encoding_order(k, pairs):
        ieKey = ie["ie_type"]
        ieName = v_cap(ie["ie_value"])
        ieTag = v_upper(ieKey) + "_TYPE"
        customType = custom_type_of(ieKey)
        if repeats(k, ie):
            write_field_multi_ie(f, ieName, ieTag, customType, ieKey in group_list, ie["size"])
        else:
            write_field_ie(f, ieName, ieTag, customType, ieKey in group_list, ie["size"], groupPresence(k, ie, customType))
    f.write("return tlv.TLV(" + v_upper(k)+ "_TYPE, fields...)\n")
    f.write("}\n")
    f.write("\n")

    #binary unmashaller
    f.write("func (t *" + groupType + ") UnmarshalBinary(buf []byte) (err error) {\n")
    f.write("*t = " + groupType + "{}\n")
    f.write("elements, err := tlv.Elements(buf)\nif err != nil {\nreturn\n}\n")
    singles = [ie for ie in rows if not repeats(k, ie)]
    if singles:
        f.write("var seen [%d]bool\n" % len(singles))
    f.write("for _, de := range elements {\n")
    f.write("switch de.Type {\n")
    for ie in rows:
        ieKey = ie["ie_type"]
        ieName = v_cap(ie["ie_value"])
        customType = custom_type_of(ieKey)
        ieTag = v_upper(ieKey) + "_TYPE"
        if ieKey in pairs:
            first, second = pairs[ieKey]
            if ie is not first:
                continue
            # the first occurrence fills the first row, a second the second
            a, b = singles.index(first), singles.index(second)
            f.write("case " + ieTag + ":\n")
            f.write("if !seen[%d] {\nseen[%d] = true\n" % (a, a))
            write_unmarshal_ie(f, v_cap(first["ie_value"]), customType, first["size"], groupPresence(k, first, customType), ieKey in group_list, ieKey)
            f.write("} else if !seen[%d] {\nseen[%d] = true\n" % (b, b))
            write_unmarshal_ie(f, v_cap(second["ie_value"]), customType, second["size"], groupPresence(k, second, customType), ieKey in group_list, ieKey)
            f.write("} else {\ncontinue\n}\n")
            continue
        f.write("case " + ieTag + ":\n")
        if repeats(k, ie):
            write_unmarshal_multi_ie(f, ieName, customType, ie["size"], ieKey in group_list, ieKey)
        else:
            write_first_only(f, singles.index(ie))
            write_unmarshal_ie(f, ieName, customType, ie["size"], groupPresence(k, ie, customType), ieKey in group_list, ieKey)

    f.write("default:\n")
    f.write("t.Unknown = append(t.Unknown, de.Type)\n")
    f.write("}\n")
    f.write("if err != nil {\n")
    f.write("return\n")
    f.write("}\n")
    f.write("}\n")

    write_mandatory_checks(f, [(ie, None if repeats(k, ie) else singles.index(ie),
                                 v_cap(ie["ie_value"]), v_upper(ie["ie_type"]) + "_TYPE")
                                for ie in rows if mandatory_on_n4(ie)])
    f.write("return\n")
    f.write("}\n")
    f.write("\n")

    write_go_file(ieDir + groupName + '.go', 'ie', f.getvalue())


# messages: a file per message and its response, in package message, every IE
# type and tag qualified by package ie, and i the loop variable, since ie is the
# package
msgFiles = {}
msgOrder = []
for (k, v) in sorted_msg_list:
    msgType = v_cap(k)
    msgTag = "ie.MSG_" + v_upper(k)+ "_TYPE"
    if "ies" not in msg_list[k]:
        continue
    fileName = re.sub('(Request|Response)$', '', msgType)
    if fileName not in msgFiles:
        msgFiles[fileName] = io.StringIO()
        msgOrder.append(fileName)
    f = msgFiles[fileName]

    for ie in msg_list[k]["ies"]:
        ieKey = ie["ie_type"]
        ie["customType"] = None
        if ieKey in group_list or ieKey in custom_types:
            ie["customType"] = "ie." + v_cap(ieKey)
            ie["goType"] = ie["customType"]
            ie["size"] = 0
        else:
            ieSize = 0
            if "size" in type_list[ieKey]:
                ieSize = type_list[ieKey]["size"]
            if ieSize not in (0, 1, 2, 3, 4):
                assert False, "Unknown size = %d for key = %s" % (ieSize, ieKey)
            ie["size"] = ieSize
            ie["goType"] = get_types(ieSize)[1]
        ie["isGroup"] = ieKey in group_list
        ie["tag"] = "ie." + v_upper(ieKey) + "_TYPE"
        ie["multi"] = repeats(k, ie)

    f.write("\n/* \'" + k + "\' data structure */\n")
    f.write("type " + msgType + " struct {\n")
    for ie in msg_list[k]["ies"]:
        ieName = v_cap(ie["ie_value"])
        ieType = ie["goType"]
        field_decl(f, ieName, ieType, ie["ie_type"], ie["multi"],
                   ie["presence"] == "M" or ieType == "[]byte")
    unknown_field(f)
    f.write("}\n\n")

    #fielder
    f.write("func (t *" + msgType + ") Field() tlv.Field {\n")
    f.write("fields := []tlv.Field{}\n")
    for ie in msg_list[k]["ies"]:
        ieName = v_cap(ie["ie_value"])
        if ie["multi"]:
            write_field_multi_ie(f, ieName, ie["tag"], ie["customType"], ie["isGroup"], ie["size"], "i")
        else:
            write_field_ie(f, ieName, ie["tag"], ie["customType"], ie["isGroup"], ie["size"], ie["presence"])
    f.write("return tlv.TLV(" + msgTag + ", fields...)\n")
    f.write("}\n\n")

    #binary unmashaller
    f.write("func (t *" + msgType + ") UnmarshalBinary(buf []byte) (err error) {\n")
    f.write("*t = " + msgType + "{}\n")
    f.write("elements, err := tlv.Elements(buf)\nif err != nil {\nreturn\n}\n")
    singles = [ie for ie in msg_list[k]["ies"] if not ie["multi"]]
    if singles:
        f.write("var seen [%d]bool\n" % len(singles))
    f.write("for _, de := range elements {\n")
    f.write("switch de.Type {\n")
    for ie in msg_list[k]["ies"]:
        ieName = v_cap(ie["ie_value"])
        f.write("case " + ie["tag"] + ":\n")
        if ie["multi"]:
            write_unmarshal_multi_ie(f, ieName, ie["customType"], ie["size"], ie["isGroup"], ie["ie_type"])
        else:
            write_first_only(f, singles.index(ie))
            write_unmarshal_ie(f, ieName, ie["customType"], ie["size"], ie["presence"], ie["isGroup"], ie["ie_type"])
    f.write("default:\n")
    f.write("t.Unknown = append(t.Unknown, de.Type)\n")
    f.write("}\n")
    f.write("if err != nil {\n")
    f.write("return\n")
    f.write("}\n")
    f.write("}\n")

    # TS 29.244 7.6.6 checks a request but a Heartbeat Request, and a response
    # only without a rejection Cause, one of 64 and above (Table 8.2.1-1): a
    # rejection may leave out what an acceptance carries
    checks = [(ie, None if ie["multi"] else singles.index(ie), v_cap(ie["ie_value"]), ie["tag"])
              for ie in msg_list[k]["ies"] if mandatory_on_n4(ie)]
    if k != "PFCP Heartbeat Request":
        cause = [c for c in checks if c[0]["ie_type"] == "Cause"]
        if k.endswith("Response") and cause:
            write_mandatory_checks(f, cause)
            rest = [c for c in checks if c not in cause]
            if rest:
                f.write("if t.%s < ie.CauseRequestRejected {\n" % cause[0][2])
                write_mandatory_checks(f, rest)
                f.write("}\n")
        else:
            write_mandatory_checks(f, checks)
    f.write("return\n")
    f.write("}\n\n")

    #tlv unmashaller
    f.write("func (t *" + msgType + ") UnmarshalTLV(typ uint32, buf []byte) (err error) {\n")
    f.write("if typ != " + msgTag + " {\n")
    f.write("err = tlv.ErrCritical\n")
    f.write("return\n")
    f.write("}\n\n")
    f.write("err = t.UnmarshalBinary(buf)\n")
    f.write("return\n")
    f.write("}\n\n")

    #json mashaller
    f.write("func (t *" + msgType + ") MarshalJSON() (buf []byte, err error) {\n")
    f.write("buf, err = jsonMarshal(t)\n")
    f.write("return\n")
    f.write("}\n\n")

    #json unmashaller
    f.write("func (t *" + msgType + ") UnmarshalJSON(buf []byte) (err error) {\n")
    f.write("*t = " + msgType + "{}\n")
    f.write("err = jsonUnmarshal(buf, t)\n")
    f.write("return\n")
    f.write("}\n\n")

for fileName in msgOrder:
    write_go_file(msgDir + fileName + '.go', 'message', msgFiles[fileName].getvalue())
