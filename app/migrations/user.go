package migrations

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey; not null;uniqueIndex;"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
