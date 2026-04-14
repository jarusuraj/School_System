package repository

import (
	"context"

	"github.com/jarusuraj/schoolsystem/config"
)

func AddAccountApproval(status, email string) error {
	const query string = "UPDATE users SET status=($1) WHERE email=($2)"
	_, err := config.DB.Exec(context.Background(), query, status, email)
	if err != nil {
		return err
	}

	return nil
}
