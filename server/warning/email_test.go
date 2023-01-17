package warning_test

import (
	"log"
	"strings"
	"testing"

	"github.com/brokercap/Bifrost/server/warning"
)

func TestUserSplit(t *testing.T) {
	p := make(map[string]string, 0)
	p["TO"] = "jc3wish@126.com;"
	tmp := strings.Split(p["TO"], ";")
	log.Println("len:", len(tmp))
	log.Println(tmp[len(tmp)-1])
	tmp = tmp[:len(tmp)-1]
	log.Println("tmp", tmp)
}

func TestSendWarning(t *testing.T) {
	obj := &warning.Email{}

	p := make(map[string]interface{}, 0)

	p["FROM"] = "ops@roboofood.com"
	p["TO"] = "yan.jiang@roboofood.com"
	p["Password"] = "A83VPMiKtdSAM5xh"
	p["SmtpHost"] = "smtp.feishu.cn"
	p["SmtpPort"] = 465
	p["NickName"] = "test nick name"

	err := obj.SendWarning(p, "test warning title", "it is test")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		log.Println("success")
	}
}
