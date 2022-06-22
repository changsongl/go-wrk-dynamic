package format

import (
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
)

type RangeInt struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

func (r RangeInt) GetValue() string {
	if r.To < r.From {
		return "0"
	} else if r.To == r.From {
		return strconv.FormatInt(r.To, 64)
	}

	return strconv.FormatInt(rand.Int63n(r.To-r.From)+r.From, 10)
}

func NewRangeInt(r interface{}) (Rule, error) {
	rangeRule := &RangeInt{}
	if err := InterfaceToStruct(r, rangeRule); err != nil {
		bts, _ := json.Marshal(r)
		return nil, errors.New("RangeInt struct not ok: " + string(bts))
	}

	return rangeRule, nil
}
