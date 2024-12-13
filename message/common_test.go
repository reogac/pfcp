package message

import (
	"encoding/binary"
	"encoding/json"
	"github.com/reogac/pfcp/ie"
	"fmt"
	"net"
	"testing"
)

func Test_pfcp(t *testing.T) {
	fmt.Printf("Test Pfcp message encoding/decoding\n")
	nodeId := net.ParseIP("192.168.0.100")
	seid := uint64(100)
	u32 := uint32(12)
	var seidBytes [8]byte
	binary.BigEndian.PutUint64(seidBytes[:], seid)

	createPdr := []ie.CreatePDR{
		{
			PDRID:      1,
			Precedence: 100,
			PDI: &ie.PDI{
				LocalFTEID:      ie.NewFTEID(&nodeId, nil, 12),
				SourceInterface: 1,
				NetworkInstance: []byte{1},
				UEIPaddress:     ie.NewUEIPAddress(nil, &nodeId),
				SDFFilter: []ie.SDFFilter{
					*ie.NewSDFFilter("alo", "hi", "chao", "bye", 12),
				},
			},
			OuterHeaderRemoval: &ie.OuterHeaderRemoval{
				Desc: ie.OuterHeaderRemovalGtpUUdpIpv4,
			},
			URRID: []uint32{1, 2, 4},
			QERID: &u32,
		},
		{
			PDRID:      2,
			Precedence: 100,
			PDI: &ie.PDI{
				LocalFTEID:      ie.NewFTEID(&nodeId, nil, 12),
				SourceInterface: 1,
				NetworkInstance: []byte{1},
				UEIPaddress:     ie.NewUEIPAddress(nil, &nodeId),
				SDFFilter: []ie.SDFFilter{
					*ie.NewSDFFilter("alo", "hi", "chao", "bye", 12),
				},
			},
			OuterHeaderRemoval: &ie.OuterHeaderRemoval{
				Desc: ie.OuterHeaderRemovalGtpUUdpIpv4,
			},
			URRID: []uint32{1, 2, 4},
			QERID: &u32,
		},
	}
	barId := uint8(1)
	createFar := []ie.CreateFAR{
		{
			FARID: 1,
			ApplyAction: ie.ApplyAction{
				Flags: 50,
			},
			ForwardingParameters: &ie.ForwardingParameters{
				DestinationInterface: 1,
				NetworkInstance:      []byte{1, 2},
				OuterHeaderCreation:  ie.NewOuterHeaderCreationGtpIpv4(net.ParseIP("192.168.0.100"), 1234),
				ForwardingPolicy:     []byte{1, 2},
			},
			BARID: &barId,
		},
		{
			FARID: 2,
			ApplyAction: ie.ApplyAction{
				Flags: 50,
			},
			ForwardingParameters: &ie.ForwardingParameters{
				DestinationInterface: 1,
				NetworkInstance:      []byte{1, 2},
				OuterHeaderCreation:  ie.NewOuterHeaderCreationGtpIpv4(net.ParseIP("192.168.0.100"), 1234),
				ForwardingPolicy:     []byte{1, 2},
			},
			BARID: &barId,
		},
	}
	qosFlowId := uint8(5)
	refQos := uint8(5)
	createQer := []ie.CreateQER{
		{
			QERID: 1,
			GateStatus: &ie.GateStatus{
				Ul: 100,
				Dl: 100,
			},
			QoSflowidentifier: &qosFlowId,
			ReflectiveQoS:     &refQos,
			MaximumBitrate:    &ie.MBR{},
			GuaranteedBitrate: &ie.GBR{Ul: 12, Dl: 12},
		},
	}
	createUrr := []ie.CreateURR{
		{
			URRID: 12,
		},
	}
	msg := PFCPSessionEstablishmentRequest{
		NodeID:    ie.NewNodeIDv4(nodeId),
		CPFSEID:   *ie.NewFSEID(&nodeId, nil, 100),
		CreatePDR: createPdr,
		CreateFAR: createFar,
		CreateQER: createQer,
		CreateURR: createUrr,
	}
	var buf []byte
	var err error
	if buf, err = json.Marshal(&msg); err != nil {
		t.Errorf("Encoding error: %+v", err)
	}
	newMsg := PFCPSessionEstablishmentRequest{}
	if err = json.Unmarshal(buf, &newMsg); err != nil {
		t.Errorf("Decoding error: %+v", err)
	}
	// if newMsg.CreateQER[0].QERID != 1 {
	// 	t.Errorf("Wrong decoded message")
	// }
	// if newMsg.CreateFAR[0].ApplyAction.Flags != 50 {
	// 	t.Errorf("Wrong decoded message")
	// }
	// if newMsg.CreateFAR[0].ForwardingParameters.OuterHeaderCreation.Teid() != 1234 {
	// 	t.Errorf("Wrong decoded message")
	// }
	// if *newMsg.CreateQER[0].QoSflowidentifier != qosFlowId {
	// 	t.Errorf("Wrong decoded message")
	// }
	// if *newMsg.CreateQER[0].ReflectiveQoS != refQos {
	// 	t.Errorf("Wrong decoded message")
	// }

	// if newMsg.CreatePDR[0].PDI.UEIPaddress.V6().String() != "192.168.0.100" {
	// 	t.Errorf("Wrong decoded message")
	// }

	// if newMsg.CreatePDR[1].PDI.LocalFTEID.Teid() != 100 {
	// 	t.Errorf("Wrong decoded message: CreatePDR<-PDI<-FTEID<-TEID")
	// }

}
