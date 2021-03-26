package tag

import "isso0424/racion-api/types/domain"

func (controller TagController) Delete(id string) (domain.Tag, error) {
	return controller.repo.Delete(id)
}
