package globalver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func InitConfig() {
	getEnv()
	getConfigs()
}

func getEnv() {
	envName := os.Getenv(ENVNAME)
	if len(envName) > 0 {
		RuntimeInfo.EnvName = envName
	} else {
		RuntimeInfo.EnvName = "dev"
	}
}

func getConfigs() {
	data, err := ioutil.ReadFile("./" + APPSETTINGNAME + "." + RuntimeInfo.EnvName + ".json")
	if err != nil {
		fmt.Println("not find" + APPSETTINGNAME + ".")
		return
	}
	if err := json.Unmarshal(data, &Appsetting); err != nil {
		fmt.Println("json Unmarshal " + APPSETTINGNAME + "." + RuntimeInfo.EnvName + ".json err.")
		return
	}

	getRestfulConfig()

	getServerListConfig()
}

func getRestfulConfig() {
	restfulPath := ""
	for _, val := range Appsetting.ConfigurationCenter.Files {
		if val.Name == "restful" {
			restfulPath = val.Path
			break
		}
	}

	if restfulPath == "" {
		return
	}

	data, _ := ioutil.ReadFile("./" + Appsetting.ConfigurationCenter.Local.Folder + "/" + restfulPath)
	json.Unmarshal(data, &RestFulService)
}

func getServerListConfig() {
	serverListPath := ""
	for _, val := range Appsetting.ConfigurationCenter.Files {
		if val.Name == "server-list" {
			serverListPath = val.Path
			break
		}
	}

	if serverListPath == "" {
		return
	}

	data, _ := ioutil.ReadFile("./" + Appsetting.ConfigurationCenter.Local.Folder + "/" + serverListPath)
	json.Unmarshal(data, &ServerList)
}
