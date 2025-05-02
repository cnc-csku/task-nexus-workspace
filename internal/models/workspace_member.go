package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type WorkspaceMember struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      bson.ObjectID `bson:"user_id" json:"userId"`
	WorkspaceID bson.ObjectID `bson:"workspace_id" json:"workspaceId"`
	JoinedAt    time.Time     `bson:"joined_at" json:"joinedAt"`
	RemovedAt   *time.Time    `bson:"removed_at,omitempty" json:"removedAt,omitempty"`
}
