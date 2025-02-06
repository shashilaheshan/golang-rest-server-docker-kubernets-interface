package main

import (
	"books-api/t/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPosts(res http.ResponseWriter, req *http.Request) {
	logger := NewLogger()

	client, err := ConnectMongoDB()

	if err != nil {
		logger.LogMessage(err.Error())
	}
	postRepo := NewPost("google.com", client, *logger)

	posts, err := postRepo.GetPosts()

	if err != nil {
		logger.LogMessage(err.Error())
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(posts)
}

func AddPost(res http.ResponseWriter, req *http.Request) {

	fmt.Print(req.Method)
	if req.Method != http.MethodPost {
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var post models.Post

	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		http.Error(res, "Error on message payload", http.StatusBadRequest)
		return
	}
	logger := NewLogger()

	client, err := ConnectMongoDB()

	if err != nil {
		logger.LogMessage(err.Error())
	}
	postRepo := NewPost("google.com", client, *logger)

	success, err := postRepo.AddPost(&post)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	if success {
		// Send response
		res.WriteHeader(http.StatusCreated)
		response := map[string]interface{}{
			"message": "Post created successfully",
		}
		json.NewEncoder(res).Encode(response)
	}

}
func main() {
	port := ":8787"
	http.HandleFunc("/posts", GetPosts)
	http.HandleFunc("/addpost", AddPost)
	http.ListenAndServe(port, nil)
}
