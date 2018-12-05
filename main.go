package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	_ "whois/routers"
	"whois/services"
)

func main() {

	beego.Trace("Loading language: " + services.Lang)
	if err := i18n.SetMessage(services.Lang, "conf/"+"locale_"+services.Lang+".ini"); err != nil {
		beego.Error("Fail to set message file: " + err.Error())
		return
	}

	log := logs.NewLogger()
	log.SetLogger(logs.AdapterFile, `{"filename":"logs/app.log"}`)

	beego.Run()
}
