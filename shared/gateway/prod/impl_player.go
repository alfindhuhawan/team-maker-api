package prod

import (
	"context"
	"fmt"
	"log"
	"team-maker-api/shared/infrastructure/database"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/model/entity"
	"team-maker-api/shared/model/enum"
	"team-maker-api/shared/model/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type PlayerImpl struct {
	*database.MongoWithTransaction
	Log    logger.Logger
	DbName string
}

func (r *PlayerImpl) SavePlayer(ctx context.Context, obj *entity.Player) error {
	coll := r.MongoClient.Database(r.DbName).Collection(CollectionPlayer)

	_, err := coll.InsertOne(ctx, obj)
	if err != nil {
		return err
	}

	return nil
}

func (r *PlayerImpl) UpdatePlayer(ctx context.Context, userID primitive.ObjectID, obj *repository.UpdatePlayerRequest) error {
	coll := r.MongoClient.Database(r.DbName).Collection(CollectionPlayer)

	criteria := bson.M{}
	criteria["_id"] = userID

	updated := bson.M{}
	updated["$set"] = obj

	_, err := coll.UpdateOne(ctx, criteria, updated)
	if err != nil {
		return err
	}

	return nil
}

func (r *PlayerImpl) FindOnePlayer(ctx context.Context, filterBy enum.FilterByEnum, filter interface{}) (*entity.Player, error) {
	var obj entity.Player

	coll := r.MongoClient.Database(r.DbName).Collection(CollectionPlayer)

	criteria := bson.M{}

	if filterBy == enum.IDFilterByEnum {
		criteria["_id"] = filter
	} else if filterBy == enum.NameFilterByEnum {
		criteria["name"] = filter
	}

	err := coll.FindOne(ctx, criteria).Decode(&obj)
	if err != nil {
		r.Log.Error(ctx, err.Error())

		if err.Error() == "mongo: no documents in result" {
			return nil, fmt.Errorf("player not found")
		}

		return nil, err
	}

	return &obj, nil
}

func (r *PlayerImpl) FindAllPlayer(ctx context.Context, req repository.FindAllPlayerRequest) ([]*entity.Player, int64, error) {
	var objs []*entity.Player

	coll := r.MongoClient.Database(r.DbName).Collection(CollectionPlayer)

	criteria := bson.M{}

	/* filter */

	// name
	if req.Name != "" {
		criteria["name"] = primitive.Regex{Pattern: string(req.Name), Options: "i"}
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

func (r *PlayerImpl) FindAllPlayerList(ctx context.Context, req *repository.FindAllPlayerListRequest) ([]*entity.Player, error) {
	var objs []*entity.Player

	coll := r.MongoClient.Database(r.DbName).Collection(CollectionPlayer)

	criteria := bson.M{}

	/* filter */

	// playerIDs
	if req != nil {
		// Check username first
		var objectIDs []primitive.ObjectID
		// Konversi array string ke array ObjectID
		for _, id := range req.PlayerIDs {
			objectID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				// Handle error jika string bukan ObjectID valid
				log.Fatal(err)
			}
			objectIDs = append(objectIDs, objectID)
		}

		criteria["_id"] = bson.M{"$in": objectIDs}
	}

	sort := bson.M{"updated_at": -1}

	findOpts := options.FindOptions{
		Sort: sort,
	}

	cursor, err := coll.Find(ctx, criteria, &findOpts)
	if err != nil {
		r.Log.Error(ctx, err.Error())
		return nil, err
	}

	if err := cursor.All(ctx, &objs); err != nil {
		r.Log.Error(ctx, err.Error())
		return nil, err
	}

	return objs, nil
}
