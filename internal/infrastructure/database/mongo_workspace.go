package database

import (
	"context"

	tnMongo "github.com/cnc-csku/task-nexus-go-lib/mongo"
	"github.com/cnc-csku/task-nexus-workspace/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// @WireSet("Infrastructure")
func NewWorkspaceMongoClient(ctx context.Context, configs *config.Config) *mongo.Client {
	return tnMongo.NewMongoClient(
		ctx,
		configs.Mongo.URI,
		configs.Mongo.Database,
	)
}
