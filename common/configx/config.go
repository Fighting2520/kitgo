package configx

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v2"
)

type IConfig interface {
	ReadConfig() Config
}

type confType int8

var ErrInvalidConfigType = errors.New("invalid type")

const (
	jsonConfigType confType = iota
	yamlConfigType
)

type Config struct {
	typ confType
	bs  []byte
}

func (c Config) Unmarshal(v interface{}) error {
	switch c.typ {
	case jsonConfigType:
		return json.Unmarshal(c.bs, v)
	case yamlConfigType:
		return yaml.Unmarshal(c.bs, v)
	default:
		return ErrInvalidConfigType
	}
}
