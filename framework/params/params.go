package params

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type (
	Params struct {
		json []byte
		obj  *gjson.Result
	}
)

func NewParams() *Params {
	return &Params{
		json: []byte(`{}`),
	}
}

func (p *Params) Parse() *Params {
	o := gjson.ParseBytes(p.json)
	p.obj = &o
	return p
}

func (p *Params) GetString(val string) string {
	res := p.obj.Get(val)
	if !res.Exists() {
		panic(fmt.Errorf("JSON does not contain value: '%s'", val))
	}
	if res.IsArray() || res.IsObject() {
		panic(fmt.Errorf("JSON value is object or array: '%s'", val))
	}

	return res.String()
}

func (p *Params) Set(path string, val interface{}) error {
	json, err := sjson.SetBytes(p.json, path, val)
	if err != nil {
		p.json = json
	}
	return err
}
