package template

import (
	"errors"
	"isso0424/racion-api/types/domain"
)

type TemplateRepository struct {
	Data []domain.Template
}

func(repo *TemplateRepository) Create(name, color string, tags []domain.Tag) (domain.Template, error) {
	for _, data := range repo.Data {
		if data.Name == name {
			return domain.Template{}, errors.New("duplicate name")
		}
	}
	newData := domain.Template{ Name: name, Color: color, Tags: tags }

	repo.Data = append(repo.Data, newData)

	return newData, nil
}

func(repo *TemplateRepository) Edit(name, color string, tags []domain.Tag) (domain.Template, error) {
	for index, data := range repo.Data {
		if data.Name == name {
			repo.Data[index].Color = color
			repo.Data[index].Tags = tags

			return repo.Data[index], nil
		}
	}

	return domain.Template{}, errors.New("target not found")
}

func(repo *TemplateRepository) GetAll() ([]domain.Template, error) {
	return repo.Data, nil
}

func(repo *TemplateRepository) GetByName(name string) (domain.Template, error) {
	for _, data := range repo.Data {
		if data.Name == name {
			return data, nil
		}
	}

	return domain.Template{}, errors.New("target not found")
}
