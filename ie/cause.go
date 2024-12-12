package ie

const (
	CauseNotDefined                      uint8 = 0
	CauseRequestAccepted                 uint8 = 1
	CauseRequestRejected                 uint8 = 64
	CauseSessionContextNotFound          uint8 = 65
	CauseMandatoryIEMissing              uint8 = 66
	CauseConditionalIEMissing            uint8 = 67
	CauseInvalidLength                   uint8 = 68
	CauseMandatoryIEIncorrect            uint8 = 69
	CauseInvalidForwardingPolicy         uint8 = 70
	CauseInvalidFTEIDAllocationOption    uint8 = 71
	CauseNoEstablishedPFCPAssociation    uint8 = 72
	CauseRuleCreationModificationFailure uint8 = 73
	CausePFCPEntityInCongestion          uint8 = 74
	CauseNoResourcesAvailable            uint8 = 75
	CauseServiceNotSupported             uint8 = 76
	CauseSystemFailure                   uint8 = 77
	CauseRedirectionRequested            uint8 = 78
)
