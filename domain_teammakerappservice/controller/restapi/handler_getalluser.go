package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"team-maker-api/domain_teammakerappservice/usecase/getalluser"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/infrastructure/util"
	"team-maker-api/shared/model/entity"
	"team-maker-api/shared/model/payload"
)

// getAllUserHandler ...
func (r *Controller) getAllUserHandler(inputPort getalluser.Inport) gin.HandlerFunc {

	type request struct {
	}

	type response struct {
		Count int64          `json:"count"`
		Items []*entity.User `json:"items"`
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

		var req getalluser.InportRequest

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Count = res.Count
		jsonRes.Items = res.Items

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
