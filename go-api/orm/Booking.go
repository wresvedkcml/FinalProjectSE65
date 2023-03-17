package orm

import (
	"time"

	"gorm.io/gorm" //framwork ต่อกะบ database ภาษา go
)

type Booking struct {
	gorm.Model
	UserID string
	CarID string
	Start time.Time
	End time.Time
}