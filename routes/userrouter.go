package routes

import (
	doctors "hospital-mgt-sys/doctors"
	"hospital-mgt-sys/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users/:user_id", doctors.GetDoctor())

}
