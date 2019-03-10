// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// text HTTP client types
//
// Command:
// $ goa gen goa.design/examples/encodings/design -o
// $(GOPATH)/src/goa.design/examples/encodings

package client

import (
	text "goa.design/examples/encodings/gen/text"
)

// NewConcatstringfieldMyConcatenationOK builds a "text" service
// "concatstringfield" endpoint result from a HTTP "OK" response.
func NewConcatstringfieldMyConcatenationOK(body string) *text.MyConcatenation {
	v := body
	res := &text.MyConcatenation{
		Stringfield: &v,
	}
	return res
}

// NewConcatbytesfieldMyConcatenationOK builds a "text" service
// "concatbytesfield" endpoint result from a HTTP "OK" response.
func NewConcatbytesfieldMyConcatenationOK(body []byte) *text.MyConcatenation {
	v := body
	res := &text.MyConcatenation{
		Bytesfield: v,
	}
	return res
}