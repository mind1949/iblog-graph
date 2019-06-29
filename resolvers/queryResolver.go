package resolvers

import (
	"context"
	"github.com/mind1949/iblog/models"
)

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
