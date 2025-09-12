package validate

import "app/internal/domain"

type Validator interface {
	Validate(cmd domain.Command) error
}
