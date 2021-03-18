package action

import (
	"errors"
	"isso0424/racion-api/types/domain"
	"time"
)

type MockActionDB struct {
	Data []domain.Action
}

func(db *MockActionDB) Create(title, color string, tags []domain.Tag, startAt, endAt time.Time) (domain.Action, error) {
	newData := domain.Action{
		Title: title,
		Color: color,
		Tags: tags,
		StartAt: startAt,
		EndAt: endAt,
	}

	db.Data = append(db.Data, newData)

	return newData, nil
}

func(db *MockActionDB) Edit(id, title, color string, tags []domain.Tag, startAt, endAt time.Time) (domain.Action, error) {
	for index, data := range db.Data {
		if data.ID == id {
			db.Data[index].Title = title
			db.Data[index].Color = color
			db.Data[index].Tags = tags
			db.Data[index].StartAt = startAt
			db.Data[index].EndAt = endAt

			return db.Data[index], nil
		}
	}

	return domain.Action{}, errors.New("target not found")
}

func(db MockActionDB) GetAll() ([]domain.Action, error) {
	return db.Data, nil
}

func(db MockActionDB) GetByTitle(title string) ([]domain.Action, error) {
	var result []domain.Action
	for _, data := range db.Data {
		if data.Title == title {
			result = append(result, data)
		}
	}

	if len(result) != 0 {
		return result, nil
	}
	return []domain.Action{}, errors.New("target not found")
}

func(db MockActionDB) GetByID(id string) (domain.Action, error) {
	for _, data := range db.Data {
		if data.ID == id {
			return data, nil
		}
	}
	return domain.Action{}, errors.New("target not found")
}
