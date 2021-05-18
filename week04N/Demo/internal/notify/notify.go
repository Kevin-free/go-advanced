package notify

import (
	pb "biohouse/api"
	"git.huoys.com/middle-end/library/pkg/net/comet"
	"github.com/gogo/protobuf/proto"
)

type Notify interface {
	//AsyncPush 非阻塞式PUSH
    AsyncPush(sessionID string, ops pb.GameCommand, msg proto.Message)
    //SyncPush 阻塞式PUSH
    SyncPush(sessionID string, ops pb.GameCommand, msg proto.Message)
	Close(sessionID string)
}

type notify struct {
	pushChan  chan *comet.PushData
	closeChan chan string
	s *comet.Server
}

func New(pc chan *comet.PushData, cc chan string, cs *comet.Server) (n Notify) {
	n = &notify{
		pushChan:  pc,
		closeChan: cc,
		s: cs,
	}
	return
}

func (n *notify) Close(sessionID string) {
	n.closeChan <- sessionID
}

func (n *notify) AsyncPush(sessionID string, ops pb.GameCommand, msg proto.Message) {
	data, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	p := &comet.PushData{
		Mid:  sessionID,
		Data: data,
		Ops:  int32(ops),
	}
	n.pushChan <- p
}

func (n *notify) SyncPush(sessionID string, ops pb.GameCommand, msg proto.Message) {
	data, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	n.s.PushByChannel(sessionID, int32(ops), data)
}
