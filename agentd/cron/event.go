package cron

import (
	"github.com/GeorgeWang1994/cicada/agentd/cc"
	"github.com/GeorgeWang1994/cicada/agentd/gg"
	"github.com/GeorgeWang1994/cicada/pkg/model"
	"github.com/google/uuid"
	"log"
	"time"
)

// SendEvent 不间断发送数据
func SendEvent() {
	var err error
	duration := time.Duration(cc.Config().Interval) * time.Second

	for {
		time.Sleep(duration)

		events := make([]model.HoneypotEvent, 0)
		perCnt := cc.Config().PerCount
		if perCnt == 0 {
			perCnt = gg.DefaultPerCount
		}
		for i := 0; i < perCnt; i++ {
			events = append(events, mockEventInfo())
		}
		req := model.HoneypotEventRequest{
			Events: events,
		}

		var resp model.RpcResponse
		err = gg.EventRpcClient().Call("HoneypotEvent.ReceiveEvent", req, &resp)
		if err != nil {
			log.Println("call HoneypotEvent.RecvEvent fail:", err)
			continue
		}
		log.Println("send event success")
	}
}

func mockEventInfo() model.HoneypotEvent {
	return model.HoneypotEvent{
		Proto:     "tcp",
		Honeypot:  uuid.New().String(),
		Agent:     cc.Config().ID,
		StartTime: time.Now(),
		EndTime:   time.Now(),
		SrcIp:     "127.0.0.1",
		SrcPort:   10000,
		DestIp:    "127.0.0.1",
		DestPort:  20000,
		RiskLevel: 1,
	}
}
