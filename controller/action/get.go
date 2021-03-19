package action

import "isso0424/racion-api/types/domain"

func(controller ActionController) GetAll() ([]domain.Action, error) {
	return controller.actionRepo.GetAll()
}

func(controller ActionController) GetByTitle(title string) ([]domain.Action, error) {
	return controller.actionRepo.GetByTitle(title)
}

func(controller ActionController) GetByID(id string) (domain.Action, error) {
	return controller.actionRepo.GetByID(id)
}
