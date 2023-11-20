package util

import (
	gpc "github.com/fstar-dev/go-playground-converter"
	"github.com/go-playground/validator/v10"
)

func GoValidator(s interface{}, config []gpc.ErrorMetaConfig) (interface{}, int) {
	var validate *validator.Validate
	validators := gpc.NewValidator(validate)
	bind := gpc.NewBindValidator(validators)

	errResponse, errCount := bind.BindValidator(s, config)
	return errResponse, errCount
}
