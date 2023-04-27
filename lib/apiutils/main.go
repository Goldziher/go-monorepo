package apiutils

import (
	"net/http"

	"github.com/go-chi/render"
)

type ApiError struct {
	Message    string `json:"message"`
	StatusText string `json:"status"`
	StatusCode int    `json:"statusCode"`
}

func (e *ApiError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func BadRequest(message string) render.Renderer {
	return &ApiError{
		StatusCode: 400,
		StatusText: "Bad Request",
		Message:    message,
	}
}

func Unauthorized(message string) render.Renderer {
	return &ApiError{
		StatusCode: 401,
		StatusText: "Unauthorized",
		Message:    message,
	}
}

func UnprocessableContent(message string) render.Renderer {
	return &ApiError{
		StatusCode: 422,
		StatusText: "Unprocessable Content",
		Message:    message,
	}
}

func InternalServerError(message string) render.Renderer {
	return &ApiError{
		StatusCode: 500,
		StatusText: "Internal Server Error",
		Message:    message,
	}
}
