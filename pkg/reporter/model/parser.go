package model

import (
	"bjdaos/pegasus/pkg/reporter/types"
	strutil "github.com/1851616111/util/strings"
	"github.com/golang/glog"
)

//[^[{^{尿素（Urea）^,^2^,^生化检验}^}^,^{^{肌酐（Cr）^,^2^,^生化检验}^}^,^{^{尿酸（UA）^,^2^,^生化检验}^}^,^{^{甲胎蛋白(AFP)（发光）^,^2^,^肿瘤检测}^}^,^{^{一般检查^,^2^,^一般检查}^}^,^{^{血常规五分类^,^2^,^血常规}^}^,^{^{内科检查^,^2^,^内科}^}^,^{^{外科(女)^,^2^,^外科}^}^,^{^{妇科检查^,^2^,^妇科}^}^,^{^{胸部正位DR片^,^3^,^胸部摄片}^}^,^{^{免费早餐^,^3^,^免费早餐}^}^,^{^{蛋白4项^,^2^,^生化检验}^}^,^{^{丙氨酸氨基转移酶（ALT）^,^2^,^生化检验}^}^,^{^{天门冬氨酸氨基转移酶（AST）^,^2^,^生化检验}^}^,^{^{碱性磷酸酶（ALP）^,^2^,^生化检验}^}^,^{^{γ-谷氨酰转移酶（GGT）^,^2^,^生化检验}^}^,^{^{血脂2项^,^2^,^生化检验}^}^,^{^{心电图^,^2^,^心电图}^}^,^{^{眼科常规检查(含裂隙灯)^,^2^,^眼科}^}^,^{^{耳鼻咽喉常规^,^2^,^耳鼻喉科}^}^,^{^{口腔常规^,^2^,^口腔科}^}^,^{^{乳腺彩超^,^2^,^乳腺彩超}^}^,^{^{尿常规^,^2^,^尿常规}^}^,^{^{彩超经阴道^,^2^,^盆腔超声}^}^,^{^{腹部彩超^,^2^,^腹部超声}^}^,^{^{AST/ALT比值^,^2^,^生化检验}^}^,^{^{胆红素3项^,^2^,^生化检验}^}^,^{^{空腹血糖^,^2^,^生化检验}^}^,^{^{脂蛋白2项^,^2^,^生化检验}^}^,^{^{癌胚抗原(CEA)（发光）^,^2^,^肿瘤检测}^}]^]
func parseSalesData(sales *string) []types.Sale {
	sl := strutil.ClipDBArray(sales)

	ret := make([]types.Sale, len(sl))
	for i, str := range sl {
		strArr := strutil.ClipDBObject(&str)
		if len(strArr) == 3 {
			ret[i] = types.Sale{
				Department_name: unquote(strArr[0]),
				Checkup_name:    unquote(strArr[1]),
				Checkup_status:  unquote(strArr[2]),
			}
		} else {
			glog.Errorf("parseSalesData err with wrong format sales %s\n", str)
		}
	}

	return ret
}
