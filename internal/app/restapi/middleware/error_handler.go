package middleware

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/danielmesquitta/api-pet-curiosities/internal/app/restapi/dto"
	"github.com/danielmesquitta/api-pet-curiosities/internal/domain/errs"
	"github.com/labstack/echo/v4"
)

var mapErrTypeToStatusCode = map[errs.Type]int{
	errs.ErrTypeForbidden:    http.StatusForbidden,
	errs.ErrTypeUnauthorized: http.StatusUnauthorized,
	errs.ErrTypeValidation:   http.StatusBadRequest,
	errs.ErrTypeUnknown:      http.StatusInternalServerError,
	errs.ErrTypeNotFound:     http.StatusNotFound,
}

func (m *Middleware) ErrorHandler(
	defaultErrorHandler echo.HTTPErrorHandler,
) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		if appErr, ok := err.(*errs.Err); ok {
			statusCode := mapErrTypeToStatusCode[appErr.Type]
			if internalServerError := statusCode >= 500 || statusCode == 0; internalServerError {
				req := c.Request()

				requestData := map[string]any{}
				_ = c.Bind(&requestData)

				slog.Error(
					appErr.Error(),
					slog.String("url", req.URL.Path),
					"body", requestData,
					"query", c.QueryParams(),
					"params", c.ParamValues(),
					slog.String("stacktrace", appErr.StackTrace),
				)

				err = c.JSON(
					statusCode,
					dto.ErrorResponseDTO{Message: "internal server error"},
				)
				if err != nil {
					log.Println(err)
				}
				return
			}

			err = c.JSON(
				statusCode,
				dto.ErrorResponseDTO{Message: appErr.Error()},
			)
			if err != nil {
				log.Println(err)
			}
			return
		}

		defaultErrorHandler(err, c)
	}
}
