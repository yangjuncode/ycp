package ycp

import (
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type FnOnYCPData func(pkt []byte)
type TYCPClient struct {
	sync.Mutex

	TCPCon *net.TCPConn
	UDPCon *net.UDPConn

	//收到udp 数据
	OnUDPData FnOnYCPData
	//收到ycp 数据
	OnYCPData FnOnYCPData
	//收到tcp 数据
	OnTCPData FnOnYCPData

	LastYCPSendTime time.Time
	LastYCPRecvTime time.Time
	LastTCPSendTime time.Time
	LastTCPRecvTime time.Time

	pktNo uint32
}

func NewYCPClient(host string, port uint16) (ycpCli *TYCPClient) {
	nowTime := time.Now()
	ycpCli = &TYCPClient{
		TCPCon:          nil,
		UDPCon:          nil,
		OnUDPData:       nil,
		OnYCPData:       nil,
		OnTCPData:       nil,
		LastYCPSendTime: nowTime,
		LastYCPRecvTime: nowTime,
		LastTCPSendTime: nowTime,
		LastTCPRecvTime: nowTime,
	}
	return
}

func (this *TYCPClient) WriteUDP(pkt []byte) (int, error) {
	return this.UDPCon.Write(pkt)
}

func (this *TYCPClient) WriteYCP(pkt []byte, ordered bool) {

}
func (this *TYCPClient) WriteYCPTimeout(pkt []byte, deadlineWait time.Duration, ordered bool) {

}
func (this *TYCPClient) WriteTCP(pkt []byte) {

}

func (this *TYCPClient) NewPktNo() uint32 {
	return atomic.AddUint32(&this.pktNo, 1)
}
