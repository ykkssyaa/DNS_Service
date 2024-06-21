package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/ykkssyaa/DNS_Service/server/gen"
	"github.com/ykkssyaa/DNS_Service/server/internal/consts"
	"github.com/ykkssyaa/DNS_Service/server/pkg/logger"
	"net"
	"os"
	"os/exec"
	"strings"
)

type Server struct {
	gen.UnimplementedDnsServiceServer
	logger *logger.Logger
}

func NewServer(logger *logger.Logger) *Server {
	return &Server{logger: logger}
}

func (s Server) GetHostname(ctx context.Context, empty *gen.Empty) (*gen.Hostname, error) {

	s.logger.Info.Println("GetHostname called")

	out, err := exec.Command("hostname").Output()
	if err != nil {
		s.logger.Err.Println("GetHostname error:", err)
		return nil, err
	}

	return &gen.Hostname{Name: string(out)[:len(out)-1]}, nil
}

func (s Server) SetHostname(ctx context.Context, hostname *gen.Hostname) (*gen.Empty, error) {

	s.logger.Info.Println("SetHostname called, hostname:", hostname.Name)

	cmd := exec.Command("hostnamectl", "set-hostname", hostname.Name)
	err := cmd.Run()
	if err != nil {
		s.logger.Err.Println("SetHostname error:", err)
		return nil, err
	}
	return &gen.Empty{}, nil
}

func (s Server) ListDnsServers(ctx context.Context, empty *gen.Empty) (*gen.DnsListResponse, error) {

	s.logger.Info.Println("ListDnsServers called")

	addresses, err := getDnsList()
	if err != nil {
		s.logger.Err.Println("ListDnsServers error:", err)
		return nil, err
	}

	return &gen.DnsListResponse{Addresses: addresses}, nil
}

func (s Server) AddDnsServer(ctx context.Context, dns *gen.DNS) (*gen.Empty, error) {

	s.logger.Info.Println("AddDnsServer called, dns address:", dns.Address)

	ip := net.ParseIP(dns.Address)
	if ip == nil {
		s.logger.Err.Println("invalid IP address: ", dns.Address)
		return nil, errors.New("invalid IP address")
	}

	addresses, err := getDnsList()
	if err != nil {
		s.logger.Err.Println("AddDnsServer error with getting dns list:", err)
		return nil, err
	}

	if findInSlice(dns.Address, addresses) {
		s.logger.Err.Printf("AddDnsServer error: DNS server already exists (%s)", dns.Address)
		return nil, errors.New("DNS server already exists")
	}

	f, err := os.OpenFile(consts.ResolvConfPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		s.logger.Err.Println("AddDnsServer error:", err)
		return nil, err
	}
	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("nameserver %s\n", dns.Address)); err != nil {
		return nil, err
	}

	return &gen.Empty{}, nil
}

func (s Server) RemoveDnsServer(ctx context.Context, dns *gen.DNS) (*gen.Empty, error) {

	s.logger.Info.Println("RemoveDnsServer called, dns address:", dns.Address)

	ip := net.ParseIP(dns.Address)
	if ip == nil {
		s.logger.Err.Println("invalid IP address: ", dns.Address)
		return nil, errors.New("invalid IP address")
	}

	addresses, err := getDnsList()
	if err != nil {
		s.logger.Err.Println("RemoveDnsServer error with getting dns list:", err)
		return nil, err
	}

	if !findInSlice(dns.Address, addresses) {
		s.logger.Err.Printf("RemoveDnsServer error: DNS server doesn't exist (%s)", dns.Address)
		return nil, errors.New("DNS server doesn't exist")
	}

	content, err := os.ReadFile(consts.ResolvConfPath)
	if err != nil {
		s.logger.Err.Println("RemoveDnsServer error:", err)
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
		s.logger.Err.Println("RemoveDnsServer error:", err)
		return nil, err
	}

	return &gen.Empty{}, nil
}

func (s Server) mustEmbedUnimplementedDnsServiceServer() {

}

func getDnsList() ([]string, error) {
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

	return addresses, nil
}

func findInSlice(str string, list []string) bool {

	for _, v := range list {
		if v == str {
			return true
		}
	}

	return false
}
