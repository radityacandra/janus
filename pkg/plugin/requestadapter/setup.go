package requestadapter

import (
	"github.com/hellofresh/janus/pkg/plugin"
	"github.com/hellofresh/janus/pkg/proxy"
)

func init() {
	plugin.RegisterPlugin("request_adapter", plugin.Plugin{
		Action: setupRequestAdapter,
	})
}

func setupRequestAdapter(def *proxy.RouterDefinition, rawConfig plugin.Config) error {
	var config Config
	err := plugin.Decode(rawConfig, &config)
	if err != nil {
		return err
	}

	def.AddMiddleware(NewRequestAdapter(config))
	return nil
}
