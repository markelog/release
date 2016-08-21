// Package hollow gets hollow type/name/version
package hollow

// Hollow struct
type Hollow struct{}

// New creates Hollow instance
func New() *Hollow {
	return &Hollow{}
}

// Name gets hollow name
func (hollow Hollow) Name() string {
	return ""
}

// Version gets hollow version
func (hollow Hollow) Version() string {
	return ""
}

// Type gets hollow version
func (hollow Hollow) Type() string {
	return ""
}
