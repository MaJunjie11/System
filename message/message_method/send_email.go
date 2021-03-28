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
		err     error
	)
	email = s.Req.GetEmail()

	//如果不是邮件
	if isEmail = util.VerifyEmail(email); !isEmail {
		s.Resp.Status = pb_gen.ErrNo_Email_ForMat_Error
		s.Resp.Msg = "请输入正确的邮箱"
	}

	// 如果是个合理的邮件地址 发送邮件 并且存入本地缓存中一分钟 等待后续校验
	rand_num := util.GenRandNum(6)
	if err = util.SendEmail(email, rand_num); err != nil {
		s.Resp.Status = pb_gen.ErrNo_Internal_Err
		s.Resp.Msg = "发送邮件失败"
	}
	LocalCache.Set(email, rand_num, cache.DefaultExpiration)
	s.Resp.Status = pb_gen.ErrNo_Success
	s.Resp.Msg = "发送邮件成功啦"
}
