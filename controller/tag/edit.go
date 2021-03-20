package tag

import "isso0424/racion-api/types/domain"

func(controller TagController) Edit(id, title, description, color string) (domain.Tag, error) {
	return controller.repo.Edit(id, title, description, color)
}
