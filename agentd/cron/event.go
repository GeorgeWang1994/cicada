package cron

import (
	"cicada/agentd/cc"
	"cicada/agentd/gg"
	"cicada/pkg/model"
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

		events := make([]model.EventBaseInfo, 0)
		perCnt := cc.Config().PerCount
		if perCnt == 0 {
			perCnt = gg.DefaultPerCount
		}
		for i := 0; i < perCnt; i++ {
			events = append(events, mockEventInfo())
		}
		req := model.EventRequest{
			Events: events,
		}

		var resp model.RpcResponse
		err = gg.EventRpcClient().Call("Event.RecvEvent", req, &resp)
		if err != nil {
			log.Println("call Event.RecvEvent fail:", err)
			continue
		}
		log.Println("send event success")
	}
}

func mockEventInfo() model.EventBaseInfo {
	return model.EventBaseInfo{
		Proto:      "tcp",
		Honeypot:   uuid.New().String(),
		Agent:      cc.Config().ID,
		StartTime:  time.Now(),
		EndTime:    time.Now(),
		SrcIp:      "127.0.0.1",
		SrcPort:    10000,
		SrcMac:     "aa:bb:cc:dd",
		DestIp:     "127.0.0.1",
		DestPort:   20000,
		EventTypes: []string{"password_login", "web_attacker"},
		RiskLevel:  1,
	}
}
