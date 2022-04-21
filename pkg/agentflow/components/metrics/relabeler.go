package metrics

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/grafana/agent/pkg/agentflow/config"
)

type Relabeler struct {
	cfg  config.Relabeler
	self *actor.PID
	outs []*actor.PID
	name string
}
