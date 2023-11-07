package routes

import (
	"github.com/RAFAT-DEVELOPER/underwatersensors/api/controllers"
	"github.com/gin-gonic/gin"
)

type SensorRouteController struct {
	sensorController controllers.SensorController
}

func NewSensorRouteController(sensorController controllers.SensorController) SensorRouteController {
	return SensorRouteController{sensorController}
}

func (sc *SensorRouteController) SensorRoute(rg *gin.RouterGroup) {

	router := rg.Group("/sensor")
	router.GET("/", sc.sensorController.GetSensors)

}
