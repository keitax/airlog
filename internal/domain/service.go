package domain

type PostService interface {
}

type PostServiceImpl struct {
}

func (ps *PostServiceImpl) Get() (*Post, error) {
	return nil, nil
}
