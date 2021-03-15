package usecase

import "isso0424/racion-api/types/domain"

type TagInteractor interface {
	Create(title, description, color string) domain.Tag
	Edit(title, description, color string) domain.Tag
	GetAll() []domain.Tag
	GetByTitle(title string) domain.Tag
}
