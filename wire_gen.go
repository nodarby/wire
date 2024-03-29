// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"wiresample/model"
)

// Injectors from wire.go:

func InitializeEvent() model.Event {
	message := model.NewMessage()
	greeter := model.NewGreeter(message)
	event := model.NewEvent(greeter)
	return event
}
