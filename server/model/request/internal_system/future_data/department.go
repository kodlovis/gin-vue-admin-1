package future_data

import "gin-vue-admin/model/internal_system/future_data"

type DepartmentSearch struct {
	future_data.Department
	PageInfo
}
