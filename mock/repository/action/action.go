package action

import (
	"isso0424/racion-api/types/client_error"
	"isso0424/racion-api/types/domain"
	"time"

	"github.com/google/uuid"
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
		ID: uuid.NewString(),
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

	return domain.Action{}, client_error.CreateNotFound("Action", "ID", id)
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
	return []domain.Action{}, client_error.CreateNotFound("Action", "title", title)
}

func(db MockActionDB) GetByTag(tagID string) ([]domain.Action, error) {
	var results []domain.Action
	for _, data := range db.Data {
		for _, tag := range data.Tags {
			if tag.ID == tagID {
				results = append(results, data)
			}
		}
	}

	if len(results) == 0 {
		return results, client_error.CreateNotFound("Action", "Tag ID", tagID)
	}

	return results, nil
}

func(db MockActionDB) GetByID(id string) (domain.Action, error) {
	for _, data := range db.Data {
		if data.ID == id {
			return data, nil
		}
	}
	return domain.Action{}, client_error.CreateNotFound("Action", "ID", id)
}

func(db *MockActionDB) Delete(id string) (domain.Action, error) {
	for index, data := range db.Data {
		if data.ID == id {
			db.Data = append(db.Data[:index], db.Data[index+1:]...)

			return data, nil
		}
	}

	return domain.Action{}, client_error.CreateNotFound("Action", "ID", id)
}
