package repository

import (
	"context"

	"gorm.io/gorm"
	"shotenLinker/pkg/domain"
	interfaces "shotenLinker/pkg/repository/interface"
)

type linkDatabase struct {
	DB *gorm.DB
}

func NewLinkRepository(DB *gorm.DB) interfaces.LinkRepository {
	return &linkDatabase{DB}
}

func (c *linkDatabase) FindByShortLink(ctx context.Context, shortLink string) (domain.Links, error) {
	var link domain.Links
	err := c.DB.WithContext(ctx).Where("short_link = ?", shortLink).First(&link).Error

	return link, err
}

func (c *linkDatabase) FindByLink(ctx context.Context, fullLink string) (domain.Links, error) {
	var link domain.Links
	err := c.DB.WithContext(ctx).Where("link = ?", fullLink).First(&link).Error

	return link, err
}

func (c *linkDatabase) Save(ctx context.Context, link domain.Links) error {
	return c.DB.WithContext(ctx).Save(&link).Error
}
