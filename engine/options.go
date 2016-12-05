package engine

import (
	"crypto/rand"
	"fmt"

	"github.com/samuelngs/hyper/router"
)

// Option func
type Option func(*Options)

// Options is the engine server options
type Options struct {

	// engine server unique id
	ID string

	// server bind to address [host:port]
	Addr string

	// HTTP protocol 1.1 / 2.0
	Protocol Protocol

	// Router
	Router router.Service
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
		Protocol: HTTP,
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

// Proto to change http network protocol
func Proto(p Protocol) Option {
	return func(o *Options) {
		o.Protocol = p
	}
}

// Router to bind router to engine server
func Router(r router.Service) Option {
	return func(o *Options) {
		o.Router = r
	}
}