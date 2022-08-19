package public

import "gin-vue-admin/model/internal_system/public"

type DepartmentSearch struct {
	public.Department
	PageInfo
}
