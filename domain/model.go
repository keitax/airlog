package domain

import "time"

type Post struct {
	Filename  string
	Timestamp time.Time
	Hash      string
	Title     string
	Body      string
	Labels    []string
}

type ErrNotFound struct {
	error
}

type PushEvent struct {
	BeforeCommitID string `json:"before"`
	AfterCommitID  string `json:"after"`
}

type File struct {
	Path    string
	Content string
}
