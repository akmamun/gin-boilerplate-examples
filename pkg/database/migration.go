package database

import (
	examples "github.com/akmamun/gin-boilerplate-examples/examples/ex_models"
	"github.com/akmamun/gin-boilerplate-examples/models"
)

//Add list of model add for migrations
//var migrationModels = []interface{}{&ex_models.Example{}, &model.Example{}, &model.Address{})}
var migrationModels = []interface{}{&models.Example{}, &examples.User{}, &examples.CreditCard{}}
