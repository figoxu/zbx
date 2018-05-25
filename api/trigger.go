package api

import (
	"github.com/quexer/utee"
	"encoding/json"
	"github.com/figoxu/Figo"
)

func (p *Zpi) FindByHost(host string) []Data {
	p.api.Login()
	param := Param()
	param["host"] = host
	resp, err := p.api.ZabbixRequest("trigger.get", param)
	utee.Chk(err)
	ds := make([]Data, 0)
	v := Figo.JsonString(resp.Result)
	json.Unmarshal([]byte(v), &ds)
	return ds
}
