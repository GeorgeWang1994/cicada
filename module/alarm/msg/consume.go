package msg

import (
	"encoding/json"
	"github.com/GeorgeWang1994/cicada/module/alarm/cc"
	"github.com/GeorgeWang1994/cicada/module/alarm/gg"
	"github.com/GeorgeWang1994/cicada/module/pkg/model"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"time"
)

type Message struct {
	Receiver []string
	Content  string
}

type MsgHandler interface {
	GenerateIMContent(event *model.HoneypotEvent) string
	Send(msg *Message) error
}

func Consume() {
	queues := cc.Config().Redis.Queues
	if len(queues) == 0 {
		return
	}

	// todo: 提高读取的并发能力
	for {
		event, err := popEvent(queues)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		// todo: 获取用户列表
		var users []model.User
		ParseUserIm(event, users)
	}
}

func popEvent(queues []string) (*model.HoneypotEvent, error) {
	count := len(queues)

	params := make([]interface{}, count+1)
	for i := 0; i < count; i++ {
		params[i] = queues[i]
	}
	// set timeout 0
	params[count] = 0

	rc := gg.RedisConnPool().Get()

	// 从多个队列中读取单个事件
	reply, err := redis.Strings(rc.Do("BRPOP", params...))
	if err != nil {
		log.Errorf("get alarm event from redis fail: %v", err)
		return nil, err
	}

	var event model.HoneypotEvent
	err = json.Unmarshal([]byte(reply[1]), &event)
	if err != nil {
		log.Errorf("parse alarm event fail: %v", err)
		return nil, err
	}

	return &event, nil
}

func ParseUserIm(event *model.HoneypotEvent, users []model.User) {
	var imH ImHandler
	content := imH.GenerateIMContent(event)
	var userIDs []string
	for _, user := range users {
		userIDs = append(userIDs, user.ID)
	}
	msg := &Message{
		Receiver: userIDs,
		Content:  content,
	}
	err := imH.Send(msg)
	if err != nil {
		log.Errorf("send im fail, im:%v, error:%v", *msg, err)
	}
}
