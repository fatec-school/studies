package service

import (
	"email/internal/adapter/dto"
	"email/internal/domain/campaign/repository"
)

type Service struct {
	Repository repository.RepositoryInterface
}

func (s *Service) Create(request *dto.NewCampaignRequest) error {
	return nil
}
