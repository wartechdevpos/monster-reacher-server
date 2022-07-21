package database

import "context"

type DBDriver interface {
	SelectOne(ctx context.Context, filter interface{}) interface{}
	PushOne(ctx context.Context, data interface{}) (interface{}, error)
	DeleteOne(ctx context.Context, filter interface{}) error
	UpdateOne(ctx context.Context, filter interface{}, data interface{}) error
	UpdateSpecific(ctx context.Context, filter interface{}, key string, value interface{}) error
	IncrementValue(ctx context.Context, filter interface{}, key string, value interface{}) error
	Close() error
}
