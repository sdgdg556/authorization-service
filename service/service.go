package service

import (
	"auth-service/dao/storage"
	"auth-service/model"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Service struct {
	dao        *storage.Dao
	authConfig *model.AuthConfig
}

func InitService() *Service {
	return &Service{
		dao:        storage.InitDao(),
		authConfig: getAuthConfig(),
	}
}

func getAuthConfig() (config *model.AuthConfig) {
	config = &model.AuthConfig{}
	file, err := os.Open("configs/auth.json")
	if err != nil {
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, config); err != nil {
		return
	}
	return
}
