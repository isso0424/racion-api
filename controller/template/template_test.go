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
				Name:  "template1",
				Color: "#000000",
				ID:    "id",
				Tags: []domain.Tag{
					{
						Title:       "tag1",
						Description: "desc",
						ID:          "tag_id",
						Color:       "#ffffff",
					},
				},
			},
		},
	}

	tagRepo := tag.MockTagDB{
		Data: []domain.Tag{
			{
				Title:       "tag1",
				Description: "desc",
				ID:          "tag_id",
				Color:       "#ffffff",
			},
		},
	}

	return controller.New(&tagRepo, &templateRepo)
}

func TestCreate(t *testing.T) {
	ctrl := setup()
	template, err := ctrl.Create("test", "color", []string{"tag_id"})

	assert.Equal(t, nil, err)
	assert.Equal(t, "test", template.Name)
	assert.Equal(t, "color", template.Color)
	assert.Equal(t, "tag_id", template.Tags[0].ID)
}

func TestEdit(t *testing.T) {
	ctrl := setup()
	template, err := ctrl.Edit("id", "edited", "color", []string{"tag_id"})

	assert.Equal(t, nil, err)
	assert.Equal(t, "edited", template.Name)
	assert.Equal(t, "color", template.Color)
	assert.Equal(t, "tag_id", template.Tags[0].ID)
}

func TestGet(t *testing.T) {
	ctrl := setup()
	templates, err := ctrl.GetAll()
	assert.Equal(t, nil, err)
	assert.Equal(t, "template1", templates[0].Name)
	assert.Equal(t, "#000000", templates[0].Color)
	assert.Equal(t, "id", templates[0].ID)
	assert.Equal(t, "tag_id", templates[0].Tags[0].ID)

	templates, err = ctrl.GetByName("template1")
	assert.Equal(t, nil, err)
	assert.Equal(t, "template1", templates[0].Name)
	assert.Equal(t, "#000000", templates[0].Color)
	assert.Equal(t, "id", templates[0].ID)
	assert.Equal(t, "tag_id", templates[0].Tags[0].ID)

	template, err := ctrl.GetByID("id")
	assert.Equal(t, nil, err)
	assert.Equal(t, "template1", template.Name)
	assert.Equal(t, "#000000", template.Color)
	assert.Equal(t, "id", template.ID)
	assert.Equal(t, "tag_id", template.Tags[0].ID)
}

func TestDelete(t *testing.T) {
	ctrl := setup()
	template, err := ctrl.Delete("id")

	assert.Equal(t, nil, err)
	assert.Equal(t, "template1", template.Name)
	assert.Equal(t, "#000000", template.Color)
	assert.Equal(t, "id", template.ID)
	assert.Equal(t, "tag_id", template.Tags[0].ID)

	_, err = ctrl.GetByID("id")
	assert.NotEqual(t, nil, err)
}

func TestFail(t *testing.T) {
	ctrl := setup()
	_, err := ctrl.Create("hoge", "fuga", []string{"invalid"})
	assert.NotEqual(t, nil, err)

	_, err = ctrl.Edit("invalid", "hoge", "fuga", []string{"invalid"})
	assert.NotEqual(t, nil, err)
}
