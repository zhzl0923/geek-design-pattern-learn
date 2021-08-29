package command

import (
	"fmt"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	eventChan := make(chan string)
	defer close(eventChan)
	go func() {
		events := []string{"start", "end", "start", "end", "start", "end", "start", "end"}
		for _, event := range events {
			eventChan <- event
		}
	}()
	commands := make(chan Command, 10)
	defer close(commands)
	go func() {
		for {
			event, ok := <-eventChan
			if !ok {
				return
			}

			var command Command
			switch event {
			case "start":
				command = &StartCommand{}
			case "end":
				command = &EndCommand{}
			}

			// 将命令入队
			commands <- command
		}
	}()
	for {
		select {
		case c := <-commands:
			c.Execute()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")
			return
		}
	}
}
