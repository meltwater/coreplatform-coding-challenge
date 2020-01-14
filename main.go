package main

import (
	"github.com/Jeffail/benthos/v3/lib/log"
	"github.com/Jeffail/benthos/v3/lib/metrics"
	"github.com/Jeffail/benthos/v3/lib/processor"
	"github.com/Jeffail/benthos/v3/lib/service"
	"github.com/Jeffail/benthos/v3/lib/types"
	"github.com/meltwater/coreplatform-coding-challenge/lib/challenge"
)

func main() {
	processor.RegisterPlugin(
		"coding_challenge",
		func() interface{} {
			s := challenge.ChallengePlugin{}
			return &s
		},
		func(
			iconf interface{},
			mgr types.Manager,
			logger log.Modular,
			stats metrics.Type,
		) (types.Processor, error) {
			return iconf.(*challenge.ChallengePlugin), nil
		},
	)

	service.Run()
}
