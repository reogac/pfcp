package message

import (
	"encoding/json"

	"github.com/usnistgov/ndn-dpdk/ndn/tlv"
)

func jsonMarshal(t tlv.Fielder) (buf []byte, err error) {
	tmp := struct {
		Buf []byte
	}{}
	if tmp.Buf, err = tlv.EncodeFrom(t); err != nil {
		return
	}
	buf, err = json.Marshal(&tmp)
	return
}

func jsonUnmarshal(buf []byte, v tlv.Unmarshaler) (err error) {
	tmp := struct {
		Buf []byte
	}{}
	if err = json.Unmarshal(buf, &tmp); err != nil {
		return
	}
	err = tlv.Decode(tmp.Buf, v)
	return
}
