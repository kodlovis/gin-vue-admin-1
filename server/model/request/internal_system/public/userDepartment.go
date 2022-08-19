package public

import "gin-vue-admin/model/internal_system/public"

type UserDepartmentSearch struct {
	public.UserDepartment
	PageInfo
}

type GetDepartmentByUserId struct {
	UID int64 `json:"uid" form:"uid"` // 主键ID
	PageInfo
}
