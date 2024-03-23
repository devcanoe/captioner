package workspace

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workspace struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id"`
	Name      string               `json:"name" bson:"name"`
	OwnerID   primitive.ObjectID   `json:"owner_id" bson:"owner_id"`
	MemberID  []primitive.ObjectID `json:"members_id" bson:"members_id"`
	CreatedAt time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time            `json:"updated_at" bson:"updated_at"`
}

type CreateWorkspace struct {
	Name    string `json:"name" validate:"required"`
	OwnerID string `json:"owner_id" validate:"required"`
}

type UpdateWorkspace struct {
	Name      string               `json:"name" validate:"required"`
	MemberID  []primitive.ObjectID `json:"member_id" validate:"dive, mongodb"`
	UpdatedAt time.Time
}
