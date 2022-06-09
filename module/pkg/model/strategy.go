package model

type AlarmStrategy struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`       // 自定义规则名称
	Type      string  `json:"types"`      // 支持具体的告警类型，空表示没有设置过告警类型
	Enable    bool    `json:"enable"`     // 是否开启策略
	Risk      int     `json:"risk"`       // 最低风险级别，0表示没有设置过告警级别
	Internal  int64   `json:"internal"`   // 告警事件间隔，0表示没有设置时间间隔
	Value     float64 `json:"value"`      // 告警阈值，0表示没有设置时间阈值
	Step      int64   `json:"step"`       // 告警次数，0表示没有设置最低的告警次数
	Priority  int64   `json:"priority"`   // 优先级别，0表示没有区分优先
	StartTime int64   `json:"start_time"` // 策略生效的起始时间，0表示没有设置起始时间
	EndTime   int64   `json:"end_time"`   // 策略生效的结束时间，0表示没有设置结束时间
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
	Risk            int                 `json:"risk"`              // 最低风险级别
	Internal        int64               `json:"internal"`          // 告警事件间隔
	Value           float64             `json:"value"`             // 告警阈值
	Step            int64               `json:"step"`              // 告警次数
	Priority        int64               `json:"priority"`          // 优先级别
	StartTime       int64               `json:"start_time"`        // 策略起始时间
	EndTime         int64               `json:"end_time"`          // 策略结束时间
	Receives        []map[string]string `json:"receives"`          // 接受者数据
}
