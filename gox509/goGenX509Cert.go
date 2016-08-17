package main

import (
	"crypto/x509"
	"fmt"
	"os"
	"time"
)

func main() {
	certCerFile, err := os.Open("oneoaas.cer")
	checkError(err)
	derBytes := make([]byte, 1000) // bigger than the file
	count, err := certCerFile.Read(derBytes)
	checkError(err)
	certCerFile.Close()

	// trim the bytes to actual length in call
	cert, err := x509.ParseCertificate(derBytes[0:count])
	checkError(err)

	fmt.Printf("Name %s\n", cert.Subject.CommonName)
	fmt.Printf("Not before %s\n", cert.NotBefore.String())//开始时间
	fmt.Printf("Not after %s\n", cert.NotAfter.String())//结束时间

	before := cert.NotBefore.Unix()
	fmt.Println(" before :", before)
	if before == time.Now().Unix() {
		println("验证通过")
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}