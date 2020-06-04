package models

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/orm"
	"naruto/common"
	"time"
)

type NarutoUser struct {
	Id         int64     `json:"id" orm:"pk;auto"`
	Name       string    `json:"name"`
	Cname      string    `json:"cname"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Type       string    `json:"type"`  // ldap, local
	State      int64     `json:"state"` // 0-停用，1-正常
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(timestamp)"`
	UpdateTime time.Time `json:"update_time" orm:"auto_now;type(timestamp)"`
	IsAdmin    int64     `json:"is_admin"` // 0-不是，1-是
	Staffid    string    `json:"staffid"`  // 工号
	Weibo      string    `json:"weibo"`    // 微博ID
	Weixin     string    `json:"weixin"`   // 微信ID
}

var (
	// 可查询字段
	UserCondField = CondFields{
		IContains: []string{"name", "cname", "email", "phone", "staffid", "weibo", "weixin"},
		Exact:     []string{"state", "type"},
	}
)

func (o *NarutoUser) ToString() string {
	s, _ := json.Marshal(o)
	return string(s)
}

// 检测用户名密码
func CheckUserAuth(cname string, password string) (NarutoUser, bool) {
	o := GetOrmer()
	user := NarutoUser{
		Cname:    cname,
		Password: common.GetMd5(password),
	}
	err := o.Read(&user, "Cname", "Password")
	if err != nil {
		return user, false
	}
	return user, true
}

func ListUser(condition map[string]string, page int64, offset int64) ([]NarutoUser, int64, error) {
	o := GetOrmer()
	cond := orm.NewCondition()
	qs := o.QueryTable(TABLE_NARUTO_USER)
	qs = OrmCondition(qs, UserCondField, condition, cond)
	start := (page - 1) * offset
	var d []NarutoUser
	_, err := qs.OrderBy("id").Limit(offset, start).All(&d)
	total, _ := qs.Count()
	if err != nil {
		return nil, total, err
	}
	return d, total, nil
}

func GetUserByName(name string) (*NarutoUser, error) {
	o := GetOrmer()
	var user NarutoUser
	qs := o.QueryTable(TABLE_NARUTO_USER).Filter("name", name)
	if !qs.Exist() {
		return nil, errors.New("没有查询到该用户")
	}
	err := qs.One(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func GetUserById(id int64) (*NarutoUser, error) {
	o := GetOrmer()
	var user NarutoUser
	qs := o.QueryTable(TABLE_NARUTO_USER).Filter("id", id)
	if !qs.Exist() {
		return nil, errors.New("没有查询到该用户")
	}
	err := qs.One(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func CreateUser(d *NarutoUser) (int64, error) {
	o := GetOrmer()
	return o.Insert(d)
}
