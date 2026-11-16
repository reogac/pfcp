package ie

import (
	"github.com/reogac/pfcp/v2/tlv"
)

/* 'L2TP Session Information' data structure */
type L2TPSessionInformation struct {
	CallingNumber          []byte   //Calling Number
	CalledNumber           []byte   //Called Number
	MaximumReceiveUnit     []byte   //Maximum Receive Unit
	L2TPSessionIndications []byte   //L2TP Session Indications
	L2TPUserAuthentication []byte   //L2TP User Authentication IE
	Unknown                []uint32 `json:"-"` //IE types skipped: unknown, or not defined here (TS 29.244 7.6.9)
}

func (t *L2TPSessionInformation) Field() tlv.Field {
	fields := []tlv.Field{}
	if len(t.CallingNumber) > 0 {
		fields = append(fields, tlv.TLVFrom(CALLING_NUMBER_TYPE, tlv.Bytes(t.CallingNumber)))
	}
	if len(t.CalledNumber) > 0 {
		fields = append(fields, tlv.TLVFrom(CALLED_NUMBER_TYPE, tlv.Bytes(t.CalledNumber)))
	}
	if len(t.MaximumReceiveUnit) > 0 {
		fields = append(fields, tlv.TLVFrom(MAXIMUM_RECEIVE_UNIT_TYPE, tlv.Bytes(t.MaximumReceiveUnit)))
	}
	if len(t.L2TPSessionIndications) > 0 {
		fields = append(fields, tlv.TLVFrom(L2TP_SESSION_INDICATIONS_TYPE, tlv.Bytes(t.L2TPSessionIndications)))
	}
	if len(t.L2TPUserAuthentication) > 0 {
		fields = append(fields, tlv.TLVFrom(L2TP_USER_AUTHENTICATION_IE_TYPE, tlv.Bytes(t.L2TPUserAuthentication)))
	}
	return tlv.TLV(L2TP_SESSION_INFORMATION_TYPE, fields...)
}

func (t *L2TPSessionInformation) UnmarshalBinary(buf []byte) (err error) {
	*t = L2TPSessionInformation{}
	elements, err := tlv.Elements(buf)
	if err != nil {
		return
	}
	var seen [5]bool
	for _, de := range elements {
		switch de.Type {
		case CALLING_NUMBER_TYPE:
			if seen[0] {
				continue
			}
			seen[0] = true
			t.CallingNumber = de.Value
		case CALLED_NUMBER_TYPE:
			if seen[1] {
				continue
			}
			seen[1] = true
			t.CalledNumber = de.Value
		case MAXIMUM_RECEIVE_UNIT_TYPE:
			if seen[2] {
				continue
			}
			seen[2] = true
			t.MaximumReceiveUnit = de.Value
		case L2TP_SESSION_INDICATIONS_TYPE:
			if seen[3] {
				continue
			}
			seen[3] = true
			t.L2TPSessionIndications = de.Value
		case L2TP_USER_AUTHENTICATION_IE_TYPE:
			if seen[4] {
				continue
			}
			seen[4] = true
			t.L2TPUserAuthentication = de.Value
		default:
			t.Unknown = append(t.Unknown, de.Type)
		}
		if err != nil {
			return
		}
	}
	return
}
