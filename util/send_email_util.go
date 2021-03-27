package util

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/utils"
)

func SendEmail(to_email, msg string) {
	username := "2028427172@qq.com" // 发送者的邮箱地址
	password := "apcjaxcbmdcddeec"  // 授权密码
	host := "smtp.qq.com"           // 邮件协议
	port := "587"                   // 端口号

	emailConfig := fmt.Sprintf(`{"username":"%s","password":"%s","host":"%s","port":%s}`, username, password, host, port)
	fmt.Println("emailConfig", emailConfig)
	emailConn := utils.NewEMail(emailConfig) // beego下的
	emailConn.From = strings.TrimSpace(username)
	emailConn.To = []string{strings.TrimSpace(to_email)}
	emailConn.Subject = "知了传课邮箱验证码"
	//注意这里我们发送给用户的是激活请求地址
	emailConn.Text = msg
	err := emailConn.Send()
	fmt.Println(err)

}
