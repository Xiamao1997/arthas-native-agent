package agent

import "log"

// @author Xiamao1997 2022-06-24

// NativeAgent communicates with the work node and the tunnel server.
type NativeAgent struct {
	tunnelServerIp   string
	tunnelServerPort string
	name             string
	stop             chan struct{}
}

func NewNativeAgent(ip string, port string, name string) *NativeAgent {
	return &NativeAgent{
		tunnelServerIp:   ip,
		tunnelServerPort: port,
		name:             name,
		stop:             nil,
	}
}

// Start the NativeAgent
func (na *NativeAgent) Start() error {
	log.Println("Starting the NativeAgent:", na.name)
	for {
	}
	return nil
}

// Stop the NativeAgent
func (na *NativeAgent) Stop() error {
	log.Println("Stopping the NativeAgent:", na.name)
	return nil
}
