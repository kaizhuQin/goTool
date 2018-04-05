/**
 * 通用工具类
 */
package goTool

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"math/rand"
)

/**
 * @brief 生成32位MD5
 * @param string data 需要加密的字符
 * @return mixed
 */
func Md5(data string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

/**
 * @brief 生成随机字符串-数字加字母
 * @param length int 固定长度
 * @return mixed
 */
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/**
 * @brief 生成随机字符串-纯数字
 * @param length int 固定长度
 * @return mixed
 */
func GetRandomIntString(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}