package app

import (
	addressModel "golang-blueprint/app/address"
	userModel "golang-blueprint/app/user"
)

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: userModel.User{}},
		{Model: addressModel.Address{}},
	}
}
