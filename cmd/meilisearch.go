package cmd

import "github.com/meilisearch/meilisearch-go"

// MeiliSearchConfig meili 搜索引擎配置文件
type MeiliSearchConfig struct {
	Host      string
	Port      string
	MasterKey string
}

func NewMeiliSearch(host, port, masterKey string) *MeiliSearchConfig {

	return &MeiliSearchConfig{
		Host:      host,
		Port:      port,
		MasterKey: masterKey,
	}
}

// OpenClient 返回客户端
func (m *MeiliSearchConfig) OpenClient() *meilisearch.Client {

	return meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   m.Host + m.Port,
		APIKey: m.MasterKey,
	})
}
