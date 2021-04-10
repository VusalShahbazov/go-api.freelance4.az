package v1

import (
	"github.com/VusalShahbazov/api.freelance4.az/internal/helper"
	"github.com/VusalShahbazov/api.freelance4.az/internal/response"
	"github.com/gin-gonic/gin"
)

func  ProjectsIndex(ctx *gin.Context)  {
	user, err := helper.GetUser(ctx)
	if err != nil {
		response.UnexpectedError(ctx , nil)
		return
	}

	response.SuccessResponse(ctx , user)
}

//func Index(ctx *gin.Context)  {
//
//	var projectFilter requests.ProjectFilter
//	if err :=  ctx.ShouldBindQuery(&projectFilter) ; err != nil{
//		response.GenerateValidationResponse(ctx , err)
//		return
//	}
//	var data []models.Project
//
//	filterMap :=  make(map[string]interface{})
//	fmt.Println(projectFilter)
//	if projectFilter.IsActive != nil {
//		filterMap["is_active"] = *projectFilter.IsActive
//	}
//	fmt.Println(filterMap)
//
//	err := models.DB.Debug().Where(filterMap).Preload("Creator" , func(db *gorm.DB) *gorm.DB {
//		return db.Select("id ", "last_name"  , "first_name")
//	}).Limit(projectFilter.Limit()).Offset(projectFilter.Offset()).Find(&data).Error
//	if err != nil {
//		response.InternalServerError(ctx , err.Error())
//		return
//	}
//
//	result := struct {
//		Total int64
//	}{Total: 0}
//
//	models.DB.Table("projects").Where(filterMap).Select("count(*) as total").Find(&result)
//
//	response.PaginationResponse(ctx  , resoursec.NewProjectsResponse(data), result.Total , projectFilter.Limit())
//}


//func Store(ctx *gin.Context)  {
//
//}