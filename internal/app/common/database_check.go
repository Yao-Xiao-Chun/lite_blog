package common

import (
	"errors"
	"mywork/internal/app/admin"
)

type DatabaseCheckController struct {
	admin.AdminBaseController
}

func (dc *DatabaseCheckController) Check() error {

	return errors.New("can't connect database")

}
