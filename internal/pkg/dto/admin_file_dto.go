package dto

import "mywork/models"

type FictionList struct {
	models.LiteFiction
	TagsName    string
	Times       string
	DownloadNum int
}
