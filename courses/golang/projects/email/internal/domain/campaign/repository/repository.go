package repository

import "email/internal/domain/campaign/model"

type RepositoryInterface interface {
	Save(campaign *model.Campaign) error
}
