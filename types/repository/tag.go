package repository

import "isso0424/racion-api/types/domain"

type TagRepository interface {
	Create(title, description, color string) (domain.Tag, error)
	Edit(title, newTitle, description, color string) (domain.Tag, error)
	GetAll() ([]domain.Tag, error)
	GetByTitle(title string) (domain.Tag, error)
}
