package bone

import (
	"context"
	"net/url"

	"github.com/apus-run/bone/registry"
	"github.com/apus-run/bone/server"
)

type Option func(*Options)

type Options struct {
	// ID is the unique identifier for the service instance.
	id string
	// Name is the name of the service.
	name string
	// Version is the version of the service.
	version string
	// Metadata is the metadata associated with the service.
	metadata map[string]string

	addr *url.URL

	ctx context.Context

	registry registry.Registry

	servers []server.Server
}

func NewOptions() *Options {
	return &Options{
		metadata: make(map[string]string),
	}
}

func Apply(opts ...Option) *Options {
	options := NewOptions()
	for _, opt := range opts {
		opt(options)
	}
	return options
}

func WithID(id string) Option {
	return func(o *Options) {
		o.id = id
	}
}

func WithName(name string) Option {
	return func(o *Options) {
		o.name = name
	}
}

func WithVersion(version string) Option {
	return func(o *Options) {
		o.version = version
	}
}

// WithMetadata with service metadata.
func WithMetadata(md map[string]string) Option {
	return func(o *Options) {
		o.metadata = md
	}
}

// WithEndpoint with service endpoint.
func WithEndpoint(addr *url.URL) Option {
	return func(o *Options) {
		o.addr = addr
	}
}

// WithRegistry with service registry.
func WithRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.registry = r
	}
}
