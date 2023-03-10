package qx

import "C"
import log "github.com/sirupsen/logrus"

var Info func() string

//export Qx_Info
func Qx_Info() *C.char {
	return CString(Info())
}

//export Qx_EventEnable
func Qx_EventEnable(prt C.int) C.int {
	return cInt(EVENT_AGREE)
}

//export Qx_EventDisable
func Qx_EventDisable() C.int {
	return cInt(EVENT_AGREE)
}

//export Qx_EventUninstall
func Qx_EventUninstall() C.int {
	return cInt(EVENT_AGREE)
}

//export Qx_EventFriendReq
func Qx_EventFriendReq(botId *C.char, str *C.char) C.int {
	return cInt(EVENT_CONTINUE)
}

//export Qx_EventGroupMsg
func Qx_EventGroupMsg(botId *C.char, str *C.char) C.int {
	s := goString(str)
	log.Info("收到群消息:", s)
	return cInt(EVENT_CONTINUE)
}

//export Qx_EventGroupReq
func Qx_EventGroupReq(botId *C.char, str *C.char) C.int {
	return cInt(EVENT_CONTINUE)
}

//export Qx_EventPluginIns
func Qx_EventPluginIns(ptr *C.int) C.int {
	return cInt(EVENT_AGREE)
}

//export Qx_EventPrivateMsg
func Qx_EventPrivateMsg(botId *C.char, str *C.char) C.int {
	s := goString(str)
	log.Info("收到私聊消息:", s)
	return cInt(EVENT_CONTINUE)
}

//export Qx_EventQrcPay
func Qx_EventQrcPay(botId *C.char, str *C.char) C.int {
	return cInt(EVENT_CONTINUE)
}

//export Qx_EventRecall
func Qx_EventRecall(botId *C.char, str *C.char) C.int {
	return cInt(EVENT_CONTINUE)
}

//export Qx_EventSendMsg
func Qx_EventSendMsg(botId *C.char, str *C.char) C.int {
	return cInt(EVENT_CONTINUE)
}

//export Qx_EventTransfer
func Qx_EventTransfer(botId *C.char, str *C.char) C.int {
	return cInt(EVENT_CONTINUE)
}

//export Qx_EventWeChat
func Qx_EventWeChat(botId *C.char, str *C.char) C.int {
	return cInt(EVENT_CONTINUE)
}

//export Qx_Set
func Qx_Set(botId *C.char, str *C.char) C.int {
	return cInt(EVENT_CONTINUE)
}
