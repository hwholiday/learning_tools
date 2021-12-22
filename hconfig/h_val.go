package hconfig

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type HVal []byte

func (h HVal) FormatJson(v interface{}) error {
	return json.Unmarshal(h, v)
}

func (h HVal) FormatYaml(v interface{}) error {
	return yaml.Unmarshal(h, v)
}

func (h HVal) String() string {
	return string(h)
}

func (h HVal) Bytes() []byte {
	return h
}
