package db

import (
	"cicada/ev/gg"
	"cicada/pkg/model"
	"context"
	"fmt"
)

type QueryHoneypotEventParam struct {
	Limit  int32
	Offset int32
}

// BatchInsertHoneypotEvent 批量添加事件
func BatchInsertHoneypotEvent(ctx context.Context, events []*model.HoneypotEvent) error {
	batch, err := gg.ClickhouseClient.PrepareBatch(ctx, "INSERT INTO honeypot_event")
	if err != nil {
		return err
	}
	for _, e := range events {
		err := batch.Append(
			e.ID, e.Proto, e.Honeypot, e.Agent, e.StartTime, e.EndTime, e.SrcIp,
			e.SrcPort, e.DestIp, e.DestPort, e.RiskLevel,
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
				"(%s, %s, %s, %s, %d, %d, %s, %d, %s, %d, %d)",
				e.ID, e.Proto, e.Honeypot, e.Agent, e.StartTime.Second(), e.EndTime.Second(),
				e.SrcIp, e.SrcPort, e.DestIp, e.DestPort, e.RiskLevel,
			),
			wait)
		if err != nil {
			return err
		}
	}
	return nil
}

// QueryHoneypotEvents 查询事件列表
func QueryHoneypotEvents(ctx context.Context, param QueryHoneypotEventParam) (events []model.HoneypotEvent, err error) {
	if param.Offset < 0 {
		param.Offset = 0
	}
	if param.Limit <= 0 {
		param.Limit = 20
	}
	start, end := param.Offset, param.Offset+param.Limit
	if err = gg.ClickhouseClient.Select(ctx,
		&events,
		"SELECT * FROM honeypot_event LIMIT $1, $2", start, end,
	); err != nil {
		return
	}
	return
}

// GetHoneypotEvent 查询事件
func GetHoneypotEvent(ctx context.Context, eventID string) (event model.HoneypotEvent, err error) {
	if err = gg.ClickhouseClient.Select(ctx,
		&event,
		"SELECT * FROM honeypot_event WHERE id=$1 LIMIT 1",
		eventID,
	); err != nil {
		return
	}
	return
}
