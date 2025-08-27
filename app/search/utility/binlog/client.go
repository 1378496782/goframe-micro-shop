package binlog

import (
	"context"
	"github.com/go-mysql-org/go-mysql/mysql"
	"shop-goframe-micro-service-refacotor/app/search/utility/elasticsearch"
	"time"

	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// StartBinlogSyncer 启动 MySQL Binlog 同步
func StartBinlogSyncer(ctx context.Context) {
	// 从配置获取 MySQL 连接信息
	mysqlHost := g.Cfg().MustGet(ctx, "binlog.goods.mysql.host").String()
	mysqlPort := g.Cfg().MustGet(ctx, "binlog.goods.mysql.port").Int()
	mysqlUser := g.Cfg().MustGet(ctx, "binlog.goods.mysql.username").String()
	mysqlPass := g.Cfg().MustGet(ctx, "binlog.goods.mysql.password").String()

	// 创建 Binlog 同步器
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100, // 需要唯一
		Flavor:   "mysql",
		Host:     mysqlHost,
		Port:     uint16(mysqlPort),
		User:     mysqlUser,
		Password: mysqlPass,
	}

	syncer := replication.NewBinlogSyncer(cfg)
	defer syncer.Close()

	// 获取当前位点（可以从保存的位置读取，这里简单从最新开始）
	position := mysql.Position{
		Name: "", // 空字符串表示从当前位点开始
		Pos:  0,
	}

	// 开始同步
	streamer, err := syncer.StartSync(position)
	if err != nil {
		g.Log().Errorf(ctx, "启动 Binlog 同步失败: %v", err)
		return
	}

	g.Log().Info(ctx, "开始监听 MySQL Binlog...")

	for {
		// 获取 Binlog 事件
		ev, err := streamer.GetEvent(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "获取 Binlog 事件失败: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		// 处理事件
		processBinlogEvent(ctx, ev)
	}
}

// processBinlogEvent 处理 Binlog 事件
func processBinlogEvent(ctx context.Context, ev *replication.BinlogEvent) {
	switch e := ev.Event.(type) {
	case *replication.RowsEvent:
		// 只处理指定数据库的 goods_info 表
		if string(e.Table.Schema) != "goods" || string(e.Table.Table) != "goods_info" {
			return
		}

		g.Log().Debugf(ctx, "收到Binlog事件: 数据库=%s, 表=%s", e.Table.Schema, e.Table.Table)

		// 根据事件类型处理
		switch ev.Header.EventType {
		case replication.WRITE_ROWS_EVENTv1, replication.WRITE_ROWS_EVENTv2:
			handleInsert(ctx, e.Rows)
		case replication.UPDATE_ROWS_EVENTv1, replication.UPDATE_ROWS_EVENTv2:
			handleUpdate(ctx, e.Rows)
		case replication.DELETE_ROWS_EVENTv1, replication.DELETE_ROWS_EVENTv2:
			handleDelete(ctx, e.Rows)
		default:
		}
	}
}

// handleInsert 处理插入事件
func handleInsert(ctx context.Context, rows [][]interface{}) {
	for _, row := range rows {
		// 将行数据转换为 map
		columnMap := parseRowData(row)
		upsertToES(ctx, columnMap)
	}
}

// handleUpdate 处理更新事件
func handleUpdate(ctx context.Context, rows [][]interface{}) {
	// 更新事件的行数据格式为 [旧行数据, 新行数据]
	for i := 0; i < len(rows); i += 2 {
		if i+1 < len(rows) {
			columnMap := parseRowData(rows[i+1]) // 取新数据
			upsertToES(ctx, columnMap)
		}
	}
}

// handleDelete 处理删除事件
func handleDelete(ctx context.Context, rows [][]interface{}) {
	for _, row := range rows {
		columnMap := parseRowData(row)
		deleteFromES(ctx, columnMap)
	}
}

// parseRowData 解析行数据为 map
func parseRowData(row []interface{}) map[string]interface{} {
	// 这里需要根据你的表结构定义字段名
	fields := []string{
		"id", "name", "images", "price", "level1_category_id",
		"level2_category_id", "level3_category_id", "brand",
		"stock", "sale", "tags", "detail_info",
		"created_at", "updated_at", "deleted_at",
	}

	result := make(map[string]interface{})
	for i, value := range row {
		if i < len(fields) {
			result[fields[i]] = value
		}
	}
	return result
}

// upsertToES 插入或更新文档到 ES
func upsertToES(ctx context.Context, data map[string]interface{}) {
	client := elasticsearch.GetClient()
	if client == nil {
		g.Log().Error(ctx, "ES客户端未初始化")
		return
	}

	esIndexGoods := g.Cfg().MustGet(ctx, "elasticsearch.indices.goods").String()

	_, err := client.Index().
		Index(esIndexGoods).
		Id(gconv.String(data["id"])).
		BodyJson(map[string]interface{}{
			"id":                 gconv.Uint32(data["id"]),
			"name":               gconv.String(data["name"]),
			"images":             gconv.String(data["images"]),
			"price":              gconv.Uint64(data["price"]),
			"level1_category_id": gconv.Uint32(data["level1_category_id"]),
			"level2_category_id": gconv.Uint32(data["level2_category_id"]),
			"level3_category_id": gconv.Uint32(data["level3_category_id"]),
			"brand":              gconv.String(data["brand"]),
			"stock":              gconv.Uint32(data["stock"]),
			"sale":               gconv.Uint32(data["sale"]),
			"tags":               gconv.String(data["tags"]),
			"detail_info":        gconv.String(data["detail_info"]),
			"created_at":         gconv.String(data["created_at"]),
			"updated_at":         gconv.String(data["updated_at"]),
			"deleted_at":         gconv.String(data["deleted_at"]),
		}).
		Do(ctx)

	if err != nil {
		g.Log().Errorf(ctx, "同步商品到ES失败: %v", err)
	} else {
		g.Log().Debugf(ctx, "成功同步商品到ES: ID=%s", gconv.String(data["id"]))
	}
}

// deleteFromES 从 ES 删除文档
func deleteFromES(ctx context.Context, data map[string]interface{}) {
	client := elasticsearch.GetClient()
	if client == nil {
		g.Log().Error(ctx, "ES客户端未初始化")
		return
	}

	id := gconv.String(data["id"])
	if id == "" {
		g.Log().Error(ctx, "删除操作未找到ID")
		return
	}

	esIndexGoods := g.Cfg().MustGet(ctx, "elasticsearch.indices.goods").String()

	_, err := client.Delete().
		Index(esIndexGoods).
		Id(id).
		Do(ctx)

	if err != nil {
		g.Log().Errorf(ctx, "从ES删除商品失败: ID=%s, error=%v", id, err)
	} else {
		g.Log().Debugf(ctx, "成功从ES删除商品: ID=%s", id)
	}
}
