package ui

import (
	"gitee.com/wuntsong-auth/backend/src/config"
	"gitee.com/wuntsong-auth/backend/src/service/v1/src/respmsg"
	"gitee.com/wuntsong-auth/backend/src/utils"
	errors "github.com/wuntsong-org/wterrors"
	"net/http"

	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/logic/ui"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/svc"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/types"
	"github.com/wuntsong-org/go-zero-plus/rest/httpx"
)

func GetFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ui.NewGetFileLogic(r.Context(), svcCtx)
		err := l.GetFile(&req, w, r)
		if err != nil {
			resp := respmsg.GetRespByError(r.Context(), respmsg.NotFound, errors.WarpQuick(err), "获取媒体文件错误")
			utils.NotFound(w, r, err, config.BackendConfig.GetModeFromRequests(r) == config.RunModeDevelop, resp.RequestsID)
		}
	}
}
