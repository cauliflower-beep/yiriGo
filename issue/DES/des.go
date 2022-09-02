package main

/*
业务背景：
给汇丰写PMsg消息中心服务的部署文档，汇丰要求不能使用明文传输；

解决方案：
配置文件增设pwdtype字段，若pwdtype=1，表示为 密码密文。
加密方式为： 先DES，后Base64

遇到问题：golang加密的密文，在线并不能解析出来

要澄清一个概念：加密的对象不是字符串，而是长度固定的一段二进制字节
*/
import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"log"
	"strconv"
	"strings"
)

var key = []byte("0My@Sql1")
var iv = key

// ZeroPadding 填充函数，跟下面的 PKCS5Padding 补码只是填充的方式不同
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

// ZeroUnPadding 去填充函数
func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

///////////////////////////////////////////////////////////

// PKCS5Padding 明文补码算法
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding 明文减码算法
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// DesEncrypt 加密
func DesEncrypt(text string, key []byte) (string, error) {
	src := []byte(text)

	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	bs := block.BlockSize()
	// 按8个字节对DES加密。如果最后一段字节不足8位，需要对明文数据补码
	//src = PKCS5Padding(src, bs)
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("need a multiple of the blocksize")
	}

	out := make([]byte, len(src))

	/*
		ECB 模式编码（电子密码本）
		即将加密的数据分成若干组，每组的大小个加密密钥长度相同，这样做可以使密文输出完全由明文和密钥决定。
	*/
	// 对明文按照blocksize进行分块加密
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs]) // 加密第一个块，结果保存在 dst 中
		src = src[bs:]
		dst = dst[bs:]
	}

	// CBC 模式编码（加密块链模式）
	//blockMode := cipher.NewCBCEncrypter(block, iv)
	//blockMode.CryptBlocks(out, src)

	fmt.Println(out)
	//fmt.Println(u16To8(out))
	return base64.StdEncoding.EncodeToString(out), nil
}

// DesDecrypt 解密
func DesDecrypt(decrypted string, key []byte) (string, error) {
	src, err := base64.StdEncoding.DecodeString(decrypted)
	//src, err := base64.RawStdEncoding.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	//src, err = base64.RawStdEncoding.DecodeString(decrypted)
	//if err != nil {
	//	return "", err
	//}
	//src := []byte{153, 91, 175, 183, 79, 93, 220, 56, 30, 15, 68, 215, 20, 3, 103, 79}

	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	//out = ZeroUnPadding(out)
	out = PKCS5UnPadding(out)
	return string(out), nil
}

// unicode2utf8 unicode 转 utf-8
func unicode2utf8(s string) string {
	res := []string{""}
	sUnicode := strings.Split(s, "\\u")
	var ctx = ""
	for _, v := range sUnicode {
		var additional = ""
		if len(v) < 1 {
			continue
		}
		if len(v) > 4 {
			rs := []rune(v)
			v = string(rs[:4])
			additional = string(rs[4:])
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			ctx += v
		}
		ctx += fmt.Sprintf("%c", temp)
		ctx += additional
	}
	res = append(res, ctx)
	return strings.Join(res, "")
}

//GBK转utf8的方法
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// unicode形态的[]byte 转 utf8 形态的string
func u16To8(u16 []byte) string {
	log.Println("u16 = ", u16)
	if len(u16)%2 != 0 {
		log.Println("err len", len(u16))
		return ""
	}

	var body bytes.Buffer

	for i := 0; i < len(u16)/2; i++ {
		v := int(u16[2*i]) + int(u16[2*i+1])<<8
		log.Println(int(u16[2*i]), int(u16[2*i+1])<<8)
		log.Println("v = ", v)
		if v <= 127 {

			body.WriteByte(byte(v))
		} else if v <= 2047 {
			a1 := byte(v&63) + 128

			v = v >> 6
			a2 := byte(v&31) + 192
			body.WriteByte(a2)
			body.WriteByte(a1)

		} else if v <= 65535 {
			a1 := byte(v&63) + 128

			v = v >> 6
			a2 := byte(v&63) + 128

			v = v >> 6
			a3 := byte(v&15) + 224
			body.WriteByte(a3)
			body.WriteByte(a2)
			body.WriteByte(a1)
		}
	}
	return string(body.Bytes())
}

func main() {
	//pwd := "upchina2017"
	//pwd := "123"
	//GBK2UTF8
	//keys := ConvertToString("0My@Sql1", "gbk", "utf-8")
	//fmt.Println(keys)
	//key = []byte(keys)
	//
	//res, _ := DesEncrypt(pwd, key)
	//fmt.Println(res)

	cryptPwd := "BVc5sIR3GJkHjE6uFrnqOA=="
	decrypt, _ := DesDecrypt(cryptPwd, key)
	fmt.Println(decrypt)

	//encrypt := "VTJGc2RHVmtYMS9FWnUrTlNjNStTZEpsREhXMDl5NUk0ZmlUWkdxYkVzdz0="
	//rsp, _ := DesDecrypt(encrypt, key)
	//fmt.Println(rsp)

	// 第三方包加密 ECB 模式
	//cipher := goencrypt.NewDESCipher(key, []byte(""), goencrypt.ECBMode, goencrypt.Pkcs7, goencrypt.PrintBase64)
	//txt, err := cipher.DESEncrypt([]byte(pwd))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(txt)

	// 第三方包加密 CBC 模式
	//cipher := goencrypt.NewDESCipher(key, iv, goencrypt.CBCMode, goencrypt.Pkcs7, goencrypt.PrintBase64)
	//txt, err := cipher.DESEncrypt([]byte(pwd))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(txt)

	//uniCode := "\\u0056\\u0054\\u004a\\u0047\\u0063\\u0032\\u0052\\u0048\\u0056\\u006d\\u0074\\u0059\\u004d\\u0053\\u0039\\u0046\\u0057\\u006e\\u0055\\u0072\\u0054\\u006c\\u004e\\u006a\\u004e\\u0053\\u0074\\u0054\\u005a\\u0045\\u0070\\u0073\\u0052\\u0045\\u0068\\u0058\\u004d\\u0044\\u006c\\u0035\\u004e\\u0055\\u006b\\u0030\\u005a\\u006d\\u006c\\u0055\\u0057\\u006b\\u0064\\u0078\\u0059\\u006b\\u0056\\u007a\\u0064\\u007a\\u0030\\u003d"
	//fmt.Println(unicode2utf8(uniCode))

	//cryptPwd := "U2FsdGVkX1/EZu+NSc5+SdJlDHW09y5I4fiTZGqbEsw="
	//decrypt, _ := DesDecrypt(cryptPwd, key)
	//fmt.Println(decrypt)

}
