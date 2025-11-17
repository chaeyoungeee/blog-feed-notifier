package service

import (
	"github.com/chaeyoungeee/blog-feed-notifier/domain"
	"github.com/chaeyoungeee/blog-feed-notifier/repository"
)

type BlogService struct {
	Repo *repository.BlogRepo
}

func NewBlogService(repo *repository.BlogRepo) *BlogService {
	return &BlogService{Repo: repo}
}

func (s *BlogService) GetBlogs() ([]*domain.Blog, error) {
	return s.Repo.GetAll()
}

func (s *BlogService) UpdateLastID(blogID uint, lastID string) error {
	return s.Repo.UpdateLastID(blogID, lastID)
}
