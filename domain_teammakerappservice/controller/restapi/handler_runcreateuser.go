package restapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"team-maker-api/domain_teammakerappservice/usecase/runcreateuser"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/infrastructure/util"
	"team-maker-api/shared/model/payload"
)

// runCreateUserHandler ...
func (r *Controller) runCreateUserHandler(inputPort runcreateuser.Inport) gin.HandlerFunc {

	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
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

		var req runcreateuser.InportRequest
		req.Username = jsonReq.Username
		req.Password = jsonReq.Password
		req.Name = jsonReq.Name
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
