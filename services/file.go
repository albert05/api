package services

import (
	"api/common"
	"api/util/dates"
	"strconv"
)

func GenerateFileName() string {
	r := common.GenerateRangeNum(10000, 99999)
	n := dates.CurrentMicro()

	return common.Md5(strconv.Itoa(r)) + common.Md5(strconv.Itoa(int(n)))
}

func SaveFile() {

}