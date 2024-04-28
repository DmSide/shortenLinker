package interfaces

import (
	"context"

	domains "shotenLinker/pkg/domain"
)

type LinkUseCase interface {
	FindByShortLink(ctx context.Context, shortLink string) (domains.Links, error)
	FindByLink(ctx context.Context, fullLink string) (domains.Links, error)
	Save(ctx context.Context, link domains.Links) error
}
