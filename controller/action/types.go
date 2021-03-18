package action

import "isso0424/racion-api/types/repository"

type ActionController struct {
	actionRepo repository.ActionRepository
	tagRepo repository.TagRepository
	templateRepo repository.TemplateRepository
}

func New(
	actionRepo repository.ActionRepository,
	tagRepo repository.TagRepository,
	templateRepo repository.TemplateRepository,
) ActionController {
	return ActionController{ actionRepo, tagRepo, templateRepo }
}
