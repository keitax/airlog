package domain

import "time"

type Post struct {
	Timestamp time.Time
	Title     string
	Body      string
}
