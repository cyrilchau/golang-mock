package delivery

import (
	"context"
	"myapp/pkg/middleware"
	"myapp/pkg/utils/response"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *handlers) Detail(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
	defer cancel()

	userID := c.Get("identity").(*middleware.CustomClaims).UserID

	data, code, err := h.uc.Detail(ctx, userID)
	if err != nil {
		return c.JSON(code, response.NewResponseError(code, response.MsgFailed, err.Error()))
	}

	return c.JSON(code, response.NewResponse(code, response.MsgSuccess, data))
}
