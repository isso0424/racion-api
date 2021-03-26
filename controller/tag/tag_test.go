package tag_test

import (
	controller "isso0424/racion-api/controller/tag"
	"isso0424/racion-api/mock/repository/tag"
	"isso0424/racion-api/types/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() controller.TagController {
	repo := tag.MockTagDB{
		Data: []domain.Tag{
			{
				Title:       "tag1",
				Color:       "#000000",
				Description: "desc",
				ID:          "id",
			},
		},
	}

	return controller.New(&repo)
}

func TestCreate(t *testing.T) {
	ctrl := setup()
	tag, err := ctrl.Create("test", "desc", "color")

	assert.Equal(t, nil, err)
	assert.Equal(t, "test", tag.Title)
	assert.Equal(t, "desc", tag.Description)
	assert.Equal(t, "color", tag.Color)
}

func TestEdit(t *testing.T) {
	ctrl := setup()
	tag, err := ctrl.Edit("id", "edited1", "edited2", "edited3")

	assert.Equal(t, nil, err)
	assert.Equal(t, "edited1", tag.Title)
	assert.Equal(t, "edited2", tag.Description)
	assert.Equal(t, "edited3", tag.Color)
}

func TestGet(t *testing.T) {
	ctrl := setup()
	tags, err := ctrl.GetAll()
	assert.Equal(t, nil, err)
	assert.Equal(t, "tag1", tags[0].Title)
	assert.Equal(t, "#000000", tags[0].Color)
	assert.Equal(t, "desc", tags[0].Description)

	tags, err = ctrl.GetByTitle("tag1")
	assert.Equal(t, nil, err)
	assert.Equal(t, "tag1", tags[0].Title)
	assert.Equal(t, "#000000", tags[0].Color)
	assert.Equal(t, "desc", tags[0].Description)

	tag, err := ctrl.GetByID("id")
	assert.Equal(t, nil, err)
	assert.Equal(t, "tag1", tag.Title)
	assert.Equal(t, "#000000", tag.Color)
	assert.Equal(t, "desc", tag.Description)
}

func TestDelete(t *testing.T) {
	ctrl := setup()
	tag, err := ctrl.Delete("id")

	assert.Equal(t, nil, err)
	assert.Equal(t, "tag1", tag.Title)
	assert.Equal(t, "#000000", tag.Color)
	assert.Equal(t, "desc", tag.Description)

	_, err = ctrl.GetByID("id")
	assert.NotEqual(t, nil, err)
}
