package common

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/ldap.v2"
)

type UserLdap struct {
	Account string `json:"account"`
	Cn      string `json:"cn"`
	Mail    string `json:"mail"`
	Mobile  string `json:"mobile"`
	Phone   string `json:"phone"`
}

var (
	LdapHost   string
	LdapPort   string
	LdapUser   string
	LdapPass   string
	LdapBind   string
	LdapSearch string
)

func (o UserLdap) String() string {
	s, _ := json.Marshal(o)
	return string(s)
}

func AuthLdap(user, pass string) (*UserLdap, error) {
	// 连接LDAP
	ds, err := ldap.Dial("tcp", fmt.Sprintf("%s:%s", LdapHost, LdapPort))
	if err != nil {
		beego.Error("connect to ldap server err: ", err)
		return nil, err
	}
	defer ds.Close()

	// 使用TLS重连
	err = ds.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		beego.Error("reconnect to ldap with TLS err: ", err)
		return nil, err
	}

	// BIND admin account
	err = ds.Bind(fmt.Sprintf("CN=%s,%s", LdapUser, LdapBind), LdapPass)
	if err != nil {
		beego.Error("bind to ldap server err: ", err)
		return nil, err
	}

	// 查找用户
	searchRequest := ldap.NewSearchRequest(
		LdapSearch,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(sAMAccountName=%s))", user),
		[]string{"dn", "cn", "sn", "name", "sAMAccountName", "mail", "mobile", "telephoneNumber"},
		nil)
	sr, err := ds.Search(searchRequest)
	if err != nil {
		beego.Error("ldap search err: ", err)
		return nil, err
	}
	// 未找到用户
	if len(sr.Entries) == 0 {
		err = errors.New("username or password wrong")
		beego.Error("ldap search err :", err)
		return nil, err
	}
	// 找到多个用户
	if len(sr.Entries) > 1 {
		err = errors.New("more than one user returned")
		beego.Error("ldap search err :", err)
		return nil, err
	}

	err = ds.Bind(sr.Entries[0].DN, pass)
	if err != nil {
		beego.Error("check username and password err: ", err)
		return nil, err
	}

	uInfo := UserLdap{
		Account: user,
		Cn:      sr.Entries[0].GetAttributeValue("cn"),
		Mail:    sr.Entries[0].GetAttributeValue("mail"),
		Mobile:  sr.Entries[0].GetAttributeValue("mobile"),
		Phone:   sr.Entries[0].GetAttributeValue("telephoneNumber"),
	}
	return &uInfo, nil
}
