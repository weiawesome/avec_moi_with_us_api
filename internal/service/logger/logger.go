package logger

import (
	"avec_moi_with_us_api/api/utils"
	"avec_moi_with_us_api/internal/repository/redis"
)

type ServiceLogger struct {
	Repository redis.Repository
}

func (s *ServiceLogger) Debug(data utils.LogData) {
	if hashValue, err := s.Repository.GetHashValue(data.User); err != nil {
		utils.LogError(utils.LogData{Event: "Convert to hash User error", Details: err.Error(), User: "system"})
	} else {
		data.User = hashValue
		utils.LogDebug(data)
	}

}
func (s *ServiceLogger) Info(data utils.LogData) {
	if hashValue, err := s.Repository.GetHashValue(data.User); err != nil {
		utils.LogError(utils.LogData{Event: "Convert to hash User error", Details: err.Error(), User: "system"})
	} else {
		data.User = hashValue
		utils.LogInfo(data)
	}
}
func (s *ServiceLogger) Warn(data utils.LogData) {
	if hashValue, err := s.Repository.GetHashValue(data.User); err != nil {
		utils.LogError(utils.LogData{Event: "Convert to hash User error", Details: err.Error(), User: "system"})
	} else {
		data.User = hashValue
		utils.LogWarn(data)
	}
}
func (s *ServiceLogger) Error(data utils.LogData) {
	if hashValue, err := s.Repository.GetHashValue(data.User); err != nil {
		utils.LogError(utils.LogData{Event: "Convert to hash User error", Details: err.Error(), User: "system"})
	} else {
		data.User = hashValue
		utils.LogError(data)
	}
}
func (s *ServiceLogger) Fatal(data utils.LogData) {
	if hashValue, err := s.Repository.GetHashValue(data.User); err != nil {
		utils.LogError(utils.LogData{Event: "Convert to hash User error", Details: err.Error(), User: "system"})
	} else {
		data.User = hashValue
		utils.LogFatal(data)
	}
}
func (s *ServiceLogger) Panic(data utils.LogData) {
	if hashValue, err := s.Repository.GetHashValue(data.User); err != nil {
		utils.LogError(utils.LogData{Event: "Convert to hash User error", Details: err.Error(), User: "system"})
	} else {
		data.User = hashValue
		utils.LogPanic(data)
	}
}
