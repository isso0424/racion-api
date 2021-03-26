package template

import "isso0424/racion-api/types/repository"

type TemplateController struct {
	tagRepo      repository.TagRepository
	templateRepo repository.TemplateRepository
}

func New(tagRepo repository.TagRepository, templateRepo repository.TemplateRepository) TemplateController {
	return TemplateController{tagRepo, templateRepo}
}
