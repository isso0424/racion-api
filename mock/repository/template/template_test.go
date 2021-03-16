package template_test

import (
	"isso0424/racion-api/mock/repository/template"
	"isso0424/racion-api/types/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTemplate(t *testing.T) {
	repository := template.TemplateRepository{}
	repository.Data = []domain.Template{
		{
			Name: "test",
			Color: "#ffffff",
			Tags: []domain.Tag{
				{
					Title: "tags1",
					Description: "desc",
					Color: "#000000",
				},
			},
		},
	}

	template, err := repository.Create("example", "#ff00ff", []domain.Tag{
		{
			Title: "example",
			Description: "desc",
			Color: "#012345",
		},
	})
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "example", template.Name)
	assert.Equal(t, "#ff00ff", template.Color)
	assert.Equal(t, "example", template.Tags[0].Title)
	assert.Equal(t, "desc", template.Tags[0].Description)
	assert.Equal(t, "#012345", template.Tags[0].Color)

	assert.Equal(t, 2, len(repository.Data))
	assert.Equal(t, "example", repository.Data[1].Name)
	assert.Equal(t, "#ff00ff", repository.Data[1].Color)
	assert.Equal(t, "example", repository.Data[1].Tags[0].Title)
	assert.Equal(t, "desc", repository.Data[1].Tags[0].Description)
	assert.Equal(t, "#012345", repository.Data[1].Tags[0].Color)
}

func TestEditTemplate(t *testing.T) {
	repository := template.TemplateRepository{}
	repository.Data = []domain.Template{
		{
			Name: "test",
			Color: "#ffffff",
			Tags: []domain.Tag{
				{
					Title: "tags1",
					Description: "desc",
					Color: "#000000",
				},
			},
		},
	}
	template, err := repository.Edit("test", "#012345", []domain.Tag{
		{
			Title: "hoge",
			Description: "fuga",
			Color: "#543210",
		},
	})
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "test", template.Name)
	assert.Equal(t, "#012345", template.Color)
	assert.Equal(t, "hoge", template.Tags[0].Title)
	assert.Equal(t, "fuga", template.Tags[0].Description)
	assert.Equal(t, "#543210", template.Tags[0].Color)

	assert.Equal(t, 1, len(repository.Data))
	assert.Equal(t, "test", repository.Data[0].Name)
	assert.Equal(t, "#012345", repository.Data[0].Color)
	assert.Equal(t, "hoge", repository.Data[0].Tags[0].Title)
	assert.Equal(t, "fuga", repository.Data[0].Tags[0].Description)
	assert.Equal(t, "#543210", repository.Data[0].Tags[0].Color)
}

func TestGetAll(t *testing.T) {
	repository := template.TemplateRepository{}
	repository.Data = []domain.Template{
		{
			Name: "test",
			Color: "#ffffff",
			Tags: []domain.Tag{
				{
					Title: "tags1",
					Description: "desc",
					Color: "#000000",
				},
			},
		},
	}
	templates, err := repository.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test", templates[0].Name)
	assert.Equal(t, "#ffffff", templates[0].Color)
	assert.Equal(t, "tags1", templates[0].Tags[0].Title)
	assert.Equal(t, "desc", templates[0].Tags[0].Description)
	assert.Equal(t, "#000000", templates[0].Tags[0].Color)
}

func TestGetByName(t *testing.T) {
	repository := template.TemplateRepository{}
	repository.Data = []domain.Template{
		{
			Name: "test",
			Color: "#ffffff",
			Tags: []domain.Tag{
				{
					Title: "tags1",
					Description: "desc",
					Color: "#000000",
				},
			},
		},
	}
	template, err := repository.GetByName("test")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test", template.Name)
	assert.Equal(t, "#ffffff", template.Color)
	assert.Equal(t, "tags1", template.Tags[0].Title)
	assert.Equal(t, "desc", template.Tags[0].Description)
	assert.Equal(t, "#000000", template.Tags[0].Color)
}

func TestFail(t *testing.T) {
	repository := template.TemplateRepository{}
	repository.Data = []domain.Template{
		{
			Name: "test",
			Color: "#ffffff",
			Tags: []domain.Tag{
				{
					Title: "tags1",
					Description: "desc",
					Color: "#000000",
				},
			},
		},
	}
	const errorMsg = "error should occur in here"

	_, err := repository.Create("test", "#ffffff", []domain.Tag{})
	if err == nil {
		t.Fatal(errorMsg)
	}

	_, err = repository.Edit("invalid", "#ffffff", []domain.Tag{})
	if err == nil {
		t.Fatal(errorMsg)
	}

	_, err = repository.GetByName("invalid")
	if err == nil {
		t.Fatal(errorMsg)
	}
}
