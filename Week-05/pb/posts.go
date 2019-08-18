package pb

import (
	. "../services"
)

type PostWithCommentsResponseImpl struct {
	Posts []PostWithComments `json:"posts"`
}

func (p PostWithCommentsResponseImpl) GetPosts() []PostWithComments {
	return p.Posts
}

type PostWithCommentsResponse interface {
	GetPosts() []PostWithComments
}

type PostWithCommentsImpl struct {
	ID       int64     `json:"id"`
	Title    string    `json:"string"`
	Comments []Comment `json:"comments,omitempty"`
}

type PostWithComments interface {
	GetID() int64
	GetTitle() string
	GetComments() []Comment
}


func (p PostWithCommentsImpl) GetID() int64 {
	return p.ID
}

func (p PostWithCommentsImpl) GetTitle() string {
	return p.Title
}

func (p PostWithCommentsImpl) GetComments() []Comment {
	return p.Comments
}
