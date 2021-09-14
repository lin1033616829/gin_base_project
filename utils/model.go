package utils

import (
	"bytes"
	"encoding/base32"
	"errors"
	"fmt"
	"github.com/pborman/uuid"
	"mmrights/global"
	"strings"
	"time"
)

var reservedName = []string{
	"signup",
	"login",
	"admin",
	"channel",
	"post",
	"api",
	"oauth",
	"error",
	"help",
}

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

// NewId is a globally unique identifier.  It is a [A-Z0-9] string 26
// characters long.  It is a UUID version 4 Guid that is zbased32 encoded
// with the padding stripped off.
func NewId() string {
	var b bytes.Buffer
	var encoder = base32.NewEncoder(encoding, &b)
	_,_ = encoder.Write(uuid.NewRandom())
	_ = encoder.Close()
	b.Truncate(26) // removes the '==' padding
	return b.String()
}

// NewRandomTeamName is a NewId that will be a valid team name.
func NewRandomTeamName() string {
	teamName := NewId()

	for IsReservedTeamName(teamName) {
		teamName = NewId()
	}

	return teamName
}

func IsReservedTeamName(s string) bool {
	s = strings.ToLower(s)

	for _, value := range reservedName {
		if strings.Index(s, value) == 0 {
			return true
		}
	}

	return false
}

// GetMillis is a convenience method to get milliseconds since epoch.
func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetInSqlBySlice(itemSlice []string) (string, error) {
	if len(itemSlice) == 0 {
		global.ServerLog.Debug("in查询中 param 元素不能空")
		return "", errors.New("参数禁止为空")
	}

	var ret string

	for _, item := range itemSlice {
		ret = fmt.Sprintf("%v,'%v'", ret, item)
	}

	ret = strings.TrimLeft(ret, ",")
	return ret, nil
}