package model

import (
	"strconv"
	"time"
)

type HttpConfig struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type AuthConfig struct {
	TokenExpire string `json:"token_expire_time"`
}

func ParseTime(timeStr string) time.Duration {
	if len(timeStr) == 0 {
		return 0
	}
	unit := timeStr[len(timeStr)-1]
	num, _ := strconv.ParseInt(timeStr[:len(timeStr)-1], 10, 64)
	switch string(unit) {
	case "s":
		return time.Duration(num) * time.Second
	case "m":
		return time.Duration(num) * time.Minute
	case "h":
		return time.Duration(num) * time.Hour
	default:
		return 0
	}
}
