package common

import (
	"strings"
	"strconv"
	"fmt"
	"crypto/md5"
	"math/rand"
	"time"
	"os"
)

func CheckParams(params map[string]string, keys []string) bool {
	for _, k := range keys {
		if val, ok := params[k]; !ok || strings.TrimSpace(val) == "" {
			return false
		}
	}

	return true
}

func Str2Int(s string) int {
	s1, _ := strconv.ParseInt(s, 10, 64)

	return int(s1)
}

func Str2Float(s string) float64 {
	s1, _ := strconv.ParseFloat(s, 10)

	return s1
}

func Float2Money(v float64) int {
	return int(v * 100)
}

func Money2Float(v int) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(v) / 100), 64)
	return value
}

func Md5(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min) + min
	return randNum
}

// 检查文件或目录是否存在
// param string filename  文件或目录
// return bool
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 创建目录
// param string dirPath 目录
// return error
func MakeDir(dirPath string) error {
	if !IsExist(dirPath) {
		return os.MkdirAll(dirPath, 0755)
	}

	return nil
}
