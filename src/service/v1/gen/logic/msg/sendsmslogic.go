package msg

import (
	"context"
	"gitee.com/wuntsong-auth/backend/src/model/db"
	"gitee.com/wuntsong-auth/backend/src/mysql"
	utils2 "gitee.com/wuntsong-auth/backend/src/service/v1/gen/logic/utils"
	"gitee.com/wuntsong-auth/backend/src/service/v1/src/respmsg"
	"gitee.com/wuntsong-auth/backend/src/service/v1/src/warp"
	"gitee.com/wuntsong-auth/backend/src/sms"
	errors "github.com/wuntsong-org/wterrors"
	"net/http"

	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/svc"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/types"

	"github.com/wuntsong-org/go-zero-plus/core/logx"
)

type SendSMSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSMSLogic {
	return &SendSMSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSMSLogic) SendSMS(req *types.SendSMSReq, r *http.Request) (resp *types.SendMsgResp, err error) {
	phoneModel := db.NewPhoneModel(mysql.MySQLConn)

	user, err := utils2.FindUser(l.ctx, req.UserID, false)
	if errors.Is(err, utils2.UserNotFound) {
		return &types.SendMsgResp{
			Resp: respmsg.GetRespByMsg(l.ctx, respmsg.UserNotFound, "用户未找到"),
		}, nil
	} else if err != nil {
		return nil, respmsg.MySQLSystemError.WarpQuick(err)
	}

	web, ok := l.ctx.Value("X-Src-Website").(warp.Website)
	if !ok {
		return nil, respmsg.BadContextError.New("X-Src-Website")
	}

	bannedModel := db.NewOauth2BanedModel(mysql.MySQLConn)
	allow, err := bannedModel.CheckAllow(r.Context(), user.Id, web.ID, db.AllowMsg)
	if err != nil || !allow {
		return &types.SendMsgResp{
			Resp: respmsg.GetRespByMsg(l.ctx, respmsg.WebsiteNotAllow, "用户关闭了通信授权许可"),
		}, nil
	}

	p, err := phoneModel.FindByUserID(l.ctx, user.Id)
	if errors.Is(err, db.ErrNotFound) {
		return &types.SendMsgResp{
			Resp: respmsg.GetRespSuccess(l.ctx),
			Data: types.SendMsgData{
				Success: false,
				Have:    false,
			},
		}, nil
	} else if err != nil {
		return nil, respmsg.MySQLSystemError.WarpQuick(err)
	}

	data := make(map[string]string, len(req.Data))
	for _, v := range req.Data {
		data[v.Label] = v.Value
	}

	err = sms.Send(l.ctx, data, req.Sig, req.Template, p.Phone, web.ID)
	if err != nil {
		return &types.SendMsgResp{
			Resp: respmsg.GetRespByError(l.ctx, respmsg.Success, errors.WarpQuick(err), "发送短信失败"),
			Data: types.SendMsgData{
				Success: false,
				Have:    true,
			},
		}, nil
	}

	return &types.SendMsgResp{
		Resp: respmsg.GetRespSuccess(l.ctx),
		Data: types.SendMsgData{
			Success: true,
			Have:    true,
		},
	}, nil
}