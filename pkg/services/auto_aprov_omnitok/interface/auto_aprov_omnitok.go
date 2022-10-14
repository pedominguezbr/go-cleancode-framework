package services

import (
	"context"

	domain "framework-auto-aprov-go/pkg/domain"
)

type Auto_Aprov_OmnitokUseCase interface {
	FindAll(ctx context.Context) ([]domain.Auto_Aprov_Omnitok, error)
	FindByID(ctx context.Context, id uint) (domain.Auto_Aprov_Omnitok, error)
	FindByServicio(ctx context.Context, idservicio uint) ([]domain.Auto_Aprov_Omnitok, error)
	Save(ctx context.Context, auto_Aprov_Omnitok domain.Auto_Aprov_Omnitok) (domain.Auto_Aprov_Omnitok, error)
	SaveAll(ctx context.Context, auto_Aprov_Omnitok []domain.Auto_Aprov_Omnitok) (string, error)
	Update(ctx context.Context, auto_Aprov_Omnitok domain.Auto_Aprov_Omnitok) (domain.Auto_Aprov_Omnitok, error)
	Delete(ctx context.Context, auto_Aprov_Omnitok domain.Auto_Aprov_Omnitok) error
}
