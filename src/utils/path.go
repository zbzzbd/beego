package utils

import "os"

func EnsurePath(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, 0744)
	}
	if err != nil {
		GetLog().Error("helpers.MakeSurePath: MkdirAll error=%v, path=%v", err, path)
	}

	return err
}

func BaseUrl() string {
	baseUrl := GetConf().String("app::base_url")
	httpport := GetConf().String("app::httpport")
	if httpport != "80" {
		baseUrl += ":" + httpport
	}
	return baseUrl
}
