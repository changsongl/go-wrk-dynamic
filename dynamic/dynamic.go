package dynamic

import (
	"encoding/json"
	"github.com/changsongl/go-wrk-dynamic/dynamic/format"
	"strings"
)

type FieldType string

type FieldName string

const (
	FieldTypePath FieldType = "path"
	FieldTypeBody FieldType = "body"
)

type Dynamic interface {
	Generate(t FieldType, data string) string
}

type dynamic struct {
	fields []*field
}

type field struct {
	Name     FieldName     `json:"name"`
	Type     FieldType     `json:"type"`
	Format   format.Format `json:"-"`
	Rule     interface{}   `json:"rule"`
	RuleType format.Type   `json:"rule-type"`
}

func NewDynamic(d string) (Dynamic, error) {
	dy := &dynamic{}
	if d == "" {
		return dy, nil
	}

	if err := json.Unmarshal([]byte(d), &dy.fields); err != nil {
		return nil, err
	}

	if err := dy.parseAndValidate(); err != nil {
		return nil, err
	}

	return dy, nil
}

func (d *dynamic) parseAndValidate() error {
	for i, f := range d.fields {
		fm, err := format.NewFormat(f.RuleType, f.Rule)
		if err != nil {
			return err
		}

		d.fields[i].Format = fm
	}

	return nil
}

func (d *dynamic) Generate(t FieldType, data string) string {
	for _, f := range d.fields {
		if f.Type == t {
			data = strings.Replace(data, f.Name.String(), f.Format.GetValue(), -1)
		}
	}

	return data
}

func (n FieldName) String() string {
	return string(n)
}
