package msg

import (
	"net/http"

	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/logic/msg"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/svc"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/types"
	"github.com/wuntsong-org/go-zero-plus/rest/httpx"
)

func SendAuditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendAuditReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := msg.NewSendAuditLogic(r.Context(), svcCtx)
		resp, err := l.SendAudit(&req, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
