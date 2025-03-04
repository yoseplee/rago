package v1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONDocumentLoader_Load(t *testing.T) {
	loader := JSONDocumentLoader{
		FilePath: "../data/sample_shop_items.json",
	}

	load, err := loader.Load()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Println(load)
	assert.Equal(t, jsonResult, fmt.Sprintf("%v", load))
}

const jsonResult = `[{"_id":"01JACVJ9H1RR4GF2QG8WTXKDXG","chainId":"01J4K8SQF6Z80DC20FXZ9JG2SD","chainItemId":"01JACVJ9FVBN79PCA7116TKQPV","displayRule":{"description":"水分、イオンの補給がスムーズ発汗により失われた水分、イオン（電解質）をスムーズに補給する健康飲料です。スポーツやお風呂上りなどに最適適切な濃度と体液に近い組成の電解質溶液のため、すばやく吸収されます。そのためスポーツ、仕事、お風呂上り、寝起きなど、発汗状態におかれている方に最も適した飲料です。原材料砂糖、果糖ぶどう糖液糖、果汁、食塩、酸味料、香料、塩化K、乳酸Ca、調味料（アミノ酸）、塩化Mg、酸化防止剤（ビタミンC）内容量500ml","itemImage":{"imageUrl":"https://retail-item.line-scdn.net/r/retail/itemreal/cj0tMzY2dTRkdTI4MDNxNyZzPWpwNiZ0PW0mdT0xZjA2YTJsZHMzaGcwJmk9MA/main","obsOid":"cj0tMzY2dTRkdTI4MDNxNyZzPWpwNiZ0PW0mdT0xZjA2YTJsZHMzaGcwJmk9MA"},"sortOrder":null},"itemClassification":"GENERAL","itemName":"大塚製薬　ポカリスエット　500ml（45019517）","shopId":"01J4KJWZYYWZ549PG8RAN98QAF"} {"_id":"01JACVE0S4CKRM18V1ZB2VZ9JS","chainId":"01J4K8SQF6Z80DC20FXZ9JG2SD","chainItemId":"01JACVE0QVKAG1S8G78H56GK1B","displayRule":{"description":"・髪をまとめたり、髪を留める以外の目的には使用しないでください。・変形する恐れがありますので、本体にドライヤーの熱を加えないでください。・必要以上の力を加えると破損する場合があります。乳幼児の手の届かない所に保管してください。","itemImage":{"imageUrl":"https://retail-item.line-scdn.net/r/retail/itemreal/cj0xczR0NzQwMWg5OHYxJnM9anA2JnQ9bSZ1PTFmMDY5aGllZzNoZzAmaT0w/main","obsOid":"cj0xczR0NzQwMWg5OHYxJnM9anA2JnQ9bSZ1PTFmMDY5aGllZzNoZzAmaT0w"},"sortOrder":null},"itemClassification":"GENERAL","itemName":"イノウエ　ヘップリング　L　黒　2本（4510073166085）","shopId":"01J4KJWZYYWZ549PG8RAN98QAF"} {"_id":"01JACVE12JSR12NY3JWPYHCPCD","chainId":"01J4K8SQF6Z80DC20FXZ9JG2SD","chainItemId":"01JACVE10SJEW97BQW3S0HF8TM","displayRule":{"description":"・髪をまとめたり、髪を留める以外の目的には使用しないでください。・変形する恐れがありますので、本体にドライヤーの熱を加えないでください。・必要以上の力を加えると破損する場合があります。乳幼児の手の届かない所に保管してください。","itemImage":{"imageUrl":"https://retail-item.line-scdn.net/r/retail/itemreal/cj0tN2xpN2o1ODg3OGdyaiZzPWpwNiZ0PW0mdT0xZjA2OWhqZzQzaTAwJmk9MA/main","obsOid":"cj0tN2xpN2o1ODg3OGdyaiZzPWpwNiZ0PW0mdT0xZjA2OWhqZzQzaTAwJmk9MA"},"sortOrder":null},"itemClassification":"GENERAL","itemName":"イノウエ　ヘップリング　L　茶　2本（4510073166092）","shopId":"01J4KJWZYYWZ549PG8RAN98QAF"} {"_id":"01JACVE1DDHVCE1B86KVN89PHZ","chainId":"01J4K8SQF6Z80DC20FXZ9JG2SD","chainItemId":"01JACVE1CBM63H42XKZSMK1YF2","displayRule":{"description":"・髪をまとめたり、髪を留める以外の目的には使用しないでください。・変形する恐れがありますので、本体にドライヤーの熱を加えないでください。・必要以上の力を加えると破損する場合があります。乳幼児の手の届かない所に保管してください。","itemImage":{"imageUrl":"https://retail-item.line-scdn.net/r/retail/itemreal/cj0ya2VxZGRzdjBwMTRvJnM9anA2JnQ9bSZ1PTFmMDY5aGt0ZzNoMDAmaT0w/main","obsOid":"cj0ya2VxZGRzdjBwMTRvJnM9anA2JnQ9bSZ1PTFmMDY5aGt0ZzNoMDAmaT0w"},"sortOrder":null},"itemClassification":"GENERAL","itemName":"イノウエ　髪に絡まないゴム黒60本LD7　60本（4510073166108）","shopId":"01J4KJWZYYWZ549PG8RAN98QAF"} {"_id":"01JACVE1R883E7FKANP7H58F5M","chainId":"01J4K8SQF6Z80DC20FXZ9JG2SD","chainItemId":"01JACVE1Q5416NK8ZKGMH2HKQP","displayRule":{"description":"・髪をまとめたり、髪を留める以外の目的には使用しないでください。・変形する恐れがありますので、本体にドライヤーの熱を加えないでください。・必要以上の力を加えると破損する場合があります。乳幼児の手の届かない所に保管してください。","itemImage":{"imageUrl":"https://retail-item.line-scdn.net/r/retail/itemreal/cj0tMmNndGFxMDFtbGpjYSZzPWpwNiZ0PW0mdT0xZjA2OWhtODgzaGcwJmk9MA/main","obsOid":"cj0tMmNndGFxMDFtbGpjYSZzPWpwNiZ0PW0mdT0xZjA2OWhtODgzaGcwJmk9MA"},"sortOrder":null},"itemClassification":"GENERAL","itemName":"イノウエ　髪に絡まないゴム　茶　60本（4510073166115）","shopId":"01J4KJWZYYWZ549PG8RAN98QAF"}]`
