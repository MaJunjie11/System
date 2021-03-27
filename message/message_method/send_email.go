package message_method

import (
	"System/pb_gen"
	"System/util"
	"time"

	"github.com/patrickmn/go-cache"
)

type SendEmailHandler struct {
	Req  *pb_gen.SendEmailRequest
	Resp *pb_gen.SendEmailResponse
}

var LocalCache = cache.New(time.Second*60, time.Second*10)

func NewSendEmailHandler(req *pb_gen.SendEmailRequest, resp *pb_gen.SendEmailResponse) *SendEmailHandler {
	return &SendEmailHandler{
		Req:  req,
		Resp: resp,
	}
}

func (s *SendEmailHandler) Run() {
	var (
		isEmail bool
		email   string
	)
	email = s.Req.GetEmail()

	//如果不是邮件
	if isEmail = util.VerifyEmail(email); !isEmail {
		//TODO: 后续定义一些错误码
		s.Resp.Status = -1
		s.Resp.Msg = "请输入正确的邮箱"
	}

	// 如果是个合理的邮件地址 发送邮件 并且存入本地缓存中一分钟 等待后续校验
	rand_num := util.GenRandNum(6)
	util.SendEmail(email, rand_num)
	LocalCache.Set(email, rand_num, cache.DefaultExpiration)
	s.Resp.Status = 0
	s.Resp.Msg = "发送成功"
}
