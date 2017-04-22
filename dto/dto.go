package dto

import (
	"encoding/gob"
	"time"
)

//SensorMessage struct is the dto of sensorMessages
type SensorMessage struct {
	Name      string
	Value     float64
	Timestamp time.Time
}

func init() {
	gob.Register(SensorMessage{})
}
