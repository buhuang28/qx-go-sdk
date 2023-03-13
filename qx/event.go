package qx

import "C"
import log "github.com/sirupsen/logrus"

var Info func() string
var EventEnable func() int64
var EventDisable func() int64
var EventUninstall func() int64
var EventFriendReq func(string, string) int64
var EventGroupMsg func(string, string) int64
var EventGroupReq func(string, string) int64
var EventPluginIns func(int64) int64
var EventPrivateMsg func(string, string) int64
var EventQrcPay func(string, string) int64
var EventRecall func(string, string) int64
var EventSendMsg func(string, string) int64
var EventTransfer func(string, string) int64
var EventWeChat func(string, string) int64
var Set func()

//export Qx_Info
func Qx_Info() *C.char {
	return CString(Info())
}

//export Qx_EventEnable
func Qx_EventEnable(prt C.int) C.int {
	return cInt(EventEnable())
}

//export Qx_EventDisable
func Qx_EventDisable() C.int {
	return cInt(EventDisable())
}

//export Qx_EventUninstall
func Qx_EventUninstall() C.int {
	return cInt(EventUninstall())
}

//export Qx_EventFriendReq
func Qx_EventFriendReq(botId *C.char, str *C.char) C.int {
	return cInt(EventFriendReq(goString(botId), goString(str)))
}

//export Qx_EventGroupMsg
func Qx_EventGroupMsg(botId *C.char, str *C.char) C.int {
	s := goString(str)
	log.Info("收到群消息:", s)
	return cInt(EventGroupMsg(goString(botId), goString(str)))
}

//export Qx_EventGroupReq
func Qx_EventGroupReq(botId *C.char, str *C.char) C.int {
	return cInt(EventGroupReq(goString(botId), goString(str)))
}

//export Qx_EventPluginIns
func Qx_EventPluginIns(ptr C.int) C.int {
	return cInt(EventPluginIns(goInt(ptr)))
}

//export Qx_EventPrivateMsg
func Qx_EventPrivateMsg(botId *C.char, str *C.char) C.int {
	return cInt(EventPrivateMsg(goString(botId), goString(str)))
}

//export Qx_EventQrcPay
func Qx_EventQrcPay(botId *C.char, str *C.char) C.int {
	return cInt(EventQrcPay(goString(botId), goString(str)))
}

//export Qx_EventRecall
func Qx_EventRecall(botId *C.char, str *C.char) C.int {
	return cInt(EventRecall(goString(botId), goString(str)))
}

//export Qx_EventSendMsg
func Qx_EventSendMsg(botId *C.char, str *C.char) C.int {
	return cInt(EventSendMsg(goString(botId), goString(str)))
}

//export Qx_EventTransfer
func Qx_EventTransfer(botId *C.char, str *C.char) C.int {
	return cInt(EventTransfer(goString(botId), goString(str)))
}

//export Qx_EventWeChat
func Qx_EventWeChat(botId *C.char, str *C.char) C.int {
	return cInt(EventWeChat(goString(botId), goString(str)))
}

//export Qx_Set
func Qx_Set() {
	Set()
	return
}
