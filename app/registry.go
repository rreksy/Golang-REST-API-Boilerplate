package app

import (
	migrations "golang-blueprint/app/migrations"
)

type Migration struct {
	Migration interface{}
}

func RegisterMigrations() []Migration {
	return []Migration{
		{Migration: migrations.User{}},
		{Migration: migrations.Address{}},
	}
}
