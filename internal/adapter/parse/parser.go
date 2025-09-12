package parse

import "app/internal/domain"

type Parser interface {
	Parse(raw string) (domain.Command, error)
}
