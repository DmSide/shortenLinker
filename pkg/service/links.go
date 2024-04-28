package service

import (
	"context"

	domains "shotenLinker/pkg/domain"
	interfaces "shotenLinker/pkg/repository/interface"
	services "shotenLinker/pkg/service/interface"
)

type linkUseCase struct {
	linkRepo interfaces.LinkRepository
}

func NewLinkUseCase(repo interfaces.LinkRepository) services.LinkUseCase {
	return &linkUseCase{
		linkRepo: repo,
	}
}

func (c *linkUseCase) FindByShortLink(ctx context.Context, shortLink string) (domains.Links, error) {
	link, err := c.linkRepo.FindByShortLink(ctx, shortLink)
	return link, err
}

func (c *linkUseCase) FindByLink(ctx context.Context, fullLink string) (domains.Links, error) {
	link, err := c.linkRepo.FindByLink(ctx, fullLink)
	return link, err
}

func (c *linkUseCase) Save(ctx context.Context, link domains.Links) error {
	err := c.linkRepo.Save(ctx, link)

	return err
}
