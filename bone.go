package bone

import (
	"context"
	"sync"

	"github.com/apus-run/bone/registry"
)

type App interface {
	ID() string
	Name() string
	Version() string
	Metadata() map[string]string
	Endpoint() string
}

type Bone struct {
	opts   *Options
	ctx    context.Context
	cancel func()

	mux      sync.Mutex
	instance *registry.ServiceInstance
}

func New(opts ...Option) *Bone {
	options := Apply(opts...)
	ctx, cancel := context.WithCancel(context.Background())
	return &Bone{
		opts:   options,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (b *Bone) RegisterService() *Options {
	return b.opts
}

// ID returns app instance id.
func (b *Bone) ID() string { return b.opts.id }

// Name returns service name.
func (b *Bone) Name() string { return b.opts.name }

// Version returns app version.
func (b *Bone) Version() string { return b.opts.version }

// Metadata returns service metadata.
func (b *Bone) Metadata() map[string]string { return b.opts.metadata }

// Endpoint returns endpoints.
func (b *Bone) Endpoint() string {
	if b.instance != nil {
		return b.instance.Addr
	}
	return ""
}
