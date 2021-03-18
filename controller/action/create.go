package action

import (
	"isso0424/racion-api/types/domain"
	"time"
)

func(controller ActionController) Create(title, color string, tags []string, startAt, endAt time.Time) (domain.Action, error) {
	var tagsArray []domain.Tag
	for _, tagID := range tags {
		result, err := controller.tagRepo.GetByID(tagID)
		if err != nil {
			return domain.Action{}, err
		}

		tagsArray = append(tagsArray, result)
	}

	return controller.actionRepo.Create(title, color, tagsArray, startAt, endAt)
}
