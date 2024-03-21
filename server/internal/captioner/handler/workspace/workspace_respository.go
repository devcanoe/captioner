package workspace

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "captioner"
	COLLECTION = "workspaces"
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
	Name    string             `validate:"required"`
	OwnerID primitive.ObjectID `validate:"required, mongodb"`
}

type UpdateWorkspace struct {
	Name      string               `validate:"required"`
	MemberID  []primitive.ObjectID `validate:"dive, mongodb"`
	UpdatedAt time.Time
}

type IWorkspace interface {
	GetAllWorkpaces() (*[]Workspace, error)
	GetOneWorkpace(id primitive.ObjectID) (*Workspace, error)
	CreateWorkspace(worksapce CreateWorkspace) (*Workspace, error)
	UpdateWorkspace(id primitive.ObjectID, workspace UpdateWorkspace) (*Workspace, error)
	DeleteWorkspace(id primitive.ObjectID) error
}

type WorkspaceRepository struct {
	client    *mongo.Collection
	Workspace Workspace
}

func NewWorkspaceRepsository(client *mongo.Client) *WorkspaceRepository {
	return &WorkspaceRepository{
		client: client.Database(DATABASE).Collection(COLLECTION),
	}
}

func (w *WorkspaceRepository) GetAllWorkpaces() (*[]Workspace, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := w.client
	var workspace []Workspace
	defer cancel()

	result, err := m.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer result.Close(ctx)
	for result.Next(ctx) {
		var singleWorkspace Workspace
		result.Decode(&singleWorkspace)

		workspace = append(workspace, singleWorkspace)
	}
	return &workspace, nil
}

func (w *WorkspaceRepository) GetOneWorkpace(id primitive.ObjectID) (*Workspace, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := w.client
	var worksapce Workspace
	defer cancel()

	result := m.FindOne(ctx, bson.M{"_id": id})

	err := result.Decode(&worksapce)
	if err != nil {
		return nil, err
	}
	return &worksapce, nil
}

func (w *WorkspaceRepository) CreateWorkspace(worksapce CreateWorkspace) (*Workspace, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := w.client
	var newWorkspace *Workspace
	defer cancel()

	newWorkspace = &Workspace{
		ID:        primitive.NewObjectID(),
		Name:      worksapce.Name,
		OwnerID:   worksapce.OwnerID,
		MemberID:  []primitive.ObjectID{},
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	_, err := m.InsertOne(ctx, newWorkspace)
	if err != nil {
		return nil, err
	}
	return newWorkspace, nil
}

func (w *WorkspaceRepository) UpdateWorkspace(id primitive.ObjectID, workspace UpdateWorkspace) (*Workspace, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := w.client
	var updatedWorkspace *Workspace
	defer cancel()
	updatedParams := bson.M{
		"name":       workspace.Name,
		"member_id":  workspace.MemberID,
		"updated_at": time.Now().UTC(),
	}
	result := m.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": updatedParams}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	err := result.Decode(&updatedWorkspace)

	if err != nil {
		return nil, err
	}
	return updatedWorkspace, nil
}

func (w *WorkspaceRepository) DeleteWorkspace(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := w.client
	defer cancel()

	result, err := m.DeleteOne(ctx, bson.M{"_id": id})

	if result.DeletedCount < 1 {
		return errors.New("user id not found")
	}
	if err != nil {
		return err
	}
	return nil
}
