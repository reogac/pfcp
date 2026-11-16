package ie

// Constructors written by hand: the generator writes the data structures alone.

func NewForwardingParameters(dest uint8, in string, oc *OuterHeaderCreation, pid string) (far *ForwardingParameters) {
	far = &ForwardingParameters{}
	far.DestinationInterface = dest
	far.NetworkInstance = []byte(in)
	if oc != nil {
		far.OuterHeaderCreation = oc
	}
	far.ForwardingPolicy = []byte(pid)
	return
}

func NewUpdateForwardingParameters(dest uint8, in string, oc *OuterHeaderCreation, pid string, flag uint8) (far *UpdateForwardingParameters) {
	far = &UpdateForwardingParameters{}
	far.DestinationInterface = &dest
	far.NetworkInstance = []byte(in)
	if oc != nil {
		far.OuterHeaderCreation = oc
	}
	far.ForwardingPolicy = []byte(pid)
	far.PFCPSMReqFlags = &flag
	return
}
