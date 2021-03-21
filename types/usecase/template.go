package usecase

import "isso0424/racion-api/types/domain"

type TemplateInteractor interface {
	Create(name, color string, tags []string) (domain.Template, error)
	Edit(id, name, color string, tags []string) (domain.Template, error)
	GetAll() ([]domain.Template, error)
	GetByName(name string) ([]domain.Template, error)
	GetByID(id string) (domain.Template, error)
	Delete(id string) (domain.Template, error)
}
