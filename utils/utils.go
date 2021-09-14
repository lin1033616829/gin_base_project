package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type CommonUtil struct{}

func (c *CommonUtil) IntStrToIntArr(idStr string) []int {
	idStrArr := strings.Split(idStr, ",")
	idIntArr := make([]int, 1)
	for _, tagId := range idStrArr {
		tmpInt, _ := strconv.Atoi(tagId)
		idIntArr = append(idIntArr, tmpInt)
	}
	return idIntArr
}

func (c *CommonUtil) Md5Str(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// RandString 生成随机字符串
func (c *CommonUtil) RandString(len int) string {
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
