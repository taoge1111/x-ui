package xray

import (
	"x-ui/util/json_util"
)

// Protocol
// SendThrough
// Tag
// Settings
// StreamSetting
// ProxySettings
// MuxSettings
type OutboundConfig struct {
	Protocol    string                 `json:"protocol"`
	Tag         string                 `json:"tag"`
	SendThrough json_util.RawString    `json:"sendThrough,omitempty"`
	Settings    json_util.PrettyString `json:"settings"`
}

func (c *OutboundConfig) Equals(other *OutboundConfig) bool {
	if c.Protocol != other.Protocol {
		return false
	}
	if c.Settings != other.Settings {
		return false
	}
	if c.Tag != other.Tag {
		return false
	}
	if c.SendThrough != other.SendThrough {
		return false
	}
	return true
}
