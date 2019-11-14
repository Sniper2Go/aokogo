/*
Copyright (this) <year> <copyright holders>

"Anti 996" License Version 1.0 (Draft)

Permission is hereby granted to any individual or legal entity
obtaining a copy of this licensed work (including the source code,
documentation and/or related items, hereinafter collectively referred
to as the "licensed work"), free of charge, to deal with the licensed
work for any purpose, including without limitation, the rights to use,
reproduce, modify, prepare derivative works of, distribute, publish
and sublicense the licensed work, subject to the following conditions:

1. The individual or the legal entity must conspicuously display,
without modification, this License and the notice on each redistributed
or derivative copy of the Licensed Work.

2. The individual or the legal entity must strictly comply with all
applicable laws, regulations, rules and standards of the jurisdiction
relating to labor and employment where the individual is physically
located or where the individual was born or naturalized; or where the
legal entity is registered or is operating (whichever is stricter). In
case that the jurisdiction has no such laws, regulations, rules and
standards or its laws, regulations, rules and standards are
unenforceable, the individual or the legal entity are required to
comply with Core International Labor Standards.

3. The individual or the legal entity shall not induce, metaphor or force
its employee(s), whether full-time or part-time, or its independent
contractor(s), in any methods, to agree in oral or written form, to
directly or indirectly restrict, weaken or relinquish his or her
rights or remedies under such laws, regulations, rules and standards
relating to labor and employment as mentioned above, no matter whether
such written or oral agreement are enforceable under the laws of the
said jurisdiction, nor shall such individual or the legal entity
limit, in any methods, the rights of its employee(s) or independent
contractor(s) from reporting or complaining to the copyright holder or
relevant authorities monitoring the compliance of the license about
its violation(s) of the said license.

THE LICENSED WORK IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE COPYRIGHT HOLDER BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN ANY WAY CONNECTION WITH THE
LICENSED WORK OR THE USE OR OTHER DEALINGS IN THE LICENSED WORK.
*/

package tcpNet

import (
	"common/Define"
	"common/Log"
	"common/msgProto/MSG_MainModule"
	"common/msgProto/MSG_Server"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"reflect"
	"sync/atomic"
	"time"

	//"common/S2SMessage"
	"context"
	"sync"

	"github.com/golang/protobuf/proto"
	//. "common/Define"
)

type TcpSession struct {
	host    string
	isAlive bool
	// The net connection.
	conn *net.TCPConn
	// Buffered channel of outbound messages.
	send chan []byte
	// send/recv
	sw  sync.WaitGroup
	ctx context.Context
	// receive message call back
	recvCb MessageCb
	// person offline flag
	off chan *TcpSession
	//message pack
	pack IMessagePack
	//tcp session manager
	Engine ITcpEngine
	// session id
	SessionID uint64
	//Dest point
	SvrType Define.ERouteId
	//src point
	RegPoint Define.ERouteId
	//person StrIdentify
	StrIdentify string
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 4096
	//offline session
	maxOfflineSize = 1024
)

func (this *TcpSession) Connect() {
	if !this.isAlive {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", this.host)
		if err != nil {
			Log.FmtPrintln("session failed: ", err)
			return
		}

		this.conn, err = net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			return
		}

		this.isAlive = true
	}

}

func NewSession(addr string,
	conn *net.TCPConn,
	ctx context.Context,
	SvrType Define.ERouteId,
	newcb MessageCb,
	off chan *TcpSession,
	pack IMessagePack,
	Engine ITcpEngine) *TcpSession {
	return &TcpSession{
		host:    addr,
		conn:    conn,
		send:    make(chan []byte, maxMessageSize),
		isAlive: false,
		ctx:     ctx,
		recvCb:  newcb,
		pack:    pack,
		off:     make(chan *TcpSession, maxOfflineSize),
		Engine:  Engine,
		SvrType: SvrType,
	}
}

