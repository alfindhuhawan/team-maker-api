package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"team-maker-api/domain_teammakerappservice/usecase/getoneplayer"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/infrastructure/util"
	"team-maker-api/shared/model/entity"
	"team-maker-api/shared/model/payload"
)

// getOnePlayerHandler ...
func (r *Controller) getOnePlayerHandler(inputPort getoneplayer.Inport) gin.HandlerFunc {

	type request struct {
	}

	type response struct {
		Item *entity.Player `json:"item"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.Bind(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req getoneplayer.InportRequest
		req.PlayerID = c.Param("player_id")

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Item = res.Item

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
