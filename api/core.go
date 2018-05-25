package api

import "github.com/rday/zabbix"

type Zpi struct {
	api *zabbix.API
}

func NewZpi(server, user, passwd string) (*Zpi, error) {
	api, err := zabbix.NewAPI(server, user, passwd)
	zpi := &Zpi{
		api: api,
	}
	return zpi, err
}

type Data map[string]string

func Param() Data {
	return make(Data)
}
