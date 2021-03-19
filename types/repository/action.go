package repository

import (
	"isso0424/racion-api/types/domain"
	"time"
)

type ActionRepository interface {
	Create(title, color string, tags []domain.Tag, startAt, endAt time.Time) (domain.Action, error)
	Edit(id, title, color string, tags []domain.Tag, startAt, endAt time.Time) (domain.Action, error)
	GetAll() ([]domain.Action, error)
	GetByTitle(title string) ([]domain.Action, error)
	GetByTag(tagID string) ([]domain.Action, error)
	GetByID(id string) (domain.Action, error)
}
