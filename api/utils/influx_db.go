package utils

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/rs/zerolog"
	"time"
)

type InfluxDBClient struct {
	Client influxdb2.Client
}

var influxDBClient InfluxDBClient

func NewInfluxDBClient() InfluxDBClient {
	if influxDBClient.Client != nil {
		return influxDBClient
	}
	url := EnvInfluxDbAddress()
	token := EnvInfluxDbToken()
	influxDBClient.Client = influxdb2.NewClient(url, token)
	return influxDBClient
}

type InfluxDBHook struct {
	Client InfluxDBClient
}

func (h InfluxDBHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	point := influxdb2.
		NewPointWithMeasurement("logs").
		SetTime(time.Now()).
		AddTag("level", level.String()).
		AddField("message", message)

	org := EnvInfluxDbOrg()
	bucket := EnvInfluxDbBucket()
	writeAPI := h.Client.Client.WriteAPIBlocking(org, bucket)

	err := writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		logData := LogData{Event: "Can't upload to influxdb", Details: err.Error(), User: "system"}
		LogError(logData)
		return
	}
}
