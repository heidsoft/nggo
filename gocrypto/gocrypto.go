package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"crypto/md5"
	"encoding/base64"
	"time"
	"math/rand"
	"crypto/hmac"
	"github.com/dgrijalva/jwt-go"
)

func test() {
	//sha1
	h := sha1.New()
	io.WriteString(h, "aaaaaa")
	fmt.Printf("%x\n", h.Sum(nil))

	//hmac ,use sha1
	key := []byte("123456")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("aaaaaa"))
	fmt.Printf("%x\n", mac.Sum(nil))
}





func mybase64(src string) string {
	b1 := []byte(src)
	return base64.StdEncoding.EncodeToString(b1)
}

//对字符串进行MD5哈希
func CreateAccessKeyId(data string) string {
	t := md5.New();
	io.WriteString(t,data);
	return fmt.Sprintf("%x",t.Sum(nil));
}

//对字符串进行SHA1哈希
func CreateAccessKeySecret(data string) string {
	t := sha1.New();
	io.WriteString(t,data);
	return fmt.Sprintf("%x",t.Sum(nil));
}

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

//创建token
func CreateToke(signedString []byte)(string){
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims["foo"] = "bar"
	//设置过期时间
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, _ := token.SignedString(signedString)

	return tokenString
}

func Auth(authString string) bool {
	token, err := jwt.Parse(authString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors & jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return false
			} else if ve.Errors & (jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return false
			} else {
				// Couldn't handle this token
				return false
			}
		} else {
			// Couldn't handle this token
			return false
		}
	}
	if !token.Valid {
		return false
	}
	return true

}

func CreateSignature(stringToSignature, accessKeySecret string) string {
	// Crypto by HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, []byte(accessKeySecret))
	hmacSha1.Write([]byte(stringToSignature))
	sign := hmacSha1.Sum(nil)

	// Encode to Base64
	base64Sign := base64.StdEncoding.EncodeToString(sign)

	return base64Sign
}

func main(){
	var data string = RandomString(5);
	fmt.Printf("apiSecret : %s\n",CreateAccessKeySecret(data))
	fmt.Printf("apiKey : %s\n",CreateAccessKeyId(data))

	fmt.Println(CreateSignature("aa","bb"))
}

