package main

import (
	"flag"
	"github.com/go-kit/kit/log"
	"github.com/oklog/ulid/v2"
	"github.com/spachava753/go-fiber-todo/todo"
	"github.com/spachava753/go-fiber-todo/user"
	"math/rand"
	"os"
	"time"
)

func main() {

	flag.Parse()

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)

	t := time.Unix(1000000, 0)
	userEntropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	var userService user.Service
	userService = user.NewMemUserService(userEntropy, t)
	userService = user.NewBasicLoggingService(log.With(logger, "component", "user"), userService)

	todoEntropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	var todoService todo.Service
	todoService = todo.NewMemTodoService(todoEntropy, t)
	todoService = todo.NewBasicLoggingService(log.With(logger, "component", "todo"), todoService)
}
