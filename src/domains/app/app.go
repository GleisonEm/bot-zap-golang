package app

import (
	"context"
	"time"

	"go.mau.fi/whatsmeow"
)

type IAppService interface {
	Login(ctx context.Context) (response LoginResponse, err error)
	Logout(ctx context.Context) (err error)
	Reconnect(ctx context.Context) (err error)
	FirstDevice(ctx context.Context) (response DevicesResponse, err error)
	FetchDevices(ctx context.Context) (response []DevicesResponse, err error)
	GetWaCli(ctx context.Context) *whatsmeow.Client
}

type DevicesResponse struct {
	Name   string `json:"name"`
	Device string `json:"device"`
}

type LoginResponse struct {
	ImagePath string        `json:"image_path"`
	Duration  time.Duration `json:"duration"`
	Code      string        `json:"code"`
}
