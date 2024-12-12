package message

import (
	"github.com/reogac/sbi/models"
)

type AssociationRequest struct {
	Slice        models.Snssai
	DataNetworks []string
}

type AssociationResponse struct {
	TeidRange    TeidRange
	DataNetworks []DnnInfo
}

type TeidRange struct {
	LowerBound uint32
	UpperBound uint32
}

type DnnInfo struct {
	Dnn     string
	Cidr    string
	IpRange IpRange
}

type IpRange struct {
	LowerBound int64
	UpperBound int64
}
