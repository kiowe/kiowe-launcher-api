package utils

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgconn"
)

func ParsePgError(err error) error {
	var pgErr *pgconn.PgError

	if errors.Is(err, pgErr) {
		pgErr = err.(*pgconn.PgError)
		return fmt.Errorf("database error. message: %s, details: %s, where: %s, sqlstate: %s",
			pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.SQLState())
	}

	return err
}

type QueryParam struct {
	Name  string
	Value string
}

func GetQueryParams(c *fiber.Ctx, paramName ...string) map[string]string {
	p := make(map[string]string)

	for _, name := range paramName {
		p[name] = c.Query(name)
	}

	return p
}
