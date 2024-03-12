package json_util

import (
	"errors"
	"strconv"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type XrayConfigUtil struct{}

func (xray XrayConfigUtil) Vaild(rawStr string) bool {

	return gjson.Valid(rawStr)
}

func (xray XrayConfigUtil) RouterIsExist(rawStr string) bool {

	obj := gjson.Parse(rawStr)

	return obj.Get("routing").Exists()
}

func (xray XrayConfigUtil) AddRouterRules(rawStr string, ruleObj string) (string, error) {

	if !xray.Vaild(ruleObj) || !xray.Vaild(rawStr) {
		return "", errors.New("invaild json")
	}

	if !xray.RouterIsExist(rawStr) {
		return "", errors.New("no reoutes exist")
	}

	obj := gjson.Parse(rawStr)

	ruleObjects := obj.Get("routing.rules")

	res, err := sjson.SetRaw(rawStr, "routing.rules"+"."+strconv.Itoa(len(ruleObjects.Array())), ruleObj)

	return res, err
}

func (xray XrayConfigUtil) DeleteRouterRules(rawStr string, ruleObj string) (string, error) {

	if !xray.Vaild(ruleObj) || !xray.Vaild(rawStr) {
		return "", errors.New("invalid json")
	}
	if !xray.RouterIsExist(rawStr) {
		return "", errors.New("no reoutes exist")
	}

	routerRuleObj := gjson.Parse(ruleObj)

	index := -1

	for pos, item := range gjson.Get(rawStr, "routing.rules").Array() {
		inboundTags := item.Get("inboundTag")
		outboundTags := item.Get("outboundTag")
		if inboundTags.Exists() && outboundTags.Exists() {
			if EqualsInboundTag(inboundTags.Array(), routerRuleObj.Get("inboundTag").Array()) &&
				EqualsOutboundTag(outboundTags.String(), routerRuleObj.Get("outboundTag").String()) {
				index = pos
				break
			}
		}
	}

	if index == -1 {
		return "", errors.New("the router rules not found" + ruleObj)
	}

	res, err := sjson.Delete(rawStr, "routing.rules"+"."+strconv.Itoa(index))

	return res, err

}

func EqualsInboundTag(l []gjson.Result, r []gjson.Result) bool {
	if len(l) != len(r) {
		return false
	}
	for i := 0; i < len(r); i++ {
		if l[i].String() != r[i].String() {
			return false
		}
	}
	return true
}

func EqualsOutboundTag(l, r string) bool {
	return l == r
}
