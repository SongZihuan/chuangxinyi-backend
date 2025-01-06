package defray

import (
	"net/http"

	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/logic/defray"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/svc"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/types"
	"github.com/wuntsong-org/go-zero-plus/rest/httpx"
)

func CreateDefrayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateDefrayReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := defray.NewCreateDefrayLogic(r.Context(), svcCtx)
		resp, err := l.CreateDefray(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}