package clinicaldepartmentsconverter

import (
	"strings"
	"testing"
)

func TestGetShortName(t *testing.T) {
	want := "リハ"
	if got := GetShortName("リハビリテーション科"); got != want {
		t.Errorf("GetShortName(\"リハビリテーション科\") = %s, want %s", got, want)
	}
}

func TestGetShortNameEmpty(t *testing.T) {
	want := ""
	if got := GetShortName("存在しない科"); got != want {
		t.Errorf("GetShortName(\"リハビリテーション科\") = %s, want %s", got, want)
	}
}

func TestGetLongName(t *testing.T) {
	want := "リハビリテーション科"
	if got := GetLongName("リハ"); got != want {
		t.Errorf("GetLongName(\"リハ\") = %s, want %s", got, want)
	}
}

func TestGetLongNameEmpty(t *testing.T) {
	want := ""
	if got := GetLongName("存在しない科"); got != want {
		t.Errorf("GetLongName(\"存在しない科\") = %s, want %s", got, want)
	}
}

func TestShortNameList(t *testing.T) {
	want := "こう,ひ,アレ,リウ,リハ,内,呼,呼内,呼外,和漢,外,婦,小,小外,小歯,形外,循,心内,心外,性,放,救命,整外,歯,歯外,気食,消／消胃,産,産婦,病理,皮,皮ひ,眼,矯歯,神,神内,精,美外,耳い／耳,胃／消胃,胸外,脳内,脳外,臨床,透析,麻"
	if got := ShortNameList(); strings.Join(got, ",") != want {
		t.Errorf("ShortNameList() = %s\nwant\n%s", strings.Join(got, ","), want)
	}
}

func TestLongNameList(t *testing.T) {
	want := "こう門科,アレルギー科,リウマチ科,リハビリテーション科,人工透析内科,内科,呼吸器内科,呼吸器外科,呼吸器科,外科,婦人科,小児外科,小児歯科,小児科,形成外科,循環器科,心療内科,心臓血管外科,性病科,放射線科,救急科,整形外科,歯科,歯科口腔外科,気管食道科,泌尿器科,消化器科,漢方内科,産婦人科,産科,病理診断科,皮膚泌尿器科,皮膚科,眼科,矯正歯科,神経内科,神経科,精神科,美容外科,耳鼻いんこう科,胃腸科,胸部外科,脳神経内科,脳神経外科,臨床検査科,麻酔科"
	if got := LongNameList(); strings.Join(got, ",") != want {
		t.Errorf("LongNameList() = %s\nwant\n%s", strings.Join(got, ","), want)
	}
}
