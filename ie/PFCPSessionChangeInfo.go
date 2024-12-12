package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'PFCP Session Change Info' data structure */
type PFCPSessionChangeInfo struct {
	PGWCSMFFQCSID               *FQCSID //FQ-CSID
	GroupId                     []byte  //Group ID
	CPIPAddress                 []byte  //CP IP Address
	AlternativeSMFPGWCIPAddress []byte  //Alternative SMF IP Address
}

func (t *PFCPSessionChangeInfo) Field() tlv.Field {
	fields := []tlv.Field{}
	if t.PGWCSMFFQCSID != nil {
		fields = append(fields, tlv.TLVFrom(FQ_CSID_TYPE, tlv.Bytes(t.PGWCSMFFQCSID.Bytes())))
	}
	if len(t.GroupId) > 0 {
		fields = append(fields, tlv.TLVFrom(GROUP_ID_TYPE, tlv.Bytes(t.GroupId)))
	}
	if len(t.CPIPAddress) > 0 {
		fields = append(fields, tlv.TLVFrom(CP_IP_ADDRESS_TYPE, tlv.Bytes(t.CPIPAddress)))
	}
	fields = append(fields, tlv.TLVFrom(ALTERNATIVE_SMF_IP_ADDRESS_TYPE, tlv.Bytes(t.AlternativeSMFPGWCIPAddress)))
	return tlv.TLV(PFCP_SESSION_CHANGE_INFO_TYPE, fields...)
}

func (t *PFCPSessionChangeInfo) UnmarshalBinary(buf []byte) (err error) {
	*t = PFCPSessionChangeInfo{}
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case FQ_CSID_TYPE:
			var v FQCSID
			if err = de.UnmarshalValue(&v); err == nil {
				t.PGWCSMFFQCSID = &v
			}
		case GROUP_ID_TYPE:
			t.GroupId = de.Value
		case CP_IP_ADDRESS_TYPE:
			t.CPIPAddress = de.Value
		case ALTERNATIVE_SMF_IP_ADDRESS_TYPE:
			t.AlternativeSMFPGWCIPAddress = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
