package tag

import "isso0424/racion-api/types/repository"

type TagContoller struct {
	repo repository.TagRepository
}

func New(repo repository.TagRepository) TagContoller {
	return TagContoller{ repo }
}
