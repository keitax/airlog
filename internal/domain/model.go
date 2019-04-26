package domain

import "time"

type Post struct {
	Filename  string
	Timestamp time.Time
	Hash      string
	Title     string
	Body      string
}

type ErrNotFound struct {
	error
}
