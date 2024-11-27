package controller

import (
	"net/http"

	"github.com/poteto0/poteto"
	"github.com/poteto0/poteto-sample-api/model"
)

func UserHandler(ctx poteto.Context) error {
	user := model.User{
		Name: "user",
	}
	return ctx.JSON(http.StatusOK, user)
}

func UserIdHandler(ctx poteto.Context) error {
	name, _ := ctx.PathParam("name")
	user := model.User{
		Name: name,
	}
	return ctx.JSON(http.StatusOK, user)
}
