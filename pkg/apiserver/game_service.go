package apiserver

import "github.com/finarfin/q3party/pkg/gamelister"

type GameService struct {
	repository GameRepository
}

func (s *GameService) List() ([]*gamelister.Game, error) {
	return s.repository.List()
}

func NewGameService(repository GameRepository) (*GameService, error) {
	return &GameService{
		repository: repository,
	}, nil
}
