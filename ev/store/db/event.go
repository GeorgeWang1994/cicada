package db

import (
	"cicada/ev/gg"
	"cicada/pkg/model"
	"context"
	"fmt"
)

// BatchInsertHoneypotEvent 批量添加事件
func BatchInsertHoneypotEvent(ctx context.Context, events []*model.HoneypotEvent) error {
	batch, err := gg.ClickhouseClient.PrepareBatch(ctx, "INSERT INTO honeypot_event")
	if err != nil {
		return err
	}
	for _, e := range events {
		err := batch.Append(
			e.ID, e.Proto, e.Honeypot, e.Agent, e.StartTime, e.EndTime, e.SrcIp,
			e.SrcPort, e.SrcMac, e.DestIp, e.DestPort, e.EventTypes, e.RiskLevel,
		)
		if err != nil {
			return err
		}
	}
	err = batch.Send()
	if err != nil {
		return err
	}
	return nil
}

// AsyncBatchInsertHoneypotEvent 异步批量写入
func AsyncBatchInsertHoneypotEvent(ctx context.Context, events []*model.HoneypotEvent, wait bool) error {
	for _, e := range events {
		err := gg.ClickhouseClient.AsyncInsert(
			ctx,
			fmt.Sprintf("INSERT INTO honeypot_event VALUES "+
				"(%s, %s, %s, %s, %d, %d, %s, %d, %s, %s, %d, %s, %d)",
				e.ID, e.Proto, e.Honeypot, e.Agent, e.StartTime.Second(), e.EndTime.Second(),
				e.SrcIp, e.SrcPort, e.SrcMac, e.DestIp, e.DestPort, e.EventTypes, e.RiskLevel,
			),
			wait)
		if err != nil {
			return err
		}
	}
	return nil
}

// QueryHoneypotEvents 查询事件列表
func QueryHoneypotEvents(ctx context.Context) (events []model.HoneypotEvent, err error) {
	if err = gg.ClickhouseClient.Select(ctx,
		&events,
		"SELECT "+
			"id, proto, honeypot, agent, start_time, end_time, src_ip, src_port, "+
			"src_mac, dest_ip, dest_port, event_types, risk_level"+
			"FROM honeypot_event",
	); err != nil {
		return
	}
	return
}
