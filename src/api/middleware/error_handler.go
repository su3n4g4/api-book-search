// middleware/error_handler.go
package middleware

import (
	"api-book-search/internal/apperrors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			status := mapAppErrorToStatus(appErr.Type)
			c.JSON(status, gin.H{"error": appErr.Message})
		} else {
			// 想定外のエラー
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		c.Abort()
	}
}

func mapAppErrorToStatus(t apperrors.ErrorType) int {
	switch t {
	case apperrors.ValidationError:
		return http.StatusBadRequest
	case apperrors.NotFoundError:
		return http.StatusNotFound
	case apperrors.TimeoutError:
		return http.StatusGatewayTimeout
	case apperrors.ExternalAPIError:
		return http.StatusBadGateway
	default:
		return http.StatusInternalServerError
	}
}
