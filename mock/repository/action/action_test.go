package action_test

import (
	"isso0424/racion-api/mock/repository/action"
	"isso0424/racion-api/types/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var location = time.FixedZone("Asia/Tokyo", 9*60*60)

func setup() action.MockActionDB {
	return action.MockActionDB{
		Data: []domain.Action{
			{
				Title: "title",
				Tags: []domain.Tag{
					{
						Title: "tag",
						Color: "#ffffff",
						Description: "desc",
						ID: "id",
					},
				},
				Color: "#123456",
				StartAt: time.Date(2000, time.January, 1, 0, 0, 0, 0, location),
				EndAt: time.Date(2000, time.April, 1, 0, 0, 0, 0, location),
				ID: "id",
			},
		},
	}
}

func TestCreate(t *testing.T) {
	repo := setup()
	startAt := time.Date(2010, time.April, 1, 0, 0, 0, 0, location)
	endAt := time.Date(2010, time.August, 1, 0, 0, 0, 0, location)
	action, err := repo.Create("hoge", "fuga", []domain.Tag{
			{
				Title: "tag1",
				Description: "desc1",
				Color: "#f0f0f0",
			},
		},
		startAt,
		endAt,
	)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "hoge", action.Title)
	assert.Equal(t, "fuga", action.Color)
	assert.Equal(t, "tag1", action.Tags[0].Title)
	assert.Equal(t, "desc1", action.Tags[0].Description)
	assert.Equal(t, "#f0f0f0", action.Tags[0].Color)
	assert.Equal(t, startAt, action.StartAt)
	assert.Equal(t, endAt, action.EndAt)
}

func TestEdit(t *testing.T) {
	repo := setup()
	startAt := time.Date(2010, time.April, 1, 0, 0, 0, 0, location)
	endAt := time.Date(2010, time.August, 1, 0, 0, 0, 0, location)
	action, err := repo.Edit(
		repo.Data[0].ID,
		"hoge",
		"fuga",
		[]domain.Tag{
			{
				Title: "tag1",
				Description: "desc1",
				Color: "#f0f0f0",
			},
		},
		startAt,
		endAt,
	)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "hoge", action.Title)
	assert.Equal(t, "fuga", action.Color)
	assert.Equal(t, "tag1", action.Tags[0].Title)
	assert.Equal(t, "desc1", action.Tags[0].Description)
	assert.Equal(t, "#f0f0f0", action.Tags[0].Color)
	assert.Equal(t, startAt, action.StartAt)
	assert.Equal(t, endAt, action.EndAt)
}

func TestGetAll(t *testing.T) {
	repo := setup()
	actions, err := repo.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "title", actions[0].Title)
	assert.Equal(t, "#123456", actions[0].Color)
	assert.Equal(t, "id", actions[0].ID)
	assert.Equal(t, "tag", actions[0].Tags[0].Title)
	assert.Equal(t, "#ffffff", actions[0].Tags[0].Color)
	assert.Equal(t, "desc", actions[0].Tags[0].Description)
}

func TestGetByTitle(t *testing.T) {
	repo := setup()
	actions, err := repo.GetByTitle("title")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "title", actions[0].Title)
	assert.Equal(t, "#123456", actions[0].Color)
	assert.Equal(t, "id", actions[0].ID)
	assert.Equal(t, "tag", actions[0].Tags[0].Title)
	assert.Equal(t, "#ffffff", actions[0].Tags[0].Color)
	assert.Equal(t, "desc", actions[0].Tags[0].Description)
}

func TestGetByID(t *testing.T) {
	repo := setup()
	action, err := repo.GetByID("id")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "title", action.Title)
	assert.Equal(t, "#123456", action.Color)
	assert.Equal(t, "id", action.ID)
	assert.Equal(t, "tag", action.Tags[0].Title)
	assert.Equal(t, "#ffffff", action.Tags[0].Color)
	assert.Equal(t, "desc", action.Tags[0].Description)
}

func TestGetByTag(t *testing.T) {
	repo := setup()
	actions, err := repo.GetByTag("id")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "title", actions[0].Title)
	assert.Equal(t, "#123456", actions[0].Color)
	assert.Equal(t, "id", actions[0].ID)
	assert.Equal(t, "tag", actions[0].Tags[0].Title)
	assert.Equal(t, "#ffffff", actions[0].Tags[0].Color)
	assert.Equal(t, "desc", actions[0].Tags[0].Description)
}

func TestDelete(t *testing.T) {
	repo := setup()
	action, err := repo.Delete("id")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "title", action.Title)
	assert.Equal(t, "#123456", action.Color)
	assert.Equal(t, "id", action.ID)
	assert.Equal(t, "tag", action.Tags[0].Title)

	_, err = repo.GetByID("id")
	assert.NotEqual(t, nil, err)
}

func TestFail(t *testing.T) {
	repo := setup()

	_, err := repo.Edit("invalid", "", "", []domain.Tag{}, time.Now(), time.Now())
	assert.NotEqual(t, nil, err)

	_, err = repo.GetByID("invalid")
	assert.NotEqual(t, nil, err)

	_, err = repo.GetByTitle("invalid")
	assert.NotEqual(t, nil, err)

	_, err = repo.GetByTag("invalid")
	assert.NotEqual(t, nil, err)

	_, err = repo.Delete("invalid")
	assert.NotEqual(t, nil, err)
}
