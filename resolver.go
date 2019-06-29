//go:generate go run github.com/99designs/gqlgen init -v
package iblog

import (
	"context"
	"github.com/mind1949/iblog/models"
	"time"
)

type Resolver struct {}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePost(ctx context.Context, input *NewPost) (*models.Post, error) {
	post := &models.Post{
		Title:   input.Title,
		Content: input.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	post.TagsID = make([]int,len(input.TagsTitle))
	for i, tagTitle := range input.TagsTitle {
		tag := &models.Tag{Title: tagTitle}
		err := tag.FindOrCreateByTitle()
		if err != nil {
			return nil, err
		}
		post.TagsID[i] = tag.ID
	}

	err := post.Save()
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (r *mutationResolver) UpdatePost(ctx context.Context, input *UpdatePost) (*models.Post, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Post(ctx context.Context, id int) (*models.Post, error) {
	post := &models.Post{}
	err := post.Find()
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	return models.Posts()
}
func (r *queryResolver) Tag(ctx context.Context, id int) (*models.Tag, error) {
	tag := &models.Tag{ID: id}
	err := tag.Find()
	if err != nil {
		return nil, err
	}
	return tag, nil
}
func (r *queryResolver) Tags(ctx context.Context) ([]*models.Tag, error) {
	return models.Tags()
}
