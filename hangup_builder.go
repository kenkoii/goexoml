// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// Autogenerated by buildergenerator

package goexoml

import (
	"errors"
)

var _ = errors.New("_")

// Setter returns setter function for the field given
func (hangupReceiver *Hangup) Setter(field string) (setter func(interface{}) (*Hangup, error)) {
	switch field {
	}
	return
}

// NewHangup return a new Hangup pointer
func NewHangup() *Hangup {
	return new(Hangup)
}

// IHangup The interface that satisfies all the methods for this struct
// IHangup asserts implementation of setters for all the fields of Hangup
type IHangup interface {
	Setter(string) func(interface{}) (*Hangup, error)
}

// AddHangup appends the verb to response
func (r *Response) AddHangup(hangup IHangup) *Response {
	r.Response = append(r.Response, hangup)
	return r
}
