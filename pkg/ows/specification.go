package ows

// Specification interface for loading the
// ows service configurations
type Specification interface {
	Service() string
	Version() string
	Validate() bool
}
