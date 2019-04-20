package domain

type Post struct {
	Filename string
	Hash     string
	Title    string
	Body     string
}

type ErrNotFound struct {
	error
}
