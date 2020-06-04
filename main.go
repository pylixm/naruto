package main

import (
	"encoding/gob"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	"naruto/common"
	"naruto/models"
	_ "naruto/routers"
	"syscall"
	"time"
)

func init() {
	// 注册session的保存对象，防止应用重启无法解析
	// see doc: https://beego.me/docs/mvc/controller/session.md#%E7%89%B9%E5%88%AB%E6%B3%A8%E6%84%8F%E7%82%B9
	gob.Register(models.NarutoUser{})

	// LDAP 初始化
	InitLdap()

	// 初始化DB
	err := initOrm4Db()
	if err != nil {
		goto end
	}

end:
	if err != nil {
		beego.Error(err)
		syscall.Exit(1)
	}

}

func initOrm4Db() error {
	// 是否打印sql 语句，默认不打印
	orm.Debug = beego.AppConfig.DefaultBool("ORMDebug", false)

	// 注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册默认数据库，orm 默认使用了golang 自己的链接池
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",
		beego.AppConfig.String("mysql_user"),
		beego.AppConfig.String("mysql_pass"),
		beego.AppConfig.String("mysql_host"),
		beego.AppConfig.String("mysql_port"),
		beego.AppConfig.String("mysql_db"))
	maxIdle, _ := beego.AppConfig.Int("mysql_maxIdle") // 最大空闲链接
	maxConn, _ := beego.AppConfig.Int("mysql_maxConn") // 最大链接
	err := orm.RegisterDataBase("default", "mysql", conn, maxIdle, maxConn)
	if err != nil {
		return err
	}

	// 设置链接合理的生命周期时间
	mdb, err := orm.GetDB("default")
	if err != nil {
		return err
	}
	mdb.SetConnMaxLifetime(time.Second * 28)

	return nil
}

func InitLdap() {
	common.LdapHost = beego.AppConfig.DefaultString("ldap_host", "")
	common.LdapPort = beego.AppConfig.DefaultString("ldap_port", "")
	common.LdapUser = beego.AppConfig.DefaultString("ldap_user", "")
	common.LdapPass = beego.AppConfig.DefaultString("ldap_pass", "")
	common.LdapBind = beego.AppConfig.DefaultString("ldap_bind", "")
	common.LdapSearch = beego.AppConfig.DefaultString("ldap_search", "")
}

// 生成数据表
func Syncdb() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RunSyncdb("default", false, true)
}

func main() {
	// 生成数据表
	// Syncdb()

	// 日志配置
	defer func() {
		if err := recover(); err != nil {
			beego.Error("panic error:", err)
		}
	}()
	log_conf := beego.AppConfig.String("log_conf")
	if log_conf == "" {
		file := beego.AppConfig.String("appname")
		log_conf = `{"filename":"logs/` + file + `.log", "maxsize":50MB, "daily":true, "maxdays":7，"separate":["error"]}`
	}
	logs.SetLogger(logs.AdapterMultiFile, log_conf)

	// 启动服务
	beego.Run()

}
