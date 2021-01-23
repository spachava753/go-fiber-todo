package todo

import (
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/gofiber/fiber/v2"
	"github.com/spachava753/go-fiber-todo/transport"
)

func createTodoRequestEncoder(ctx *fiber.Ctx) (interface{}, error) {
	var req createTodoRequest
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return createTodoRequest{}, err
	}
	return req, nil
}

func deleteTodoRequestEncoder(ctx *fiber.Ctx) (interface{}, error) {
	todoId := ctx.Params("id", "")
	if todoId == "" {
		return deleteTodoRequest{}, errors.New("missing todo id in path")
	}
	return deleteTodoRequest{todoId}, nil
}

func updateTodoRequestEncoder(ctx *fiber.Ctx) (interface{}, error) {
	var todo Todo
	if err := json.Unmarshal(ctx.Body(), &todo); err != nil {
		return updateTodoRequest{}, err
	}
	return updateTodoRequest{todo}, nil
}

func getTodoByIdRequestEncoder(ctx *fiber.Ctx) (interface{}, error) {
	todoId := ctx.Params("id", "")
	if todoId == "" {
		return getTodoByIdRequest{}, errors.New("missing todo id in path")
	}
	return getTodoByIdRequest{todoId}, nil
}

func getTodosByUserIdRequestEncoder(ctx *fiber.Ctx) (interface{}, error) {
	userId := ctx.Params("userId", "")
	if userId == "" {
		return getTodosByUserIdRequest{}, errors.New("missing user id in path")
	}
	return getTodosByUserIdRequest{userId}, nil
}

func makeResponseDecoder() func(ctx *fiber.Ctx, resp interface{}) error {
	return func(ctx *fiber.Ctx, resp interface{}) error {
		if err := ctx.JSON(resp); err != nil {
			return err
		}
		return nil
	}
}

func MakeRoutes(s Service, logger log.Logger, app *fiber.App) {
	createTodoHandler := transport.MakeFiberHandler(
		makeHttpCreateTodoEndpoint(s),
		createTodoRequestEncoder,
		makeResponseDecoder(),
	)
	app.Put("/todo", createTodoHandler)

	deleteTodoHandler := transport.MakeFiberHandler(
		makeHttpDeleteTodoEndpoint(s),
		deleteTodoRequestEncoder,
		makeResponseDecoder(),
	)
	app.Delete("/todo/:id", deleteTodoHandler)

	updateTodoHandler := transport.MakeFiberHandler(
		makeHttpUpdateTodoEndpoint(s),
		updateTodoRequestEncoder,
		makeResponseDecoder(),
	)
	app.Post("/todo", updateTodoHandler)

	getTodoByIdHandler := transport.MakeFiberHandler(
		makeHttpGetTodoByIdEndpoint(s),
		getTodoByIdRequestEncoder,
		makeResponseDecoder(),
	)
	app.Get("/todo/:id", getTodoByIdHandler)

	getTodoByUserIdTodoHandler := transport.MakeFiberHandler(
		makeHttpListTodosByUserIdEndpoint(s),
		getTodosByUserIdRequestEncoder,
		makeResponseDecoder(),
	)
	app.Get("/todo/list/:userId", getTodoByUserIdTodoHandler)
}
