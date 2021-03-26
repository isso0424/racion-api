package template

import "isso0424/racion-api/types/domain"

func (controller TemplateController) Edit(id, name, color string, tags []string) (domain.Template, error) {
	var tagsArray []domain.Tag
	for _, tagID := range tags {
		tag, err := controller.tagRepo.GetByID(tagID)
		if err != nil {
			return domain.Template{}, err
		}

		tagsArray = append(tagsArray, tag)
	}

	return controller.templateRepo.Edit(id, name, color, tagsArray)
}
