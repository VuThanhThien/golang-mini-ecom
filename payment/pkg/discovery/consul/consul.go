package consul

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/VuThanhThien/golang-gorm-postgres/payment/pkg/discovery"
	consul "github.com/hashicorp/consul/api"
)

type Registry struct {
	client *consul.Client
}

func NewRegistry(addr string) (*Registry, error) {
	client, err := consul.NewClient(&consul.Config{Address: addr})
	if err != nil {
		return nil, err
	}
	return &Registry{client: client}, nil
}

func (r *Registry) Register(ctx context.Context, serviceID string, serviceName string, hostPort string) error {
	parts := strings.Split(hostPort, ":")
	if len(parts) != 2 {
		return errors.New("hostPort must be in a form of <host>:<port>, example: localhost:8081")
	}
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}
	registration := &consul.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Address: parts[0],
		Port:    port,
		Check: &consul.AgentServiceCheck{
			CheckID: serviceID,
			TTL:     "5s",
		},
	}
	return r.client.Agent().ServiceRegister(registration)
}

func (r *Registry) Deregister(ctx context.Context, serviceID string, _ string) error {
	return r.client.Agent().ServiceDeregister(serviceID)
}

func (r *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	entries, _, err := r.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	} else if len(entries) == 0 {
		return nil, discovery.ErrNotFound
	}

	var addresses []string
	for _, entry := range entries {
		addresses = append(addresses, fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port))
	}
	return addresses, nil
}

func (r *Registry) ReportHealthyState(instance string, serviceName string) error {
	return r.client.Agent().PassTTL(instance, "")
}
