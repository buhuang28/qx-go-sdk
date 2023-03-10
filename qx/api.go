package qx

import "C"
import "syscall"

var (
	api                 = syscall.NewLazyDLL("QxWeChatApi.DLL")
	API_GetGroups       = getApiProc("Api_GetGroupList")
	API_GetGroupMembers = getApiProc("Api_GetGroupMemberList")
	API_GetSelfInfo     = getApiProc("Api_GetSelfInfo")
	API_GetBots         = getApiProc("Api_GetWeChatList")
	API_SendImg         = getApiProc("Api_SendImage")
	API_SendText        = getApiProc("Api_SendTextMsg")
)

func getApiProc(name string) *syscall.LazyProc {
	return api.NewProc(name)
}

func GetGroups(botId, getType string) string {
	r, _, _ := API_GetGroups.Call(str2ptr(botId), str2ptr(getType))
	return ptr2str(r)
}

func GetMembers(botId, groupId string) string {
	r, _, _ := API_GetGroupMembers.Call(str2ptr(botId), str2ptr(groupId))
	return ptr2str(r)
}

func GetSelfInfo(botId string) string {
	r, _, _ := API_GetSelfInfo.Call(str2ptr(botId))
	return ptr2str(r)
}

func GetBots() string {
	r, _, _ := API_GetBots.Call()
	return ptr2str(r)
}

func SendImg(botId, toId, img string) string {
	r, _, _ := API_SendImg.Call(str2ptr(botId), str2ptr(toId), str2ptr(img))
	return ptr2str(r)
}

func SendText(botId, toId, msg string) string {
	r, _, _ := API_SendText.Call(str2ptr(botId), str2ptr(toId), str2ptr(msg))
	return ptr2str(r)
}
