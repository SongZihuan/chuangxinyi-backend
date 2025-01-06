package respmsg

type SubCode struct {
	NumCode int64  `json:"_code"`
	Code    string `json:"code"`
}

var (
	Success               = SubCode{NumCode: 0, Code: "SUCCESS"}
	UnknownError          = SubCode{NumCode: 1, Code: "UNKNOWN_ERROR"}
	SystemError           = SubCode{NumCode: 2, Code: "SYSTEM_ERROR"}
	TooBusy               = SubCode{NumCode: 3, Code: "TOO_BUSY"}
	BadCode               = SubCode{NumCode: 4, Code: "BAD_CODE"}
	OutOfDateCode         = SubCode{NumCode: 5, Code: "OUT_OF_DATE_CODE"}
	BadTokenType          = SubCode{NumCode: 6, Code: "BAD_TOKEN_TYPE"}
	UserNotFound          = SubCode{NumCode: 7, Code: "USER_NOT_FOUND"}
	UserNotRegister       = SubCode{NumCode: 8, Code: "USER_NOT_REGISTER"}
	EmailNotBind          = SubCode{NumCode: 9, Code: "EMAIL_NOT_BIND"}
	BadDefrayAt           = SubCode{NumCode: 10, Code: "BAD_DEFRAY_AT"}
	UserRegisterAgain     = SubCode{NumCode: 11, Code: "USER_REGISTER_AGAIN"}
	PhoneHasBeenRegister  = SubCode{NumCode: 12, Code: "PHONE_HAS_BEEN_REGISTER"}
	BadUserToken          = SubCode{NumCode: 13, Code: "BAD_USER_TOKEN"}
	Bad2FACode            = SubCode{NumCode: 14, Code: "BAD_2FA_CODE"}
	UserHasNotBeenBind2FA = SubCode{NumCode: 15, Code: "USER_HAS_NOT_BIND_2FA"}
	UserHasBeenBind2FA    = SubCode{NumCode: 16, Code: "USER_HAS_BIND_2FA"}
	Bad2FASecret          = SubCode{NumCode: 17, Code: "BAD_2FA_SECRET"}
	BadDefrayStatus       = SubCode{NumCode: 18, Code: "BAD_DEFRAY_STATUS"}
	BadUncleStatus        = SubCode{NumCode: 19, Code: "BAD_UNCLE_STATUS"}
	BadFormFile           = SubCode{NumCode: 20, Code: "BAD_FORM_FILE"}
	BadPicType            = SubCode{NumCode: 21, Code: "BAD_PIC_TYPE"}
	FileTooBig            = SubCode{NumCode: 22, Code: "FILE_TOO_BIG"}
	RepeatAddUncle        = SubCode{NumCode: 23, Code: "REPEAT_UNCLE"}
	WithoutVerify         = SubCode{NumCode: 24, Code: "WITHOUT_VERIFY"}
	BadPassToken          = SubCode{NumCode: 25, Code: "BAD_PASS_TOKEN"}
	RobotDeny             = SubCode{NumCode: 26, Code: "ROBOT_DENY"}
	PasswordError         = SubCode{NumCode: 27, Code: "PASSWORD_ERROR"}
	OnlyDevelop           = SubCode{NumCode: 28, Code: "ONLY_DEVELOP"}
	UserNotHomePage       = SubCode{NumCode: 29, Code: "USER_NOT_HOME_PAGE"}
	//AdminHasBeenBind2FA          = SubCode{NumCode: 30, Code: ""}
	//AdminNotFound                = SubCode{NumCode: 31, Code: ""}
	RoleNotFound              = SubCode{NumCode: 32, Code: "ROLE_NOT_FOUND"}
	RootAdminCannotChangeRole = SubCode{NumCode: 33, Code: "ROOT_ADMIN_CAN_NOT_CHANGE_ROLE"}
	BadUserStatus             = SubCode{NumCode: 34, Code: "BAD_USER_STATUS"}
	//MenuFatherNotFound           = SubCode{NumCode: 35, Code: ""}
	MenuNotFount = SubCode{NumCode: 36, Code: "MENU_NOT_FOUND"}
	CanNotUp     = SubCode{NumCode: 37, Code: "CAN_NOT_UP"}
	CanNotDown   = SubCode{NumCode: 38, Code: "CAN_NOT_DOWN"}
	//RoleNotInMenu                = SubCode{NumCode: 39, Code: ""}
	PolicyDeny    = SubCode{NumCode: 40, Code: "POLICY_DENY"}
	BadAppID      = SubCode{NumCode: 41, Code: "BAD_APP_ID"}
	BadWeChatCode = SubCode{NumCode: 42, Code: "BAD_WECHAT_CODE"}
	WeChatNotBind = SubCode{NumCode: 43, Code: "WECHAT_NOT_BIND"}
	//BadWeChatToken               = SubCode{NumCode: 44, Code: ""}
	BadRequestsSender = SubCode{NumCode: 45, Code: "BAD_REQUESTS_SENDER"}
	BadWebsiteID      = SubCode{NumCode: 46, Code: "BAD_WEBSITE_ID"}
	BadWebsiteIPID    = SubCode{NumCode: 47, Code: "BAD_WEBSITE_IP_ID"}
	//WebsiteExists                = SubCode{NumCode: 48, Code: ""}
	//WebsiteSecretTooLow          = SubCode{NumCode: 49, Code: ""}
	//BadContentBody               = SubCode{NumCode: 50, Code: ""}
	//ContentTooLong               = SubCode{NumCode: 51, Code: ""}
	//NotUTF8                      = SubCode{NumCode: 52, Code: ""}
	NotFound       = SubCode{NumCode: 53, Code: "STATUS_NOT_FOUND"}
	MethodNotAllow = SubCode{NumCode: 54, Code: "STATUS_NOT_ALLOW"}
	//OldPhoneMustBeGiven          = SubCode{NumCode: 55, Code: ""}
	UserNotVerify      = SubCode{NumCode: 56, Code: "USER_NOT_VERIFY"}
	BadLoginInfo       = SubCode{NumCode: 57, Code: "BAD_LOGIN_INFO"}
	DomainNotFound     = SubCode{NumCode: 58, Code: "DOMAIN_NOT_FOUND"}
	BadEmail           = SubCode{NumCode: 59, Code: "BAD_EMAIL"}
	BadPhone           = SubCode{NumCode: 60, Code: "BAD_PHONE"}
	BadPassword        = SubCode{NumCode: 61, Code: "BAD_PASSWORD"}
	CAPTCHASecondCheck = SubCode{NumCode: 62, Code: "CAPTCHA_SECOND_CHECK"}
	//NicknameTooShort             = SubCode{NumCode: 63, Code: ""}
	BadUserInfo     = SubCode{NumCode: 64, Code: "BAD_USER_INFO"}
	BadCompanyInfo  = SubCode{NumCode: 65, Code: "BAD_COMPANY_INFO"}
	CreateTradeFail = SubCode{NumCode: 66, Code: "CREATE_TRADE_FAIL"}
	BadPayCNY       = SubCode{NumCode: 67, Code: "BAD_PAY_CNY"}
	//NotFinishRegister            = SubCode{NumCode: 8, Code: ""}
	OnlyRootUser = SubCode{NumCode: 69, Code: "ONLY_ROOT_USER"}
	//JustFatherCanDoIt            = SubCode{NumCode:7 0, Code: ""}
	SonNotFound = SubCode{NumCode: 71, Code: "SON_NOT_FOUND"}
	//BadPayIDGenerate             = SubCode{NumCode: 27, Code: ""}
	BadPayWay               = SubCode{NumCode: 73, Code: "BAD_PAY_WAY"}
	CreateDefrayFail        = SubCode{NumCode: 74, Code: "CREATE_DEFRAY_FAIL"}
	DefrayNotFound          = SubCode{NumCode: 75, Code: "DEFRAY_NOT_FOUND"}
	DefrayTimeout           = SubCode{NumCode: 76, Code: "DEFRAY_TIMEOUT"}
	DoublePayment           = SubCode{NumCode: 77, Code: "DOUBLE_PAY"}
	Insufficient            = SubCode{NumCode: 78, Code: "INSUFFICIENT"}
	PayFail                 = SubCode{NumCode: 79, Code: "PAY_FAIL"}
	BadWebsitePubKey        = SubCode{NumCode: 80, Code: "BAD_WEBSITE_PUBLIC_KEY"}
	DefrayInsufficient      = SubCode{NumCode: 81, Code: "DEFRAY_INSUFFICIENT"}
	UserNameHasBeenRegister = SubCode{NumCode: 82, Code: "USER_NAME_HAS_REGISTER"}
	PhoneNotRegister        = SubCode{NumCode: 83, Code: "PHONE_NOT_REGISTER"}
	BadQQ                   = SubCode{NumCode: 84, Code: "BAD_QQ"}
	BadWeChat               = SubCode{NumCode: 85, Code: "BAD_WECHAT"}
	NameTooShort            = SubCode{NumCode: 86, Code: ""}
	//TaxIDTooShort                = SubCode{NumCode: 87, Code: ""}
	InvoiceFail       = SubCode{NumCode: 88, Code: "INVOICE_FAIL"}
	InsufficientQuota = SubCode{NumCode: 89, Code: "INSUFFICIENT_INVOICE_QUOTA"}
	TitleNotFound     = SubCode{NumCode: 90, Code: "TITLE_NOT_FOUND"}
	AddressNotFound   = SubCode{NumCode: 91, Code: "ADDRESS_NOT_FOUND"}
	BadTitle          = SubCode{NumCode: 92, Code: "BAD_TITLE"}
	//BadAddress                   = SubCode{NumCode: 93, Code: ""}
	InvoiceNotFound        = SubCode{NumCode: 94, Code: "INVOICE_NOT_FOUND"}
	PayMustVerify          = SubCode{NumCode: 95, Code: "PAY_MUST_VERIFY"}
	MustVerify             = SubCode{NumCode: 96, Code: "MUST_VERIFY"}
	TradeNotFound          = SubCode{NumCode: 97, Code: "TRADE_NOT_FOUND"}
	DiscountNotFound       = SubCode{NumCode: 98, Code: "DISCOUNT_NOT_FOUND"}
	DiscountBuyFail        = SubCode{NumCode: 99, Code: "DISCOUNT_JOIN_FAIL"}
	BadFid                 = SubCode{NumCode: 100, Code: "BAD_FILE_ID"}
	BadForm                = SubCode{NumCode: 101, Code: "BAD_FORM"}
	AdminCannotDelete      = SubCode{NumCode: 102, Code: "ROOT_ADMIN_CAN_NOT_DELETE"}
	InviteUserNotFound     = SubCode{NumCode: 103, Code: "INVITE_USER_NOT_FOUND"}
	PayNotFound            = SubCode{NumCode: 104, Code: "PAY_NOT_FOUND"}
	InvoiceReturnFail      = SubCode{NumCode: 105, Code: "INVOICE_RETURN_FAIL"}
	BadQuota               = SubCode{NumCode: 106, Code: "BAD_QUOTA"}
	PhoneRepeat            = SubCode{NumCode: 107, Code: "PHONE_REPEAT"}
	UserNameRepeat         = SubCode{NumCode: 108, Code: "USER_NAME_REPEAT"}
	SystemRoleCanNotChange = SubCode{NumCode: 109, Code: "SYSTEM_ROLE_CAN_NOT_CHANGE"}
	AnnouncementNotFound   = SubCode{NumCode: 110, Code: "ANNOUNCEMENT_NOT_FOUND"}
	WebsiteNotFound        = SubCode{NumCode: 111, Code: "WEBSITE_NOT_FOUND"}
	BadKeyMap              = SubCode{NumCode: 112, Code: "BAD_KEY_MAP"}
	IsNotUncle             = SubCode{NumCode: 113, Code: "IS_NOT_UNCLE"}
	BadOauthAllow          = SubCode{NumCode: 111, Code: "BAD_OAUTH2_ALLOW"}
	BadArea                = SubCode{NumCode: 115, Code: "BAD_AREA"}
	TooManyFile            = SubCode{NumCode: 116, Code: "TOO_MANY_FILE"}
	//FailToCreateOrder            = SubCode{NumCode: 117, Code: ""}
	WorkOrderNotFound = SubCode{NumCode: 118, Code: "WORK_ODER_NOT_FOUND"}
	//MessageNotFound              = SubCode{NumCode: 119, Code: ""}
	SendLimit  = SubCode{NumCode: 120, Code: "SEND_LIMIT"}
	BadIDCard  = SubCode{NumCode: 121, Code: "BAD_ID_CARD"}
	BadLicense = SubCode{NumCode: 122, Code: "BAD_LICENSE"}
	//BadIDCardToken               = SubCode{NumCode: 123, Code: ""}
	NotCorrespond = SubCode{NumCode: 124, Code: "NOT_CORRESPOND"}
	//CompanyNotVerify             = SubCode{NumCode: 125, Code: ""}
	//BadCompanyToken              = SubCode{NumCode: 126, Code: ""}
	BadFaceIDCardOrName = SubCode{NumCode: 127, Code: "BAD_FACE_ID_CARD_OR_NAME"}
	FaceCheckFail       = SubCode{NumCode: 128, Code: "FACE_CHECK_FAIL"}
	//BadPayIP                     = SubCode{NumCode: 129, Code: ""}
	//BadFaceToken                 = SubCode{NumCode: 130, Code: ""}
	NeedVerify             = SubCode{NumCode: 131, Code: "NEED_VERIFY"}
	NeedCompany            = SubCode{NumCode: 132, Code: "NEED_COMPANY"}
	NeedUserOrigin         = SubCode{NumCode: 133, Code: "NEED_USER_ORIGIN"}
	NeedCompanyOrigin      = SubCode{NumCode: 134, Code: "NEED_COMPANY_ORIGIN"}
	NeedUserFace           = SubCode{NumCode: 135, Code: "NEED_USER_FACE"}
	NeedCompanyFace        = SubCode{NumCode: 136, Code: "NEED_COMPANY_FACE"}
	AgreementNotExists     = SubCode{NumCode: 137, Code: "AGREEMENT_NOT_EXISTS"}
	AgreementExists        = SubCode{NumCode: 138, Code: "AGREEMENT_EXISTS"}
	WorkOrderDoubleFinish  = SubCode{NumCode: 139, Code: "WORK_ORDER_DOUBLE_FINISH"}
	InvoiceOperationFail   = SubCode{NumCode: 140, Code: "INVOICE_OPERATION_FAIL"}
	ReturnTooLate          = SubCode{NumCode: 141, Code: "RETURN_TOO_LATE"}
	TimeTooShort           = SubCode{NumCode: 142, Code: "TIME_TOO_SHORT"}
	RepeatFileName         = SubCode{NumCode: 143, Code: "REPEAT_FAIL_NAME"}
	OrderHasFinish         = SubCode{NumCode: 144, Code: "ORDER_HAS_FINISH"}
	BadHomePage            = SubCode{NumCode: 145, Code: "BAD_HOME_PAGE"}
	BadHeader              = SubCode{NumCode: 146, Code: "BAD_HEADER"}
	InvoiceDoubleOperation = SubCode{NumCode: 147, Code: "INVOICE_DOUBLE_OPERATION"}
	BadLink                = SubCode{NumCode: 148, Code: "BAD_LINK"}
	//BadDistributionLevel         = SubCode{NumCode: 149, Code: ""}
	QueryTradeFail    = SubCode{NumCode: 150, Code: "QUERY_TRADE_FAIL"}
	RefundFail        = SubCode{NumCode: 151, Code: "REFUND_FAIL"}
	CorsNotAllow      = SubCode{NumCode: 152, Code: "CORS_NOT_ALLOW"}
	ChangeBalanceFail = SubCode{NumCode: 153, Code: "CHANGE_BALANCE_FAIL"}
	ChangeBilledFail  = SubCode{NumCode: 154, Code: "CHANGE_BILLED_FAIL"}
	//LogoutAll                    = SubCode{NumCode: 155, Code: ""}
	BadNickName = SubCode{NumCode: 156, Code: "BAD_NICKNAME"}
	//UploadDefaultHeaderFail      = SubCode{NumCode: 157, Code: ""}
	PurchaseLimit                = SubCode{NumCode: 158, Code: "PURCHASE_LIMIT"}
	BadUserName                  = SubCode{NumCode: 159, Code: "BAD_USERNAME"}
	NotOpenWebsite               = SubCode{NumCode: 160, Code: "NOT_OPEN_WEBSITE"}
	RequestsTooBusy              = SubCode{NumCode: 101, Code: "REQUESTS_TOO_BUSY"}
	DoubleReturn                 = SubCode{NumCode: 162, Code: "DOUBLE_RETURN"}
	ApplicationNotFound          = SubCode{NumCode: 163, Code: "APPLICATION_NOT_FOUND"}
	BadRoleBelong                = SubCode{NumCode: 164, Code: "BAD_ROLE_BELONG"}
	WebsiteNotAllow              = SubCode{NumCode: 165, Code: "WEBSITE_NOT_ALLOW"}
	BadApplicationStatus         = SubCode{NumCode: 166, Code: "BAD_APPLICATION_STATUS"}
	DoubleConnect                = SubCode{NumCode: 167, Code: "DOUBLE_CONNECT"}
	BadDeleteToken               = SubCode{NumCode: 168, Code: "BAD_DELETE_TOKEN"}
	TooMany                      = SubCode{NumCode: 169, Code: "TOO_MANY"}
	ReturnInsufficientQuota      = SubCode{NumCode: 190, Code: "RETURN_INSUFFICIENT_QUOTA"}
	FaceCheckWait                = SubCode{NumCode: 191, Code: "FACE_CHECK_WAIT"}
	CanNotGetLock                = SubCode{NumCode: 192, Code: "CAN_NOT_GET_LOCK"}
	BackNotFound                 = SubCode{NumCode: 193, Code: "BAD_NOT_FOUND"}
	AdminCanNotChangeStatus      = SubCode{NumCode: 194, Code: "ADMIN_CAN_NOT_CHANGE_STATUS"}
	BadInvoiceType               = SubCode{NumCode: 195, Code: "BAD_INVOICE_TYPE"}
	BadInvoiceStatus             = SubCode{NumCode: 196, Code: "BAD_INVOICE_STATUS"}
	BadName                      = SubCode{NumCode: 197, Code: "BAD_NAME"}
	WithdrawNotFound             = SubCode{NumCode: 198, Code: "WITHDRAW_NOT_FOUND"}
	WalletRecordNotFound         = SubCode{NumCode: 199, Code: "WALLET_RECORD_NOT_FOUND"}
	WebsiteFundingRecordNotFound = SubCode{NumCode: 200, Code: "WEBSITE_FUNDING_RECORD_NOT_FOUND"}
	CreateBackFail               = SubCode{NumCode: 201, Code: "CREATE_BACK_FAIL"}
	MustSelfDefray               = SubCode{NumCode: 202, Code: "MUST_SELF_DEFRAY"}
	BadRoleStatus                = SubCode{NumCode: 203, Code: "BAD_ROLE_STATUS"}
	BadWebsiteStatus             = SubCode{NumCode: 204, Code: "BAD_WEBSITE_STATUS"}
	BadMenuStatus                = SubCode{NumCode: 205, Code: "BAD_MENU_STATUS"}
	BadPolicyStatus              = SubCode{NumCode: 206, Code: "BAD_POLICY_STATUS"}
	PermissionNotFound           = SubCode{NumCode: 207, Code: "PERMISSION_NOT_FOUND"}
	BadPathMode                  = SubCode{NumCode: 208, Code: "BAD_PATH_MODE"}
	BadPathStatus                = SubCode{NumCode: 209, Code: "BAD_PATH_STATUS"}
	PathNotFound                 = SubCode{NumCode: 210, Code: "PATH_NOT_FOUND"}
	CreateWebsiteFail            = SubCode{NumCode: 211, Code: "CREATE_WEBSITE_FAIL"}
	MenuHasSon                   = SubCode{NumCode: 212, Code: "MENU_HAS_SON"}
	DefrayReturnFail             = SubCode{NumCode: 213, Code: "DEFRAY_RETURN_FAIL"}
	NotPeerName                  = SubCode{NumCode: 214, Code: "NOT_PEER_NAME"}
	PeerNameError                = SubCode{NumCode: 215, Code: "PEER_NAME_ERROR"}
	BadPeerIP                    = SubCode{NumCode: 216, Code: "BAD_PEER_IP"}
	CreateWorkOrderFail          = SubCode{NumCode: 217, Code: "CREATE_WORK_ORDER_FAIL"}
	SelfPayFail                  = SubCode{NumCode: 218, Code: "SELF_PAY_FAIL"}
	BadWechatInfo                = SubCode{NumCode: 219, Code: "BAD_WECHAT_INFO"}
	BadPathRegex                 = SubCode{NumCode: 220, Code: "BAD_PATH_REGEX"}
)
