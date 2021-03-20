package tag

import "isso0424/racion-api/types/repository"

type TagController struct {
	repo repository.TagRepository
}

func New(repo repository.TagRepository) TagController {
	return TagController{ repo }
}