func (this *TcpSession) exit(sw *sync.WaitGroup) {
	if this == nil {
		return
	}

	Log.FmtPrintf("session exit, svr: %v, regpoint: %v.", this.SvrType, this.RegPoint)
	this.isAlive = false
	this.off <- this
	//close(this.send)
	this.conn.CloseRead()
	this.conn.CloseWrite()
	this.conn.Close()
	sw.Wait()
}

func (this *TcpSession) SetSendCache(data []byte) {
	this.send <- data
}

func (this *TcpSession) Sendloop(sw *sync.WaitGroup) {
	defer sw.Done()
	defer func() {
		this.exit(sw)
	}()

	for {
		select {
		case <-this.ctx.Done():
			return
		case data := <-this.send:
			if !this.writeMessage(data) {
				return
			}
		}
	}
}

func (this *TcpSession) Recvloop(sw *sync.WaitGroup) {
	defer sw.Done()
	defer func() {
		this.exit(sw)
	}()

	for {
		select {
		case <-this.ctx.Done():
			return
		default:
			if !this.readMessage() {
				return
			}
		}
	}
}

func (this *TcpSession) writeMessage(data []byte) (succ bool) {
	if !this.isAlive || len(data) == 0 {
		return
	}

	this.conn.SetWriteDeadline(time.Now().Add(writeWait))
	//pack message then send.

	//send...
	Log.FmtPrintln("begin send response message to client.")
	_, err := this.conn.Write(data)
	if err != nil {
		Log.FmtPrintln("send data fail, err: ", err)
		return false
	}

	return true
}

func (this *TcpSession) readMessage() (succ bool) {
	this.conn.SetReadDeadline(time.Now().Add(pongWait))
	packLenBuf := make([]byte, EnMessage_NoDataLen)
	readn, err := io.ReadFull(this.conn, packLenBuf)
	if err != nil || readn < EnMessage_NoDataLen {
		if err.Error() == "EOF" {
			succ = true
		} else {
			Log.FmtPrintln("read data fail, err: ", err, readn)
		}
		return
	}

	packlen := binary.LittleEndian.Uint32(packLenBuf[EnMessage_DataPackLen:EnMessage_NoDataLen])
	if packlen > maxMessageSize {
		Log.FmtPrintln("error receiving packLen:", packlen)
		return
	}

	data := make([]byte, EnMessage_NoDataLen+packlen)
	readn, err = io.ReadFull(this.conn, data[EnMessage_NoDataLen:])
	if err != nil || readn < int(packlen) {
		Log.FmtPrintln("error receiving msg, readn:", readn, "packLen:", packlen, "reason:", err)
		return
	}

	//todo: unpack message then read real date.
	copy(data[:EnMessage_NoDataLen], packLenBuf[:])
	_, err = this.pack.UnPackAction(data)
	if err != nil {
		Log.FmtPrintln("unpack action err: ", err)
		return
	}

	route := this.pack.GetRouteID()
	mainID, subID := this.pack.GetMessageID()
	_cmd := EncodeCmd(mainID, subID)
	if mainID == uint16(MSG_MainModule.MAINMSG_SERVER) && Define.ERouteId(route) == Define.ERouteId_ER_ISG {
		this.Push(Define.ERouteId(route), []uint32{_cmd}) //外网关加入内网关session
		succ = true
		return
	}

	if mainID != uint16(MSG_MainModule.MAINMSG_SERVER) {
		msgroute := Define.ERouteId(route)
		if this.SvrType == Define.ERouteId_ER_ESG { //外网关转发路由
			msgroute = Define.ERouteId_ER_ISG
		} else if this.SvrType == Define.ERouteId_ER_ISG { //内网转发路由
			msgroute = Define.ERouteId(route)
		}

		session := this.Engine.GetSessionByType(Define.ERouteId(msgroute))
		if session != nil {
			succ = session.writeMessage(packLenBuf)
			if !session.isAlive {
				this.Engine.RemoveSession(session)
			}
		} else {
			Log.Error("can not find session, route: %v.", msgroute)
		}
	} else {
		succ, err = MessageCallBack(this) //路由消息注册
		if err != nil {
			Log.FmtPrintln("message pack call back: ", err)
		}
	}
	return
}

