// Package docs предоставляет OpenAPI документацию для Pet API.
// Этот файл создан с использованием пакета swaggo/swag.
//
// Схема версии: 1.0
// swagger:meta
package docs

import "pet/repository/petRepo"

//go:generate swag init -g docs.go

// Питомец.
// swagger:response petResponse
type petResponse struct {
	// in: body
	Body petRepo.Pet
}

// Ошибка.
// swagger:response errorResponse
type errorResponse struct {
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}
