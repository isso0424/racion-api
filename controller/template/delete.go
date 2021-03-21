package template

import "isso0424/racion-api/types/domain"

func(controller TemplateController) Delete(id string) (domain.Template, error) {
	return controller.templateRepo.Delete(id)
}
