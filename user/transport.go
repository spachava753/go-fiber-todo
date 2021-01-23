package user

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gofiber/fiber/v2"
	"github.com/spachava753/go-fiber-todo/transport"
)

func createUserRequestDecoder(ctx *fiber.Ctx) (interface{}, error) {
	ctx.Accepts("json", "application/json")
	var req createUserRequest
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return createUserRequest{}, err
	}
	return req, nil
}

func createUserResponseEncoder(ctx *fiber.Ctx, resp interface{}) error {
	r := resp.(createUserResponse)
	if err := ctx.JSON(r); err != nil {
		return err
	}
	return nil
}

func deleteUserRequestDecoder(ctx *fiber.Ctx) (interface{}, error) {
	userId := ctx.Params("id", "")
	if userId == "" {
		return deleteUserRequest{}, fmt.Errorf("user id is not provided in path")
	}
	return deleteUserRequest{userId}, nil
}

func deleteUserResponseEncoder(ctx *fiber.Ctx, resp interface{}) error {
	r := resp.(deleteUserResponse)
	if err := ctx.JSON(r); err != nil {
		return err
	}
	return nil
}

func updateUserRequestDecoder(ctx *fiber.Ctx) (interface{}, error) {
	var user User
	if err := json.Unmarshal(ctx.Body(), &user); err != nil {
		return updateUserRequest{}, err
	}
	return updateUserRequest{user}, nil
}

func updateUserResponseEncoder(ctx *fiber.Ctx, resp interface{}) error {
	r := resp.(updateUserResponse)
	if err := ctx.JSON(r); err != nil {
		return err
	}
	return nil
}

func getUserByIdRequestDecoder(ctx *fiber.Ctx) (request interface{}, err error) {
	userId := ctx.Params("id", "")
	if userId == "" {
		return getUserByIdRequest{}, fmt.Errorf("user id is not provided in path")
	}
	return getUserByIdRequest{userId}, nil
}

func getUserByIdResponseEncoder(ctx *fiber.Ctx, resp interface{}) error {
	r := resp.(getUserByIdResponse)
	if err := ctx.JSON(r); err != nil {
		return err
	}
	return nil
}

func listUsersRequestDecoder(ctx *fiber.Ctx) (request interface{}, err error) {
	return listUsersRequest{}, err
}

func listUsersResponseEncoder(ctx *fiber.Ctx, resp interface{}) error {
	r := resp.(listUsersResponse)
	if err := ctx.JSON(r); err != nil {
		return err
	}
	return nil
}

func MakeRoutes(s Service, logger log.Logger, app *fiber.App) {
	createUserHandler := transport.MakeFiberHandler(
		makeHttpCreateUserEndpoint(s),
		createUserRequestDecoder,
		createUserResponseEncoder,
	)
	app.Put("/user", createUserHandler)

	deleteUserHandler := transport.MakeFiberHandler(
		makeHttpDeleteUserEndpoint(s),
		deleteUserRequestDecoder,
		deleteUserResponseEncoder,
	)
	app.Delete("/user/:id", deleteUserHandler)

	updateUserHandler := transport.MakeFiberHandler(
		makeHttpUpdateUserEndpoint(s),
		updateUserRequestDecoder,
		updateUserResponseEncoder,
	)
	app.Post("/user", updateUserHandler)

	getUserHandler := transport.MakeFiberHandler(
		makeHttpGetUserByIdEndpoint(s),
		getUserByIdRequestDecoder,
		getUserByIdResponseEncoder,
	)
	app.Get("/user/:id", getUserHandler)

	listUserHandler := transport.MakeFiberHandler(
		makeHttpListUsersEndpoint(s),
		listUsersRequestDecoder,
		listUsersResponseEncoder,
	)
	app.Get("/user", listUserHandler)
}
