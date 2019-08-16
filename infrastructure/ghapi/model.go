package ghapi

type CompareResponse struct {
	Files []*File `json:"files"`
}

type File struct {
	Filename string `json:"filename"`
	RawURL   string `json:"raw_url"`
}
