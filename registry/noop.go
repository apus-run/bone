package registry

import "context"

// NoopRegistry is an empty implement of Registry
var NoopRegistry Registry = &noopRegistry{}

// NoopRegistry
type noopRegistry struct{}

func (r *noopRegistry) Register(ctx context.Context, si ServiceInstance) error {
	return nil
}

func (r *noopRegistry) Deregister(ctx context.Context, si ServiceInstance) error {
	return nil
}

func (r *noopRegistry) ListServices(ctx context.Context, name string) ([]ServiceInstance, error) {
	return nil, nil
}

func (r *noopRegistry) Subscribe(name string) <-chan Event {
	return nil
}

func (r *noopRegistry) Close() error {
	return nil
}
