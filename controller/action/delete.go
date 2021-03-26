package action

import "isso0424/racion-api/types/domain"

func (controller ActionController) Delete(id string) (domain.Action, error) {
	return controller.actionRepo.Delete(id)
}
