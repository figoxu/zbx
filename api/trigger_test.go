package api

import (
	"testing"
	"github.com/quexer/utee"
	"log"
	"github.com/figoxu/Figo"
)

const (
	server=""
	user=""
	passwd=""
	host4test=""
)


func TestZpi_FindByHost(t *testing.T) {
	zpi,err:=NewZpi(server,user,passwd)
	utee.Chk(err)
	ds:=zpi.FindByHost(host4test)
	log.Println(Figo.JsonString(ds))
}


