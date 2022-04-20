package model

import (
	"fmt"
	"time"
)

type HoneypotEvent struct {
	ID        string    `json:"id"`
	Proto     string    `json:"proto"`
	Honeypot  string    `json:"honeypot"`
	Agent     string    `json:"agent"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	SrcIp     string    `json:"src_ip"`
	SrcPort   int       `json:"src_port"`
	DestIp    string    `json:"dest_ip"`
	DestPort  int       `json:"dest_port"`
	RiskLevel int       `json:"risk_level"`
}

func (p *HoneypotEvent) PK() string {
	return fmt.Sprintf("%s", p.ID)
}

type HoneypotEventRequest struct {
	Events []HoneypotEvent
}
