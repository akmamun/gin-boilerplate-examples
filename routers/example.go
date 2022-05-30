package routers

import (
	"github.com/akmamun/gin-boilerplate-examples/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ExamplesRoutes(route *gin.Engine, db *gorm.DB) {
	ctrl := controllers.Controller{DB: db}
	v1 := route.Group("/v1/examples")
	v1.POST("test/", ctrl.CreateExample)
	v1.GET("test/", ctrl.GetExampleData)
	v1.GET("test/relational", ctrl.GetHasManyRelationUserData)
	v1.GET("test/card", ctrl.GetHasManyRelationCreditCardData)
	v1.GET("test/user", ctrl.GetUserDetails)

}
