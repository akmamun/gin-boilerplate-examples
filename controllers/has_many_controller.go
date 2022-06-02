package controllers

import (
	"net/http"
	"strconv"

	examples "github.com/akmamun/gin-boilerplate-examples/models"
	"github.com/akmamun/gin-boilerplate-examples/pkg/helpers/pagination"
	"github.com/gin-gonic/gin"
)

type CreditCardData struct {
	Number string `json:"number"`
}

//GetHasManyRelationUserData fetch user data with preload
func (base *Controller) GetHasManyRelationUserData(ctx *gin.Context) {
	var user []examples.User
	// ctx.JSON(http.StatusOK, &user)
	// db :=base.DB.Preload("CreditCards").Find(&user)
	paginate := pagination.Pagging(&pagination.Param{
		//DB:    base.DB,
		Page:  1,
		Limit: 3,
	}, &user)

	ctx.JSON(http.StatusOK, &paginate)

}

//GetHasManyRelationCreditCardData fetch credit-card data with preload
func (base *Controller) GetHasManyRelationCreditCardData(ctx *gin.Context) {
	var creditCards []examples.CreditCard
	base.DB.Find(&creditCards)
	ctx.JSON(http.StatusOK, &creditCards)

}

// GetUserDetails based on user_id
func (base *Controller) GetUserDetails(ctx *gin.Context) {
	var user []examples.User
	userId, _ := strconv.Atoi(ctx.DefaultQuery("user_id", ""))

	err := base.DB.Preload("CreditCards").First(&user, userId).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"user_id": "Enter valid user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
