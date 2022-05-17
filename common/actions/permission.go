package actions

import (
	"errors"
	"github.com/geiqin/gotools/helper"
	"github.com/geiqin/micro-kit/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func newDataPermission(tx *gorm.DB, userId interface{}) (*auth.UserPermission, error) {
	var err error
	p := &auth.UserPermission{}

	err = tx.Table("sys_authority_users").
		Select("sys_authority_users.id", "sys_authority_users.role_id", "sys_authority_users.dept_id", "sys_authority_roles.data_scope").
		Joins("left join sys_authority_roles on sys_authority_roles.id = sys_authority_users.role_id").
		Where("sys_authority_users.id = ?", userId).
		Scan(p).Error
	if err != nil {
		err = errors.New("获取用户数据出错 msg:" + err.Error())
		return nil, err
	}
	return p, nil
}

func Permission(tableName string, p *auth.UserPermission) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		/*
			if !config.ApplicationConfig.EnableDP {
				return db
			}
		*/
		switch p.DataScope {
		case "2":
			return db.Where(tableName+".creator_id in (select sys_authority_users.id from sys_authority_roles_dept left join sys_authority_users on sys_authority_users.dept_id=sys_authority_roles_dept.dept_id where sys_authority_roles_dept.role_id = ?)", p.RoleId)
		case "3":
			return db.Where(tableName+".creator_id in (SELECT id from sys_authority_users where dept_id = ? )", p.DeptId)
		case "4":
			return db.Where(tableName+".creator_id in (SELECT id from sys_authority_users where sys_authority_users.dept_id in(select dept_id from sys_dept where dept_path like ? ))", "%/"+helper.ToString(p.DeptId)+"/%")
		case "5":
			return db.Where(tableName+".creator_id = ?", p.UserId)
		default:
			return db
		}
	}
}

func getPermissionFromContext(c *gin.Context) *auth.UserPermission {
	p := new(auth.UserPermission)
	if pm, ok := c.Get(PermissionKey); ok {
		switch pm.(type) {
		case *auth.UserPermission:
			p = pm.(*auth.UserPermission)
		}
	}
	return p
}

// GetPermissionFromContext 提供非action写法数据范围约束
func GetPermissionFromContext(c *gin.Context) *auth.UserPermission {
	return getPermissionFromContext(c)
}
