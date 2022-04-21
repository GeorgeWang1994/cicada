package judge

import (
	"cicada/judge/cc"
	"cicada/judge/gg"
	"cicada/pkg/model"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

/**

需求：
1. 在一段时间内都不允许触发告警（利用redis进行存储，利用过期时间）
2. 某个特定的事件都不允许触发告警（同上,可以根据事件的各个字段来决定，比如类型、风险级别等等）
3. 支持自定义处理告警事件（支持自己处理逻辑）
4. 达到某个特定的数字才触发告警（直接判断数据即可）
5. 持续一段时间达到某个特定的数字才触发告警（利用redis的zset数据结构存储，然后利用zrange查询 ？）

额外（通常来说，如果数据量比较大的，一般都是根据统计时间段内的数据，数据量比较小的情况下才会考虑到事件的个数）
1. 只要超过最近n次数据的平均值才触发告警（利用redis的list数据结构存储，然后利用lrange查询出来后计算平均值）
2. 只要超过最近n次数据的最大值才触发告警（利用redis的list数据结构存储，然后利用lrange查询出来后计算最大值）
3. 只要低于最近n次数据的最小值才触发告警（利用redis的list数据结构存储，然后利用lrange查询出来后计算最小值）
4. 最近n次的数据都要超过才触发告警（利用redis的list数据结构存储，然后利用lrange查询出来后比较）
5. 超过最近n次数据的和才触发告警（利用redis的list数据结构存储，然后利用lrange查询出来后计算和）
6. 最近n次中超过m次才触发告警（利用redis的list数据结构存储，然后利用lrange查询出来后比较）
7. 自定义告警

*/

func Judge(e interface{}) error {
	strategies := gg.AlarmStrategy.Get()
	if len(strategies) < 0 {
		return errors.New("not exists strategies")
	}

	for _, s := range strategies {
		if !s.Enable {
			continue
		}
		trigger := judgeEventWithStrategy(s, e)
		if trigger {
			if err := sendEvent(e); err != nil {
				return err
			}
		}
	}

	return nil
}

// 根据策略判断事件是否触发告警
func judgeEventWithStrategy(strategy model.AlarmStrategy, event interface{}) bool {
	var isTriggered bool
	for _, fn := range Functions {
		err := fn.BeforeCompute()
		if err != nil {
			log.Errorf("before compute fn %s with strategy %d failed", fn.Name(), strategy.ID)
		}
		trigger := fn.Compute(strategy, event)
		isTriggered = trigger && isTriggered
		err = fn.AfterCompute()
		if err != nil {
			log.Errorf("after compute fn %s with strategy %d failed", fn.Name(), strategy.ID)
		}
	}
	return isTriggered
}

func sendEvent(event interface{}) error {
	bs, err := json.Marshal(event)
	if err != nil {
		log.Printf("json marshal event %v fail: %v", event, err)
		return err
	}

	// send to redis
	redisKey := fmt.Sprintf(cc.Config().Alarm.QueuePattern)
	rc := gg.RedisConnPool.Get()
	defer rc.Close()

	_, err = rc.Do("LPUSH", redisKey, string(bs))
	if err != nil {
		log.Printf("push event to redis %v fail: %v", event, err)
		return err
	}
	return nil
}
