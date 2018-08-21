package main

import (
	_ "ihome/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
	"net/http"
	_"ihome/models"
)

func main() {

	ignnoreStaticPath()
	beego.Run()
}
func ignnoreStaticPath(){
	beego.InsertFilter("/",beego.BeforeRouter,TransparentStatic)
	beego.InsertFilter("/*",beego.BeforeRouter,TransparentStatic)
}
func TransparentStatic(ctx *context.Context){
	orpath:=ctx.Request.URL.Path
	beego.Debug("request url:",orpath)
	if strings.Index(orpath,"api")>=0{
		return
	}
	beego.Info("responseWrite: ",ctx.ResponseWriter)
	beego.Info("requset:",ctx.Request)
	http.ServeFile(ctx.ResponseWriter,ctx.Request,"static/html/"+ctx.Request.URL.Path)
}
