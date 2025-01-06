package ws

import (
	"net/http"

	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/logic/ws"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/svc"
)

func WSGetInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := ws.NewWSGetInfoLogic(r.Context(), svcCtx)
		l.WSGetInfo(w, r)
	}
}