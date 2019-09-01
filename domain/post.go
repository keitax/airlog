package domain

import (
	"fmt"
	"time"
)

type Post struct {
	Filename string
	Hash     string
	Title    string
	Body     string
	Labels   []string
}

func (p *Post) Timestamp() time.Time {
	ms := filenameRegexp.FindStringSubmatch(p.Filename)
	if len(ms) < 2 {
		panic(fmt.Errorf("must not happen: %v", ms))
	}
	t, err := time.Parse("20060102", ms[1])
	if err != nil {
		panic(err) // must not happen
	}
	return t
}
