package repository

import (
	"context"

	domain "framework-auto-aprov-go/pkg/domain"

	"gorm.io/gorm"
)

type auto_Aprov_OmnitokDatabase struct {
	DB *gorm.DB
}

func NewAuto_Aprov_OmnitokRepository(DB *gorm.DB) domain.Auto_Aprov_OmnitokRepository {
	return &auto_Aprov_OmnitokDatabase{DB}
}

func (c *auto_Aprov_OmnitokDatabase) FindAll(ctx context.Context) ([]domain.Auto_Aprov_Omnitok, error) {
	var auto_Aprov_Omnitoks []domain.Auto_Aprov_Omnitok
	err := c.DB.Where("deleted = ?", "false").Find(&auto_Aprov_Omnitoks).Error

	return auto_Aprov_Omnitoks, err
}

func (c *auto_Aprov_OmnitokDatabase) FindByID(ctx context.Context, id uint) (domain.Auto_Aprov_Omnitok, error) {
	var auto_Aprov_Omnitok domain.Auto_Aprov_Omnitok
	err := c.DB.Where("deleted = ?", "false").First(&auto_Aprov_Omnitok, id).Error

	return auto_Aprov_Omnitok, err
}

func (c *auto_Aprov_OmnitokDatabase) FindByServicio(ctx context.Context, idservicio uint) ([]domain.Auto_Aprov_Omnitok, error) {
	var auto_Aprov_Omnitoks []domain.Auto_Aprov_Omnitok

	err := c.DB.Where("pro_servicio = ?", idservicio).Find(&auto_Aprov_Omnitoks).Error

	return auto_Aprov_Omnitoks, err
}

func (c *auto_Aprov_OmnitokDatabase) Save(ctx context.Context, auto_Aprov_Omnitok domain.Auto_Aprov_Omnitok) (domain.Auto_Aprov_Omnitok, error) {
	err := c.DB.Save(&auto_Aprov_Omnitok).Error

	return auto_Aprov_Omnitok, err
}

func (c *auto_Aprov_OmnitokDatabase) SaveAll(ctx context.Context, auto_Aprov_Omnitok []domain.Auto_Aprov_Omnitok) (string, error) {
	err := c.DB.Save(&auto_Aprov_Omnitok).Error

	return "save all ok", err
}
func (c *auto_Aprov_OmnitokDatabase) Update(ctx context.Context, auto_Aprov_Omnitok domain.Auto_Aprov_Omnitok) (domain.Auto_Aprov_Omnitok, error) {
	err := c.DB.Save(&auto_Aprov_Omnitok).Error

	return auto_Aprov_Omnitok, err
}

func (c *auto_Aprov_OmnitokDatabase) Delete(ctx context.Context, auto_Aprov_Omnitok domain.Auto_Aprov_Omnitok) error {
	err := c.DB.Delete(&auto_Aprov_Omnitok).Error

	return err
}
