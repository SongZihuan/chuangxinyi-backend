package admin_user

import (
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/types"
	"net/http"

	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/logic/admin_user"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/svc"
	"github.com/wuntsong-org/go-zero-plus/rest/httpx"
)

func BlueInvoiceUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadBlueInvoice
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin_user.NewBlueInvoiceUploadLogic(r.Context(), svcCtx)
		resp, err := l.BlueInvoiceUpload(&req, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}