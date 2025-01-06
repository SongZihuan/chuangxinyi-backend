package admin_accessrecord

import (
	"net/http"

	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/logic/admin_accessrecord"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/svc"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/types"
	"github.com/wuntsong-org/go-zero-plus/rest/httpx"
)

func GetTokenRecordListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTokenRecordList
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin_accessrecord.NewGetTokenRecordListLogic(r.Context(), svcCtx)
		resp, err := l.GetTokenRecordList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
