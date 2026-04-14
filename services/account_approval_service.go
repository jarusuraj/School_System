package services

import (
	"strings"

	"github.com/jarusuraj/schoolsystem/repository"
)

func AddApproval(sts, email, rol string) error {
	status := strings.TrimSpace(strings.ToLower(sts))
	if err := repository.AddAccountApproval(strings.ToLower(status), email); err != nil {
		return err
	}
	return nil
}
