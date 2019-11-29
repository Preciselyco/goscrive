package scrive

const (
	viewer          = "viewer"
	signingParty    = "signing_party"
	approver        = "approver"
	unknown         = "unknown"
	notDelivered    = "not_delivered"
	delivered       = "delivered"
	deferred        = "deferred"
	email           = "email"
	mobile          = "mobile"
	emailMobile     = "email_mobile"
	emailLink       = "email_link"
	emailLinkMobile = "email_link_mobile"
	none            = "none"
	pad             = "pad"
	api             = "api"
	standard        = "standard"
	smsPin          = "sms_pin"
	seBankID        = "se_bankid"
	noBankID        = "no_bankid"
	dkNemID         = "dk_nemid"
	fiTupas         = "fi_tupas"
	verimi          = "verimi"
)

type SignatoryRole = string

const (
	SignatoryRoleViewer       SignatoryRole = viewer
	SignatoryRoleSigningParty SignatoryRole = signingParty
	SignatoryRoleApprover     SignatoryRole = approver
)

type EmailDeliveryStatus = string

const (
	EmailDeliveryStatusUnknown      EmailDeliveryStatus = unknown
	EmailDeliveryStatusNotDelivered EmailDeliveryStatus = notDelivered
	EmailDeliveryStatusDelivered    EmailDeliveryStatus = delivered
	EmailDeliveryStatusDeferred     EmailDeliveryStatus = deferred
)

type ConfirmationEmailDeliveryStatus = string

const (
	ConfirmationEmailDeliveryStatusUnknown      ConfirmationEmailDeliveryStatus = unknown
	ConfirmationEmailDeliveryStatusNotDelivered ConfirmationEmailDeliveryStatus = notDelivered
	ConfirmationEmailDeliveryStatusDelivered    ConfirmationEmailDeliveryStatus = delivered
	ConfirmationEmailDeliveryStatusDeferred     ConfirmationEmailDeliveryStatus = deferred
)

type MobileDeliveryStatus = string

const (
	MobileDeliveryStatusUnknown      MobileDeliveryStatus = unknown
	MobileDeliveryStatusNotDelivered MobileDeliveryStatus = notDelivered
	MobileDeliveryStatusDelivered    MobileDeliveryStatus = delivered
	MobileDeliveryStatusDeferred     MobileDeliveryStatus = deferred
)

type DeliveryMethod = string

const (
	DeliveryMethodEmail       DeliveryMethod = email
	DeliveryMethodMobile      DeliveryMethod = mobile
	DeliveryMethodEmailMobile DeliveryMethod = emailMobile
	DeliveryMethodPad         DeliveryMethod = pad
	DeliveryMethodAPI         DeliveryMethod = api
)

type AuthenticationMethodToView = string

const (
	AuthenticationMethodToViewStandard AuthenticationMethodToView = standard
	AuthenticationMethodToViewSmsPin   AuthenticationMethodToView = smsPin
	AuthenticationMethodToViewSEBankID AuthenticationMethodToView = seBankID
	AuthenticationMethodToViewNOBankID AuthenticationMethodToView = noBankID
	AuthenticationMethodToViewDKNemID  AuthenticationMethodToView = dkNemID
	AuthenticationMethodToViewFITupas  AuthenticationMethodToView = fiTupas
	AuthenticationMethodToViewVerimi   AuthenticationMethodToView = verimi
)

type AuthenticationMethodToViewArchived = string

const (
	AuthenticationMethodToViewArchivedStandard AuthenticationMethodToViewArchived = standard
	AuthenticationMethodToViewArchivedSmsPin   AuthenticationMethodToViewArchived = smsPin
	AuthenticationMethodToViewArchivedSEBankID AuthenticationMethodToViewArchived = seBankID
	AuthenticationMethodToViewArchivedNOBankID AuthenticationMethodToViewArchived = noBankID
	AuthenticationMethodToViewArchivedDKNemID  AuthenticationMethodToViewArchived = dkNemID
	AuthenticationMethodToViewArchivedFITupas  AuthenticationMethodToViewArchived = fiTupas
	AuthenticationMethodToViewArchivedVerimi   AuthenticationMethodToViewArchived = verimi
)

type AuthenticationMethodToSign = string

