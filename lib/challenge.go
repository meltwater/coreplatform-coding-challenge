package challenge

import (
	"github.com/Jeffail/benthos/v3/lib/types"
	"time"
)

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
