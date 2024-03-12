package json_util_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"x-ui/util/json_util"

	"github.com/stretchr/testify/assert"
)

var config_example = `
{
  "log": {
    "loglevel": "warning"
  },
  "routing": {
    "domainStrategy": "AsIs",
    "rules": []
  },
  "dns": null,
  "inbounds": [
    {
      "listen": "127.0.0.1",
      "port": 62789,
      "protocol": "dokodemo-door",
      "settings": {
        "address": "127.0.0.1"
      },
      "streamSettings": null,
      "tag": "api",
      "sniffing": null
    }
  ],
  "outbounds": [
    {
      "protocol": "freedom",
      "tag": "direct",
      "settings": {}
    },
    {
      "protocol": "blackhole",
      "tag": "blocked",
      "settings": {}
    }
  ],
  "transport": null,
  "policy": {
    "levels": {
      "0": {
        "statsUserDownlink": true,
        "statsUserUplink": true
      }
    },
    "system": {
      "statsInboundDownlink": true,
      "statsInboundUplink": true
    }
  },
  "api": {
    "tag": "api",
    "services": [
      "HandlerService",
      "LoggerService",
      "StatsService"
    ]
  },
  "stats": {},
  "reverse": null,
  "fakedns": null,
  "observatory": null
}
`

type ExmapleRouterRule struct {
	Type        string   `json:"type"`
	InboundTag  []string `json:"inboundTag"`
	OutboundTag string   `json:"outboundTag"`
}

func TestXrayConfig(t *testing.T) {
	util := json_util.XrayConfigUtil{}

	rule := ExmapleRouterRule{
		Type:        "field",
		InboundTag:  []string{"api"},
		OutboundTag: "direct",
	}

	ruleObjStrClean, err := json.MarshalIndent(rule, "", " ")

	assert.Nil(t, err, nil)

	res, _ := util.AddRouterRules(config_example, string(ruleObjStrClean))
	// res, _ = util.AddRouterRules(res, string(ruleObjStrClean))

	res, _ = util.DeleteRouterRules(res, string(ruleObjStrClean))

	fmt.Println(res)
}
