package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"] = append(beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"] = append(beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"] = append(beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetById",
			Router: `/id/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"] = append(beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAllExcludeOneId",
			Router: `/exclude/id/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"] = append(beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetByName",
			Router: `/name/:name`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"] = append(beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetByRole",
			Router: `/role/:role`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"] = append(beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"] = append(beego.GlobalControllerRouter["dao-service/user-dao-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "DeleteById",
			Router: `/id/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