const (
	AuthenticationMethodToSignStandard AuthenticationMethodToSign = standard
	AuthenticationMethodToSignSmsPin   AuthenticationMethodToSign = smsPin
	AuthenticationMethodToSignSEBankID AuthenticationMethodToSign = seBankID
	AuthenticationMethodToSignNOBankID AuthenticationMethodToSign = noBankID
	AuthenticationMethodToSignDKNemID  AuthenticationMethodToSign = dkNemID
)

type ConfirmationDeliveryMethod = string

const (
	ConfirmationDeliveryMethodEmail           ConfirmationDeliveryMethod = email
	ConfirmationDeliveryMethodMobile          ConfirmationDeliveryMethod = mobile
	ConfirmationDeliveryMethodEmailMobile     ConfirmationDeliveryMethod = emailMobile
	ConfirmationDeliveryMethodEmailLink       ConfirmationDeliveryMethod = emailLink
	ConfirmationDeliveryMethodEmailLinkMobile ConfirmationDeliveryMethod = emailLinkMobile
	ConfirmationDeliveryMethodNone            ConfirmationDeliveryMethod = none
)

type NotificationDeliveryMethod = string

const (
	NotificationDeliveryMethodEmail           NotificationDeliveryMethod = email
	NotificationDeliveryMethodMobile          NotificationDeliveryMethod = mobile
	NotificationDeliveryMethodEmailMobile     NotificationDeliveryMethod = emailMobile
	NotificationDeliveryMethodEmailLink       NotificationDeliveryMethod = emailLink
	NotificationDeliveryMethodEmailLinkMobile NotificationDeliveryMethod = emailLinkMobile
	NotificationDeliveryMethodNone            NotificationDeliveryMethod = none
)

type Status = string

const (
	StatusPreparation   Status = "preparation"
	StatusPending       Status = "pending"
	StatusClosed        Status = "closed"
	StatusCanceled      Status = "canceled"
	StatusTimedout      Status = "timedout"
	StatusRejected      Status = "rejected"
	StatusDocumentError Status = "document_error"
)

type Lang = string

const (
	LangDA Lang = "da"
	LangDE Lang = "de"
	LangEL Lang = "el"
	LangEN Lang = "en"
	LangES Lang = "es"
	LangET Lang = "et"
	LangFI Lang = "fi"
	LangFR Lang = "fr"
	LangIS Lang = "is"
	LangIT Lang = "it"
	LangLT Lang = "lt"
	LangLV Lang = "lv"
	LangNL Lang = "nl"
	LangNO Lang = "no"
	LangPT Lang = "pt"
	LangSV Lang = "sv"
)

type ViewerRole = string

const (
	ViewerRoleCompanyShared ViewerRole = "company_shared"
	ViewerRoleCompanyAdmin  ViewerRole = "company_admin"
	ViewerRoleSignatory     ViewerRole = "signatory"
)

type Document struct {
	ID                  string              `json:"id"`
	Title               string              `json:"title"`
	Parties             []*Signatory        `json:"parties"`
	File                File                `json:"file"`
	SealedFile          *File               `json:"sealed_file"`
	AuthorAttachments   []*AuthorAttachment `json:"author_attachments"`
	CTime               string              `json:"ctime"`
	MTime               string              `json:"mtime"`
	TimeoutTime         *string             `json:"timeout_time"`
	AutoRemindTime      *string             `json:"auto_remind_time"`
	Status              Status              `json:"status"`
	DaysToSign          uint32              `json:"days_to_sign"`
	DaysToRemind        *uint32             `json:"days_to_remind"`
	DisplayOptions      DisplayOptions      `json:"display_options"`
	InvitationMessage   string              `json:"invitation_message"`
	ConfirmationMessage string              `json:"confirmation_message"`
	Lang                Lang                `json:"lang"`
	APICallbackURL      *string             `json:"api_callback_url"`
	ObjectVersion       *uint32             `json:"object_version,omitempty"`
	AccessToken         string              `json:"access_token"`
	Timezone            string              `json:"timezone"`
	Tags                []*Tag              `json:"tags"`
	IsTemplate          bool                `json:"is_template"`
	IsSaved             bool                `json:"is_saved"`
	FolderID            *string             `json:"folder_id"`
	IsShared            bool                `json:"is_shared"`
	IsTrashed           bool                `json:"is_trashed"`
	IsDeleted           bool                `json:"is_deleted"`
	Viewer              *Viewer             `json:"viewer"`
	ShareableLink       *string             `json:"shareable_link"`
	TemplateID          *string             `json:"template_id"`
	FromShareableLink   bool                `json:"from_shareable_link"`
}

