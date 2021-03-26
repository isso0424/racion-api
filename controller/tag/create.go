package tag

import "isso0424/racion-api/types/domain"

func (controller TagController) Create(title, description, color string) (domain.Tag, error) {
	return controller.repo.Create(title, description, color)
}
