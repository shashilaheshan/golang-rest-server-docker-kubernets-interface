package main

import "books-api/t/models"

type PostRepoInterface interface {
	GetPosts() (*[]models.Post, error)
	AddPost(post *models.Post) (bool, error)
}
