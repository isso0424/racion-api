package tag

import (
	"errors"
	"isso0424/racion-api/types/domain"
)

type MockTagDB struct {
	Data []domain.Tag
}

func(db *MockTagDB) Create(title, description, color string) (domain.Tag, error) {
	for _, data := range db.Data {
		if data.Title == title {
			return domain.Tag{}, errors.New("duplicate name")
		}
	}
	tag := domain.Tag{ Title: title, Description: description, Color: color }
	db.Data = append(db.Data, tag)

	return tag, nil
}

func(db *MockTagDB) Edit(title, newTitle, description, color string) (domain.Tag, error) {
	for index, data := range db.Data {
		if data.Title == title {
			db.Data[index].Title = newTitle
			db.Data[index].Description = description
			db.Data[index].Color = color

			return db.Data[index], nil
		}
	}

	return domain.Tag{}, errors.New("target not found")
}

func(db MockTagDB) GetAll() ([]domain.Tag, error) {
	return db.Data, nil
}

func(db MockTagDB) GetByTitle(title string) (domain.Tag, error) {
	for _, data := range db.Data {
		if data.Title == title {
			return data, nil
		}
	}
	return domain.Tag{}, errors.New("target not found")
}
