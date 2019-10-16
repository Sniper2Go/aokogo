package LogicMsg

import (
	"common/Define"
	"common/Log"
	"common/msgProto/MSG_MainModule"
	"common/msgProto/MSG_Server"
	"common/tcpNet"
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
)

func ExternalGatewayMessageCallBack(c net.Conn, mainID uint16, subID uint16, msg proto.Message) {
	Log.FmtPrintf("exec external gateway server message call back: %v, %v.", c.RemoteAddr(), c.LocalAddr())
}

func onServer(session *tcpNet.TcpSession, req *MSG_Server.CS_EnterServer_Req) (succ bool, err error) {
	Log.FmtPrintf("onServer recv: %v.", req.Enter)
	return
}

func onSvrRegister(session *tcpNet.TcpSession, req *MSG_Server.CS_ServerRegister_Req) (succ bool, err error) {
	Log.FmtPrintf("onSvrRegister recv: %v.", req.ServerType)
	var (
		msgfmt string
	)

	session.SrcPoint = Define.ERouteId(req.ServerType)
	session.Push(req.Msgs)
	for _, id := range req.Msgs {
		mainid, subid := tcpNet.DecodeCmd(uint32(id))
		msgfmt += fmt.Sprintf("[mainid: %v, subid: %v]\t", mainid, subid)
	}

	msgfmt += "\n"
	Log.FmtPrintln("message context: ", msgfmt)
	rsp := &MSG_Server.SC_ServerRegister_Rsp{}
	rsp.Ret = MSG_Server.ErrorCode_Success
	session.SendMsg(uint16(req.ServerType),
		uint16(MSG_MainModule.MAINMSG_SERVER),
		uint16(MSG_Server.SUBMSG_SC_ServerRegister),
		rsp)
	return
}

func init() {
	tcpNet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_SERVER), uint16(MSG_Server.SUBMSG_CS_EnterServer), onServer)
	tcpNet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_SERVER), uint16(MSG_Server.SUBMSG_CS_ServerRegister), onSvrRegister)
}
