package tag

import (
	"isso0424/racion-api/types/client_error"
	"isso0424/racion-api/types/domain"

	"github.com/google/uuid"
)

type MockTagDB struct {
	Data []domain.Tag
}

func (db *MockTagDB) Create(title, description, color string) (domain.Tag, error) {
	tag := domain.Tag{Title: title, Description: description, Color: color, ID: uuid.NewString()}
	db.Data = append(db.Data, tag)

	return tag, nil
}

func (db *MockTagDB) Edit(id, title, description, color string) (domain.Tag, error) {
	for index, data := range db.Data {
		if data.ID == id {
			db.Data[index].Title = title
			db.Data[index].Description = description
			db.Data[index].Color = color

			return db.Data[index], nil
		}
	}

	return domain.Tag{}, client_error.CreateNotFound("Tag", "ID", id)
}

func (db MockTagDB) GetAll() ([]domain.Tag, error) {
	return db.Data, nil
}

func (db MockTagDB) GetByTitle(title string) (tags []domain.Tag, err error) {
	for _, data := range db.Data {
		if data.Title == title {
			tags = append(tags, data)
		}
	}
	if len(tags) == 0 {
		err = client_error.CreateNotFound("Tag", "Title", title)

		return
	}

	return
}

func (db MockTagDB) GetByID(id string) (domain.Tag, error) {
	for _, data := range db.Data {
		if data.ID == id {
			return data, nil
		}
	}

	return domain.Tag{}, client_error.CreateNotFound("Tag", "ID", id)
}

func (db *MockTagDB) Delete(id string) (domain.Tag, error) {
	for index, data := range db.Data {
		if data.ID == id {
			db.Data = append(db.Data[:index], db.Data[index+1:]...)

			return data, nil
		}
	}

	return domain.Tag{}, client_error.CreateNotFound("Tag", "ID", id)
}
