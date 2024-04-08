package hypertables

import (
	"log/slog"
)

type HypertablesService interface {
	GetHypertables(filter *HypertablesFilter) ([]HypertableInfo, error)
}

type hypertablesService struct {
	repo   HypertablesRepository
	logger *slog.Logger
}

func NewHypertablesService(repo HypertablesRepository, logger *slog.Logger) HypertablesService {
	return &hypertablesService{
		repo:   repo,
		logger: logger,
	}
}

func (h *hypertablesService) GetHypertables(filter *HypertablesFilter) ([]HypertableInfo, error) {
	return h.repo.GetHypertables(filter)
}
