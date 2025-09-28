package prod

import (
	"team-maker-api/shared/driver"
	"team-maker-api/shared/gateway/prod"
	"team-maker-api/shared/infrastructure/config"
	"team-maker-api/shared/infrastructure/database"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/infrastructure/messaging"
)

type gateway struct {
	*database.MongoWithTransaction
	*prod.ConfigImpl
	*prod.RedisImpl
	*prod.ValidationImpl
	*prod.MessagePublisherImpl
	*prod.UserImpl
	*prod.PlayerImpl
	*prod.TeamImpl
	*prod.CriteriaImpl
}

func NewGateway(log logger.Logger, appData driver.ApplicationData, cfg *config.Config, pub messaging.Publisher) *gateway {

	// jwtToken, err := token.NewJWTToken(cfg)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// cacheClient := cache.NewRedisCacheDefault(cfg)

	cl := database.NewMongoDefault(cfg)
	mwt := database.NewMongoWithTransaction(cl)

	// masterConn := database.NewMasterConnect(cfg)
	// masterMwt := database.NewMongoWithTransaction(masterConn)

	dbName := cfg.Database.MongoDB.DbName
	// dbMasterName := cfg.Database.MasterDB.DbName

	// prod.PrepareCollection(dbName, mwt)

	return &gateway{
		MongoWithTransaction: mwt,
		ConfigImpl:           &prod.ConfigImpl{Log: log, Cfg: cfg},
		UserImpl:             &prod.UserImpl{MongoWithTransaction: mwt, Log: log, DbName: dbName},
		PlayerImpl:           &prod.PlayerImpl{MongoWithTransaction: mwt, Log: log, DbName: dbName},
		TeamImpl:             &prod.TeamImpl{MongoWithTransaction: mwt, Log: log, DbName: dbName},
		CriteriaImpl:         &prod.CriteriaImpl{MongoWithTransaction: mwt, Log: log, DbName: dbName},
	}
}
