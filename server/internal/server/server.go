package server

import (
	"context"
	"fmt"
	"github.com/ykkssyaa/DNS_Service/server/internal/consts"
	"github.com/ykkssyaa/DNS_Service/server/internal/gen"
	"os"
	"os/exec"
	"strings"
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
	content, err := os.ReadFile(consts.ResolvConfPath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var addresses []string

	for _, line := range lines {
		if strings.HasPrefix(line, "nameserver") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				addresses = append(addresses, parts[1])
			}
		}
	}

	return &gen.DnsListResponse{Addresses: addresses}, nil
}

func (s Server) AddDnsServer(ctx context.Context, dns *gen.DNS) (*gen.Empty, error) {
	f, err := os.OpenFile(consts.ResolvConfPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("nameserver %s\n", dns.Address)); err != nil {
		return nil, err
	}

	return &gen.Empty{}, nil
}

func (s Server) RemoveDnsServer(ctx context.Context, dns *gen.DNS) (*gen.Empty, error) {
	content, err := os.ReadFile(consts.ResolvConfPath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var newLines []string

	for _, line := range lines {
		if !strings.Contains(line, dns.Address) {
			newLines = append(newLines, line)
		}
	}

	err = os.WriteFile(consts.ResolvConfPath, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		return nil, err
	}

	return &gen.Empty{}, nil
}

func (s Server) mustEmbedUnimplementedDnsServiceServer() {

}
