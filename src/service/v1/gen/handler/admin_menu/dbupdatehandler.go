package admin_menu

import (
	"net/http"

	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/logic/admin_menu"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/svc"
	"github.com/wuntsong-org/go-zero-plus/rest/httpx"
)

func DBUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin_menu.NewDBUpdateLogic(r.Context(), svcCtx)
		resp, err := l.DBUpdate()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
