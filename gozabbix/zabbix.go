package main

import (
	"oneoaas_cmdb/zabbix"
	"fmt"
)

func main() {

	api, err := zabbix.NewAPI("http://10.0.1.114/api_jsonrpc.php", "admin", "zabbix")
	versionresult, err := api.Version()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(versionresult)

	_, err = api.Login()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to API")

	host, err := GetHost(api, "Zabbix server")
	fmt.Println("host", host)
}

func GetHost(api *zabbix.API, host string) (zabbix.ZabbixHost, error) {
	params := make(map[string]interface{}, 0)
	filter := make(map[string]string, 0)
	filter["host"] = host
	params["filter"] = filter
	params["output"] = "extend"
	params["select_groups"] = "extend"
	params["templated_hosts"] = 1
	ret, err := api.Host("get", params)

	// This happens if there was an RPC error
	if err != nil {
		return nil, err
	}

	// If our call was successful
	if len(ret) > 0 {
		return ret[0], err
	}

	// This will be the case if the RPC call was successful, but
	// Zabbix had an issue with the data we passed.
	return nil, &zabbix.ZabbixError{0, "", "Host not found"}
}