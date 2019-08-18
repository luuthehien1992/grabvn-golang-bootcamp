package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	. "./services"
	. "./pb"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)



//TODO: how to separate API logic, business logic and response format logic
func main() {
	http.HandleFunc("/postWithComments", postWithComments)

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getPosts() ([]Post, error) {
	resp, err := http.Get(getPostsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var postImpls []PostImpl
	if err = json.Unmarshal(body, &postImpls); err != nil {
		return nil, err
	}

	var posts = make([]Post, len(postImpls))
	for i:=0; i < len(postImpls); i++{
		posts[i] = postImpls[i]
	}

	return posts, nil
}

func getComments() ([]Comment, error) {
	resp, err := http.Get(getCommentsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var commentImpls []CommentImpl
	if err = json.Unmarshal(body, &commentImpls); err != nil {
		return nil, err
	}

	var comments = make([]Comment, len(commentImpls))
	for i:=0; i < len(commentImpls); i++{
		comments[i] = commentImpls[i]
	}

	return comments, nil
}

func combinePostWithComments(posts []Post, comments []Comment) []PostWithComments {
	commentsByPostID := map[int64][]Comment{}
	for _, comment := range comments {
		commentsByPostID[comment.GetPostID()] = append(commentsByPostID[comment.GetPostID()], comment)
	}

	result := make([]PostWithComments, 0, len(posts))
	for _, post := range posts {
		result = append(result, PostWithCommentsImpl{
			ID:       post.GetID(),
			Title:    post.GetTitle(),
			Comments: commentsByPostID[post.GetID()],
		})
	}

	return result
}

func postWithComments(writer http.ResponseWriter, request *http.Request)  {
	// Get posts from api
	posts, err := getPosts()

	if err != nil {
		log.Println("get posts failed with error: ", err)
		writer.WriteHeader(500)
		return
	}

	// Get comments from api
	comments, err := getComments()
	if err != nil {
		log.Println("get comments failed with error: ", err)
		writer.WriteHeader(500)
		return
	}

	// Combine and return response
	postWithComments := combinePostWithComments(posts, comments)
	resp := PostWithCommentsResponseImpl{Posts: postWithComments}
	buf, err := json.Marshal(resp)
	if err != nil {
		log.Println("unable to parse response: ", err)
		writer.WriteHeader(500)
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(buf)
}
