package controllers

import (
	"dao-service/user-dao-service/models"
	"dao-service/user-dao-service/service"
	"encoding/json"
	"fmt"
	"model"

	"github.com/astaxie/beego"
)

// Operations for Users
type UserController struct {
	beego.Controller
}

// @Title Create
// @Description create user
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.Response
// @Failure 403 body is empty
// @router / [post]
func (this *UserController) Create() {
	var err error
	var user model.User
	var response models.Response

	// unmarshal
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err == nil {
		var svc service.Service
		var result []byte
		err = svc.Create(&user)
		if err == nil {
			result, err = json.Marshal(&user)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug("Unmarshal data failed")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}

	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.Response
// @router / [get]
func (this *UserController) GetAll() {
	var err error
	var response models.Response

	var svc service.Service
	var users []*model.User
	var result []byte
	users, err = svc.GetAll()
	if err == nil {
		result, err = json.Marshal(users)
		if err == nil {
			response.Status = model.MSG_RESULTCODE_SUCCESS
			response.Reason = "success"
			response.Result = string(result)
		}
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}

	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title GetById
// @Description get user by id
// @Param	id		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @Failure 403 :id is invalid
// @router /id/:id [get]
func (this *UserController) GetById() {
	var err error
	var response models.Response

	var id int64
	id, err = this.GetInt64(":id")
	beego.Debug("GetById", id)
	if id > 0 && err == nil {
		var svc service.Service
		var user *model.User
		var result []byte
		user, err = svc.GetById(id)
		if err == nil {
			result, err = json.Marshal(user)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "user id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title GetAllExcludeOneId
// @Description get all users exclude one by id
// @Param	id		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @Failure 403 :id is invalid
// @router /exclude/id/:id [get]
func (this *UserController) GetAllExcludeOneId() {
	var err error
	var response models.Response

	var id int64
	id, err = this.GetInt64(":id")
	beego.Debug("GetExcludeOne", id)
	if id > 0 && err == nil {
		var svc service.Service
		var users []*model.User
		var result []byte
		users, err = svc.GetAllExcludeOneId(id)
		if err == nil {
			result, err = json.Marshal(users)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "user id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title GetByName
// @Description get user by name
// @Param	name		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @Failure 403 :name is empty
// @router /name/:name [get]
func (this *UserController) GetByName() {
	var err error
	var response models.Response

	var name string
	name = this.GetString(":name")
	//beego.Debug("GetByName", name)
	if name != "" {
		var svc service.Service
		var result []byte
		var user *model.User
		user, err = svc.GetByName(name)
		if err == nil {
			result, err = json.Marshal(user)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		err = fmt.Errorf("%s", "user id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title GetByRole
// @Description get user by role
// @Param	role		path 	int	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @Failure 403 :role is invalid
// @router /role/:role [get]
func (this *UserController) GetByRole() {
	var err error
	var response models.Response

	role, err := this.GetInt(":role")
	beego.Debug("GetByRole", role)
	if role > 0 && err == nil {
		var svc service.Service
		var user *model.User
		var result []byte
		user, err = svc.GetByRole(role)
		if err == nil {
			result, err = json.Marshal(user)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "user role is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.Response
// @Failure 403 :id is invalid
// @router / [put]
func (this *UserController) Update() {
	var err error
	var user model.User
	var response models.Response

	// unmarshal
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err == nil {
		var svc service.Service
		var result []byte
		var newUser *model.User
		newUser, err = svc.Update(&user)
		if err == nil {
			result, err = json.Marshal(newUser)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug("Unmarshal data failed")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}

	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title DeleteById
// @Description delete the user by id
// @Param	id		path 	int64	true		"The int you want to delete"
// @Success 200 {object} models.Response
// @Failure 403 id is invalid
// @router /id/:id [delete]
func (this *UserController) DeleteById() {
	var err error
	var response models.Response

	var id int64
	id, err = this.GetInt64(":id")
	beego.Debug("DeleteById", id)
	if id > 0 && err == nil {
		var svc service.Service
		err = svc.DeleteById(id)
		if err == nil {
			response.Status = model.MSG_RESULTCODE_SUCCESS
			response.Reason = "success"
			response.Result = ""
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "user id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}
