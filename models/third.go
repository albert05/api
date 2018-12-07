package models

import (
	"github.com/astaxie/beego"
	"strings"
)

var KeyList map[string]string

func init() {
	KeyList = make(map[string]string)
	thirdList := beego.AppConfig.String("thirds")

	l := strings.Split(thirdList, ",")
	for _, items := range l {
		ks := strings.Split(items, "=")
		if len(ks) == 2 {
			KeyList[ks[0]] = ks[1]
		}
	}
}
