package variables

import (
	"isso0424/racion-api/controller/action"
	"isso0424/racion-api/controller/tag"
	"isso0424/racion-api/controller/template"
	"isso0424/racion-api/types/repository"
)

func New(
	actionRepo repository.ActionRepository,
	tagRepo repository.TagRepository,
	templateRepo repository.TemplateRepository,
) {
	ActionController = action.New(
		actionRepo,
		tagRepo,
		templateRepo,
	)

	TagController = tag.New(tagRepo)

	TemplateController = template.New(
		tagRepo,
		templateRepo,
	)
}
