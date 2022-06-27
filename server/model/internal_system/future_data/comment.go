package future_data

import (
	"gin-vue-admin/global"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                                                                  // 用户UUID
	Username  string    `json:"userName" gorm:"comment:用户登录名"`                                                               // 用户登录名
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`                                                                    // 用户登录密码
	NickName  string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                   // 用户昵称
	HeaderImg string    `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/1576554439myAvatar.png;comment:用户头像"` // 用户头像

	AuthorityId string `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
}

func (SysUser) TableName() string {
	return "sys_users"
}

type UserDepartment struct {
	global.GVA_MODEL
	Userid         string `json:"userid" form:"userid" gorm:"column:user_id;comment:;type:character varying;"`
	DepartmentCode string `json:"departmentCode" form:"departmentCode" gorm:"column:department_code;comment:;type:character varying;"`
}

func (UserDepartment) TableName() string {
	return "master_data.us010_user_department"
}
