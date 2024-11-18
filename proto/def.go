package proto

type ESDSDKTarget = int

const (
	HardwareAndSoftware ESDSDKTarget = iota
	HardwareOnly
	SoftwareOnly
)
