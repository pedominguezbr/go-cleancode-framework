package usecase

import (
	"context"

	domain "framework-auto-aprov-go/pkg/domain"
	services "framework-auto-aprov-go/pkg/services/auto_Aprov_Omnitok/interface"
)

type auto_Aprov_OmnitokUseCase struct {
	auto_Aprov_OmnitokRepo domain.Auto_Aprov_OmnitokRepository
}

func NewAuto_Aprov_OmnitokUseCase(repo domain.Auto_Aprov_OmnitokRepository) services.Auto_Aprov_OmnitokUseCase {
	return &auto_Aprov_OmnitokUseCase{
		auto_Aprov_OmnitokRepo: repo,
	}
}

func (c *auto_Aprov_OmnitokUseCase) FindAll(ctx context.Context) ([]domain.Auto_Aprov_Omnitok, error) {
	auto_Aprov_Omnitoks, err := c.auto_Aprov_OmnitokRepo.FindAll(ctx)
	return auto_Aprov_Omnitoks, err
}

func (c *auto_Aprov_OmnitokUseCase) FindByID(ctx context.Context, id uint) (domain.Auto_Aprov_Omnitok, error) {
	auto_Aprov_Omnitok, err := c.auto_Aprov_OmnitokRepo.FindByID(ctx, id)
	return auto_Aprov_Omnitok, err
}

func (c *auto_Aprov_OmnitokUseCase) FindByServicio(ctx context.Context, idservicio uint) ([]domain.Auto_Aprov_Omnitok, error) {
	auto_Aprov_Omnitok, err := c.auto_Aprov_OmnitokRepo.FindByServicio(ctx, idservicio)
	return auto_Aprov_Omnitok, err
}

func (c *auto_Aprov_OmnitokUseCase) Save(ctx context.Context, auto_Aprov_Omnitok domain.Auto_Aprov_Omnitok) (domain.Auto_Aprov_Omnitok, error) {
	auto_Aprov_Omnitok, err := c.auto_Aprov_OmnitokRepo.Save(ctx, auto_Aprov_Omnitok)

	return auto_Aprov_Omnitok, err
}

func (c *auto_Aprov_OmnitokUseCase) SaveAll(ctx context.Context, auto_Aprov_Omnitok []domain.Auto_Aprov_Omnitok) (string, error) {
	_result, err := c.auto_Aprov_OmnitokRepo.SaveAll(ctx, auto_Aprov_Omnitok)

	return _result, err
}

func (c *auto_Aprov_OmnitokUseCase) Update(ctx context.Context, auto_Aprov_Omnitok domain.Auto_Aprov_Omnitok) (domain.Auto_Aprov_Omnitok, error) {
	auto_Aprov_Omnitok, err := c.auto_Aprov_OmnitokRepo.Save(ctx, auto_Aprov_Omnitok)

	return auto_Aprov_Omnitok, err
}

func (c *auto_Aprov_OmnitokUseCase) Delete(ctx context.Context, auto_Aprov_Omnitok domain.Auto_Aprov_Omnitok) error {
	err := c.auto_Aprov_OmnitokRepo.Delete(ctx, auto_Aprov_Omnitok)

	return err
}
