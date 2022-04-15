package judge

import (
	"cicada/judge/cc"
	"cicada/judge/gg"
	"cicada/pkg/model"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

/**

需求：
1. 过去一段时间内都不允许触发告警（利用redis进行存储，利用过期时间）
2. 某个特定的事件都不允许触发告警（同上）
3. 支持自定义处理告警事件
4. 达到某个特定的数字才触发告警（直接判断数据即可）
5. 持续一段时间达到某个特定的数字才触发告警（利用redis的zset数据结构存储，然后利用zrange查询）

额外
1. 只要超过最近n次数据的平均值才触发告警（利用redis的list数据结构存储，然后利用lrange查询出来后计算平均值）
2. 只要超过最近n次数据的最大值才触发告警（利用redis的list数据结构存储，然后利用lrange查询出来后计算最大值）
3. 只要低于最近n次数据的最小值才触发告警（利用redis的list数据结构存储，然后利用lrange查询出来后计算最小值）
4. 最近n次的数据都要超过才触发告警（利用redis的list数据结构存储，然后利用lrange查询出来后比较）
5. 超过最近n次数据的和才触发告警（利用redis的list数据结构存储，然后利用lrange查询出来后计算和）
6. 最近n次中超过m次才触发告警（利用redis的list数据结构存储，然后利用lrange查询出来后比较）
7. 自定义告警

*/

func Judge(e *model.HoneypotEvent) error {
	if err := sendEvent(e); err != nil {
		return err
	}
	return nil
}

func sendEvent(event *model.HoneypotEvent) error {
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
