package interfaces

import (
	"context"

	"shotenLinker/pkg/domain"
)

type LinkRepository interface {
	FindByShortLink(ctx context.Context, shortLink string) (domain.Links, error)
	FindByLink(ctx context.Context, fullLink string) (domain.Links, error)
	Save(ctx context.Context, link domain.Links) error
}
