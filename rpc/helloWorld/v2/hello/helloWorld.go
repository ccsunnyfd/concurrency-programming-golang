package hello

import "net/rpc"

// HelloServiceName is
const HelloServiceName = "path/to/pkg.HelloService"

// HelloServiceInterface is
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

// RegisterHelloService is
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

// HelloServiceClient is
type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

// DialHelloService is
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

// Hello is
func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
