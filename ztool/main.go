package main

import (
	"flag"
	"fmt"
	"github.com/quexer/utee"
	"github.com/oliveagle/jsonpath"
	"encoding/json"
	"github.com/figoxu/Figo"
)

var (
	url  = ""
	prop = ""
)

func init() {
	flag.StringVar(&url, "url", "", "Http Address For Get Json Data")
	flag.StringVar(&prop, "prop", "", "Property Key For Get Value")
}

func main() {
	flag.Parse()
	if url == "" || prop == "" {
		fmt.Println("Use Following Command For Help './ztool --help' Or './ztool -h'")
		return
	}
	bs, err := utee.HttpGet(url)
	utee.Chk(err)
	var json_data interface{}
	json.Unmarshal(bs, &json_data)
	pv := fmt.Sprint("$.", prop)
	vs, err := jsonpath.JsonPathLookup(json_data, pv)
	utee.Chk(err)
	if v, err := Figo.TpInt(vs); err == nil {
		fmt.Print(v)
	} else {
		fmt.Print(vs)
	}
}
