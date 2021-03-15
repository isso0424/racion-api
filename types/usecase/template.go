package usecase

import "isso0424/racion-api/types/domain"

type TemplateInteractor interface {
	Create(tags []domain.Tag) domain.Template
	Edit(name string, tags []domain.Tag) domain.Template
	GetAll() []domain.Template
	GetByName(name string) domain.Template
}
