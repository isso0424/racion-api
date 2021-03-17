package tag

import "isso0424/racion-api/types/domain"

type MockTagDB struct {
	Data []domain.Tag
}

func(db MockTagDB) Create(title, description, color string) (domain.Tag, error) {
	return domain.Tag{}, nil
}

func(db MockTagDB) Edit(title, description, color string) (domain.Tag, error) {
	return domain.Tag{}, nil
}

func(db MockTagDB) GetAll() ([]domain.Tag, error) {
	return []domain.Tag{}, nil
}

func(db MockTagDB) GetByTitle(title string) (domain.Tag, error) {
	return domain.Tag{}, nil
}
