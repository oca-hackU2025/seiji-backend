package db

import (
	"log"

	"github.com/KENKUN-1031/seiji-backend/models"
)

// Seed inserts dummy data into the database
func Seed() error {
	seijikas := []models.Seijika{
		{
			Name:         "佐藤 花子",
			NameFurigana: "さとう はなこ",
			Age:          52,
			PartyName:    "立憲民主党",
			Term:         4,
			CatchPhrase:  "市民の声を政治に届ける！",
			MainImg:      "https://example.com/hanako_main.jpg",
			GraphImg:     "https://example.com/hanako_graph.jpg",
			Activities:   []string{"子育て支援", "地域医療の改善", "防災訓練"},
		},
		{
			Name:         "田中 一郎",
			NameFurigana: "たなか いちろう",
			Age:          60,
			PartyName:    "自由民主党",
			Term:         5,
			CatchPhrase:  "未来を切り開くリーダー",
			MainImg:      "https://example.com/ichiro_main.jpg",
			GraphImg:     "https://example.com/ichiro_graph.jpg",
			Activities:   []string{"経済改革", "防衛政策", "外交交渉"},
		},
		{
			Name:         "鈴木 太郎",
			NameFurigana: "すずき たろう",
			Age:          39,
			PartyName:    "日本維新の会",
			Term:         2,
			CatchPhrase:  "若い力で政治を変える！",
			MainImg:      "https://example.com/taro_main.jpg",
			GraphImg:     "https://example.com/taro_graph.jpg",
			Activities:   []string{"教育改革", "テクノロジー推進", "環境問題対策"},
		},
		{
			Name:         "山本 結衣",
			NameFurigana: "やまもと ゆい",
			Age:          47,
			PartyName:    "共産党",
			Term:         3,
			CatchPhrase:  "市民に寄り添う政治を",
			MainImg:      "https://example.com/yui_main.jpg",
			GraphImg:     "https://example.com/yui_graph.jpg",
			Activities:   []string{"ジェンダー平等", "労働者の権利", "最低賃金向上"},
		},
	}

	for _, s := range seijikas {
		if err := DB.Create(&s).Error; err != nil {
			log.Printf("❌ Failed to insert seijika: %+v, error: %v", s, err)
		} else {
			log.Printf("✅ Inserted seijika: %s", s.Name)
		}
	}

	return nil
}
