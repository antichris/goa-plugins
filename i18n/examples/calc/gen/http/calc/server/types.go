// Code generated by goa v3.5.5, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen goa.design/plugins/v3/i18n/examples/calc/design -o
// $(GOPATH)/src/goa.design/plugins/i18n/examples/calc

package server

import (
	calc "goa.design/plugins/v3/i18n/examples/calc/gen/calc"
)

// NewAddPayload builds a calc service add endpoint payload.
func NewAddPayload(a int, b int) *calc.AddPayload {
	v := &calc.AddPayload{}
	v.A = a
	v.B = b

	return v
}
