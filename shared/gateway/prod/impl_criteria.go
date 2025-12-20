package prod

import (
	"context"
	"fmt"
	"team-maker-api/shared/infrastructure/database"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/model/entity"
	"team-maker-api/shared/model/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type CriteriaImpl struct {
	*database.MongoWithTransaction
	Log    logger.Logger
	DbName string
}

func (r *CriteriaImpl) SaveCriteria(ctx context.Context, obj *entity.Criteria) error {
	coll := r.MongoClient.Database(r.DbName).Collection(CollectionCriteria)

	_, err := coll.InsertOne(ctx, obj)
	if err != nil {
		return err
	}

	return nil
}

func (r *CriteriaImpl) FindOneCriteria(ctx context.Context, criteriaID string) (*entity.Criteria, error) {
	var obj entity.Criteria

	coll := r.MongoClient.Database(r.DbName).Collection(CollectionCriteria)

	criteria := bson.M{}

	oid, err := primitive.ObjectIDFromHex(criteriaID)
	if err != nil {
		return nil, fmt.Errorf("invalid player id: %w", err)
	}

	criteria["_id"] = oid

	err = coll.FindOne(ctx, criteria).Decode(&obj)
	if err != nil {
		r.Log.Error(ctx, err.Error())

		if err.Error() == "mongo: no documents in result" {
			return nil, fmt.Errorf("criteria not found")
		}

		return nil, err
	}

	return &obj, nil
}

func (r *CriteriaImpl) FindAllCriteria(ctx context.Context, req repository.FindAllCriteriaRequest) ([]entity.Criteria, int64, error) {
	var objs []entity.Criteria

	coll := r.MongoClient.Database(r.DbName).Collection(CollectionCriteria)

	criteria := bson.M{}

	/* filter */

	// title
	if req.Title != "" {
		criteria["title"] = primitive.Regex{Pattern: string(req.Title), Options: "i"}
	}

	skip := req.Size * (req.Page - 1)
	limit := req.Size

	// countOpts := options.CountOptions{
	// 	Limit: &limit,
	// 	Skip:  &skip,
	// }

	count, err := coll.CountDocuments(ctx, criteria)
	if err != nil {
		r.Log.Error(ctx, err.Error())
		return nil, 0, err
	}

	sort := bson.M{"updated_at": -1}

	findOpts := options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  sort,
	}

	cursor, err := coll.Find(ctx, criteria, &findOpts)
	if err != nil {
		r.Log.Error(ctx, err.Error())
		return nil, 0, err
	}

	if err := cursor.All(ctx, &objs); err != nil {
		r.Log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return objs, count, nil
}

func (r *CriteriaImpl) DeleteCriteria(ctx context.Context, criteriaID string) error {
	coll := r.MongoClient.Database(r.DbName).Collection(CollectionCriteria)

	oid, err := primitive.ObjectIDFromHex(criteriaID)
	if err != nil {
		return fmt.Errorf("invalid criteria id: %w", err)
	}

	criteria := bson.M{"_id": oid}

	_, err = coll.DeleteOne(ctx, criteria)
	if err != nil {
		return err
	}

	return nil
}

func (r *CriteriaImpl) UpdateCriteria(ctx context.Context, userID string, obj *repository.UpdatePlayerRequest) error {
	coll := r.MongoClient.Database(r.DbName).Collection(CollectionPlayer)

	// convert user id to string
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid player id: %w", err)
	}

	criteria := bson.M{}
	criteria["_id"] = oid

	updated := bson.M{}
	updated["$set"] = obj

	_, err = coll.UpdateOne(ctx, criteria, updated)
	if err != nil {
		return err
	}

	return nil
}
