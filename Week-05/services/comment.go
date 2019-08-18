package services


type CommentImpl struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	PostID int64  `json:"postId"`
}

type Comment interface {
	GetID() int64
	GetBody() string
	GetPostID() int64
}

func (c CommentImpl) GetID() int64{
	return c.ID
}

func (c CommentImpl) GetBody() string{
	return c.Body
}

func (c CommentImpl) GetPostID() int64{
	return c.PostID
}
