package hyper

import (
	"crypto/rand"
	"fmt"

	"github.com/samuelngs/hyper/engine"
)

// Option func
type Option func(*Options)

// Options is the hyper server options
type Options struct {

	// hyper server unique id
	ID string

	// server bind to address [host:port]
	Addr string

	// HTTP protocol 1.1 / 2.0
	Protocol engine.Protocol

	// before and after funcs
	BeforeStart []func() error
	AfterStop   []func() error
}

func newID() string {
	b := new([16]byte)
	rand.Read(b[:])
	b[8] = (b[8] | 0x40) & 0x7F
	b[6] = (b[6] & 0xF) | (4 << 4)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func newOptions(opts ...Option) Options {
	opt := Options{
		ID:       newID(),
		Addr:     ":0",
		Protocol: engine.HTTP,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// ID to change server reference id
func ID(s string) Option {
	return func(o *Options) {
		o.ID = s
	}
}

// Addr to change server bind address
func Addr(s string) Option {
	return func(o *Options) {
		o.Addr = s
	}
}

// HTTP to use 1.1 network protocol
func HTTP() Option {
	return func(o *Options) {
		o.Protocol = engine.HTTP
	}
}

// HTTP2 to use 2.0 network protocol
func HTTP2() Option {
	return func(o *Options) {
		o.Protocol = engine.HTTP2
	}
}

// BeforeStart to add before start action to hyper server
func BeforeStart(f func() error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, f)
	}
}

// AfterStop to add after stop action to hyper server
func AfterStop(f func() error) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, f)
	}
}