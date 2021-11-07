//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"
import "wiresample/model"

func InitializeEvent() model.Event {
	wire.Build(model.NewEvent, model.NewGreeter, model.NewMessage)
	return model.Event{}
}
