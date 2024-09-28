package broker

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/gorm"

	"telemetry-service/models"

	"go-micro.dev/v5/broker"
)

func SetupBroker(db *gorm.DB) {
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker connect error: %v", err)
	}

	_, err := broker.Subscribe("device.data", func(p broker.Event) error {
		var telemetry models.TelemetryData
		if err := json.Unmarshal(p.Message().Body, &telemetry); err != nil {
			return fmt.Errorf("error unmarshalling telemetry data: %v", err)
		}

		if err := db.Create(&telemetry).Error; err != nil {
			return fmt.Errorf("error saving telemetry data: %v", err)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error subscribing: %v", err)
	}
}
