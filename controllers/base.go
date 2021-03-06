package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/yunkaiyueming/MonShark/helpers"
	"github.com/yunkaiyueming/MonShark/models"
	"gopkg.in/mgo.v2"
)

type BaseController struct {
	beego.Controller

	layoutFile  string
	headerFile  string
	sidebarFile string
	footerFile  string

	mgoSession *mgo.Session
}

func (this *BaseController) Prepare() {
	fmt.Println("base controller prepare")

	this.headerFile = "include/header.html"
	this.footerFile = "include/footer.html"
	this.layoutFile = "include/layout/classic.html"
	this.sidebarFile = "include/sidebar/classic_sidebar.html"

	this.ConnMongoDB()
}

func (this *BaseController) MyRender(viewFile string) {
	this.Layout = this.layoutFile
	this.TplName = viewFile

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["headerFile"] = this.headerFile
	this.LayoutSections["footerFile"] = this.footerFile
	this.LayoutSections["sidebarFile"] = this.sidebarFile

	this.PrepareViewData()
	this.Render()
}

func (this *BaseController) MyRedirect(baseUrl string, code int) {
	realUrl := helpers.SiteUrl(baseUrl)
	this.Redirect(realUrl, code)
}

func (this *BaseController) PrepareViewData() {
	staticUrl := beego.AppConfig.String("static_url")
	siteUrl := beego.AppConfig.String("siteUrl")

	this.Data["staticUrl"] = staticUrl
	this.Data["siteUrl"] = siteUrl
}

func (this *BaseController) CheckLogin() bool {
	email := this.GetSession("email")
	if email != nil {
		return true
	} else {
		this.Redirect("user/login", 302)
		return false
	}
}

func (this *BaseController) GetSessionUser() interface{} {
	return this.GetSession("email")
}

func (this *BaseController) ConnMongoDB() {
	url := beego.AppConfig.String("mongoUrl")
	this.mgoSession = models.GetDbConn(url)
}

func (this *BaseController) CloseMongoDB() {
	this.mgoSession.Close()
}

func (this *BaseController) GetMgoDbs() []string {
	dbs, err := this.mgoSession.DatabaseNames()
	helpers.CheckError(err)
	return dbs
}

func (this *BaseController) GetColsByDb(dbName string) []string {
	cols, err := this.mgoSession.DB(dbName).CollectionNames()
	helpers.CheckError(err)
	return cols
}
