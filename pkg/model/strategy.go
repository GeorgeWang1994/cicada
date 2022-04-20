package model

type AlarmStrategy struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`       // 自定义规则名称
	Type      string  `json:"type"`       // 告警类型
	Enable    bool    `json:"enable"`     // 是否开启策略
	Risk      int64   `json:"risk"`       // 最低风险级别
	Internal  int64   `json:"internal"`   // 告警事件间隔
	Value     float64 `json:"value"`      // 最大告警数目
	MaxStep   int64   `json:"maxStep"`    // 最大告警次数
	Priority  int64   `json:"priority"`   // 优先级别
	StartTime int64   `json:"start_time"` // 策略起始时间
	EndTime   int64   `json:"end_time"`   // 策略结束时间
	Note      string  `json:"note"`       // 描述信息
}

type StrategiesResponse struct {
	Strategies []AlarmStrategy `json:"strategies"`
}

type SubscribeStrategy struct {
	ID              int64               `json:"id"`
	AlarmStrategyID int64               `json:"alarm_strategy_id"` // 默认绑定策略
	Name            string              `json:"name"`              // 自定义规则名称
	Type            string              `json:"type"`              // 订阅类型，包括邮件、短信、IM等等
	Enable          bool                `json:"enable"`            // 是否开启策略
	Risk            int64               `json:"risk"`              // 最低风险级别
	Internal        int64               `json:"internal"`          // 告警事件间隔
	Value           float64             `json:"value"`             // 最大告警数目
	MaxStep         int64               `json:"maxStep"`           // 最大告警次数
	Priority        int64               `json:"priority"`          // 优先级别
	Receives        []map[string]string `json:"receives"`          // 接受者数据
}
