package service

import (
	"errors"

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
	exists, err := s.Repo.ExistsByUsername(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("이미 존재하는 유저 이름입니다")
	}

	return s.Repo.Create(user)
}

func (s *UserService) Login(username, password string) (*domain.User, error) {
	user, err := s.Repo.GetByUsername(username)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, errors.New("유저 이름 혹은 비밀번호가 일치하지 않습니다")
		}
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("유저 이름 혹은 비밀번호가 일치하지 않습니다")
	}

	return user, nil
}

func (s *UserService) SetDiscordWebhook(userID uint, webhookURL string) error {
	return s.Repo.UpdateDiscordWebhook(userID, webhookURL)
}
