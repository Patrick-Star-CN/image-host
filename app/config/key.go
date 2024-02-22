package config

const webpUrlKey = "webpUrlKey"

func GetWebpUrlKey() string {
	return getConfig(webpUrlKey)
}

func SetWebpUrlKey(url string) error {
	return setConfig(webpUrlKey, url)
}

func DelWebpUrlKey() error {
	return delConfig(webpUrlKey)
}
