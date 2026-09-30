package message

type AssociationRequest struct {
	Slice        Snssai
	DataNetworks []string
}

// Snssai is the slice an association is for. It is declared here rather than
// taken from the SBI models so that this library imports no SBI code: the SBI
// library's N4 service carries these messages, and a dependency back would tie
// the two into a cycle. The JSON tags are the models' own, so the encoding on
// the wire is unchanged.
type Snssai struct {
	Sd  string `json:"sd,omitempty"`
	Sst int    `json:"sst"`
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
