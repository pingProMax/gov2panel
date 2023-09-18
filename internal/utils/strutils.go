package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// md5 盐加密
func MD5V(password, salt string) string {
	combined := password + salt
	hasher := md5.New()
	io.WriteString(hasher, combined)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

// 消费订单生成 时间戳-套餐id-原价-优惠码-随机数
func UseOrderNo(planId int, price float64, code string) string {
	return fmt.Sprintf("%v-%v-%v-%v-%v", time.Now().Unix(), planId, price, code, strings.Split(uuid.New().String(), "-")[0])
}

// 充值订单生成 时间戳-充值金额(实际支付的)-payID-随机数
func RechargeOrderNo(price float64, payId int) string {
	return fmt.Sprintf("%v-%v-%v-%v", time.Now().Unix(), price, payId, strings.Split(uuid.New().String(), "-")[0])
}

// bytes 转 GB
func BytesToGB(bytes int64) float64 {
	gigabytes := float64(bytes) / 1073741824
	return gigabytes
}

// GB 转 bytes
func GBToBytes(gigabytes float64) int64 {
	bytes := int64(gigabytes * 1073741824)
	return bytes
}

// 2个字符后所有显示*号
func MaskString(input string) string {
	if len(input) <= 2 {
		return input
	}
	// 使用 strings.Repeat 函数来生成星号(*)的部分
	masked := input[:2] + strings.Repeat("*", len(input)-2)
	return masked
}

// float64 只保留两位小数
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// no rows in result set 错误判单
func IgnoreErrNoRows(err error) error {
	if err == sql.ErrNoRows {
		return nil
	} else {
		return err
	}
}
