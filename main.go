package main

import (
	"github.com/astaxie/beego"
	"github.com/anselal/webcron/app/controllers"
	"github.com/anselal/webcron/app/jobs"
	_ "github.com/anselal/webcron/app/mail"
	"github.com/anselal/webcron/app/models"
	"html/template"
	"net/http"
)

const VERSION = "1.0.0"

func main() {
	models.Init()
	jobs.InitJobs()

	// Set the default 404 page
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})

	// The debug log is not output in the production environment
	if beego.AppConfig.String("runmode") == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}
	beego.AppConfig.Set("version", VERSION)

	// Route settings
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/profile", &controllers.MainController{}, "*:Profile")
	beego.Router("/gettime", &controllers.MainController{}, "*:GetTime")
	beego.Router("/help", &controllers.HelpController{}, "*:Index")
	beego.AutoRouter(&controllers.TaskController{})
	beego.AutoRouter(&controllers.GroupController{})

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
