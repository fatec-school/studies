package service

import (
	"email/internal/adapter/dto"
	internalerror "email/internal/domain/campaign/internal-error"
	"email/internal/domain/campaign/model"
	"email/internal/domain/campaign/repository"
)

type Service struct {
	Repository repository.RepositoryInterface
}

func (s *Service) Create(request *dto.NewCampaignRequest) (string, error) {

	campaign, err := model.NewCampaign(request.Name, request.Content, request.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerror.ErrInternalError
	}

	return campaign.ID, nil
}
