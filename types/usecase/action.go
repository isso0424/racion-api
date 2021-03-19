package usecase

import (
	"isso0424/racion-api/types/domain"
	"time"
)

type ActionInteractor interface {
	Create(title, color string, tags []string, startAt, endAt time.Time) (domain.Action, error)
	CreateFromTemplate(title, templateID string, startAt, endAt time.Time) (domain.Action, error)
	Edit(id, title, color string, tags []string, startAt, endAt time.Time) (domain.Action, error)
	GetAll() ([]domain.Action, error)
	GetByTitle(title string) ([]domain.Action, error)
	GetByID(id string) (domain.Action, error)
}
