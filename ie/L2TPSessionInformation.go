package ie

import (
	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

/* 'L2TP Session Information' data structure */
type L2TPSessionInformation struct {
	CallingNumber          []byte //Calling Number
	CalledNumber           []byte //Called Number
	MaximumReceiveUnit     []byte //Maximum Receive Unit
	L2TPSessionIndications []byte //L2TP Session Indications
	L2TPUserAuthentication []byte //L2TP User Authentication IE
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
	d := tlv.DecodingBuffer(buf)
	for _, de := range d.Elements() {
		switch de.Type {
		case CALLING_NUMBER_TYPE:
			t.CallingNumber = de.Value
		case CALLED_NUMBER_TYPE:
			t.CalledNumber = de.Value
		case MAXIMUM_RECEIVE_UNIT_TYPE:
			t.MaximumReceiveUnit = de.Value
		case L2TP_SESSION_INDICATIONS_TYPE:
			t.L2TPSessionIndications = de.Value
		case L2TP_USER_AUTHENTICATION_IE_TYPE:
			t.L2TPUserAuthentication = de.Value
		default:
			err = tlv.ErrCritical
		}
		if err != nil {
			return
		}
	}
	return
}
