package elasticsearch

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

// Init 初始化ES客户端
func Init(ctx context.Context) error {
	esAddress := g.Cfg().MustGet(ctx, "elasticsearch.address").String()
	sniff := g.Cfg().MustGet(ctx, "elasticsearch.sniff").Bool()
	healthcheck := g.Cfg().MustGet(ctx, "elasticsearch.healthcheck").Bool()

	// 构建客户端选项
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(esAddress),
		elastic.SetSniff(sniff),
		elastic.SetHealthcheck(healthcheck),
	}

	// 创建客户端
	var err error
	client, err = elastic.NewClient(options...)
	if err != nil {
		return fmt.Errorf("未能创建 ES client: %v", err)
	}

	// 测试连接
	_, _, err = client.Ping(esAddress).Do(ctx)
	if err != nil {
		return fmt.Errorf("ES ping 失败: %v", err)
	}

	// 自动创建商品索引
	if err := createGoodsIndex(ctx); err != nil {
		return fmt.Errorf("创建商品索引失败: %v", err)
	}

	g.Log().Info(ctx, "Elasticsearch客户端和索引初始化成功")
	return nil
}

// GetClient 获取ES客户端
func GetClient() *elastic.Client {
	return client
}

// createGoodsIndex 创建商品索引
func createGoodsIndex(ctx context.Context) error {
	esIndexGoods := g.Cfg().MustGet(ctx, "elasticsearch.indices.goods").String()
	// 检查索引是否存在
	exists, err := client.IndexExists(esIndexGoods).Do(ctx)
	if err != nil {
		return err
	}

	if exists {
		g.Log().Info(ctx, "商品索引已存在")
		return nil
	}

	// 创建索引映射
	mapping := `{
		"mappings": {
			"properties": {
				"id": {"type": "long"},
				"name": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"images": {"type": "keyword"},
				"price": {"type": "long"},
				"level1_category_id": {"type": "long"},
				"level2_category_id": {"type": "long"},
				"level3_category_id": {"type": "long"},
				"brand": {
					"type": "keyword",
					"fields": {
						"text": {"type": "text"}
					}
				},
				"stock": {"type": "long"},
				"sale": {"type": "long"},
				"tags": {"type": "keyword"},
				"detail_info": {"type": "text"},
				"created_at": {"type": "text"},
				"updated_at": {"type": "text"},
				"deleted_at": {"type": "text"}
			}
		}
	}`

	// 创建索引
	createIndex, err := client.CreateIndex(esIndexGoods).Body(mapping).Do(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "创建索引失效:%s", err)
		return err
	}

	if !createIndex.Acknowledged {
		return fmt.Errorf("创建索引未被确认")
	}

	g.Log().Info(ctx, "商品索引创建成功")
	return nil
}
