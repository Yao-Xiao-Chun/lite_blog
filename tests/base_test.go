package test

import (
	"mywork/internal/app/common"
	"testing"
)

func TestBaseController_SetMd5Pwd(t *testing.T) {

	var base common.BaseController

	md5 := base.SetMd5Pwd("1234@abcd")

	t.Log(md5)
}
