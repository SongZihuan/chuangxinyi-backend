package admin_user

import (
	"context"
	"gitee.com/wuntsong-auth/backend/src/model/db"
	"gitee.com/wuntsong-auth/backend/src/mysql"
	"gitee.com/wuntsong-auth/backend/src/oss"
	"gitee.com/wuntsong-auth/backend/src/service/v1/src/respmsg"
	errors "github.com/wuntsong-org/wterrors"

	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/svc"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/types"

	"github.com/wuntsong-org/go-zero-plus/core/logx"
)

type DownloadInvoiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadInvoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadInvoiceLogic {
	return &DownloadInvoiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadInvoiceLogic) DownloadInvoice(req *types.AdminDownloadInvoiceReq) (resp *types.AdminDownloadInvoiceResp, err error) {
	invoiceModel := db.NewInvoiceModel(mysql.MySQLConn)
	invoice, err := invoiceModel.FindByInvoiceID(l.ctx, req.InvoiceID)
	if errors.Is(err, db.ErrNotFound) {
		return &types.AdminDownloadInvoiceResp{
			Resp: respmsg.GetRespByMsg(l.ctx, respmsg.InvoiceNotFound, "发票未找到"),
		}, nil
	} else if err != nil {
		return nil, respmsg.MySQLSystemError.WarpQuick(err)
	}

	invoiceUrl := ""
	if invoice.Status == db.InvoiceOK || invoice.Status == db.InvoiceWaitReturn || invoice.Status == db.InvoiceReturn || invoice.Status == db.InvoiceRedFlush {
		invoiceUrl, err = oss.GetInvoice(invoice.InvoiceId, false)
		if err != nil {
			invoiceUrl = "" // 表示没有发票
		}
	}

	redInvoiceUrl := ""
	if invoice.Status == db.InvoiceRedFlush {
		redInvoiceUrl, err = oss.GetInvoice(invoice.InvoiceId, true)
		if err != nil {
			redInvoiceUrl = "" // 表示没有发票
		}
	}

	return &types.AdminDownloadInvoiceResp{
		Resp: respmsg.GetRespSuccess(l.ctx),
		Data: types.AdminDownloadInvoiceData{
			BlueInvoice: invoiceUrl,
			RedInvoice:  redInvoiceUrl,
		},
	}, nil
}
