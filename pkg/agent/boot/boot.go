package boot

import (
	"errors"
	"github.com/Xiamao1997/arthas-native-agent/pkg/agent/agent"
	"github.com/satori/go.uuid"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @author Xiamao1997 2022-06-20

// Start Native Agent
func Start(ip string, port string, name string) error {
	if ip == "" {
		return errors.New("tunnel server ip is empty")
	}
	if port == "" {
		return errors.New("tunnel server port is empty")
	}
	if name == "" {
		name = uuid.NewV4().String()
		log.Println("Randomly generated the native agent name:", name)
	}

	nativeAgent := agent.NewNativeAgent(ip, port, name)

	// Monitor system signals
	sigChan := NewOSWatcher(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case sig := <-sigChan:
			nativeAgent.Stop()
			log.Fatalf("Received signal %v, shutting down.", sig)
		}
	}()

	if err := nativeAgent.Start(); err != nil {
		return err
	}

	return nil
}

func NewOSWatcher(sigs ...os.Signal) chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, sigs...)

	return sigChan
}
