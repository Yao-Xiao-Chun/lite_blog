package home

import "testing"

func TestIndexController_ArticleClick(t *testing.T) {
	type fields struct {
		HomeBaseController HomeBaseController
	}
	var tests []struct {
		name   string
		fields fields
	}

	//子测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &IndexController{
				HomeBaseController: tt.fields.HomeBaseController,
			}
			c.ArticleClick()
		})
	}
}
