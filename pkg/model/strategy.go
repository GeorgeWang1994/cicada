package model

type Strategy struct {
	Id       int64   `json:"id"`
	Type     string  `json:"type"`     // 告警类型
	Enable   bool    `json:"enable"`   // 是否开启策略
	Internal int64   `json:"internal"` // 告警事件间隔
	Value    float64 `json:"value"`    // 最大告警数目
	MaxStep  int64   `json:"maxStep"`  // 最大告警次数
	Priority int     `json:"priority"` // 优先级别
	Note     string  `json:"note"`     // 描述信息
}

type HostStrategy struct {
	Hostname   string     `json:"hostname"`
	Strategies []Strategy `json:"strategies"`
}

type HostStrategiesResponse struct {
	HostStrategies []*HostStrategy `json:"hostStrategies"`
}
