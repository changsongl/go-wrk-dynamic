package format

import (
	"encoding/json"
	"fmt"
)

type Type string

const (
	TypeRangeInt       Type = "range-int"
	TypeRangeDouble    Type = "range-double"
	TypeOption         Type = "option"
	TypeOptionFromFile Type = "option-from-file"
)

type format struct {
	Type     Type        `json:"type"`
	Rule     interface{} `json:"rule"`
	RuleFunc Rule        `json:"-"`
}

type Rule interface {
	GetValue() string
}

type Format interface {
	GetValue() string
}

func NewFormat(t Type, r interface{}) (Format, error) {
	formatInst := format{
		Type: t,
		Rule: r,
	}

	if err := formatInst.ParseRule(); err != nil {
		return nil, err
	}

	return formatInst.RuleFunc, nil
}

func (f *format) ParseRule() (err error) {
	switch f.Type {
	case TypeRangeInt:
		f.RuleFunc, err = NewRangeInt(f.Rule)
		if err != nil {
			return err
		}

		return nil
	case TypeOptionFromFile:
		return nil
	//case TypeRangeDouble:
	//	fallthrough
	//case TypeOption:
	//	fallthrough
	default:
		return fmt.Errorf("rule is not support.(%s)", f.Type)
	}
}

func InterfaceToStruct(from, to interface{}) error {
	bts, err := json.Marshal(from)
	if err != nil {
		return err
	}

	return json.Unmarshal(bts, to)
}
