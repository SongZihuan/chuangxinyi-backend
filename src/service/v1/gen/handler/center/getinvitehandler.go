package center

import (
	"net/http"

	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/logic/center"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/svc"
	"github.com/wuntsong-org/go-zero-plus/rest/httpx"
)

func GetInviteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := center.NewGetInviteLogic(r.Context(), svcCtx)
		resp, err := l.GetInvite()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
