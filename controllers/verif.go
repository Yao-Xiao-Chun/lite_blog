package controllers


import (
	"github.com/dchest/captcha"
)

// operations for Captcha
type CaptchaController struct {
	BaseController
}

/**
	验证码是否正确
	@param id value
	@return bool
 */
func VerifyCaptcha(captchaId string,captchaValue string) bool{

	if !captcha.VerifyString(captchaId, captchaValue) {
		return  false
	} else {
		return true
	}

}

