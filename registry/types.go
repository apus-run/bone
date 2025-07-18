package registry

import (
	"context"
	"fmt"
	"io"
)

type Registry interface {
	Register(ctx context.Context, si ServiceInstance) error
	Deregister(ctx context.Context, si ServiceInstance) error

	ListServices(ctx context.Context, name string) ([]ServiceInstance, error)
	Subscribe(name string) <-chan Event

	io.Closer
}

type ServiceInstance struct {
	// ID is the unique instance ID as registered.
	ID string `json:"id"`
	// Name is the service name as registered.
	Name string `json:"name"`
	// Address is the service instance address.
	Addr string `json:"addr"`
	// Version is the version of the compiled.
	Version string `json:"version"`
	// Metadata is the kv pair metadata associated with the service instance.
	Metadata map[string]string `json:"metadata"`

	InitCapacity int64
	MaxCapacity  int64
	IncreaseStep int64
	GrowthRate   float64
}

func (si ServiceInstance) String() string {
	return fmt.Sprintf("%s-%s", si.Name, si.ID)
}

type EventType int

const (
	EventTypeUnknown EventType = iota
	EventTypeAdd
	EventTypeDelete
)

func (e EventType) IsAdd() bool {
	return e == EventTypeAdd
}

func (e EventType) IsDelete() bool {
	return e == EventTypeDelete
}

type Event struct {
	Type     EventType
	Instance ServiceInstance
}
