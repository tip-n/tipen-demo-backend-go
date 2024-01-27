package main

import (
	"context"
	envconf "tipen-demo/config/env"
	"tipen-demo/pkg"

	sq "github.com/Masterminds/squirrel"
)

func main() {
	env := envconf.ReadEnv()
	db, err := env.Database.Init()
	if err != nil {
		panic(err.Error())
	}

	t := pkg.NowUTC()

	db.Exec(context.Background(), "")
	sq.Insert("restaurant_menu_type").Columns(
		"created_at",
		"updated_at",
		"type",
	).Values(
		[]any{
			t, t, "Coffee",
			t, t, "Non-Coffee",
		},
	).ToSql()
}
