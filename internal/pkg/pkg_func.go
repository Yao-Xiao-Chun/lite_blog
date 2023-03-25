package pkg

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/dchest/captcha"
	"lite_blog/internal/pkg/entity"
)

// 公共的方法

// SuperCategory /**
func SuperCategory(allCate []entity.Cat, pid int) []entity.Cat {
	var arr []entity.Cat
	for _, v := range allCate {
		if pid == v.MenuParent {
			arr = append(arr, v)
			sonCate := SuperCategory(allCate, v.ID)
			arr = append(arr, sonCate...)
		}
	}
	return arr
}

// GetSummary /**
func GetSummary(content string) (string, error) {

	// bytes.Buffer，非常常用。
	var buf bytes.Buffer

	buf.Write([]byte(content))

	// 用goquery来解析content
	doc, err := goquery.NewDocumentFromReader(&buf)

	if err != nil {

		return "", err

	}

	// Text() 得到body元素下的文本内容（去掉html元素）
	str := doc.Find("body").Text()

	// 截取字符串
	if len(str) > 300 {

		str = str[0:300] + "..."
	}

	return str, nil

}

/**
验证码是否正确
@param id value
@return bool
*/
func VerifyCaptcha(captchaId string, captchaValue string) bool {

	if !captcha.VerifyString(captchaId, captchaValue) {
		return false
	} else {
		return true
	}

}
