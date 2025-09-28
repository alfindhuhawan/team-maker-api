package prod

import (
	"context"
	"fmt"
	"team-maker-api/shared/infrastructure/database"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/model/entity"
	"team-maker-api/shared/model/enum"
	"team-maker-api/shared/model/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type UserImpl struct {
	*database.MongoWithTransaction
	Log    logger.Logger
	DbName string
}

func (r *UserImpl) SaveUser(ctx context.Context, obj *entity.User) error {
	coll := r.MongoClient.Database(r.DbName).Collection(CollectionUser)

	_, err := coll.InsertOne(ctx, obj)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserImpl) UpdateUser(ctx context.Context, userID primitive.ObjectID, obj *repository.UpdateRequest) error {
	coll := r.MongoClient.Database(r.DbName).Collection(CollectionUser)

	criteria := bson.M{}
	criteria["_id"] = userID

	updated := bson.M{}
	updated["$set"] = obj

	_, err := coll.UpdateOne(ctx, criteria, updated)
	if err != nil {
		return err
	}

	// filter := bson.D{{"_id", userID}}
	// update := bson.D{{"$set", obj}}
	// opts := options.Update().SetUpsert(true)

	// result, err := coll.UpdateOne(ctx, filter, update, opts)
	// if err != nil {
	// 	fmt.Println("masuk sini")
	// 	fmt.Println(err.Error())
	// 	return err
	// }

	// fmt.Println("result >>")
	// fmt.Println(result)

	return nil
}

func (r *UserImpl) FindOneUser(ctx context.Context, filterBy enum.FilterByEnum, filter interface{}) (*entity.User, error) {
	var obj entity.User

	coll := r.MongoClient.Database(r.DbName).Collection(CollectionUser)

	criteria := bson.M{}

	if filterBy == enum.IDFilterByEnum {
		criteria["_id"] = filter
	} else if filterBy == enum.UsernameFilterByEnum {
		criteria["username"] = filter
	}

	err := coll.FindOne(ctx, criteria).Decode(&obj)
	if err != nil {
		r.Log.Error(ctx, err.Error())

		if err.Error() == "mongo: no documents in result" {
			return nil, fmt.Errorf("user not found")
		}

		return nil, err
	}

	return &obj, nil
}

