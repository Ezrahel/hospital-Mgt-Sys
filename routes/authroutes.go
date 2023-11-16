package routes

import (
	"github.com/gin-gonic/gin"
	"hospital-mgt-sys/doctors"
)

func AuthRoutes(incoming Routes *gin.Engine){
	incomingRoutes.POST("user/signup", Doctors.signup())
	incomingRoutes.POST("user/signup", Doctors.signin())
}