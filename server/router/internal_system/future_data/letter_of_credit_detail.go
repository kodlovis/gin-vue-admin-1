package future_data

import (
	v1 "gin-vue-admin/api/v1/internal_system/future_data"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitLetterOfCreditDetailRouter(Router *gin.RouterGroup) {
	LetterOfCreditDetailRouter := Router.Group("letterOfCreditDetail").Use(middleware.OperationRecord())
	{
		LetterOfCreditDetailRouter.POST("createLetterOfCreditDetail", v1.CreateLetterOfCreditDetail)             // 新建LetterOfCreditDetail
		LetterOfCreditDetailRouter.DELETE("deleteLetterOfCreditDetail", v1.DeleteLetterOfCreditDetail)           // 删除LetterOfCreditDetail
		LetterOfCreditDetailRouter.DELETE("deleteLetterOfCreditDetailByIds", v1.DeleteLetterOfCreditDetailByIds) // 批量删除LetterOfCreditDetail
		LetterOfCreditDetailRouter.PUT("updateLetterOfCreditDetail", v1.UpdateLetterOfCreditDetail)              // 更新LetterOfCreditDetail
		LetterOfCreditDetailRouter.GET("findLetterOfCreditDetail", v1.FindLetterOfCreditDetail)                  // 根据ID获取LetterOfCreditDetail
		LetterOfCreditDetailRouter.GET("getLetterOfCreditDetailList", v1.GetLetterOfCreditDetailList)            // 获取LetterOfCreditDetail列表
		LetterOfCreditDetailRouter.GET("getLetterOfCreditDetailListWithNoPurchaseRate", v1.GetLetterOfCreditDetailListWithNoPurchaseRate)
		LetterOfCreditDetailRouter.PUT("updateLetterOfCreditPurchaseRate", v1.UpdateLetterOfCreditPurchaseRate)

	}
}
