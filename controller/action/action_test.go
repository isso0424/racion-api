package action_test

import (
	controller "isso0424/racion-api/controller/action"
	"isso0424/racion-api/mock/repository/action"
	"isso0424/racion-api/mock/repository/tag"
	"isso0424/racion-api/mock/repository/template"
	"isso0424/racion-api/types/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	location = time.FixedZone("Asia/Tokyo", 9*60*60)
	startAt = time.Date(2000, time.April, 1, 1, 0, 0, 0, location)
	endAt = time.Date(2001, time.April, 1, 1, 0, 0, 0, location)
)

func setup() controller.ActionController {
	actionRepo := action.MockActionDB{
		Data: []domain.Action{
			{
				ID: "id",
				Title: "action1",
				Color: "#ffffff",
				Tags: []domain.Tag{
					{
						Title: "tag1",
						Description: "desc",
						Color: "#123456",
						ID: "id",
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
				Color: "#123456",
				ID: "id",
			},
		},
	}

	templateRepo := template.MockTemplateDB{
		Data: []domain.Template{
			{
				ID: "id",
				Name: "template1",
				Color: "#000000",
				Tags: []domain.Tag{
					{
						Title: "tag1",
						Description: "desc",
						Color: "#123456",
						ID: "id",
					},
				},
			},
			{
				ID: "broken",
				Name: "broken",
				Color: "#000000",
				Tags: []domain.Tag{
					{
						Title: "tag1",
						Description: "desc",
						Color: "#123456",
						ID: "broken",
					},
				},
			},
		},
	}

	return controller.New(&actionRepo, &tagRepo, &templateRepo)
}

func TestCreate(t *testing.T) {
	ctrl := setup()
	action, err := ctrl.Create("title", "color", []string{"id"}, startAt, endAt)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "title", action.Title)
	assert.Equal(t, "color", action.Color)
	assert.Equal(t, "tag1", action.Tags[0].Title)
	assert.Equal(t, startAt, action.StartAt)
	assert.Equal(t, endAt, action.EndAt)

	action, err = ctrl.CreateFromTemplate("title2", "id", startAt, endAt)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "title2", action.Title)
	assert.Equal(t, "#000000", action.Color)
	assert.Equal(t, "tag1", action.Tags[0].Title)
	assert.Equal(t, startAt, action.StartAt)
	assert.Equal(t, endAt, action.EndAt)
}

func TestEdit(t *testing.T) {
	ctrl := setup()
	action, err := ctrl.Edit("id", "title", "color", []string{"id"}, startAt, endAt)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "title", action.Title)
	assert.Equal(t, "color", action.Color)
	assert.Equal(t, "tag1", action.Tags[0].Title)
	assert.Equal(t, startAt, action.StartAt)
	assert.Equal(t, endAt, action.EndAt)
}

func TestGet(t *testing.T) {
	ctrl := setup()
	actions, err := ctrl.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "action1", actions[0].Title)
	assert.Equal(t, "#ffffff", actions[0].Color)
	assert.Equal(t, "id", actions[0].ID)

	actions, err = ctrl.GetByTitle("action1")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "action1", actions[0].Title)
	assert.Equal(t, "#ffffff", actions[0].Color)
	assert.Equal(t, "id", actions[0].ID)

	action, err := ctrl.GetByID("id")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "action1", action.Title)
	assert.Equal(t, "#ffffff", action.Color)
	assert.Equal(t, "id", action.ID)
}

func TestFail(t *testing.T) {
	ctrl := setup()
	_, err := ctrl.Create("invalid", "invalid", []string{"invalid"}, startAt, endAt)
	assert.NotEqual(t, nil, err)

	_, err = ctrl.CreateFromTemplate("invalid", "invalid", startAt, endAt)
	assert.NotEqual(t, nil, err)

	_, err = ctrl.CreateFromTemplate("invalid", "broken", startAt, endAt)
	assert.NotEqual(t, nil, err)

	_, err = ctrl.Edit("invalid", "invalid", "invalid", []string{"id"}, startAt, endAt)
	assert.NotEqual(t, nil, err)

	_, err = ctrl.Edit("id", "id", "invalid", []string{"invalid"}, startAt, endAt)
	assert.NotEqual(t, nil, err)
}
