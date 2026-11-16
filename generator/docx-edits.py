"""docx-edits.py - the hand edits that turn 29244-h71.docx into
29244-h71-modified.docx, the generator's input, as a script.

    python3 docx-edits.py 29244-h71.docx 29244-h71-modified.docx
    python3 docx-edits.py --check 29244-h71.docx 29244-h71-modified.docx

The first writes the edited copy. --check edits a copy in memory and compares
every table's cell text with the second file, which is how these edits were
recovered from a copy edited by hand: the check passes when they are all of it.

Each table is found by what it holds - a caption, a header, a row's name - never
by its index, and each edit asserts what it expects to find, so a release that
moved or changed a table stops the script rather than editing the wrong one.
"""

import copy
import re
import sys

from docx import Document
from docx.table import _Row

# User Plane IP Resource Information: removed by Release 17, and still used by
# Open5GS, which this generator was ported from. Its row as Release 16 had it
UPIRI = [
    "User Plane IP Resource Information ",
    "O",
    "This IE may be present if the UP function sends this message.\n\n"
    "When present, this IE shall contain an IPv4 and/or an IPv6 address, together with a TEID "
    "range that the CP function shall use to allocate GTP-U F-TEID in the UP function.\n"
    "Several IEs with the same IE type may be present to represent multiple User Plane IP Resources.",
    "X", "X", "X", "X", "X",
    "User Plane IP Resource Information",
]

# the four messages it goes back into, and the row it goes before in each
UPIRI_AT = [
    ("PFCP Association Setup Request", "Alternative SMF IP Address"),
    ("PFCP Association Setup Response", "Alternative SMF IP Address"),
    ("PFCP Association Update Request", "PFCP Association Release Request"),
    ("PFCP Association Update Response", "UE IP Address Usage Information"),
]

# IE type names spelled as the rest of the specification spells them
IE_TYPE_NAMES = {
    "Update Non 3GPP Access Forwarding Action Information": "Update Non-3GPP Access Forwarding Action Information",
    "Provide RDS configuration information": "Provide RDS Configuration Information",
    "RDS configuration information": "RDS Configuration Information",
}

# grouped IE headers that disagree with the IE type table: a footnote marker in
# one, and type 303 named "MBS Session N4 Information" where the type table has
# "MBS Session N4mb Information"
HEADERS = [
    ("3GPP Access Forwarding Action Information 1 IE Type", "3GPP Access Forwarding Action Information IE Type"),
    ("MBS Session N4 Information IE Type = 303", "MBS Session N4mb Information IE Type = 303"),
]


def caption_key(name):
    name = re.sub(r'^(a|an) ', '', name.strip())
    return re.sub(r'^PFCP ', '', name).lower()


def message_table(document, key):
    found, caption, i = [], "", 0
    tables = list(document.tables)
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


def row_named(table, name):
    rows = [r for r in table.rows if r.cells[0].text.strip() == name]
    assert len(rows) == 1, "%d rows named %r" % (len(rows), name)
    return rows[0]


def set_row(row, texts):
    assert len(row.cells) == len(texts), "a row of %d cells, %d texts" % (len(row.cells), len(texts))
    for cell, text in zip(row.cells, texts):
        cell.text = text


def ie_type_table(document):
    tables = [t for t in document.tables if t.rows and 'IE Type value' in t.rows[0].cells[0].text]
    assert len(tables) == 1, "%d IE type tables" % len(tables)
    return tables[0]


def edit(document):
    # User Plane IP Resource Information back into four messages
    for key, before in UPIRI_AT:
        table = message_table(document, key)
        anchor = row_named(table, before)
        new = copy.deepcopy(anchor._tr)
        anchor._tr.addprevious(new)
        set_row(_Row(new, table), UPIRI)

    # and into the IE type table, in its reserved place
    types = ie_type_table(document)
    rows = [r for r in types.rows if r.cells[0].text.strip() == "116"]
    assert len(rows) == 1 and rows[0].cells[1].text.strip() == "Reserved", "type 116 is not reserved"
    rows[0].cells[1].text = "User Plane IP Resource Information"
    rows[0].cells[2].text = "Extendable / Clause 8.2.82"
    rows[0].cells[3].text = "l+1-4"
    for old, new in IE_TYPE_NAMES.items():
        rows = [r for r in types.rows if r.cells[1].text.strip() == old]
        assert len(rows) == 1, "%d IE types named %r" % (len(rows), old)
        rows[0].cells[1].text = new

    # the two grouped IE headers
    for old, new in HEADERS:
        tables = [t for t in document.tables if t.rows and any(old in c.text for c in t.rows[0].cells)]
        assert len(tables) == 1, "%d tables headed %r" % (len(tables), old)
        for cell in tables[0].rows[0].cells:
            if old in cell.text:
                cell.text = cell.text.replace(old, new)

    # IP Multicast Addressing Info has two Source IP Address rows, O on N4 and M
    # on N4mb; this library is for N4, so the N4mb one goes
    tables = [t for t in document.tables if t.rows and any(
        c.text.startswith("IP Multicast Addressing Info IE Type") for c in t.rows[0].cells)]
    assert len(tables) == 1, "%d IP Multicast Addressing Info tables" % len(tables)
    rows = [r for r in tables[0].rows if r.cells[0].text.strip() == "Source IP Address"]
    assert len(rows) == 2 and rows[1].cells[1].text.strip() == "M", "not the two Source IP Address rows"
    rows[1]._tr.getparent().remove(rows[1]._tr)


def tables_text(document):
    return [[[c.text for c in r.cells] for r in t.rows] for t in document.tables]


def main(argv):
    check = argv[:1] == ["--check"]
    if check:
        argv = argv[1:]
    if len(argv) != 2:
        sys.exit(__doc__)
    document = Document(argv[0])
    edit(document)
    if not check:
        document.save(argv[1])
        return
    got, want = tables_text(document), tables_text(Document(argv[1]))
    assert len(got) == len(want), "%d tables, want %d" % (len(got), len(want))
    bad = [i for i, (g, w) in enumerate(zip(got, want)) if g != w]
    if bad:
        for i in bad[:5]:
            for r, (g, w) in enumerate(zip(got[i], want[i])):
                if g != w:
                    print("table %d row %d:\n  got  %r\n  want %r" % (i, r, g, w))
        sys.exit("%d tables differ" % len(bad))
    print("every table matches")


if __name__ == "__main__":
    main(sys.argv[1:])
