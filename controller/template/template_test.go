package template_test

import (
	controller "isso0424/racion-api/controller/template"
	"isso0424/racion-api/mock/repository/tag"
	"isso0424/racion-api/mock/repository/template"
	"isso0424/racion-api/types/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() controller.TemplateController {
	templateRepo := template.MockTemplateDB{
		Data: []domain.Template{
			{
				Name: "template1",
				Color: "#000000",
				ID: "id",
				Tags: []domain.Tag{
					{
						Title: "tag1",
						Description: "desc",
						ID: "id",
						Color: "#ffffff",
					},
				},
			},
		},
	}

	tagRepo := tag.MockTagDB{
		Data: []domain.Tag{
			{
				Title: "tag1",
				Description: "desc",
				ID: "id",
				Color: "#ffffff",
			},
		},
	}

	return controller.New(&tagRepo, &templateRepo)
}

func TestCreate(t *testing.T) {
	ctrl := setup()
	template, err := ctrl.Create("test", "color", []string{"id"})

	assert.Equal(t, nil, err)
	assert.Equal(t, "test", template.Name)
	assert.Equal(t, "color", template.Color)
	assert.Equal(t, "id", template.Tags[0].ID)
}

func TestEdit(t *testing.T) {
	ctrl := setup()
	template, err := ctrl.Edit("id", "edited", "color", []string{"id"})

	assert.Equal(t, nil, err)
	assert.Equal(t, "edited", template.Name)
	assert.Equal(t, "color", template.Color)
	assert.Equal(t, "id", template.Tags[0].ID)
}
