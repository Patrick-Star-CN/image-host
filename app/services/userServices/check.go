package userServices

func CheckUsername(username string) bool {
	user, _ := GetUserByUsername(username)
	return user != nil
}

func CheckWechatOpenID(wechatOpenID string) bool {
	user := GetUserByWechatOpenID(wechatOpenID)
	return user != nil
}
