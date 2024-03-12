package json_util

import (
	"encoding/json"
	"errors"
)

type RawMessage []byte

// MarshalJSON: Customize json.RawMessage default behavior
func (m RawMessage) MarshalJSON() ([]byte, error) {
	if len(m) == 0 {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON: sets *m to a copy of data.
func (m *RawMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

type RawString string

func (m RawString) MarshalJSON() ([]byte, error) {
	if m == "" {
		return []byte("null"), nil
	}
	return json.Marshal(string(m))
}

func (m *RawString) UnmarshalJSON(data []byte) error {
	// 如果 JSON 值是 null，设置 MyString 为默认的空字符串并返回。
	if string(data) == "null" {
		*m = RawString("")
		return nil
	}
	// 尝试反序列化为普通字符串。
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*m = RawString(s)
	return nil
}

// PrettyString 定义了一个自定义的字符串类型。
type PrettyString string

// MarshalJSON 自定义了 PrettyString 的 JSON 序列化行为。
func (p PrettyString) MarshalJSON() ([]byte, error) {
	if p == "" {
		// 当 PrettyString 为空时，序列化为 {}
		return []byte("{}"), nil
	}
	// 非空时，序列化为带有换行符和缩进的 JSON 字符串
	// 注意：这里我们把 PrettyString 当作普通字符串处理，并加上缩进。
	// 你可能需要根据实际情况调整这里的逻辑。
	b, err := json.MarshalIndent(string(p), "", "  ")
	if err != nil {
		return nil, err
	}
	// 为了符合你的要求，我们可能需要在序列化的字符串后添加换行符。
	return append(b, '\n'), nil
}

// UnmarshalJSON 自定义了 PrettyString 的 JSON 反序列化行为。
func (p *PrettyString) UnmarshalJSON(data []byte) error {
	// 反序列化逻辑
	// 这里只是简单地将 JSON 字符串转换回 PrettyString。
	// 注意检查 data 是否为 "{}"，以正确处理空值的情况。
	if string(data) == "{}" {
		*p = ""
		return nil
	}
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*p = PrettyString(s)
	return nil
}
