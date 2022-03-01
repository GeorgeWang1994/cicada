package model

import "time"

type HoneypotEvent struct {
	ID         string    `json:"id"`
	Proto      string    `json:"proto"`
	Honeypot   string    `json:"honeypot"`
	Agent      string    `json:"agent"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	SrcIp      string    `json:"src_ip"`
	SrcPort    int       `json:"src_port"`
	SrcMac     string    `json:"src_mac"`
	DestIp     string    `json:"dest_ip"`
	DestPort   int       `json:"dest_port"`
	EventTypes []string  `json:"event_types"`
	RiskLevel  int       `json:"risk_level"`
}

type HoneypotEventRequest struct {
	Events []HoneypotEvent
}
