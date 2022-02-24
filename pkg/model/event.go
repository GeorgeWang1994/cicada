package model

import "time"

type EventBaseInfo struct {
	Proto      string
	Honeypot   string
	Agent      string
	StartTime  time.Time
	EndTime    time.Time
	SrcIp      string
	SrcPort    int
	SrcMac     string
	DestIp     string
	DestPort   int
	EventTypes []string
	RiskLevel  int
}

type EventRequest struct {
	Events []EventBaseInfo
}
