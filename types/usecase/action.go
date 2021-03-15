package usecase

import (
	"isso0424/racion-api/types/domain"
	"time"
)

type ActionInteractor interface {
	Create(title, color string, tags []string, startAt, endAt time.Time) domain.Action
	CreateFromTemplate(title string, template domain.Template, startAt, endAt time.Time) domain.Action
	Edit(title, color string, tags []string) domain.Action
	GetAll() []domain.Action
	GetByTitle() domain.Action
}
