package encoding

// PackageType represents a package type
type PackageType uint8

// All the different existing
// package type.
const (
	FileDescriptorType                PackageType = 0x01
	PhysicalDescriptorType                        = 0x011
	RateResolutionType                            = 0x21
	CalibrationConfigurationType                  = 0x22
	CalibrationConfigurationFixedType             = 0x26
	A4BilinearType                                = 0x23
	TimeDateType                                  = 0x31
	PositionDelayType                             = 0x41
	PositionDistanceType                          = 0x52
	PositionCoordinationType                      = 0x61
	DataTestPointType                             = 0x02
	ThrownPointsType                              = 0xa1
	TimeoutType                                   = 0xb1
	StatusChangeType                              = 0xf1
	UnknownType                                   = 0xff
)

// Record represents an
// ELI Generic Record
type Record struct {
	Type    PackageType
	Length  uint8
	Payload []byte
}

type FileDescriptor struct {
	FileType     byte
	PositionData byte
	CancelFlag   byte
	_            byte
}

type PhysicalParameters struct {
	MiscSepr  byte
	LeftBias  uint16
	RightBias uint16
	AirTemp   byte
	_         byte
}

type RateResolution struct {
	FSamp           uint32
	Log2NFFT        byte
	DistResolution  uint16
	CoordResolution uint16
	Timing          uint32
	_               [2]byte
}

type CalibrationConfiguration struct {
	Padding [9]byte
	_       [3]byte
	AL1     [3]byte
	BL1     [3]byte
	AR1     [3]byte
	BR1     [3]byte
	AL2     [3]byte
	BL2     [3]byte
	AR2     [3]byte
	BR2     [3]byte
}

type A4Bilinear struct {
	Algorithm [32]byte
	_         [2]byte
}

type DateTime struct {
	TimeH byte
	TimeM byte
	TimeS byte
	DateD byte
	DateM byte
	DateY byte
	_     byte
}
