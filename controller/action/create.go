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

func(controller ActionController) CreateFromTemplate(title string, templateID string, startAt, endAt time.Time) (domain.Action, error) {
	template, err := controller.templateRepo.GetByID(templateID)
	if err != nil {
		return domain.Action{}, err
	}

	var tags []domain.Tag
	for _, tag := range template.Tags {
		t, err := controller.tagRepo.GetByID(tag.ID)
		if err != nil {
			return domain.Action{}, err
		}

		tags = append(tags, t)
	}

	return controller.actionRepo.Create(title, template.Color, template.Tags, startAt, endAt)
}
