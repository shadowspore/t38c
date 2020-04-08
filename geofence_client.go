package t38c

import (
	"log"
)

type FenceClient struct {
	debug bool
	exec  FenceExecutor
}

func NewFence(dialer FenceExecutorDialer, debug bool) (*FenceClient, error) {
	executor, err := dialer()
	if err != nil {
		return nil, err
	}

	client := &FenceClient{
		exec:  executor,
		debug: debug,
	}

	return client, nil
}

func (client *FenceClient) FenceExecute(command string, args ...string) (FenceChan, error) {
	ch, err := client.exec.Fence(command, args...)
	if client.debug {
		cmd := command
		if len(args) > 0 {
			for _, arg := range args {
				cmd += " " + arg
			}
		}
		if err != nil {
			log.Printf("[%s]: %s", cmd, err)
		} else {
			log.Printf("[%s]", cmd)
		}
	}

	return ch, err
}
