package metrics

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/grafana/agent/pkg/agentflow/config"
	"github.com/grafana/agent/pkg/agentflow/types/actorstate"
	"github.com/grafana/agent/pkg/agentflow/types/exchange"
)

type Relabeler struct {
	cfg  config.Relabeler
	self *actor.PID
	outs []*actor.PID
	name string
}

func (m *Relabeler) AllowableInputs() []actorstate.InOutType {
	return []actorstate.InOutType{actorstate.Metrics}
}

func (m *Relabeler) Output() actorstate.InOutType {
	return actorstate.Metrics
}

func NewRelabeler(name string, cfg config.Relabeler) (actorstate.FlowActor, error) {
	return &Relabeler{
		cfg:  cfg,
		name: name,
	}, nil
}

func (m *Relabeler) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case actorstate.Init:
		m.outs = msg.Children
	case actorstate.Start:
		m.self = c.Self()
	case []exchange.Metric:
		metrics := make([]exchange.Metric, 0)
		for _, m := range msg {
			ls := append(m.Labels(), exchange.Label{Key: "FOO", Value: "BAR"})
			newM := exchange.NewMetric(m.Name(), m.Value(), m.Timestamp(), ls)
			metrics = append(metrics, newM)
		}
		for _, out := range m.outs {
			c.Send(out, metrics)
		}
	}
}

func (m *Relabeler) Name() string {
	return m.name
}

func (m *Relabeler) PID() *actor.PID {
	return m.self
}
