package routes

import (
	"github.com/RAFAT-DEVELOPER/underwatersensors/api/controllers"
	"github.com/gin-gonic/gin"
)

type SensorObservationRouteController struct {
	sensorObservationController controllers.SensorObservationController
}

func NewSensorObservationRouteController(sensorObservationController controllers.SensorObservationController) SensorObservationRouteController {
	return SensorObservationRouteController{sensorObservationController}
}

func (soc *SensorObservationRouteController) SensorObservationRoute(rg *gin.RouterGroup) {

	router := rg.Group("/observation")
	router.POST("/", soc.sensorObservationController.CreateSensorObservation)
	router.POST("/bulk", soc.sensorObservationController.CreateBulkSensorObservationHandler)

}
