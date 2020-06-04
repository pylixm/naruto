package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"sync"
)

const (
	TABLE_NARUTO_USER = "naruto_user" // 数据库中表的名字
)

var (
	Once      sync.Once
	GlobalOrm orm.Ormer
)

// 定义查询条件，用于ORM查询语句的拼接
type CondFields struct {
	Exact       []string
	IExact      []string
	Contains    []string
	IContains   []string
	In          []string
	Gt          []string
	Gte         []string
	Lt          []string
	Lte         []string
	StartsWith  []string
	IStartsWith []string
	EndsWith    []string
	IEndsWith   []string
	IsNull      []string
}

func init() {
	orm.RegisterModel(
		new(NarutoUser),
	)
}

// 使用单例，防止协程死锁
func GetOrmer() orm.Ormer {
	Once.Do(func() {
		GlobalOrm = orm.NewOrm()
	})
	return GlobalOrm
}

func OrmCondition(qs orm.QuerySeter, fields CondFields, condArr map[string]string, cond *orm.Condition) orm.QuerySeter {
	if len(fields.Exact) > 0 {
		flag := false
		for _, v := range fields.Exact {
			if condArr[v] != "" {
				cond = cond.And(v, condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.IExact) > 0 {
		flag := false
		for _, v := range fields.IExact {
			if condArr[v] != "" {
				cond = cond.And(v+"__iexact", condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.Contains) > 0 {
		flag := false
		for _, v := range fields.Contains {
			if condArr[v] != "" {
				cond = cond.And(v+"__contains", condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.IContains) > 0 {
		flag := false
		for _, v := range fields.IContains {
			if condArr[v] != "" {
				cond = cond.And(v+"__icontains", condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.In) > 0 {
		flag := false
		for _, v := range fields.In {
			if condArr[v] != "" {
				cond = cond.And(v+"__in", strings.Split(condArr[v], ","))
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.Gt) > 0 {
		flag := false
		for _, v := range fields.Gt {
			if condArr[v] != "" {
				cond = cond.And(v+"__gt", condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.Gte) > 0 {
		flag := false
		for _, v := range fields.Gte {
			if condArr[v] != "" {
				cond = cond.And(v+"__gte", condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.Lt) > 0 {
		flag := false
		for _, v := range fields.Lt {
			if condArr[v] != "" {
				cond = cond.And(v+"__lt", condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.Lte) > 0 {
		flag := false
		for _, v := range fields.Lte {
			if condArr[v] != "" {
				cond = cond.And(v+"__lte", condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.StartsWith) > 0 {
		flag := false
		for _, v := range fields.StartsWith {
			if condArr[v] != "" {
				cond = cond.And(v+"__startswith", condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.IStartsWith) > 0 {
		flag := false
		for _, v := range fields.IStartsWith {
			if condArr[v] != "" {
				cond = cond.And(v+"__istartswith", condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.EndsWith) > 0 {
		flag := false
		for _, v := range fields.EndsWith {
			if condArr[v] != "" {
				cond = cond.And(v+"__endswith", condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.IEndsWith) > 0 {
		flag := false
		for _, v := range fields.IEndsWith {
			if condArr[v] != "" {
				cond = cond.And(v+"__iendswith", condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	if len(fields.IsNull) > 0 {
		flag := false
		for _, v := range fields.IsNull {
			if condArr[v] != "" {
				cond = cond.And(v+"__IsNull", condArr[v])
				flag = true
			}
		}
		if flag {
			qs = qs.SetCond(cond)
		}
	}

	return qs
}
