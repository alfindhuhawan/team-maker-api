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

func (r *CriteriaImpl) FindOneCriteria(ctx context.Context, id string) (*entity.Criteria, error) {
	var obj entity.Criteria

	coll := r.MongoClient.Database(r.DbName).Collection(CollectionCriteria)

	criteria := bson.M{}

	criteria["_id"] = bson.ObjectId(id)

	err := coll.FindOne(ctx, criteria).Decode(&obj)
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

// func (r *PlayerImpl) FindAllPlayerList(ctx context.Context, req *repository.FindAllPlayerListRequest) ([]*entity.Player, error) {
// 	var objs []*entity.Player

// 	coll := r.MongoClient.Database(r.DbName).Collection(CollectionPlayer)

// 	criteria := bson.M{}

// 	/* filter */

// 	// playerIDs
// 	if req != nil {
// 		// Check username first
// 		var objectIDs []primitive.ObjectID
// 		// Konversi array string ke array ObjectID
// 		for _, id := range req.PlayerIDs {
// 			objectID, err := primitive.ObjectIDFromHex(id)
// 			if err != nil {
// 				// Handle error jika string bukan ObjectID valid
// 				log.Fatal(err)
// 			}
// 			objectIDs = append(objectIDs, objectID)
// 		}

// 		criteria["_id"] = bson.M{"$in": objectIDs}
// 	}

// 	sort := bson.M{"updated_at": -1}

// 	findOpts := options.FindOptions{
// 		Sort: sort,
// 	}

// 	cursor, err := coll.Find(ctx, criteria, &findOpts)
// 	if err != nil {
// 		r.Log.Error(ctx, err.Error())
// 		return nil, err
// 	}

// 	if err := cursor.All(ctx, &objs); err != nil {
// 		r.Log.Error(ctx, err.Error())
// 		return nil, err
// 	}

// 	return objs, nil
// }
