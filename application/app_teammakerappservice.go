package application

import (
	"fmt"
	"team-maker-api/domain_teammakerappservice/controller/restapi"
	"team-maker-api/domain_teammakerappservice/gateway/prod"
	"team-maker-api/domain_teammakerappservice/usecase/getallcriteria"
	"team-maker-api/domain_teammakerappservice/usecase/getallplayer"
	"team-maker-api/domain_teammakerappservice/usecase/getallplayerlist"
	"team-maker-api/domain_teammakerappservice/usecase/getalluser"
	"team-maker-api/domain_teammakerappservice/usecase/getonecriteria"
	"team-maker-api/domain_teammakerappservice/usecase/getoneplayer"
	"team-maker-api/domain_teammakerappservice/usecase/getoneuser"
	"team-maker-api/domain_teammakerappservice/usecase/runcreatecriteria"
	"team-maker-api/domain_teammakerappservice/usecase/runcreateplayer"
	"team-maker-api/domain_teammakerappservice/usecase/runcreateteam"
	"team-maker-api/domain_teammakerappservice/usecase/runcreateuser"
	"team-maker-api/domain_teammakerappservice/usecase/rundeletecriteria"
	"team-maker-api/domain_teammakerappservice/usecase/rundeleteplayer"
	"team-maker-api/domain_teammakerappservice/usecase/runloginuser"
	"team-maker-api/domain_teammakerappservice/usecase/runsaveteam"
	"team-maker-api/domain_teammakerappservice/usecase/runupdateplayer"
	"team-maker-api/domain_teammakerappservice/usecase/runupdateuser"
	"team-maker-api/shared/driver"
	"team-maker-api/shared/infrastructure/config"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/infrastructure/server"
	"team-maker-api/shared/infrastructure/util"
)

type appteammaker struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
	// Messaging   messaging.Subscriber
	// consumer    driver.Controller
}

func (c appteammaker) RunApplication() {
	// c.consumer.RegisterRouter()
	c.controller.RegisterRouter()
	// go c.Messaging.Run()
	c.httpHandler.RunApplication()
}

func NewAppTeamMaker() func() driver.RegistryContract {
	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData("appteammakerapi", appID)

		// pub := messaging.NewPublisher(cfg)

		// log := logger.NewSimpleJSONLogger(appData, pub)
		log := logger.NewSimpleJSONLogger(appData, nil)

		appAddress := fmt.Sprintf(":%d", cfg.ApplicationServer.AppTeamMakerService.Port)
		httpHandler := server.NewGinHTTPHandler(log, appAddress, appData)

		// datasource := prod.NewGateway(log, appData, cfg, pub)
		datasource := prod.NewGateway(log, appData, cfg, nil)

		// msg := messaging.NewSubscriber(appData.AppName, cfg)

		return &appteammaker{
			httpHandler: &httpHandler,
			controller: &restapi.Controller{
				Log:                     log,
				Config:                  cfg,
				Router:                  httpHandler.Router,
				GetAllPlayerInport:      getallplayer.NewUsecase(datasource),
				GetAllUserInport:        getalluser.NewUsecase(datasource),
				GetOnePlayerInport:      getoneplayer.NewUsecase(datasource),
				GetOneUserInport:        getoneuser.NewUsecase(datasource),
				RunCreatePlayerInport:   runcreateplayer.NewUsecase(datasource),
				RunCreateTeamInport:     runcreateteam.NewUsecase(datasource),
				RunCreateUserInport:     runcreateuser.NewUsecase(datasource),
				RunLoginUserInport:      runloginuser.NewUsecase(datasource),
				RunSaveTeamInport:       runsaveteam.NewUsecase(datasource),
				RunUpdatePlayerInport:   runupdateplayer.NewUsecase(datasource),
				RunUpdateUserInport:     runupdateuser.NewUsecase(datasource),
				GetAllPlayerListInport:  getallplayerlist.NewUsecase(datasource),
				RunCreateCriteriaInport: runcreatecriteria.NewUsecase(datasource),
				RunDeletePlayerInport:   rundeleteplayer.NewUsecase(datasource),
				RunDeleteCriteriaInport: rundeletecriteria.NewUsecase(datasource),
				GetAllCriteriaInport:    getallcriteria.NewUsecase(datasource),
				GetOneCriteriaInport:    getonecriteria.NewUsecase(datasource),
			},
		}
	}
}
