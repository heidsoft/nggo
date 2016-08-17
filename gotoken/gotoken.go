package main

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
)
//HMAC是密钥相关的哈希运算消息认证码（Hash-based Message Authentication Code）,
//HMAC运算利用哈希算法，以一个密钥和一个消息为输入，生成一个消息摘要作为输出。
//
//JSON 网络令牌是什么？
//JSON 网络令牌（JWT）是一个带有某些数量要求的令牌创建标准。
//比如，一台服务器能够生成一个要求是“以管理员身份登录”并提供给客户端。
//客户端随后可利用这个令牌来证明他们确实以管理员身份登录。
//这些令牌被服务器的密钥签署，因此服务器可验证这个令牌是合法的。
//JWT一般由三部分组成：一个头信息、一个有效负载、以及一个签名。
//头信息识别生成签名的算法，并查看如下一些信息：
func main() {

	token := jwt.New(jwt.GetSigningMethod(jwt.SigningMethodHS256.Alg()))

	token.SignedString("123456")
	fmt.Println(token.Signature)
	fmt.Println(token.Raw)

}