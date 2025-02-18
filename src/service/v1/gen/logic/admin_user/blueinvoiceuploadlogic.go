package admin_user

import (
	"context"
	"database/sql"
	"gitee.com/wuntsong-auth/backend/src/audit"
	"gitee.com/wuntsong-auth/backend/src/model/db"
	"gitee.com/wuntsong-auth/backend/src/mysql"
	"gitee.com/wuntsong-auth/backend/src/oss"
	"gitee.com/wuntsong-auth/backend/src/service/v1/src/respmsg"
	"gitee.com/wuntsong-auth/backend/src/utils"
	errors "github.com/wuntsong-org/wterrors"
	"io"
	"net/http"
	"time"

	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/svc"
	"gitee.com/wuntsong-auth/backend/src/service/v1/gen/types"

	"github.com/wuntsong-org/go-zero-plus/core/logx"
)

type BlueInvoiceUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBlueInvoiceUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlueInvoiceUploadLogic {
	return &BlueInvoiceUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BlueInvoiceUploadLogic) BlueInvoiceUpload(req *types.UploadBlueInvoice, r *http.Request) (resp *types.RespEmpty, err error) {
	user, ok := l.ctx.Value("X-Token-User").(*db.User)
	if !ok {
		return nil, respmsg.BadContextError.New("X-Token-User")
	}

	issuerDate, err := time.Parse("2006-01-02", req.IssuerDate)
	if err != nil {
		return &types.RespEmpty{
			Resp: respmsg.GetRespByError(l.ctx, respmsg.BadForm, errors.WarpQuick(err), "读取issuer-date错误"),
		}, nil
	}

	fileHeader, ok := r.MultipartForm.File["file"]
	if !ok || len(fileHeader) == 0 {
		return &types.RespEmpty{
			Resp: respmsg.GetRespByMsg(l.ctx, respmsg.BadForm, "没有上传发票file参数"),
		}, nil
	}

	if fileHeader[0].Size >= FileMaxSize {
		return &types.RespEmpty{
			Resp: respmsg.GetRespByMsg(l.ctx, respmsg.FileTooBig, "发票文件太大: %d >= %d", fileHeader[0].Size, FileMaxSize),
		}, nil
	}

	file, err := fileHeader[0].Open()
	if err != nil {
		return &types.RespEmpty{
			Resp: respmsg.GetRespByError(l.ctx, respmsg.BadFormFile, errors.WarpQuick(err), "打开发票文件错误"),
		}, nil
	}

	defer utils.Close(file)

	fileByte, err := io.ReadAll(file)
	if err != nil {
		return &types.RespEmpty{
			Resp: respmsg.GetRespByError(l.ctx, respmsg.BadFormFile, errors.WarpQuick(err), "读取发票文件错误"),
		}, nil
	}

	fileType := utils.GetMediaType(fileByte)

	if !utils.IsPdf(fileType) {
		return &types.RespEmpty{
			Resp: respmsg.GetRespByMsg(l.ctx, respmsg.BadInvoiceType, "错误的发票文件类型"),
		}, nil
	}

	invoiceModel := db.NewInvoiceModel(mysql.MySQLConn)
	invoice, err := invoiceModel.FindByInvoiceID(l.ctx, req.ID)
	if errors.Is(err, db.ErrNotFound) {
		return &types.RespEmpty{
			Resp: respmsg.GetRespByMsg(l.ctx, respmsg.InvoiceNotFound, "发票未找到"),
		}, nil
	} else if err != nil {
		return nil, respmsg.MySQLSystemError.WarpQuick(err)
	}

	if invoice.Status != db.InvoiceRedFlush && invoice.Status != db.InvoiceReturn && invoice.Status != db.InvoiceOK && invoice.Status != db.InvoiceWaitReturn {
		return &types.RespEmpty{
			Resp: respmsg.GetRespByMsg(l.ctx, respmsg.BadInvoiceStatus, "发票状态错误"),
		}, nil
	}

	key, err := oss.UploadInvoice(invoice.InvoiceId, fileByte, false, false)
	if err != nil {
		return nil, respmsg.OSSError.WarpQuick(err)
	}

	invoice.InvoiceCode = sql.NullString{
		Valid:  len(req.InvoiceCode) != 0,
		String: req.InvoiceCode,
	}

	invoice.InvoiceNumber = sql.NullString{
		Valid:  true,
		String: req.InvoiceNumber,
	}

	invoice.InvoiceCheckCode = sql.NullString{
		Valid:  len(req.InvoiceCheckCode) != 0,
		String: req.InvoiceCheckCode,
	}

	invoice.IssuerAt = sql.NullTime{
		Valid: true,
		Time:  issuerDate,
	}

	invoice.InvoiceKey = sql.NullString{
		Valid:  true,
		String: key,
	}

	err = invoiceModel.Update(l.ctx, invoice)

	audit.NewAdminAudit(user.Id, "管理员上传蓝字发票（%s）成功", req.ID)

	return &types.RespEmpty{
		Resp: respmsg.GetRespSuccess(l.ctx),
	}, nil
}
