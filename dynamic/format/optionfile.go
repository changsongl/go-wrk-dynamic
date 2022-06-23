package format

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type OptionFile struct {
	Options []string
	Num     int
	File    string
}

func (r OptionFile) GetValue() string {
	if r.Num == 0 {
		return ""
	}

	return r.Options[rand.Intn(r.Num)]
}

func NewOptionFile(r interface{}) (Rule, error) {
	var file string
	if err := InterfaceToStruct(r, &file); err != nil {
		bts, _ := json.Marshal(r)
		return nil, errors.New("OptionFile struct not ok: " + string(bts))
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	contentSlice := strings.Split(string(content), "\n")

	opt := &OptionFile{Options: contentSlice, Num: len(contentSlice), File: file}

	return opt, nil
}
