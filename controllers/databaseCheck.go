package controllers

import "errors"

type DatabaseCheckController struct {
	AdminBaseController
}

func (dc *DatabaseCheckController) Check() error {

	return errors.New("can't connect database")

}