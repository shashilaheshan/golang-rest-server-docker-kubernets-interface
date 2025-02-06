package main

import (
	"books-api/t/mocks"
	"books-api/t/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	// Update this with the correct import path
)

// Test_GetPosts_Success tests GetPosts with mock response
func Test_GetPosts_Success(t *testing.T) {
	mockRepo := new(mocks.PostRepoInterface)

	expectedPosts := []models.Post{
		{ID: "1", Name: "Post1", Author: "Author1", Client: "Client1"},
		{ID: "2", Name: "Post2", Author: "Author2", Client: "Client2"},
	}

	mockRepo.On("GetPosts").Return(&expectedPosts, nil)

	posts, err := mockRepo.GetPosts()

	assert.NoError(t, err)
	assert.NotNil(t, posts)
	assert.Len(t, *posts, 2)
	assert.Equal(t, "Post1", (*posts)[0].Name)
	mockRepo.AssertExpectations(t)
}

// Test_GetPosts_Error tests GetPosts returning an error
func Test_GetPosts_Error(t *testing.T) {
	mockRepo := new(mocks.PostRepoInterface)

	mockRepo.On("GetPosts").Return(nil, errors.New("database error"))

	posts, err := mockRepo.GetPosts()

	assert.Error(t, err)
	assert.Nil(t, posts)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}

// Test_AddPost_Success tests AddPost with a successful insert
func Test_AddPost_Success(t *testing.T) {
	mockRepo := new(mocks.PostRepoInterface)

	post := &models.Post{Name: "New Post", Author: "AuthorX", Client: "ClientX"}
	mockRepo.On("AddPost", post).Return(true, nil)

	success, err := mockRepo.AddPost(post)

	assert.NoError(t, err)
	assert.True(t, success)
	mockRepo.AssertExpectations(t)
}

// Test_AddPost_Error tests AddPost returning an error
func Test_AddPost_Error(t *testing.T) {
	mockRepo := new(mocks.PostRepoInterface)

	post := &models.Post{Name: "New Post", Author: "AuthorX", Client: "ClientX"}
	mockRepo.On("AddPost", post).Return(false, errors.New("insert error"))

	success, err := mockRepo.AddPost(post)

	assert.Error(t, err)
	assert.False(t, success)
	assert.Equal(t, "insert error", err.Error())
	mockRepo.AssertExpectations(t)
}
