package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"team-maker-api/domain_teammakerappservice/usecase/getallcriteria"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/infrastructure/util"
	"team-maker-api/shared/model/entity"
	"team-maker-api/shared/model/payload"
)

// getAllCriteriaHandler ...
func (r *Controller) getAllCriteriaHandler(inputPort getallcriteria.Inport) gin.HandlerFunc {

	type request struct {
		Page  int64  `form:"page,omitempty,default=0"`
		Size  int64  `form:"size,omitempty,default=0"`
		Title string `form:"title,omitempty,default=0"`
	}

	type response struct {
		Count int64             `json:"count"`
		Items []entity.Criteria `json:"items"`
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

		var req getallcriteria.InportRequest
		req.Page = jsonReq.Page
		req.Size = jsonReq.Size
		req.Title = jsonReq.Title

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
