package entity

import "lite_blog/models"

type FictionList struct {
	models.LiteFiction
	TagsName    string
	Times       string
	DownloadNum int
}
