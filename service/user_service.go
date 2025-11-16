package service

import (
	"fmt"

	"github.com/chaeyoungeee/blog-feed-notifier/domain"
	"github.com/chaeyoungeee/blog-feed-notifier/repository"
)

type UserService struct {
	Repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	if user.Username == "" {
		return fmt.Errorf("유저 이름은 비어 있을 수 없습니다")
	}
	if user.Password == "" {
		return fmt.Errorf("비밀번호는 비어 있을 수 없습니다")
	}

	exists, err := s.Repo.ExistsByUsername(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("이미 존재하는 유저 이름입니다")
	}

	return s.Repo.Create(user)
}
