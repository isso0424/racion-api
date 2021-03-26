package template

import "isso0424/racion-api/types/domain"

func (controller TemplateController) GetAll() ([]domain.Template, error) {
	return controller.templateRepo.GetAll()
}

func (controller TemplateController) GetByName(name string) ([]domain.Template, error) {
	return controller.templateRepo.GetByName(name)
}

func (controller TemplateController) GetByID(id string) (domain.Template, error) {
	return controller.templateRepo.GetByID(id)
}
