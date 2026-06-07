package config

import (
	"context"
	"fmt"

	"github.com/igortuchel/sec-esp-server/internal/handlers"
)

func Startup(cfg *handlers.Config) {
	path := "data/app.db"
	sqlDb, err := OpenSQLite(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("database connected: ", path)

	ctx := context.Background()
	redisDb, err := NewRedisClient(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("redis connected: ", redisDb.Options().Addr)

	cfg.Db = sqlDb
	cfg.Redis = redisDb
}
