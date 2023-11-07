package main

import (
	"fmt"
	"log"

	"github.com/RAFAT-DEVELOPER/underwatersensors/api/models"
	"github.com/RAFAT-DEVELOPER/underwatersensors/initializers"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectPostgres(&config)
}

func main() {
	err := initializers.DB.AutoMigrate(&models.SensorGroup{}, &models.Sensor{}, &models.FishSpecies{}, &models.SensorObservation{})
	if err != nil {
		log.Fatal("Could not migrate the Database", err)
		return
	}
	fmt.Println("Migration complete")
}
