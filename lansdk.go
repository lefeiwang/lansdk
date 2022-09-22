package lansdk

import (
	"time"

	"github.com/lefeiwang/lansdk/tool"
)

// any类型转字符串
func ArgsToS(args ...interface{}) []string {
	return tool.ArgsToS(args)
}

// 花费时长
func TimeCost(start time.Time, mark string) {
	tool.TimeCost(start, mark)
}

// 加解密
// 获取aes加密key
func GetAesKey(n int) (aesKey []byte) {
	return tool.GetAesKey(n)
}

// aes加密
func AesEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
	return tool.AesEncrypt(plaintext, key, iv)
}

// aes解密
func AesDecrypt(ciphertext []byte, key, iv []byte) (plaintext []byte, err error) {
	return tool.AesDecrypt(ciphertext, key, iv)
}

// rsa 密钥对
func GetRsaKeys() (pubKey, priKey []byte) {
	return tool.GetRsaKeys()
}

// 公钥加密函数-分段
func RsaPubEncryptBlock(plaintext, publicKeyByte []byte) (ciphertext []byte, err error) {
	return tool.RsaPubEncryptBlock(plaintext, publicKeyByte)
}

// 公钥解密函数-分段
func RsaPubDecryptBlock(ciphertext, publicKeyByte []byte) (plaintext []byte, err error) {
	return tool.RsaPubDecryptBlock(ciphertext, publicKeyByte)
}

// 私钥加密函数-分段
func RsaPrivEncryptBlock(plaintext, privateKeyByte []byte) (ciphertext []byte, err error) {
	return tool.RsaPrivEncryptBlock(plaintext, privateKeyByte)
}

// 私钥解密函数-分段
func RsaPrivDecryptBlock(ciphertext, privateKeyByte []byte) (plaintext []byte, err error) {
	return tool.RsaPrivDecryptBlock(ciphertext, privateKeyByte)
}

// 读文件
func ReadFile(srcFile string) (content []byte, err error) {
	return tool.ReadFile(srcFile)
}

// 写文件，不存在自动创建
func WriteFile(srcFile string, content []byte) (err error) {
	return tool.WriteFile(srcFile, content)
}

// 判断所给路径文件/文件夹是否存在
func FileIsExist(srcFile string) bool {
	return tool.FileIsExist(srcFile)
}

// 删除所给路径文件/文件夹
func RemoveFile(srcFile string) error {
	return tool.RemoveFile(srcFile)
}

// 创建嵌套文件夹
func MkDir(srcFile string) error {
	return tool.MkDir(srcFile)
}

// 拷贝文件，适用于小文件拷贝，先全部加载到内存
func CopyFile(srcFile, destFile string) (int, error) {
	return tool.CopyFile(srcFile, destFile)
}

// 拷贝文件，适用于大文件拷贝，分段拷贝，段大小BUFFERSIZE字节
func CopyFileBuff(src, dst string, BufferSize int64) (cn int, err error) {
	return tool.CopyFileBuff(src, dst, BufferSize)
}

// 图片url转[]byte
func UrlImgByte(imgUrl string) ([]byte, error) {
	return tool.UrlImgByte(imgUrl)
}

// 图片url转base64字符串
func UrlImgBase64(imgUrl string) (string, error) {
	return tool.UrlImgBase64(imgUrl)
}

// 获取本地ip
func GetOutBoundIP() (ip string, err error) {
	return tool.GetOutBoundIP()
}

// 获取计算机名称
func GetHostName() string {
	return tool.GetHostName()
}

// 返回从[min, max)1个随机整数
func RandInt(min, max int) int {
	return tool.RandInt(min, max)
}

// 返回从[min, max)不重复的n个随机从小到大有序整数切片
func RandNeSortInts(min, max, n int) []int {
	return tool.RandInts(min, max, n)
}

// 返回从[min, max)n个随机浮点数
func RandFloats(min, max float64, n int) []float64 {
	return tool.RandFloats(min, max, n)
}

// 打印报错栈信息
func PrintStackTrace(err interface{}) string {
	return tool.PrintStackTrace(err)
}

// http请求
type Request tool.Request
