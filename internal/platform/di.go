package platform

import (
	"app/internal/adapter/http"
	"app/internal/adapter/parse"
	"app/internal/adapter/validate"
	"app/internal/domain"
	"app/internal/usecase"
	"fmt"
)

type App struct {
	Handler *http.Handler
	Config  Config
}

func Build() (*App, error) {
	// domain
	decoder := domain.DefaultDecoder{}

	// adapters
	parser := parse.WhitespaceParser{}
	validator := validate.DefaultValidator{}

	// usecase
	uc := usecase.NewDecodeInteractor(parser, validator, decoder)

	// http handler
	h := http.NewHandler(uc)

	cfg := Load()
	_ = fmt.Sprintf("") // ここでcfgを必要に応じて検証

	return &App{Handler: h, Config: cfg}, nil
}
