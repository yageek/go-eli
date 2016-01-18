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
	Type   PackageType
	Length uint8
}
