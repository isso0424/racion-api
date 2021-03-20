package usecase

import "isso0424/racion-api/types/domain"

type TagInteractor interface {
	Create(title, description, color string) (domain.Tag, error)
	Edit(id, title, description, color string) (domain.Tag, error)
	GetAll() ([]domain.Tag, error)
	GetByTitle(title string) ([]domain.Tag, error)
	GetByID(id string) (domain.Tag, error)
}
