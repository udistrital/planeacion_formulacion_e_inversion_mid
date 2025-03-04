package main

import (
	_ "github.com/udistrital/planeacion_formulacion_mid/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	xray "github.com/udistrital/utils_oas/xray"
	auditoria "github.com/udistrital/utils_oas/auditoria"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/udistrital/utils_oas/customerrorv2"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	AllowedOrigins := []string{"*.udistrital.edu.co"}
	if beego.BConfig.RunMode == "dev" {
		AllowedOrigins = []string{"*"}
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: AllowedOrigins,
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	xray.InitXRay()
	beego.ErrorController(&customerrorv2.CustomErrorController{})
	apistatus.Init()
	auditoria.InitMiddleware()
	beego.Run()
}
