package restful

import (
	"github.com/gin-gonic/gin"
)

type App interface {
	AddGet(path string, handler func (*gin.Context))
	AddPost(path string, handler func (*gin.Context))
	AddDelete(path string, handler func (*gin.Context))
	AddPut(path string, handler func (*gin.Context))
	Run()
}

type restfulapp struct {
	engine *gin.Engine
	App
}

func CreateRestfulApp() App {
	engine := gin.Default()

	return &restfulapp{
		engine: engine,
	}
}

func (r *restfulapp) AddGet(path string, handler func (*gin.Context)) {
	r.engine.GET(path, handler)
}

func (r *restfulapp) AddPost(path string, handler func (*gin.Context)) {
	r.engine.POST(path, handler)
}

func (r *restfulapp) AddDelete(path string, handler func (*gin.Context)) {
	r.engine.DELETE(path, handler)
}

func (r *restfulapp) AddPut(path string, handler func (*gin.Context)) {
	r.engine.PUT(path, handler)
}

func (r *restfulapp) Run() {
	r.engine.Run()
}