type Viewer struct {
	Role        ViewerRole `json:"role"`
	SignatoryID string     `json:"signatory_id"`
}

type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DisplayOptions struct {
	ShowHeader        bool `json:"show_header"`
	ShowPDFDownload   bool `json:"show_pdf_download"`
	ShowRejectOption  bool `json:"show_reject_option"`
	AllowRejectReason bool `json:"allow_reject_reason"`
	ShowFooter        bool `json:"show_footer"`
	DocumentIsReceipt bool `json:"document_is_receipt"`
	ShowArrow         bool `json:"show_arrow"`
}

type AuthorAttachment struct {
	Name            string `json:"name"`
	Required        bool   `json:"required"`
	AddToSealedFile bool   `json:"add_to_sealed_file"`
	FileID          string `json:"file_id"`
}

type Signatory struct {
	ID                                 string                             `json:"id"`
	UserID                             *string                            `json:"user_id"`
	IsAuthor                           bool                               `json:"is_author"`
	IsSignatory                        bool                               `json:"is_signatory"`
	SignatoryRole                      SignatoryRole                      `json:"signatory_role"`
	Fields                             []*SignatoryField                  `json:"fields"`
	ConsentModule                      *ConsentModule                     `json:"consent_module"`
	SignOrder                          uint32                             `json:"sign_order"`
	SignTime                           *string                            `json:"sign_time,omitempty"`
	SeenTime                           *string                            `json:"seen_time,omitempty"`
	ReadInvitationTime                 *string                            `json:"read_invitation_time,omitempty"`
	RejectedTime                       *string                            `json:"rejected_time,omitempty"`
	RejectionReason                    *string                            `json:"rejection_reason,omitempty"`
	SignSuccessRedirectURL             *string                            `json:"sign_success_redirect_url,omitempty"`
	RejectRedirectURL                  *string                            `json:"reject_redirect_url,omitempty"`
	EmailDeliveryStatus                EmailDeliveryStatus                `json:"email_delivery_status"`
	MobileDeliveryStatus               MobileDeliveryStatus               `json:"mobile_delivery_status"`
	ConfirmationEmailDeliveryStatus    ConfirmationEmailDeliveryStatus    `json:"confirmation_email_delivery_status"`
	HasAuthenticatedToView             bool                               `json:"has_authenticated_to_view"`
	CSV                                *[]string                          `json:"csv,omitempty"`
	DeliveryMethod                     DeliveryMethod                     `json:"delivery_method"`
	AuthenticationMethodToView         AuthenticationMethodToView         `json:"authentication_method_to_view"`
	AuthenticationMethodToViewArchived AuthenticationMethodToViewArchived `json:"authentication_method_to_view_archived"`
	AuthenticationMethodToSign         AuthenticationMethodToSign         `json:"authentication_method_to_sign"`
	ConfirmationDeliveryMethod         ConfirmationDeliveryMethod         `json:"confirmation_delivery_method"`
	NotificationDeliveryMethod         NotificationDeliveryMethod         `json:"notification_delivery_method"`
	AllowsHighlighting                 bool                               `json:"allows_highlighting"`
	HidePersonalNumber                 bool                               `json:"hide_personal_number"`
	CanForward                         bool                               `json:"can_forward"`
	Attachments                        *[]*Attachment                     `json:"attachments"`
	HighlightedPages                   *[]*HighlightedPage                `json:"highlighted_pages"`
	APIDeliveryURL                     *string                            `json:"api_delivery_url"`
}

type ConsentModule struct {
	Title     string                   `json:"title"`
	Questions []*ConsentModuleQuestion `json:"questions"`
}

type ConsentModuleQuestion struct {
	Title               string                           `json:"title"`
	PositiveOption      string                           `json:"positive_option"`
	NegativeOption      string                           `json:"negative_option"`
	Response            bool                             `json:"response"`
	DetailedDescription ConsentModuleDetailedDescription `json:"detailed_description"`
}

type ConsentModuleDetailedDescription struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type HighlightedPage struct {
	Page   uint32 `json:"page"`
	FileID string `json:"file_id"`
}

type Attachment struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	FileID      string `json:"file_id"`
	FileName    string `json:"file_name"`
}

type SignatoryFieldType = string

