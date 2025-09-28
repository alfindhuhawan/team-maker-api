package restapi

import (
	"team-maker-api/domain_teammakerappservice/usecase/getallplayer"
	"team-maker-api/domain_teammakerappservice/usecase/getallplayerlist"
	"team-maker-api/domain_teammakerappservice/usecase/getalluser"
	"team-maker-api/domain_teammakerappservice/usecase/getoneplayer"
	"team-maker-api/domain_teammakerappservice/usecase/getoneuser"
	"team-maker-api/domain_teammakerappservice/usecase/runcreateplayer"
	"team-maker-api/domain_teammakerappservice/usecase/runcreateteam"
	"team-maker-api/domain_teammakerappservice/usecase/runcreateuser"
	"team-maker-api/domain_teammakerappservice/usecase/runloginuser"
	"team-maker-api/domain_teammakerappservice/usecase/runsaveteam"
	"team-maker-api/domain_teammakerappservice/usecase/runupdateplayer"
	"team-maker-api/domain_teammakerappservice/usecase/runupdateuser"

	"github.com/gin-gonic/gin"

	"team-maker-api/shared/helper"
	"team-maker-api/shared/infrastructure/config"
	"team-maker-api/shared/infrastructure/logger"
)

type Controller struct {
	Router gin.IRouter
	Config *config.Config
	Log    logger.Logger
	Helper helper.HTTPHelper

	GetAllPlayerInport     getallplayer.Inport
	GetAllUserInport       getalluser.Inport
	GetOnePlayerInport     getoneplayer.Inport
	GetOneUserInport       getoneuser.Inport
	RunCreatePlayerInport  runcreateplayer.Inport
	RunCreateTeamInport    runcreateteam.Inport
	RunCreateUserInport    runcreateuser.Inport
	RunLoginUserInport     runloginuser.Inport
	RunSaveTeamInport      runsaveteam.Inport
	RunUpdatePlayerInport  runupdateplayer.Inport
	RunUpdateUserInport    runupdateuser.Inport
	GetAllPlayerListInport getallplayerlist.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	// privateAdminOnly := r.Router.Group("/api/v1", r.authenticatedAdmin())
	public := r.Router.Group("/api/v1")
	// private := r.Router.Group("/api/v1", r.authenticated())

	// private.GET("/profile", r.authorized(), r.getUserProfileHandler(r.GetUserProfileInport))
	// private.GET("/admin-hub", r.authorized(), r.getListAdminHubHandler(r.GetListAdminHubInport))
	// private.GET("/sla", r.authorized(), r.getSLAHandler(r.GetSLAInport))
	// private.GET("/sla-detail", r.authorized(), r.getSLADetailHandler(r.GetSLADetailInport))
	// private.GET("/sla-detail-list", r.authorized(), r.getSLADetailListHandler(r.GetSLADetailListInport))
	// private.POST("/export-csv", r.authorized(), r.runExportCsvHandler(r.RunExportCsvInport))
	// private.GET("/download-list", r.authorized(), r.getDownloadListHandler(r.GetDownloadListInport))
	// private.GET("/download-file/:name", r.authorized(), r.getDownloadFileHandler(r.GetDownloadFileInport))

	// private.GET("/getallhero", r.authorized(), r.getAllHeroHandler(r.GetAllHeroInport))
	// private.GET("/getalllocker", r.authorized(), r.getAllLockerHandler(r.GetAllLockerInport))
	// private.GET("/getexportcsvperformance", r.authorized(), r.getExportCSVPerformanceHandler(r.GetExportCSVPerformanceInport))
	// private.GET("/getexportcsvsla", r.authorized(), r.getExportCSVSLAHandler(r.GetExportCSVSLAInport))
	// private.GET("/getheroperformance", r.authorized(), r.getHeroPerformanceHandler(r.GetHeroPerformanceInport))

	// public.POST("/login", r.authorized(), r.runLoginHandler(r.RunLoginInport))

	// private.POST("/runtestingexportcsv", r.authorized(), r.runTestingExportCsvHandler(r.RunTestingExportCsvInport))

	// User
	public.GET("/user", r.authorized(), r.getAllUserHandler(r.GetAllUserInport))
	public.GET("/user/:user_id", r.authorized(), r.getOneUserHandler(r.GetOneUserInport))
	public.PUT("/user/:user_id", r.authorized(), r.runUpdateUserHandler(r.RunUpdateUserInport))
	public.POST("/user", r.authorized(), r.runCreateUserHandler(r.RunCreateUserInport))

	// Player
	public.GET("/player", r.authorized(), r.getAllPlayerHandler(r.GetAllPlayerInport))
	public.GET("/player/:player_id", r.authorized(), r.getOnePlayerHandler(r.GetOnePlayerInport))
	public.POST("/player", r.authorized(), r.runCreatePlayerHandler(r.RunCreatePlayerInport))
	public.PUT("/player/:player_id", r.authorized(), r.runUpdatePlayerHandler(r.RunUpdatePlayerInport))

	// Team
	public.POST("/runcreateteam", r.authorized(), r.runCreateTeamHandler(r.RunCreateTeamInport))
	public.POST("/runsaveteam", r.authorized(), r.runSaveTeamHandler(r.RunSaveTeamInport))

	// login
	public.POST("/login", r.authorized(), r.runLoginUserHandler(r.RunLoginUserInport))
	public.GET("/playerlist", r.authorized(), r.getAllPlayerListHandler(r.GetAllPlayerListInport))
}
