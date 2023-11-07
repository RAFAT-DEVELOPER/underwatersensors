package main

import (
	"fmt"
	"os"

	"github.com/RAFAT-DEVELOPER/underwatersensors/generatedata"
)

func main() {
	// Generating the sensor groups and sensors
	sensorGroups := generatedata.GenerateSensors()
	fmt.Println(sensorGroups)
	// Generating the SQL statements
	sqlDataGenerationStatements := generatedata.GenerateSQLStatements(sensorGroups)
	err := os.WriteFile("assets/initialSQL.txt", []byte(sqlDataGenerationStatements), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("String successfully written to file.")
}