const (
	SignatoryFieldTypeName           SignatoryFieldType = "name"
	SignatoryFieldTypeEmail          SignatoryFieldType = "email"
	SignatoryFieldTypeMobile         SignatoryFieldType = "mobile"
	SignatoryFieldTypeSignature      SignatoryFieldType = "signature"
	SignatoryFieldTypeCompanyName    SignatoryFieldType = "company"
	SignatoryFieldTypeCompanyNumber  SignatoryFieldType = "company_number"
	SigantoryFieldTypePersonalNumber SignatoryFieldType = "personal_number"
	SignatoryFieldTypeCheckbox       SignatoryFieldType = "checkbox"
	SignatoryFieldTypeRadiogroup     SignatoryFieldType = "radiogroup"
	SignatoryFieldTypeCustomText     SignatoryFieldType = "text"
)

type SignatoryField struct {
	Type      SignatoryFieldType `json:"type"`
	Name      *string            `json:"name,omitempty"`
	IsChecked *bool              `json:"is_checked,omitempty"`
	Signature *string            `json:"signature,omitempty"`
	Order     *uint32            `json:"order,omitempty"`
	Value     *string            `json:"value,omitempty"`
	// values must be equal in length to placements
	Values                 *[]string                          `json:"values,omitempty"`
	IsObligatory           *bool                              `json:"is_obligatory"`
	ShouldBeFilledBySender *bool                              `json:"should_be_filled_by_sender,omitempty"`
	EditableBySignatory    *bool                              `json:"editable_by_signatory,omitempty"`
	Placements             *[]*SignatoryFieldPlacement        `json:"placements,omitempty"`
	CustomValidation       *[]*SignatoryFieldCustomValidation `json:"custom_validation,omitempty"`
}

type SignatoryFieldPlacementTip = string

const (
	SignatoryFieldPlacementTipLeft  SignatoryFieldPlacementTip = "left"
	SignatoryFieldPlacementTipRight SignatoryFieldPlacementTip = "right"
)

type SignatoryFieldPlacement struct {
	XRel    float32                          `json:"xrel"`
	YRel    float32                          `json:"yrel"`
	WRel    float32                          `json:"wrel"`
	HRel    float32                          `json:"hrel"`
	FSRel   float32                          `json:"fsrel"`
	Page    uint32                           `json:"page"`
	Tip     *SignatoryFieldPlacementTip      `json:"tip"`
	Anchors []*SignatoryFieldPlacementAnchor `json:"anchors"`
}

type SignatoryFieldPlacementAnchor struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

type SignatoryFieldCustomValidation struct {
	Pattern         string `json:"pattern"`
	PositiveExample string `json:"positive_example"`
	Tooltip         string `json:"tooltip"`
}

type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserGroup struct {
	ID             string             `json:"id"`
	ParentID       string             `json:"parent_id"`
	Name           string             `json:"name"`
	Children       []*UserGroupChild  `json:"children"`
	Settings       *UserGroupSettings `json:"settings"`
	ContactDetails *ContactDetails    `json:"contact_details"`
}

type UserGroupChild struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserGroupSettings struct {
	InheritedFrom       *string              `json:"inherited_from"`
	DataRetentionPolicy *DataRetentionPolicy `json:"data_retention_policy"`
	InheritablePreview  *UserGroupSettings   `json:"inheritable_preview"`
}

type DataRetentionPolicy struct {
	IdleDocTimeoutPreparation *uint32 `json:"idle_doc_timeout_preparation"`
	IdleDocTimeoutClosed      *uint32 `json:"idle_doc_timeout_closed"`
	IdleDocTimeoutCanceled    *uint32 `json:"idle_doc_timeout_canceled"`
	IdleDocTimeoutTimedout    *uint32 `json:"idle_doc_timeout_timedout"`
	IdleDocTimeoutRejected    *uint32 `json:"idle_doc_timeout_rejected"`
	IdleDocTimeoutError       *uint32 `json:"idle_doc_timeout_error"`
	ImmediateTrash            *bool   `json:"immediate_trash"`
}

type ContactDetails struct {
	InheritedFrom      *string         `json:"inherited_from"`
	Address            *Address        `json:"address"`
	InheritablePreview *ContactDetails `json:"inheritable_preview"`
}

type Address struct {
	CompanyNumber *string `json:"company_number"`
	Address       *string `json:"address"`
	Zip           *string `json:"zip"`
	City          *string `json:"city"`
	Country       *string `json:"country"`
}
