package repository

import (
	"ec-notification-management/src/core/globalver"
	"ec-notification-management/src/core/models/configs"
	"errors"
	"io/ioutil"
	"net/http"
)

type RestFulService struct {
}

func (restfulService *RestFulService) GetMsgCount(customerNumber string, configName string) (string, error) {
	var resourceInfo configs.Resource = configs.Resource{}
	for _, val := range globalver.RestFulService.Resources {
		if val.Resource.Name == configName {
			resourceInfo = val.Resource
			break
		}
	}
	if len(resourceInfo.Name) <= 0 {
		return "", errors.New("rest config:\"" + configName + "\" not exists.")
	}

	serverInfo := configs.Server{}
	for _, val := range globalver.ServerList.Servers {
		if val.Name == resourceInfo.Host {
			serverInfo = val
			break
		}
	}
	if len(serverInfo.Name) <= 0 {
		return "", errors.New("list config:\"" + resourceInfo.Name + "\" not exists.")
	}

	hostInfo := configs.Host{}
	for _, val := range serverInfo.HostList {
		if val.Channel == globalver.RuntimeInfo.EnvName {
			hostInfo = val
			break
		}
	}
	if len(hostInfo.Host) <= 0 {
		return "", errors.New("host config:\"" + serverInfo.Name + "\" not exists.")
	}

	if resourceInfo.Method == "GET" {
		url := hostInfo.Host + resourceInfo.Path
		resp, err := http.Get(url)
		if err != nil {
			return "", errors.New("get url \"" + url + "\" err. " + err.Error())
		}

		defer resp.Body.Close()

		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", errors.New("get url \"" + url + "\" err. " + err.Error())
		}

		return string(buf), nil
	}

	return "", nil
}
