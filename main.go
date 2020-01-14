package main

import (
	"time"

	"github.com/Jeffail/benthos/v3/lib/log"
	"github.com/Jeffail/benthos/v3/lib/metrics"
	"github.com/Jeffail/benthos/v3/lib/processor"
	"github.com/Jeffail/benthos/v3/lib/service"
	"github.com/Jeffail/benthos/v3/lib/types"
)

func main() {
	processor.RegisterPlugin(
		"coding_challenge",
		func() interface{} {
			s := ChallengePlugin{}
			return &s
		},
		func(
			iconf interface{},
			mgr types.Manager,
			logger log.Modular,
			stats metrics.Type,
		) (types.Processor, error) {
			return iconf.(*ChallengePlugin), nil
		},
	)

	service.Run()
}

//
// ChallengePlugin
//

// ChallengePlugin ...
type ChallengePlugin struct{}

// ProcessMessage returns messages mutated with their sarcasm level.
func (c *ChallengePlugin) ProcessMessage(msg types.Message) ([]types.Message, types.Response) {
	newMsg := msg.Copy()

	newMsg.Iter(func(i int, p types.Part) error {
		// Call DoSomething and store the result as the full content of the part
		result := c.DoSomething(p.Get())
		p.Set(result)
		return nil
	})

	return []types.Message{newMsg}, nil
}

// CloseAsync does nothing.
func (c *ChallengePlugin) CloseAsync() {}

// WaitForClose does nothing.
func (c *ChallengePlugin) WaitForClose(timeout time.Duration) error {
	return nil
}

// DoSomething ...
func (c *ChallengePlugin) DoSomething(data []byte) []byte {
	//
	//
	//
	// This is where you will implement the plugin
	//
	//
	//

	return data
}
