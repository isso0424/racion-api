package repository

import "isso0424/racion-api/types/domain"

type TemplateRepository interface {
	Create(name, color string, tags []domain.Tag) (domain.Template, error)
	Edit(name, color string, tags []domain.Tag) (domain.Template, error)
	GetAll() ([]domain.Template, error)
	GetByName(name string) (domain.Template, error)
}
