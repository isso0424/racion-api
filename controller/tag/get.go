package tag

import "isso0424/racion-api/types/domain"

func(controller TagController) GetAll() ([]domain.Tag, error) {
	return controller.repo.GetAll()
}

func(controller TagController) GetByTitle(title string) ([]domain.Tag, error) {
	return controller.repo.GetByTitle(title)
}
