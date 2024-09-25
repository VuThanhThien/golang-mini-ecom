package memory

import (
	"context"
	"errors"
	"sync"
	"time"
)

type serviceName string
type instanceID string

// Registry defines an in-memory service registry.
type Registry struct {
	sync.RWMutex
	serviceAddrs map[serviceName]map[instanceID]*serviceInstance
}
type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

// NewRegistry creates a new in-memory service
// registry instance.
func NewRegistry() *Registry {
	return &Registry{serviceAddrs: map[serviceName]map[instanceID]*serviceInstance{}}
}

// Register creates a service record in the registry.
func (r *Registry) Register(ctx context.Context,
	instance string, sn string, hostPort string) error {
	r.Lock()
	defer r.Unlock()
	sName := serviceName(sn)
	iID := instanceID(instance)
	if _, ok := r.serviceAddrs[sName]; !ok {
		r.serviceAddrs[sName] = map[instanceID]*serviceInstance{}
	}
	r.serviceAddrs[sName][iID] = &serviceInstance{hostPort: hostPort, lastActive: time.Now()}
	return nil
}

// Deregister removes a service record from the
// registry.
func (r *Registry) Deregister(ctx context.Context, instance string, sn string) error {
	r.Lock()
	defer r.Unlock()
	sName := serviceName(sn)
	iID := instanceID(instance)
	if _, ok := r.serviceAddrs[sName]; !ok {
		return nil
	}
	delete(r.serviceAddrs[sName], iID)
	return nil
}

func (r *Registry) ServiceAddresses(ctx context.Context, sn string) ([]string, error) {
	r.RLock()
	defer r.RUnlock()
	sName := serviceName(sn)
	if _, ok := r.serviceAddrs[sName]; !ok {
		return nil, errors.New("service not found")
	}
	var addrs []string
	for _, instance := range r.serviceAddrs[sName] {
		addrs = append(addrs, instance.hostPort)
	}
	return addrs, nil
}

func (r *Registry) ReportHealthyState(instance string, sn string) error {
	r.Lock()
	defer r.Unlock()
	sName := serviceName(sn)
	iID := instanceID(instance)
	if _, ok := r.serviceAddrs[sName]; !ok {
		return errors.New("service is not registered yet")
	}
	if _, ok := r.serviceAddrs[sName][iID]; !ok {
		return errors.New("service instance is not registered yet")
	}
	r.serviceAddrs[sName][iID].lastActive = time.Now()
	return nil
}
