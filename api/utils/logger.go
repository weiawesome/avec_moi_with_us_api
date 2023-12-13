package utils

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var logger zerolog.Logger

func InitLogger() {
	logFile := &lumberjack.Logger{
		Filename:   "logs/server.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}
	localLogger := zerolog.New(logFile).With().Timestamp().Logger()
	influxdbLogger := zerolog.Logger{}
	logger = localLogger
	isUsingInfluxDB := false

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				influxClient := NewInfluxDBClient()
				if result, err := influxDBClient.Client.Ping(ctx); result == true {
					if !isUsingInfluxDB {
						logger = influxdbLogger
						logger = logger.Hook(InfluxDBHook{Client: influxClient})
						isUsingInfluxDB = true
						uploadRotatedLogsToInfluxDB(influxClient.Client)
					}
				} else {
					if isUsingInfluxDB {
						logger = localLogger
						isUsingInfluxDB = false
					}
					logData := LogData{Event: "InfluxDB connection failed, using local logs", User: "system", Details: err.Error()}
					LogError(logData)
				}
				time.Sleep(time.Minute)
			}
		}
	}()

	time.Sleep(time.Second)
}

func uploadRotatedLogsToInfluxDB(influxClient influxdb2.Client) {
	org := EnvInfluxDbOrg()
	bucket := EnvInfluxDbBucket()

	writeAPI := influxClient.WriteAPIBlocking(org, bucket)

	files, err := filepath.Glob("logs/server.log*")
	if err != nil {
		data := LogData{User: "system", Event: "Fail to find log file", Details: "error: " + err.Error()}
		LogError(data)
	}

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			data := LogData{User: "system", Event: "Fail to open log file", Details: "File name: " + filename + " error: " + err.Error()}
			LogError(data)
			continue
		}
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			level := getField(line, "level")
			logTime := getField(line, "time")
			logMessage := getField(line, "message")
			point := influxdb2.NewPointWithMeasurement("logs").
				AddTag("level", level).
				AddField("message", logMessage).
				SetTime(parseTime(logTime))

			err := writeAPI.WritePoint(context.Background(), point)
			if err != nil {
				data := LogData{User: "system", Event: "Fail to upload log file", Details: "File name: " + filename + " error: " + err.Error()}
				LogError(data)
			}
		}
		err = file.Close()
		if err != nil {
			data := LogData{User: "system", Event: "Fail to close log file", Details: "File name: " + filename + " error: " + err.Error()}
			LogError(data)
			continue
		}
		err = os.Remove(filename)
		if err != nil {
			data := LogData{User: "system", Event: "Fail to delete log file", Details: "File name: " + filename + " error: " + err.Error()}
			LogError(data)
		}
	}
}

func getField(line, field string) string {
	start := strings.Index(line, `"`+field+`":"`) + len(field) + 4
	if field == "message" {
		return line[start : len(line)-2]
	}
	end := strings.Index(line[start:], `","`)
	return line[start : start+end]
}

func parseTime(logTime string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, logTime)
	if err != nil {
		return time.Now()
	}
	return parsedTime
}

type LogData struct {
	Event   string      `json:"event"`
	Method  string      `json:"method"`
	Path    string      `json:"path"`
	Header  interface{} `json:"header"`
	User    string      `json:"user"`
	Details string      `json:"details"`
}

func LogDebug(obj LogData) {
	jsonData, err := json.Marshal(obj)
	if err == nil {
		logger.Debug().Msg(string(jsonData))
	} else {
		logger.Error().Msg("\"event\":\"error in debug log\",\"details\":\"" + err.Error() + "\",\"user\":\"system\"")
	}
}

func LogInfo(obj LogData) {
	jsonData, err := json.Marshal(obj)
	if err == nil {
		logger.Info().Msg(string(jsonData))
	} else {
		logger.Error().Msg("\"event\":\"error in info log\",\"details\":\"" + err.Error() + "\",\"user\":\"system\"")
	}
}

func LogWarn(obj LogData) {

	jsonData, err := json.Marshal(obj)
	if err == nil {
		logger.Warn().Msg(string(jsonData))
	} else {
		logger.Error().Msg("\"event\":\"error in warn log\",\"details\":\"" + err.Error() + "\",\"user\":\"system\"")
	}
}

func LogError(obj LogData) {
	jsonData, err := json.Marshal(obj)
	if err == nil {
		logger.Error().Msg(string(jsonData))
	} else {
		logger.Error().Msg("\"event\":\"error in error log\",\"details\":\"" + err.Error() + "\",\"user\":\"system\"")
	}
}

func LogFatal(obj LogData) {
	jsonData, err := json.Marshal(obj)
	if err == nil {
		logger.Fatal().Msg(string(jsonData))
	} else {
		logger.Error().Msg("\"event\":\"error in fatal log\",\"details\":\"" + err.Error() + "\",\"user\":\"system\"")
	}
}

func LogPanic(obj LogData) {
	jsonData, err := json.Marshal(obj)
	if err == nil {
		logger.Panic().Msg(string(jsonData))
	} else {
		logger.Error().Msg("\"event\":\"error in panic log\",\"details\":\"" + err.Error() + "\",\"user\":\"system\"")
	}
}
