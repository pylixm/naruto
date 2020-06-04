package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"naruto/common"
	"naruto/models"
	"strconv"
)

type UserController struct {
	MainController
}

type LoginToken struct {
	// User  models.NaturoUser `json:"user"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginUser struct {
	Username string `json:"username"` // json 中的别名 username
	Password string `json:"password"`
}

// 用户登录获取token
func (c *UserController) Login() {
	resp := ApiResponse{}
	var userInfo LoginUser
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		beego.Error("parse request failed: ", err)
		c.RespInputError(err)
		return
	}
	// 检查传入
	if userInfo.Username == "" || userInfo.Password == "" {
		err = errors.New("username or password is empty")
		c.RespInputError(err)
		return
	}
	// 连接LDAP验证
	lInfo, err := common.AuthLdap(userInfo.Username, userInfo.Password)
	if err != nil {
		resp.Code = STATUS_INPUT_ERR
		resp.Msg = "username or password wrong"
		resp.Content = userInfo.Username
		beego.Error("auth by ldap err: ", err)
		c.ApiResponse = resp
		c.RespJsonWithStatus()
		return
	}
	resp.Msg = "success"

	// 检查用户是否已存在
	uInfo, _ := models.GetUserByName(userInfo.Username)
	fmt.Println(uInfo)
	fmt.Println(lInfo)
	if uInfo == nil {
		// 用户不存在, 添加用户
		dInfo := models.NarutoUser{}
		dInfo.Name = lInfo.Account
		dInfo.Password = common.GetMd5(userInfo.Password)
		dInfo.Type = "ldap"
		dInfo.Cname = lInfo.Cn
		dInfo.Email = lInfo.Mail
		dInfo.Phone = lInfo.Mobile
		dInfo.State = 1 // 默认正常
		id, err := models.CreateUser(&dInfo)
		if err != nil {
			beego.Error("Create User err: ", err)
			resp.Msg = "login success, but create user err: " + err.Error()
		}
		dInfo.Id = id
		uInfo = &dInfo
	}
	uInfo.Password = "******"
	// 判断用户状态是否正常 0 禁用 1 正常
	if uInfo.State == 0 {
		resp.Code = 1
		resp.Msg = "permission denied"
	} else {
		resp.Msg = "success"
		// 设置session
		c.SetSession("UserInfo", uInfo.ToString())
	}
	c.CurrUser.Cname = uInfo.Cname
	resp.Content = uInfo
	c.ApiResponse = resp
	c.RespJsonWithStatus()
}

// 退出
func (c *UserController) Logout() {
	c.ApiResponse = LogoutResp
	v := c.GetSession("UserInfo")
	if v == nil {
		c.ApiResponse.Msg = "Not Online"
	} else {
		c.DelSession("UserInfo")
	}
	c.RespJsonWithStatus()
}

func (c *UserController) List() {
	resp := ApiResponse{}
	page, _ := c.GetInt64("page")
	if page < 1 {
		page = 1
	}
	page_size, _ := c.GetInt64("page_size")
	if page_size < 1 {
		page_size = common.DefaultPageSize
	}
	condition := c.FormatConditionToMap(models.UserCondField)
	d, count, err := models.ListUser(condition, page, page_size)
	if err != nil {
		beego.Error("get User list err: ", err)
		c.RespServiceError(err)
		return
	}
	resp.Count = count
	resp.Content = d
	resp.PageCount = count / page_size
	if count%page_size > 0 {
		resp.PageCount++
	}
	resp.PageSize = page_size
	resp.Page = page
	c.ApiResponse = resp
	c.Status = HTTP_CODE_SERVICE_SUCCESS
	c.RespJsonWithStatus()
}

func (c *UserController) UserInfoById() {
	resp := ApiResponse{}
	id := c.Ctx.Input.Param(":id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error("parse err:", err)
		c.RespInputError(err)
		return
	}
	d, err := models.GetUserById(userId)
	if err != nil {
		beego.Error("get User info err: ", err)
		c.RespServiceError(err)
		return
	}
	resp.Content = d
	c.ApiResponse = resp
	c.RespJsonWithStatus()
}

func (c *UserController) UserInfoByName() {
	resp := ApiResponse{}
	username := c.Ctx.Input.Param(":username")
	if username != "" {
		beego.Error("parse err: username is null ")
		c.RespInputError(errors.New("username is not null"))
		return
	}
	d, err := models.GetUserByName(username)
	if err != nil {
		beego.Error("get User info err: ", err)
		c.RespServiceError(err)
		return
	}
	resp.Content = d
	c.ApiResponse = resp
	c.RespJsonWithStatus()
}
