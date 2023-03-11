package common

import (
	"errors"
)

type DatabaseCheckController struct {
	BaseController
}

func (dc *DatabaseCheckController) Check() error {

	return errors.New("can't connect database")

}
