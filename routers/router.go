package routers

import (
	"github.com/astaxie/beego"
	"html/template"
	"naruto/controllers"
	"naruto/filters"
	"net/http"
)

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

func init() {
	// 自定义错误页面
	beego.ErrorHandler("404", page_not_found)

	// 初始化 api router
	ns := beego.NewNamespace("/api/v1/",
		beego.NSRouter("/", &controllers.MainController{}, "get:ApiInfo"),

		// 用户名密码登录
		beego.NSRouter("login/", &controllers.UserController{}, "post:Login"),
		beego.NSRouter("logout/", &controllers.UserController{}, "post:Logout"),

		// beego.NSBefore(filters.AuthUserToken),
		beego.NSBefore(filters.AuthUser),
		beego.NSNamespace("/user",
			beego.NSRouter("/:id", &controllers.UserController{}, "get:UserInfoById"),
			beego.NSRouter("/:username", &controllers.UserController{}, "get:UserInfoByName"),
			beego.NSRouter("/list", &controllers.UserController{}, "get:List"),
		),
		// //此处正式版时改为验证加密请求
		// beego.NSCond(func(ctx *context.Context) bool {
		// 	if ua := ctx.Input.Request.UserAgent(); ua != "" {
		// 		return true
		// 	}
		// 	return false
		// }),
		// beego.NSNamespace("/ios",
		// 	//CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)
		// 	beego.NSNamespace("/create",
		// 		// /api/ios/create/node/
		// 		beego.NSRouter("/node", &apis.CreateNodeHandler{}),
		// 		// /api/ios/create/topic/
		// 		beego.NSRouter("/topic", &apis.CreateTopicHandler{}),
		// 	),
		// ),
	)

	beego.AddNamespace(ns)

}
