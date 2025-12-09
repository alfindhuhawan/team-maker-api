package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"team-maker-api/domain_teammakerappservice/usecase/getonecriteria"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/infrastructure/util"
	"team-maker-api/shared/model/entity"
	"team-maker-api/shared/model/payload"
)

// getOneCriteriaHandler ...
func (r *Controller) getOneCriteriaHandler(inputPort getonecriteria.Inport) gin.HandlerFunc {

	type response struct {
		Item *entity.Criteria `json:"item"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var req getonecriteria.InportRequest
		req.CriteriaID = c.Param("criteria_id")

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
