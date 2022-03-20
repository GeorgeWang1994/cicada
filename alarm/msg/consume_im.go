package msg

import (
	"cicada/alarm/cc"
	"cicada/pkg/model"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/toolkits/net/httplib"
	"strings"
	"time"
)

type ImHandler struct{}

func (h *ImHandler) GenerateIMContent(event *model.HoneypotEvent) string {
	return fmt.Sprintf("收到攻击蜜罐%s的探针%s的告警信息", event.Honeypot, event.Agent)
}

func (h *ImHandler) Send(msg *Message) error {
	url := cc.Config().Provider.IM
	r := httplib.Post(url).SetTimeout(5*time.Second, 30*time.Second)
	r.Param("tos", strings.Join(msg.Receiver, ","))
	r.Param("content", msg.Content)
	resp, err := r.String()
	if err != nil {
		return err
	}
	log.Debugf("send im:%v, resp:%v, url:%s", *msg, resp, url)
	return nil
}
