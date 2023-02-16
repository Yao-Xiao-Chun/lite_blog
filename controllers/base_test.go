package controllers

import "testing"

func TestBaseController_SetMd5Pwd(t *testing.T) {

	var base BaseController

	md5 := base.SetMd5Pwd("1234@abcd")

	t.Log(md5)
}
