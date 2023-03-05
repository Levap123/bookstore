package mongo

import (
	"context"
	"fmt"

	"github.com/Levap123/book_service/internal/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(ctx context.Context, cfg *configs.Configs) (*mongo.Client, error) {
	connStr := fmt.Sprintf("mongodb://%s:%s@%s", cfg.Mongo.Username, cfg.Mongo.Password, cfg.Mongo.Addr)
	DB, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}
	return DB, err
}

func ShutDownDB(ctx context.Context, cl *mongo.Client) error {
	return cl.Disconnect(ctx)
}

