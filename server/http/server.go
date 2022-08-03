package http

import (
	"auth-service/model"
	"auth-service/service"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var s *service.Service

func StartServer() {
	s = service.InitService()
	r := InitRoute()
	r.AddRoute("POST", "/user/create", CreateUserHandler)
	r.AddRoute("POST", "/user/delete", DeleteUserHandler)
	r.AddRoute("POST", "/authorization", AuthorizationHandler)
	r.AddRoute("POST", "/invalidate", InvalidateHandler)
	r.AddRoute("POST", "/role/create", CreateRoleHandler)
	r.AddRoute("POST", "/role/delete", DeleteRoleHandler)
	r.AddRoute("POST", "/user/add-role", UserAddRoleHandler)
	r.AddRoute("POST", "/user/check-role", UserCheckRoleHandler)
	r.AddRoute("POST", "/user/roles", UserAllRolesHandler)

	config := &model.HttpConfig{}
	getHttpConfig(config)
	addr := strings.Join([]string{config.Ip, config.Port}, ":")
	log.Printf("start http server, listing...")
	if err := http.ListenAndServe(addr, r); err != nil {
		panic(err)
	}
}

func getHttpConfig(config *model.HttpConfig) {
	file, err := os.Open("configs/http.json")
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
}
