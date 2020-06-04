package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"naruto/models"
	"strconv"
)

type MainController struct {
	beego.Controller
	ApiResponse  ApiResponse
	Status       int
	CurrUser     models.NarutoUser
	CurrApp      int64
	LogFlag      bool
	RemoteServer string
}

type ApiResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Client string `json:"client"`
	// Server    string      `json:"server"`
	Count     int64       `json:"count"`
	PageSize  int64       `json:"page_size"`
	PageCount int64       `json:"page_count"`
	Page      int64       `json:"page"`
	Content   interface{} `json:"content"`
}

const (
	// http_code
	HTTP_CODE_SERVICE_ERROR   = 500
	HTTP_CODE_BAD_REQUEST     = 400
	HTTP_CODE_SERVICE_SUCCESS = 200

	// 服务状态
	STATUS_SERVICE_SUCCESS = 0
	STATUS_SERVICE_ERR     = 10001

	// 服务状态 鉴权错误
	STATUS_USER_NOT_LOGIN = 20000
	STATUS_USER_NOT_ADMIN = 20001

	// 服务状态 Client非法请求
	STATUS_INPUT_ERR = 30001
)

var (
	ServiceErrResp = ApiResponse{
		Code: STATUS_SERVICE_ERR,
		Msg:  "Server encountered an error!",
	}
	InputErrResp = ApiResponse{
		Code: STATUS_INPUT_ERR,
		Msg:  "Input error!",
	}
	NotLoginResp = ApiResponse{
		Code:    STATUS_USER_NOT_LOGIN,
		Msg:     "Offline",
		Content: false,
	}
	LoginError = ApiResponse{
		Code: STATUS_USER_NOT_LOGIN,
		Msg:  "用户名或命名错误",
	}
	LogoutResp = ApiResponse{
		Code:    STATUS_SERVICE_SUCCESS,
		Msg:     "success",
		Content: false,
	}
	NotAdmin = ApiResponse{
		Code:    STATUS_USER_NOT_ADMIN,
		Msg:     "Not Admin",
		Content: false,
	}
)

// 返回函数 ----------------------------------------------------------------------------------------
// 登录错误
func (c *MainController) LoginError() {
	c.ApiResponse = LoginError
	c.RespJsonWithStatus()
}

// 无权限
func (c *MainController) NoPermission() {
	c.TplName = "403.html"
}

// 不是管理员
func (c *MainController) RespNotAdmin() {
	c.ApiResponse = NotAdmin
	c.RespJsonWithStatus()
}

// 服务错误
func (c *MainController) RespServiceError(err error) {
	r := ServiceErrResp
	r.Msg = err.Error()
	c.ApiResponse = r
	c.Status = HTTP_CODE_SERVICE_ERROR
	c.RespJsonWithStatus()
}

// 请求错误
func (c *MainController) RespInputError(err error) {
	c.ApiResponse = InputErrResp
	if err != nil {
		c.ApiResponse.Msg = err.Error()
	}
	c.Status = HTTP_CODE_BAD_REQUEST
	c.RespJsonWithStatus()
}

func (c *MainController) RespJsonWithStatus() {
	c.ApiResponse.Client = c.Ctx.Request.RemoteAddr
	// c.ApiResponse.Server = common.LocalHost
	c.Data["json"] = c.ApiResponse
	// 跨域支持
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")

	c.Ctx.Output.SetStatus(c.Status)
	c.Ctx.Request.Header.Add("Status", strconv.Itoa(c.Status))
	c.ServeJSON()
}

// 返回函数 end ----------------------------------------------------------------------------------------

func (c *MainController) ApiInfo() {
	c.ApiResponse.Msg = "success"
	c.ApiResponse.Content = "API V1.0"
	c.RespJsonWithStatus()
}

// 检索条件参数处理 to map
func (c *MainController) FormatConditionToMap(fields models.CondFields) map[string]string {
	r := map[string]string{}
	if len(fields.Exact) > 0 {
		for _, v := range fields.Exact {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.IExact) > 0 {
		for _, v := range fields.IExact {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.Contains) > 0 {
		for _, v := range fields.Contains {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.IContains) > 0 {
		for _, v := range fields.IContains {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.In) > 0 {
		for _, v := range fields.In {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.Gt) > 0 {
		for _, v := range fields.Gt {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.Gte) > 0 {
		for _, v := range fields.Gte {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.Lt) > 0 {
		for _, v := range fields.Lt {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.Lte) > 0 {
		for _, v := range fields.Lte {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.StartsWith) > 0 {
		for _, v := range fields.StartsWith {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.IStartsWith) > 0 {
		for _, v := range fields.IStartsWith {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.EndsWith) > 0 {
		for _, v := range fields.EndsWith {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.IEndsWith) > 0 {
		for _, v := range fields.IEndsWith {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}

	if len(fields.IsNull) > 0 {
		for _, v := range fields.IsNull {
			if t := c.GetString(v); t != "" {
				r[v] = t
			}
		}
	}
	logs.Debug("查询条件：", r)
	return r
}
