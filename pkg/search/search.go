package search

import (
	"fmt"
	"github.com/meilisearch/meilisearch-go"
	"github.com/prometheus/common/log"
	"lite_blog/internal/pkg/entity"
	"lite_blog/models"
	"lite_blog/pkg/enum"
	"lite_blog/pkg/model"
)

var (
	SearchSDK *meilisearch.Client
)

// Search 搜索引擎
type Search struct {
}

// AutoSearchData 同步搜索引擎数据
func AutoSearchData() {

	if flag := SearchSDK.IsHealthy(); flag {
		//准备更新数据
		var arr []models.LiteArticle
		err := model.Db.Select("title,descript,keywords,id").Order("id desc").Find(&arr).Error

		if err != nil {
			log.Error(fmt.Sprintf("同步数据引擎错误%v", err))
		}

		index := SearchSDK.Index(enum.ARTICLE_SEARCH) //搜索引擎的数据
		documents := make([]entity.SearchFromEntity, 0)
		for _, item := range arr {

			documents = append(documents, entity.SearchFromEntity{
				Title:       item.Title,
				Description: item.Descript,
				Keywords:    item.Keywords,
				ID:          int(item.ID), //设置ID
			})
		}

		_, err = index.AddDocuments(documents)

		log.Error(fmt.Sprintf("写入搜索引擎的结果：%v", err))
	}
}
