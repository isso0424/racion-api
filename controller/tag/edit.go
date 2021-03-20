package tag

import "isso0424/racion-api/types/domain"

func(repo TagController) Edit(title, description, color string) (domain.Tag, error) {
	return repo.Edit(title, description, color)
}
