package tag_test

import (
	"isso0424/racion-api/mock/repository/tag"
	"isso0424/racion-api/types/domain"
	"isso0424/racion-api/types/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() tag.MockTagDB {
	return tag.MockTagDB{
		Data: []domain.Tag{
			{
				Title: "hoge",
				Description: "fuga",
				Color: "#ffffff",
			},
		},
	}
}

func TestCreate(t *testing.T) {
	repo := setup()

	tag, err := repo.Create("foo", "bar", "fuga")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "foo", tag.Title)
	assert.Equal(t, "bar", tag.Description)
	assert.Equal(t, "fuga", tag.Color)

	assert.Equal(t, 2, len(repo.Data))
	assert.Equal(t, "foo", repo.Data[1].Title)
	assert.Equal(t, "bar", repo.Data[1].Description)
	assert.Equal(t, "fuga", repo.Data[1].Color)
}

func TestEdit(t *testing.T) {
	repo := setup()

	tag, err := repo.Edit(repo.Data[0].ID, "fuga", "foo", "bar")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "fuga", tag.Title)
	assert.Equal(t, "foo", tag.Description)
	assert.Equal(t, "bar", tag.Color)

	assert.Equal(t, 1, len(repo.Data))
	assert.Equal(t, "fuga", repo.Data[0].Title)
	assert.Equal(t, "foo", repo.Data[0].Description)
	assert.Equal(t, "bar", repo.Data[0].Color)
}

func TestGetAll(t *testing.T) {
	repo := setup()

	tags, err := repo.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(tags))
	assert.Equal(t, "hoge", tags[0].Title)
	assert.Equal(t, "fuga", tags[0].Description)
	assert.Equal(t, "#ffffff", tags[0].Color)

	assert.Equal(t, 1, len(repo.Data))
	assert.Equal(t, "hoge", repo.Data[0].Title)
	assert.Equal(t, "fuga", repo.Data[0].Description)
	assert.Equal(t, "#ffffff", repo.Data[0].Color)
}

func TestGetByName(t *testing.T) {
	repo := setup()

	tags, err := repo.GetByTitle("hoge")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "hoge", tags[0].Title)
	assert.Equal(t, "fuga", tags[0].Description)
	assert.Equal(t, "#ffffff", tags[0].Color)
}

func TestGetByID(t *testing.T) {
	repo := setup()

	tag, err := repo.GetByID(repo.Data[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "hoge", tag.Title)
	assert.Equal(t, "fuga", tag.Description)
	assert.Equal(t, "#ffffff", tag.Color)
}

func TestFail(t *testing.T) {
	repo := setup()

	_, err := repo.Edit("invalid", "fuga", "invalid", "invalid")
	assert.NotEqual(t, nil, err)

	_, err = repo.GetByTitle("fuga")
	assert.NotEqual(t, nil, err)

	_, err = repo.GetByID("invalid")
	assert.NotEqual(t, nil, err)
}

func TestImplInterface(t *testing.T) {
	repo := setup()
	f := func (r repository.TagRepository) {}

	f(&repo)
}
