package services

import (
	"api/common"
	"api/util/dates"
	"strconv"
	"fmt"
	"github.com/astaxie/beego"
)

func GenerateFileName() string {
	r := common.GenerateRangeNum(10000, 99999)
	n := dates.CurrentMicro()

	return common.Md5(strconv.Itoa(r) + strconv.Itoa(int(n)))
}

func GenerateFilePath(key string) string {
	year := dates.NowYearStr()
	month := dates.NowMonthStr()
	path := fmt.Sprintf("%s/%s/%s/%s/", beego.AppConfig.String("fileRootPath"), key, year, month)

	return path
}
