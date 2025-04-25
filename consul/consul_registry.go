package consul

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/hashicorp/consul/api"
	"net"
)

// ConsulRegistry 实现 registry.Registry 接口
type ConsulRegistry struct {
	client *api.Client
}

// NewConsulRegistry 创建一个 Consul 注册中心实例
func NewConsulRegistry(address string) (*ConsulRegistry, error) {
	config := api.DefaultConfig()
	config.Address = address
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &ConsulRegistry{client: client}, nil
}

// Register 实现服务注册
func (r *ConsulRegistry) Register(info *registry.Info) error {
	// 从 Info.Addr 中获取地址和端口
	var ip string
	var port int

	switch addr := info.Addr.(type) {
	case *net.TCPAddr:
		ip = addr.IP.String()
		port = addr.Port
	case *net.UDPAddr:
		ip = addr.IP.String()
		port = addr.Port
	default:
		return fmt.Errorf("unsupported address type: %T", info.Addr)
	}
	registration := &api.AgentServiceRegistration{
		ID:      info.ServiceName,
		Name:    info.ServiceName,
		Port:    port,
		Address: ip,
		Check: &api.AgentServiceCheck{
			TCP:      fmt.Sprintf("%s:%d", ip, port), // 使用 TCP 检查
			Interval: "10s",
			Timeout:  "1s",
		},
	}
	return r.client.Agent().ServiceRegister(registration)
}

// Deregister 实现服务注销
func (r *ConsulRegistry) Deregister(info *registry.Info) error {
	return r.client.Agent().ServiceDeregister(info.ServiceName)
}