func (r *UserImpl) FindAllUser(ctx context.Context, req repository.FindAllUserRequest) ([]*entity.User, int64, error) {
	var objs []*entity.User

	coll := r.MongoClient.Database(r.DbName).Collection(CollectionUser)

	criteria := bson.M{}

	/* filter */

	// name
	if req.Name != "" {
		criteria["name"] = primitive.Regex{Pattern: string(req.Name), Options: "i"}
	}

	// username
	if req.Username != "" {
		criteria["username"] = primitive.Regex{Pattern: string(req.Username), Options: "i"}
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

// func (r *BuyerImpl) FindOneBuyerByID(ctx context.Context, buyerId vo.BuyerID) (*entity.Buyer, error) {
// 	var obj entity.Buyer

// 	coll := r.MongoClient.Database(r.DbName).Collection(CollectionBuyer)

// 	err := coll.FindOne(ctx, bson.M{"_id": buyerId}).Decode(&obj)
// 	if err != nil {
// 		r.Log.Error(ctx, err.Error())

// 		if err.Error() == "mongo: no documents in result" {
// 			return nil, fmt.Errorf("buyer not found")
// 		}

// 		return nil, err
// 	}

// 	return &obj, nil
// }

// func (r *BuyerImpl) FindOneBuyerByPhone(ctx context.Context, phone vo.PhoneNumber) (*entity.Buyer, error) {
// 	var obj entity.Buyer

// 	coll := r.MongoClient.Database(r.DbName).Collection(CollectionBuyer)

// 	err := coll.FindOne(ctx, bson.M{"phone_number": phone}).Decode(&obj)
// 	if err != nil {
// 		r.Log.Error(ctx, err.Error())

// 		if err.Error() == "mongo: no documents in result" {
// 			return nil, fmt.Errorf("buyer not found")
// 		}

// 		return nil, err
// 	}

// 	return &obj, nil
// }

// func (r *BuyerImpl) FindOneBuyerByMemberCode(ctx context.Context, memberCode string) (*entity.Buyer, error) {
// 	var obj entity.Buyer

// 	coll := r.MongoClient.Database(r.DbName).Collection(CollectionBuyer)

// 	err := coll.FindOne(ctx, bson.M{"member_code": memberCode}).Decode(&obj)
// 	if err != nil {
// 		r.Log.Error(ctx, err.Error())

// 		if err.Error() == "mongo: no documents in result" {
// 			return nil, fmt.Errorf("buyer not found")
// 		}

// 		return nil, err
// 	}

// 	return &obj, nil
// }

// func (r *BuyerImpl) FindOneBuyerByApiToken(ctx context.Context, apiToken string) (*entity.Buyer, error) {
// 	var obj entity.Buyer

// 	isGetFromDB := false

// 	exist, errRedis := r.CacheClient.Exist(ctx, apiToken)
// 	if errRedis != nil {
// 		// IF REDIS ERROR THEN GET DATA FROM DB
// 		isGetFromDB = true
// 	}

// 	if exist && !isGetFromDB {
// 		value, err := r.CacheClient.Get(ctx, apiToken)
// 		if err != nil {
// 			return nil, err
// 		}

// 		err = json.Unmarshal([]byte(value), &obj)
// 		if err != nil {
// 			return nil, err
// 		}
// 	} else {
// 		isGetFromDB = true
// 	}

// 	// CHECK IF GET DATA BUYER FROM DB
// 	if isGetFromDB {
// 		coll := r.MongoClient.Database(r.DbName).Collection(CollectionBuyer)

// 		err := coll.FindOne(ctx, bson.M{"api_token": apiToken}).Decode(&obj)
// 		if err != nil {
// 			r.Log.Error(ctx, err.Error())

// 			if err.Error() == "mongo: no documents in result" {
// 				return nil, fmt.Errorf("buyer not found")
// 			}

// 			return nil, err
// 		}

// 		byteJson, err := json.Marshal(obj)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// REMOVE ERR BECAUSE DONT RETURN ERROR IF REDDIS NOT CONNECTED
// 		_ = r.CacheClient.Set(ctx, apiToken, byteJson, 1*time.Hour) // TODO : MASUKIN CONFIG 1 JAMNYA
// 	}

// 	return &obj, nil
// }

// func (r *BuyerImpl) FindAllBuyerCms(ctx context.Context, req repository.FindAllBuyerRequest) ([]*entity.Buyer, int64, error) {
// 	var objs []*entity.Buyer

// 	coll := r.MongoClient.Database(r.DbName).Collection(CollectionBuyer)

// 	arrayFilter := make([]bson.M, 0)

// 	/* filter */
// 	if req.Keyword != "" {
// 		arrayFilter = append(arrayFilter, bson.M{"_id": req.Keyword})
// 		arrayFilter = append(arrayFilter, bson.M{"name": primitive.Regex{Pattern: string(req.Keyword), Options: "i"}})
// 		arrayFilter = append(arrayFilter, bson.M{"username": primitive.Regex{Pattern: string(req.Keyword), Options: "i"}})
// 		arrayFilter = append(arrayFilter, bson.M{"phone_number": primitive.Regex{Pattern: string(req.Keyword), Options: "i"}})
// 	}

// 	arrayDate := make([]bson.M, 0)

// 	// filter by join date from
// 	if req.DateStart != "" {
// 		dateStart, err := time.Parse("2006-01-02", req.DateStart)
// 		if err != nil {
// 			return nil, 0, err
// 		}
// 		arrayDate = append(arrayDate, bson.M{"created_date": bson.M{"$gte": dateStart}})
// 	}

// 	// filter by join date to
// 	if req.DateEnd != "" {
// 		dateEnd, err := time.Parse("2006-01-02", req.DateEnd)
// 		if err != nil {
// 			return nil, 0, err
// 		}
// 		arrayDate = append(arrayDate, bson.M{"created_date": bson.M{"$lte": dateEnd.AddDate(0, 0, 1)}})
// 	}

// 	var allCriteria []bson.M

// 	var criteriaFilter bson.M
// 	if len(arrayFilter) > 0 {
// 		criteriaFilter = bson.M{"$or": arrayFilter}
// 		allCriteria = append(allCriteria, criteriaFilter)
// 	}

// 	var criteriaDate bson.M
// 	if len(arrayDate) > 0 {
// 		criteriaDate = bson.M{"$and": arrayDate}
// 		allCriteria = append(allCriteria, criteriaDate)
// 	}

// 	criteria := bson.M{}
// 	if len(allCriteria) > 0 {
// 		criteria = bson.M{"$and": allCriteria}
// 	}

// 	skip := req.Size * (req.Page - 1)
// 	limit := req.Size

// 	count, err := coll.CountDocuments(ctx, criteria)
// 	if err != nil {
// 		r.Log.Error(ctx, err.Error())
// 		return nil, 0, err
// 	}

// 	sort := bson.M{"updated_date": -1}
// 	if req.SortBy != "" {
// 		sort = bson.M{req.SortBy: req.SortType}
// 	}

// 	findOpts := options.FindOptions{
// 		Limit: &limit,
// 		Skip:  &skip,
// 		Sort:  sort,
// 	}

// 	cursor, err := coll.Find(ctx, criteria, &findOpts)
// 	if err != nil {
// 		r.Log.Error(ctx, err.Error())
// 		return nil, 0, err
// 	}

// 	if err := cursor.All(ctx, &objs); err != nil {
// 		r.Log.Error(ctx, err.Error())
// 		return nil, 0, err
// 	}

// 	return objs, count, nil
// }

// func (r *BuyerImpl) FindAllBuyerByIDs(ctx context.Context, buyerIDs []vo.BuyerID) ([]*entity.Buyer, error) {
// 	var objs []*entity.Buyer

// 	coll := r.MongoClient.Database(r.DbName).Collection(CollectionBuyer)

// 	var bsonObjectIDs bson.A
// 	for _, val := range buyerIDs {
// 		bsonObjectIDs = append(bsonObjectIDs, val)
// 	}

// 	cursor, err := coll.Find(ctx, bson.M{"_id": bson.M{"$in": bsonObjectIDs}})
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

// func (r *BuyerImpl) FindAsMapBuyerByIDs(ctx context.Context, buyerIDs []vo.BuyerID) (map[vo.BuyerID]*entity.Buyer, error) {
// 	objs := make(map[vo.BuyerID]*entity.Buyer, 0)

// 	coll := r.MongoClient.Database(r.DbName).Collection(CollectionBuyer)

// 	var bsonObjectIDs bson.A
// 	for _, val := range buyerIDs {
// 		bsonObjectIDs = append(bsonObjectIDs, val)
// 	}

// 	cursor, err := coll.Find(ctx, bson.M{"_id": bson.M{"$in": bsonObjectIDs}})
// 	if err != nil {
// 		r.Log.Error(ctx, err.Error())
// 		return nil, err
// 	}

// 	defer func(cursor *mongo.Cursor, ctx context.Context) {
// 		err := cursor.Close(ctx)
// 		if err != nil {

// 		}
// 	}(cursor, ctx)

// 	for cursor.Next(ctx) {
// 		var result entity.Buyer
// 		if err := cursor.Decode(&result); err != nil {
// 			return nil, err
// 		}

// 		objs[result.ID] = &result
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	return objs, nil
// }

// func (r *BuyerImpl) FindAllDuplicateBuyerCms(ctx context.Context) ([]vo.PhoneNumber, error) {
// 	var phones []vo.PhoneNumber
// 	// var objs []*entity.Buyer
// 	coll := r.MongoClient.Database(r.DbName).Collection(CollectionBuyer)

// 	match := bson.M{
// 		"_id":   bson.M{"$ne": ""},
// 		"count": bson.M{"$gt": 1},
// 	}

// 	group := bson.M{
// 		"_id":   "$phone_number",
// 		"count": bson.M{"$sum": 1},
// 	}

// 	pipeline := []bson.M{
// 		{"$group": group},
// 		{"$match": match},
// 		{"$sort": bson.M{
// 			"count": -1,
// 		}},
// 		{"$project": bson.M{"phone_number": "$_id", "_id": 1}},
// 	}

// 	cursor, err := coll.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		r.Log.Error(ctx, err.Error())
// 		return []vo.PhoneNumber{}, err
// 	}

// 	// if err := cursor.All(ctx, &objs); err != nil {
// 	// 	r.Log.Error(ctx, err.Error())
// 	// 	return []vo.PhoneNumber{}, err
// 	// }

// 	defer func(cursor *mongo.Cursor, ctx context.Context) {
// 		err := cursor.Close(ctx)
// 		if err != nil {

// 		}
// 	}(cursor, ctx)

// 	for cursor.Next(ctx) {
// 		var result entity.Buyer
// 		if err := cursor.Decode(&result); err != nil {
// 			return nil, err
// 		}

// 		phones = append(phones, result.PhoneNumber)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	return phones, nil
// }

// func (r *BuyerImpl) FindAsMapBuyerByPhoneNumber(ctx context.Context, phoneNumbers []vo.PhoneNumber) (map[vo.BuyerID]vo.PhoneNumber, error) {
// 	objs := make(map[vo.BuyerID]vo.PhoneNumber, 0)

// 	coll := r.MongoClient.Database(r.DbName).Collection(CollectionBuyer)

// 	var bsonObjectIDs bson.A
// 	for _, val := range phoneNumbers {
// 		bsonObjectIDs = append(bsonObjectIDs, val)
// 	}

// 	// SORTING FROM UPDATED_DATE DESC
// 	findOpts := options.FindOptions{
// 		Sort: bson.M{"updated_date": -1},
// 	}

// 	cursor, err := coll.Find(ctx, bson.M{"phone_number": bson.M{"$in": bsonObjectIDs}}, &findOpts)
// 	if err != nil {
// 		r.Log.Error(ctx, err.Error())
// 		return nil, err
// 	}

// 	defer func(cursor *mongo.Cursor, ctx context.Context) {
// 		err := cursor.Close(ctx)
// 		if err != nil {

// 		}
// 	}(cursor, ctx)

// 	for cursor.Next(ctx) {
// 		var result entity.Buyer
// 		if err := cursor.Decode(&result); err != nil {
// 			return nil, err
// 		}

// 		objs[result.ID] = result.PhoneNumber
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	return objs, nil
// }

// func (r *BuyerImpl) DeleteBulkBuyerByIDs(ctx context.Context, buyerIDs []vo.BuyerID) error {
// 	coll := r.MongoClient.Database(r.DbName).Collection(CollectionBuyer)

// 	for _, v := range buyerIDs {
// 		_, err := coll.DeleteOne(ctx, bson.D{{Key: "_id", Value: v}})
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (r *BuyerImpl) FindAsMapBuyerObjByPhoneNumber(ctx context.Context, phoneNumbers []vo.PhoneNumber) (map[vo.BuyerID]entity.Buyer, error) {
// 	objs := make(map[vo.BuyerID]entity.Buyer, 0)

// 	coll := r.MongoClient.Database(r.DbName).Collection(CollectionBuyer)

// 	var bsonObjectIDs bson.A
// 	for _, val := range phoneNumbers {
// 		bsonObjectIDs = append(bsonObjectIDs, val)
// 	}

// 	// SORTING FROM UPDATED_DATE DESC
// 	findOpts := options.FindOptions{
// 		Sort: bson.M{"updated_date": -1},
// 	}

// 	cursor, err := coll.Find(ctx, bson.M{"phone_number": bson.M{"$in": bsonObjectIDs}}, &findOpts)
// 	if err != nil {
// 		r.Log.Error(ctx, err.Error())
// 		return nil, err
// 	}

// 	defer func(cursor *mongo.Cursor, ctx context.Context) {
// 		err := cursor.Close(ctx)
// 		if err != nil {

// 		}
// 	}(cursor, ctx)

// 	for cursor.Next(ctx) {
// 		var result entity.Buyer
// 		if err := cursor.Decode(&result); err != nil {
// 			return nil, err
// 		}

// 		objs[result.ID] = result
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	return objs, nil
// }
