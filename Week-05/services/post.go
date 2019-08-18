package services

type PostImpl struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type Post interface {
	GetID() int64
	GetTitle() string
}

func (p PostImpl) GetID() int64{
	return p.ID
}

func (p PostImpl) GetTitle() string{
	return p.Title
}
