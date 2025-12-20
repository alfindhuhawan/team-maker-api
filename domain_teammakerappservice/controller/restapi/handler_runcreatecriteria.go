package restapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"team-maker-api/domain_teammakerappservice/usecase/runcreatecriteria"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/infrastructure/util"
	"team-maker-api/shared/model/payload"
)

// runCreateCriteriaHandler ...
func (r *Controller) runCreateCriteriaHandler(inputPort runcreatecriteria.Inport) gin.HandlerFunc {

	type request struct {
		Title        string   `json:"title"`
		TitleAliases []string `json:"title_aliases"`
		// Rank        []string `json:"rank"`
		// Rule        []string `json:"rule"`
		Description string `json:"description"`
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

		var req runcreatecriteria.InportRequest
		req.Title = jsonReq.Title
		req.TitleAliases = jsonReq.TitleAliases
		// req.Rank = jsonReq.Rank
		// req.Rule = jsonReq.Rule
		req.Description = jsonReq.Description
		req.Now = time.Now()

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
