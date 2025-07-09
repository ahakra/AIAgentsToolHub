package service

import (
	"AIAgentsToolHub/internal/data/model"
	"context"
)

type ToolRepoInterface interface {
	GetToolByDescription(ctx context.Context, description string) ([]model.Tool, error)
}

type ToolService struct {
	toolRepo ToolRepoInterface
}

func NewToolService(toolRepo ToolRepoInterface) ToolService {
	return ToolService{
		toolRepo: toolRepo,
	}
}

func (svc *ToolService) GetToolByDescription(ctx context.Context, description string) ([]model.Tool, error) {

	return svc.toolRepo.GetToolByDescription(ctx, description)
}
