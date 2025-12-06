package restapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"team-maker-api/domain_teammakerappservice/usecase/runupdateplayer"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/infrastructure/util"
	"team-maker-api/shared/model/enum"
	"team-maker-api/shared/model/payload"
)

// runUpdatePlayerHandler ...
func (r *Controller) runUpdatePlayerHandler(inputPort runupdateplayer.Inport) gin.HandlerFunc {

	type request struct {
		Name       string              `json:"name"`
		PlayerRank enum.PlayerRankEnum `json:"player_rank"`
		PlayerCode string              `json:"player_code"`
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req runupdateplayer.InportRequest
		req.PlayerID = c.Param("player_id")
		req.Name = jsonReq.Name
		req.PlayerRank = jsonReq.PlayerRank
		req.PlayerCode = jsonReq.PlayerCode
		req.TimeNow = time.Now()

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
