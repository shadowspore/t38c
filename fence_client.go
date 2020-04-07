package t38c

import "log"

type FenceClient struct {
	debug bool
	exec  FenceExecutor
}

func NewFence(dialer FenceExecutorDialer) (*FenceClient, error) {
	executor, err := dialer()
	if err != nil {
		return nil, err
	}

	client := &FenceClient{
		exec: executor,
	}

	return client, nil
}

func (client *FenceClient) FenceExecute(command string, args ...string) (FenceChan, error) {
	ch, err := client.exec.Fence(command, args...)
	if err != nil {
		cmd := command
		if len(args) > 0 {
			for _, arg := range args {
				cmd += " " + arg
			}
		}

		log.Printf("[%s]: %s", cmd, err)
	}
	return ch, err
}
