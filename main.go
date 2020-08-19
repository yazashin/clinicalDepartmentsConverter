package main

import (
	"fmt"
	"strings"
)

// 診療科名の略称との対応表
// https://kouseikyoku.mhlw.go.jp/kinki/tyousa/000090952.pdf

var departmentMap = map[string]string{
	"内科":         "内",
	"精神科":        "精",
	"神経科":        "神",
	"神経内科":       "神内",
	"呼吸器科":       "呼",
	"胃腸科":        "胃／消胃",
	"消化器科":       "消／消胃",
	"循環器科":       "循",
	"小児科":        "小",
	"外科":         "外",
	"整形外科":       "整外",
	"形成外科":       "形外",
	"美容外科":       "美外",
	"脳神経外科":      "脳外",
	"呼吸器外科":      "呼外",
	"心臓血管外科":     "心外",
	"小児外科":       "小外",
	"皮膚泌尿器科":     "皮ひ",
	"皮膚科":        "皮",
	"泌尿器科":       "ひ",
	"性病科":        "性",
	"こう門科":       "こう",
	"産婦人科":       "産婦",
	"産科":         "産",
	"婦人科":        "婦",
	"眼科":         "眼",
	"耳鼻いんこう科":    "耳い／耳",
	"気管食道科":      "気食",
	"リハビリテーション科": "リハ",
	"放射線科":       "放",
	"麻酔科":        "麻",
	"心療内科":       "心内",
	"リウマチ科":      "リウ",
	"アレルギー科":     "アレ",
	"病理診断科":      "病理",
	"臨床検査科":      "臨床",
	"救急科":        "救命",
	"歯科":         "歯",
	"矯正歯科":       "矯歯",
	"小児歯科":       "小歯",
	"歯科口腔外科":     "歯外",
	"呼吸器内科":      "呼内",
	"人工透析内科":     "透析",
	"脳神経内科":      "脳内",
	"漢方内科":       "和漢",
	"胸部外科":       "胸外",
}

// GetShortName 一般名から略称に変換する
func GetShortName(str string) string {
	return departmentMap[str]
}

// GetLongName 略称から一般名に変換する
func GetLongName(str string) string {
	for k, v := range departmentMap {
		if str == v {
			return k
		}
	}
	return ""
}

// ShortNameList 略称だけの一覧を返す
func ShortNameList() []string {
	vs := []string{}
	for _, v := range departmentMap {
		vs = append(vs, v)
	}
	return vs
}

// LongNameList 一般名だけの一覧を返す
func LongNameList() []string {
	ks := []string{}
	for k := range departmentMap {
		ks = append(ks, k)
	}
	return ks
}

func main() {
	fmt.Println(GetShortName("アレルギー科"))

	fmt.Println(GetLongName("アレ"))

	fmt.Println(strings.Join(LongNameList(), ", "))

	fmt.Println(strings.Join(ShortNameList(), ", "))
}
