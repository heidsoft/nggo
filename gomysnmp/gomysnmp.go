package main

import (
	"github.com/alouca/gosnmp"
	"log"
)

// snmpwalk -v 2c 10.0.1.130 -c public system
// cisco 10.0.2.1
func main() {
	s, err := gosnmp.NewGoSNMP("10.0.2.1", "oneoaas", gosnmp.Version2c, 5)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := s.Get(".1.3.6.1.2.1.1.1.0")
	if err == nil {
		for _, v := range resp.Variables {
			switch v.Type {
			case gosnmp.OctetString:
				log.Printf("Response: %s : %s : %s \n", v.Name, v.Value.(string), v.Type.String())
			}
		}
	}

}