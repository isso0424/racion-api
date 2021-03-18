package action

import (
	"isso0424/racion-api/types/domain"
	"time"
)

type MockActionDB struct {
	Data []domain.Action
}

func(db *MockActionDB) Create(title, color string, tags []domain.Tag, startAt, endAt time.Time) (domain.Action, error) {
	return domain.Action{}, nil
}

func(db *MockActionDB) Edit(id, title, color string, tags []domain.Tag, startAt, endAt time.Time) (domain.Action, error) {
	return domain.Action{}, nil
}

func(db MockActionDB) GetAll() ([]domain.Action, error) {
	return []domain.Action{}, nil
}

func(db MockActionDB) GetByTitle(title string) ([]domain.Action, error) {
	return []domain.Action{}, nil
}

func(db MockActionDB) GetByID(id string) (domain.Action, error) {
	return domain.Action{}, nil
}
