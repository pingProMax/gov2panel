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
	"golang.org/x/crypto/bcrypt"
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
	gigabytes := Decimal(float64(bytes) / 1073741824)
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

// 获取当前日期字符串 2023922
func GetDateNowStr() string {
	timeNow := time.Now()
	return fmt.Sprintf("%s%s%s", strconv.Itoa(timeNow.Year()), strconv.Itoa(int(timeNow.Month())), strconv.Itoa(timeNow.Day())) // = 2023922
}

// 获取当前日期字符串 2023922 - day
func GetDateNowMinusDayStr(day int) string {
	timeNow := time.Now()
	timeNow = timeNow.Add(-time.Duration(day) * 24 * time.Hour)

	return fmt.Sprintf("%s%s%s", strconv.Itoa(timeNow.Year()), strconv.Itoa(int(timeNow.Month())), strconv.Itoa(timeNow.Day())) // = 2023922
}

// 生成加密密码
func BcryptGeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 密码效验
func BcryptCheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
