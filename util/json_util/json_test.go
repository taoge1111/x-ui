package json_util_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"x-ui/util/json_util"
)

type Example struct {
	Name json_util.RawString `json:"name,omitempty"`
}

func TestRawString(t *testing.T) {

	var e1, e2 Example
	e1 = Example{Name: "Gopher"}
	b1, err := json.Marshal(e1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(b1)) // 输出：{"name":"Gopher"}

	// 示例：序列化空字符串（尝试忽略）
	e2 = Example{Name: ""}
	b2, err := json.Marshal(e2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(b2)) // 输出：{}

	// 示例：反序列化非空字符串
	err = json.Unmarshal([]byte(`{"name":"Gopher"}`), &e1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%#v\n", e1) // 输出：main.Example{Name:"Gopher"}

	// 示例：反序列化空字符串（尝试忽略）
	err = json.Unmarshal([]byte(`{"name":null}`), &e2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func TestPrettyString(t *testing.T) {
	// 测试序列化
	p1 := json_util.PrettyString("Hello, world!")
	b1, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(b1))

	// 测试反序列化
	var p2 json_util.PrettyString
	err = json.Unmarshal(b1, &p2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%#v\n", p2)

	// 测试空值序列化
	p3 := json_util.PrettyString("")
	b3, err := json.Marshal(p3)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(b3))

	// 测试从 "{}" 反序列化
	err = json.Unmarshal(b3, &p2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%#v\n", p2)
}
