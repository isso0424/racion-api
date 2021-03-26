package template

import (
	"errors"
	"isso0424/racion-api/types/client_error"
	"isso0424/racion-api/types/domain"

	"github.com/google/uuid"
)

type MockTemplateDB struct {
	Data []domain.Template
}

func (repo *MockTemplateDB) Create(name, color string, tags []domain.Tag) (domain.Template, error) {
	for _, data := range repo.Data {
		if data.Name == name {
			return domain.Template{}, errors.New("duplicate name")
		}
	}
	newData := domain.Template{Name: name, Color: color, Tags: tags, ID: uuid.NewString()}

	repo.Data = append(repo.Data, newData)

	return newData, nil
}

func (repo *MockTemplateDB) Edit(id, name, color string, tags []domain.Tag) (domain.Template, error) {
	for index, data := range repo.Data {
		if data.ID == id {
			repo.Data[index].Name = name
			repo.Data[index].Color = color
			repo.Data[index].Tags = tags

			return repo.Data[index], nil
		}
	}

	return domain.Template{}, client_error.CreateNotFound("Template", "ID", id)
}

func (repo *MockTemplateDB) GetAll() ([]domain.Template, error) {
	return repo.Data, nil
}

func (repo *MockTemplateDB) GetByName(name string) (templates []domain.Template, err error) {
	for _, data := range repo.Data {
		if data.Name == name {
			templates = append(templates, data)
		}
	}

	if len(templates) == 0 {
		return []domain.Template{}, client_error.CreateNotFound("Template", "Name", name)
	}
	return
}

func (repo *MockTemplateDB) GetByID(id string) (domain.Template, error) {
	for _, data := range repo.Data {
		if data.ID == id {
			return data, nil
		}
	}

	return domain.Template{}, client_error.CreateNotFound("Template", "ID", id)
}

func (db *MockTemplateDB) Delete(id string) (domain.Template, error) {
	for index, data := range db.Data {
		if data.ID == id {
			db.Data = append(db.Data[:index], db.Data[index+1:]...)

			return data, nil
		}
	}

	return domain.Template{}, client_error.CreateNotFound("Template", "ID", id)
}
