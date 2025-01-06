package center

import (
	"net/http"

	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/logic/center"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/svc"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/types"
	"github.com/wuntsong-org/go-zero-plus/rest/httpx"
)

func Bind2FAHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Bind2FAReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := center.NewBind2FALogic(r.Context(), svcCtx)
		resp, err := l.Bind2FA(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}