func (this *TcpSession) readParse(data []byte) (succ bool) {
	Log.FmtPrintf("read parse, RegPoint: %v, SvrType: %v.", this.RegPoint, this.SvrType)
	_, err := this.pack.UnPackAction(data)
	if err != nil {
		Log.FmtPrintln("unpack action err: ", err)
		return
	}

	succ, err = MessageCallBack(this)
	if err != nil {
		Log.FmtPrintln("message pack call back: ", err)
	}

	return
}

func MessageCallBack(session *TcpSession) (succ bool, err error) {
	route := session.pack.GetRouteID()
	Log.FmtPrintf("pack route: %v, sessionid: %v.", route, session.SessionID)
	mainID, subID := session.pack.GetMessageID()
	_cmd := EncodeCmd(mainID, subID)
	Log.FmtPrintf("mainid: %v, subID: %v, cmd: %v.", mainID, subID, _cmd)
	if session.Engine != nil {
		s := session.Engine.GetSessionByCmd(_cmd)
		if s != nil {
			Log.FmtPrintf("route send RegPoint: %v, mainid: %v, subID: %v, sessionid: %v.", s.RegPoint, mainID, subID, s.SessionID)
			s.SetSendCache(session.pack.GetSrcMsg())
			succ = true
			err = nil
			return
		}
	}

	switch mainID {
	case uint16(MSG_MainModule.MAINMSG_SERVER):
		Log.FmtPrintln("server message.")
		if uint16(MSG_Server.SUBMSG_SC_ServerRegister) == subID {
			succ = true
			err = nil
			return
		}
	case uint16(MSG_MainModule.MAINMSG_LOGIN):
		Log.FmtPrintln("Login message.")
	default:

	}

	if session.pack == nil {
		Log.FmtPrintln("pack is nil.")
		return
	}

	msg, cb, unpackerr, exist := session.pack.UnPackData()
	if unpackerr != nil && !exist {
		Log.FmtPrintf("direct send RegPoint: %v, mainid: %v, subID: %v, sessionid: %v.", session.RegPoint, mainID, subID, session.SessionID)
		sendsess := session.Engine.GetSessionByID(session.SessionID)
		if sendsess != nil {
			sendsess.SetSendCache(session.pack.GetSrcMsg())
			succ = true
			err = nil
		}

		return
	}

	if unpackerr != nil {
		err = unpackerr
		Log.FmtPrintln("unpack data err: ", unpackerr)
		return
	}

	params := []reflect.Value{
		//reflect.ValueOf("1"),
		reflect.ValueOf(session),
		reflect.ValueOf(msg),
	}

	ret := cb.Call(params)
	succ = ret[0].Interface().(bool)
	reterr := ret[1].Interface()
	if reterr != nil {
		Log.FmtPrintln("message return err: ", reterr.(error).Error())
	}

	return
}

func (this *TcpSession) HandleSession(sw *sync.WaitGroup) {
	this.isAlive = true
	atomic.AddUint64(&this.SessionID, 1)
	Log.FmtPrintln("handle new session: ", this.SessionID)
	sw.Add(2)
	go this.Recvloop(sw)
	go this.Sendloop(sw)
}

func (this *TcpSession) Push(RegPoint Define.ERouteId, cmds []uint32) {
	if this.Engine == nil {
		return
	}
	Log.FmtPrintf("push new sesson, reg point: %v, cmds: %v.", RegPoint, cmds)
	this.RegPoint = RegPoint
	this.Engine.PushCmdSession(this, cmds)
}

func (this *TcpSession) SetIdentify(StrIdentify string) {
	this.StrIdentify = StrIdentify
}

func (this *TcpSession) Offline() {

}

func (this *TcpSession) SendMsg(route, mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		return false, fmt.Errorf("session disconnection, route: %v, mainid: %v, subid: %v.", route, mainid, subid)
	}

	data := this.pack.PackMsg(route,
		mainid,
		subid,
		msg)
	this.SetSendCache(data)
	return true, nil
}
