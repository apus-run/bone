package resolver

import "google.golang.org/grpc/resolver"

type NoopResolver struct {
	cc resolver.ClientConn
}

func (r *NoopResolver) Close() {}

func (r *NoopResolver) ResolveNow(options resolver.ResolveNowOptions) {}
