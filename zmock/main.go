package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/figoxu/gh"
)

var (
	mockData = make(map[string]int)
)

func main() {
	log.Println("Foo Bar")
	engine := mount()
	http.Handle("/", engine)
	log.Fatal(http.ListenAndServe(":10061", nil))
}

func mount() *gin.Engine {
	r := gin.Default()
	r.GET("/zbx/vs", m_gh, h_vs)
	r.POST("/zbx/set/:key/:val", m_gh, h_set)
	return r
}

//curl http://localhost:10061/zbx/vs
func h_vs(c *gin.Context) {
	c.JSON(http.StatusOK, mockData)
}

//curl http://localhost:10061/zbx/set/hello/1 -d ""
func h_set(c *gin.Context) {
	env := c.MustGet("env").(*Env)
	ph := env.ph
	k, v := ph.String("key"), ph.Int("val")
	mockData[k] = v
	c.String(http.StatusOK, "ok")
}

type Env struct {
	fh *gh.FormHelper
	ph *gh.ParamHelper
}

func m_gh(c *gin.Context) {
	c.Set("env", &Env{
		fh: gh.NewFormHelper(c),
		ph: gh.NewParamHelper(c),
	})
	c.Next()
}
