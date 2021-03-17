package template

import (
	"errors"
	"isso0424/racion-api/types/domain"
)

type MockTemplateDB struct {
	Data []domain.Template
}

func(repo *MockTemplateDB) Create(name, color string, tags []domain.Tag) (domain.Template, error) {
	for _, data := range repo.Data {
		if data.Name == name {
			return domain.Template{}, errors.New("duplicate name")
		}
	}
	newData := domain.Template{ Name: name, Color: color, Tags: tags }

	repo.Data = append(repo.Data, newData)

	return newData, nil
}

func(repo *MockTemplateDB) Edit(id, name, color string, tags []domain.Tag) (domain.Template, error) {
	for index, data := range repo.Data {
		if data.ID == id {
			repo.Data[index].Name = name
			repo.Data[index].Color = color
			repo.Data[index].Tags = tags

			return repo.Data[index], nil
		}
	}

	return domain.Template{}, errors.New("target not found")
}

func(repo *MockTemplateDB) GetAll() ([]domain.Template, error) {
	return repo.Data, nil
}

func(repo *MockTemplateDB) GetByName(name string) (templates []domain.Template, err error) {
	for _, data := range repo.Data {
		if data.Name == name {
			templates = append(templates, data)
		}
	}

	if len(templates) == 0 {
		return []domain.Template{}, errors.New("target not found")
	}
	return
}

func(repo *MockTemplateDB) GetByID(id string) (domain.Template, error) {
	for _, data := range repo.Data {
		if data.ID == id {
			return data, nil
		}
	}

	return domain.Template{}, errors.New("target not found")
}
