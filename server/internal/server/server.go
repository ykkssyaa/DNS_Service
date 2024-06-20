package server

import (
	"context"
	"github.com/ykkssyaa/DNS_Service/server/internal/gen"
)

type Server struct {
	gen.UnimplementedDnsServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s Server) GetHostname(ctx context.Context, empty *gen.Empty) (*gen.Hostname, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) SetHostname(ctx context.Context, hostname *gen.Hostname) (*gen.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) ListDnsServers(ctx context.Context, empty *gen.Empty) (*gen.DnsListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) AddDnsServer(ctx context.Context, dns *gen.DNS) (*gen.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) RemoveDnsServer(ctx context.Context, dns *gen.DNS) (*gen.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) mustEmbedUnimplementedDnsServiceServer() {

}
