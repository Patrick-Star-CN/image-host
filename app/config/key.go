package config

const webpUrlKey = "webpUrlKey"
const pathKey = "pathKey"

func GetWebpUrlKey() string {
	return getConfig(webpUrlKey)
}

func SetWebpUrlKey(url string) error {
	return setConfig(webpUrlKey, url)
}

func GetPathKey() string {
	return getConfig(pathKey)
}

func SetPathKey(url string) error {
	return setConfig(pathKey, url)
}
