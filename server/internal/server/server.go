package server

import (
	"context"
	"github.com/ykkssyaa/DNS_Service/server/internal/gen"
	"os/exec"
)

type Server struct {
	gen.UnimplementedDnsServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s Server) GetHostname(ctx context.Context, empty *gen.Empty) (*gen.Hostname, error) {
	out, err := exec.Command("hostname").Output()
	if err != nil {
		return nil, err
	}

	return &gen.Hostname{Name: string(out)[:len(out)-1]}, nil
}

func (s Server) SetHostname(ctx context.Context, hostname *gen.Hostname) (*gen.Empty, error) {

	cmd := exec.Command("hostnamectl", "set-hostname", hostname.Name)
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return &gen.Empty{}, nil
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
