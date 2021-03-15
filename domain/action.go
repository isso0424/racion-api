package domain

import "time"

type Action struct {
	Title string
	Tags []Tag
	Color string
	StartAt time.Time
	EndAt time.Time
}
