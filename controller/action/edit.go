package action

import (
	"isso0424/racion-api/types/domain"
	"time"
)

func (controller ActionController) Edit(id, title, color string, tags []string, startAt, endAt time.Time) (domain.Action, error) {
	var tagsArray []domain.Tag
	for _, tagID := range tags {
		tag, err := controller.tagRepo.GetByID(tagID)
		if err != nil {
			return domain.Action{}, err
		}
		tagsArray = append(tagsArray, tag)
	}

	return controller.actionRepo.Edit(id, title, color, tagsArray, startAt, endAt)
}
