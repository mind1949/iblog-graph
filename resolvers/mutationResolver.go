package resolvers

import (
	"context"
	"github.com/mind1949/iblog"
	"github.com/mind1949/iblog/models"
	"time"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePost(ctx context.Context, input *iblog.NewPost) (*models.Post, error) {
	post := &models.Post{
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	post.TagsID = make([]int, len(input.TagsTitle))
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
func (r *mutationResolver) UpdatePost(ctx context.Context, input *iblog.UpdatePost) (*models.Post, error) {
	panic("not implemented")
}
