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

type strDef interface {
	strp() *string
}

type SignatoryRole string

const (
	SignatoryRoleViewer       SignatoryRole = viewer
	SignatoryRoleSigningParty SignatoryRole = signingParty
	SignatoryRoleApprover     SignatoryRole = approver
)

func (s SignatoryRole) strp() *string {
	ss := string(s)
	return &ss
}

func (s SignatoryRole) Ptr() *SignatoryRole {
	return &s
}

type EmailDeliveryStatus string

const (
	EmailDeliveryStatusUnknown      EmailDeliveryStatus = unknown
	EmailDeliveryStatusNotDelivered EmailDeliveryStatus = notDelivered
	EmailDeliveryStatusDelivered    EmailDeliveryStatus = delivered
	EmailDeliveryStatusDeferred     EmailDeliveryStatus = deferred
)

func (s EmailDeliveryStatus) strp() *string {
	ss := string(s)
	return &ss
}

type ConfirmationEmailDeliveryStatus string

const (
	ConfirmationEmailDeliveryStatusUnknown      ConfirmationEmailDeliveryStatus = unknown
	ConfirmationEmailDeliveryStatusNotDelivered ConfirmationEmailDeliveryStatus = notDelivered
	ConfirmationEmailDeliveryStatusDelivered    ConfirmationEmailDeliveryStatus = delivered
	ConfirmationEmailDeliveryStatusDeferred     ConfirmationEmailDeliveryStatus = deferred
)

func (s ConfirmationEmailDeliveryStatus) strp() *string {
	ss := string(s)
	return &ss
}

type MobileDeliveryStatus string

const (
	MobileDeliveryStatusUnknown      MobileDeliveryStatus = unknown
	MobileDeliveryStatusNotDelivered MobileDeliveryStatus = notDelivered
	MobileDeliveryStatusDelivered    MobileDeliveryStatus = delivered
	MobileDeliveryStatusDeferred     MobileDeliveryStatus = deferred
)

func (s MobileDeliveryStatus) strp() *string {
	ss := string(s)
	return &ss
}

type DeliveryMethod string

const (
	DeliveryMethodEmail       DeliveryMethod = email
	DeliveryMethodMobile      DeliveryMethod = mobile
	DeliveryMethodEmailMobile DeliveryMethod = emailMobile
	DeliveryMethodPad         DeliveryMethod = pad
	DeliveryMethodAPI         DeliveryMethod = api
)

func (s DeliveryMethod) strp() *string {
	ss := string(s)
	return &ss
}

func (s DeliveryMethod) Ptr() *DeliveryMethod {
	return &s
}

type AuthenticationMethodToView string

const (
	AuthenticationMethodToViewStandard AuthenticationMethodToView = standard
	AuthenticationMethodToViewSmsPin   AuthenticationMethodToView = smsPin
	AuthenticationMethodToViewSEBankID AuthenticationMethodToView = seBankID
	AuthenticationMethodToViewNOBankID AuthenticationMethodToView = noBankID
	AuthenticationMethodToViewDKNemID  AuthenticationMethodToView = dkNemID
	AuthenticationMethodToViewFITupas  AuthenticationMethodToView = fiTupas
	AuthenticationMethodToViewVerimi   AuthenticationMethodToView = verimi
)

func (s AuthenticationMethodToView) strp() *string {
	ss := string(s)
	return &ss
}

func (s AuthenticationMethodToView) Ptr() *AuthenticationMethodToView {
	return &s
}

type AuthenticationMethodToViewArchived string

const (
	AuthenticationMethodToViewArchivedStandard AuthenticationMethodToViewArchived = standard
	AuthenticationMethodToViewArchivedSmsPin   AuthenticationMethodToViewArchived = smsPin
	AuthenticationMethodToViewArchivedSEBankID AuthenticationMethodToViewArchived = seBankID
	AuthenticationMethodToViewArchivedNOBankID AuthenticationMethodToViewArchived = noBankID
	AuthenticationMethodToViewArchivedDKNemID  AuthenticationMethodToViewArchived = dkNemID
	AuthenticationMethodToViewArchivedFITupas  AuthenticationMethodToViewArchived = fiTupas
	AuthenticationMethodToViewArchivedVerimi   AuthenticationMethodToViewArchived = verimi
)

func (s AuthenticationMethodToViewArchived) strp() *string {
	ss := string(s)
	return &ss
}

func (s AuthenticationMethodToViewArchived) Ptr() *AuthenticationMethodToViewArchived {
	return &s
}

type AuthenticationMethodToSign string

const (
	AuthenticationMethodToSignStandard AuthenticationMethodToSign = standard
	AuthenticationMethodToSignSmsPin   AuthenticationMethodToSign = smsPin
	AuthenticationMethodToSignSEBankID AuthenticationMethodToSign = seBankID
	AuthenticationMethodToSignNOBankID AuthenticationMethodToSign = noBankID
	AuthenticationMethodToSignDKNemID  AuthenticationMethodToSign = dkNemID
)

func (s AuthenticationMethodToSign) strp() *string {
	ss := string(s)
	return &ss
}

func (s AuthenticationMethodToSign) Ptr() *AuthenticationMethodToSign {
	return &s
}

type ConfirmationDeliveryMethod string

const (
	ConfirmationDeliveryMethodEmail           ConfirmationDeliveryMethod = email
	ConfirmationDeliveryMethodMobile          ConfirmationDeliveryMethod = mobile
	ConfirmationDeliveryMethodEmailMobile     ConfirmationDeliveryMethod = emailMobile
	ConfirmationDeliveryMethodEmailLink       ConfirmationDeliveryMethod = emailLink
	ConfirmationDeliveryMethodEmailLinkMobile ConfirmationDeliveryMethod = emailLinkMobile
	ConfirmationDeliveryMethodNone            ConfirmationDeliveryMethod = none
)

func (s ConfirmationDeliveryMethod) strp() *string {
	ss := string(s)
	return &ss
}

func (s ConfirmationDeliveryMethod) Ptr() *ConfirmationDeliveryMethod {
	return &s
}

type NotificationDeliveryMethod string

const (
	NotificationDeliveryMethodEmail           NotificationDeliveryMethod = email
	NotificationDeliveryMethodMobile          NotificationDeliveryMethod = mobile
	NotificationDeliveryMethodEmailMobile     NotificationDeliveryMethod = emailMobile
	NotificationDeliveryMethodEmailLink       NotificationDeliveryMethod = emailLink
	NotificationDeliveryMethodEmailLinkMobile NotificationDeliveryMethod = emailLinkMobile
	NotificationDeliveryMethodNone            NotificationDeliveryMethod = none
)

func (s NotificationDeliveryMethod) strp() *string {
	ss := string(s)
	return &ss
}

func (s NotificationDeliveryMethod) Ptr() *NotificationDeliveryMethod {
	return &s
}

type Status string

const (
	StatusPreparation   Status = "preparation"
	StatusPending       Status = "pending"
	StatusClosed        Status = "closed"
	StatusCanceled      Status = "canceled"
	StatusTimedout      Status = "timedout"
	StatusRejected      Status = "rejected"
	StatusDocumentError Status = "document_error"
)

func (s Status) strp() *string {
	ss := string(s)
	return &ss
}

type Lang string

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

func (s Lang) strp() *string {
	ss := string(s)
	return &ss
}

type ViewerRole string

const (
	ViewerRoleCompanyShared ViewerRole = "company_shared"
	ViewerRoleCompanyAdmin  ViewerRole = "company_admin"
	ViewerRoleSignatory     ViewerRole = "signatory"
)

func (s ViewerRole) strp() *string {
	ss := string(s)
	return &ss
}

func (s ViewerRole) Ptr() *ViewerRole {
	return &s
}

type Document struct {
	ID                  *string              `json:"id,omitempty"`
	Title               *string              `json:"title,omitempty"`
	Parties             *[]*Signatory        `json:"parties,omitempty"`
	File                *File                `json:"file,omitempty"`
	SealedFile          *File                `json:"sealed_file,omitempty"`
	AuthorAttachments   *[]*AuthorAttachment `json:"author_attachments,omitempty"`
	CTime               *string              `json:"ctime,omitempty"`
	MTime               *string              `json:"mtime,omitempty"`
	TimeoutTime         *string              `json:"timeout_time,omitempty"`
	AutoRemindTime      *string              `json:"auto_remind_time,omitempty"`
	Status              *Status              `json:"status,omitempty"`
	DaysToSign          *uint32              `json:"days_to_sign,omitempty"`
	DaysToRemind        *uint32              `json:"days_to_remind,omitempty"`
	DisplayOptions      *DisplayOptions      `json:"display_options,omitempty"`
	InvitationMessage   *string              `json:"invitation_message,omitempty"`
	ConfirmationMessage *string              `json:"confirmation_message,omitempty"`
	Lang                *Lang                `json:"lang,omitempty"`
	APICallbackURL      *string              `json:"api_callback_url,omitempty"`
	ObjectVersion       *uint64              `json:"object_version,omitempty"`
	AccessToken         *string              `json:"access_token,omitempty"`
	Timezone            *string              `json:"timezone,omitempty"`
	Tags                *[]*Tag              `json:"tags,omitempty"`
	IsTemplate          *bool                `json:"is_template,omitempty"`
	IsSaved             *bool                `json:"is_saved,omitempty"`
	FolderID            *string              `json:"folder_id,omitempty"`
	IsShared            *bool                `json:"is_shared,omitempty"`
	IsTrashed           *bool                `json:"is_trashed,omitempty"`
	IsDeleted           *bool                `json:"is_deleted,omitempty"`
	Viewer              *Viewer              `json:"viewer,omitempty"`
	ShareableLink       *string              `json:"shareable_link,omitempty"`
	TemplateID          *string              `json:"template_id,omitempty"`
	FromShareableLink   *bool                `json:"from_shareable_link,omitempty"`
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
	ShowHeader        *bool `json:"show_header,omitempty"`
	ShowPDFDownload   *bool `json:"show_pdf_download,omitempty"`
	ShowRejectOption  *bool `json:"show_reject_option,omitempty"`
	AllowRejectReason *bool `json:"allow_reject_reason,omitempty"`
	ShowFooter        *bool `json:"show_footer,omitempty"`
	DocumentIsReceipt *bool `json:"document_is_receipt,omitempty"`
	ShowArrow         *bool `json:"show_arrow,omitempty"`
}

type AuthorAttachment struct {
	Name            string `json:"name"`
	Required        bool   `json:"required"`
	AddToSealedFile bool   `json:"add_to_sealed_file"`
	FileID          string `json:"file_id"`
}

type Signatory struct {
	ID                                 *string                             `json:"id,omitempty"`
	UserID                             *string                             `json:"user_id,omitempty"`
	IsAuthor                           *bool                               `json:"is_author,omitempty"`
	IsSignatory                        *bool                               `json:"is_signatory,omitempty"`
	SignatoryRole                      *SignatoryRole                      `json:"signatory_role,omitempty"`
	Fields                             *[]*SignatoryField                  `json:"fields,omitempty"`
	ConsentModule                      *ConsentModule                      `json:"consent_module,omitempty"`
	SignOrder                          *uint32                             `json:"sign_order,omitempty"`
	SignTime                           *string                             `json:"sign_time,omitempty,omitempty"`
	SeenTime                           *string                             `json:"seen_time,omitempty,omitempty"`
	ReadInvitationTime                 *string                             `json:"read_invitation_time,omitempty"`
	RejectedTime                       *string                             `json:"rejected_time,omitempty,omitempty"`
	RejectionReason                    *string                             `json:"rejection_reason,omitempty"`
	SignSuccessRedirectURL             *string                             `json:"sign_success_redirect_url,omitempty"`
	RejectRedirectURL                  *string                             `json:"reject_redirect_url,omitempty"`
	EmailDeliveryStatus                *EmailDeliveryStatus                `json:"email_delivery_status,omitempty"`
	MobileDeliveryStatus               *MobileDeliveryStatus               `json:"mobile_delivery_status,omitempty"`
	ConfirmationEmailDeliveryStatus    *ConfirmationEmailDeliveryStatus    `json:"confirmation_email_delivery_status,omitempty"`
	HasAuthenticatedToView             *bool                               `json:"has_authenticated_to_view,omitempty"`
	CSV                                *[]string                           `json:"csv,omitempty"`
	DeliveryMethod                     *DeliveryMethod                     `json:"delivery_method,omitempty"`
	AuthenticationMethodToView         *AuthenticationMethodToView         `json:"authentication_method_to_view,omitempty"`
	AuthenticationMethodToViewArchived *AuthenticationMethodToViewArchived `json:"authentication_method_to_view_archived,omitempty"`
	AuthenticationMethodToSign         *AuthenticationMethodToSign         `json:"authentication_method_to_sign,omitempty"`
	ConfirmationDeliveryMethod         *ConfirmationDeliveryMethod         `json:"confirmation_delivery_method,omitempty"`
	NotificationDeliveryMethod         *NotificationDeliveryMethod         `json:"notification_delivery_method,omitempty"`
	AllowsHighlighting                 *bool                               `json:"allows_highlighting,omitempty"`
	HidePersonalNumber                 *bool                               `json:"hide_personal_number,omitempty"`
	CanForward                         *bool                               `json:"can_forward,omitempty"`
	Attachments                        *[]*Attachment                      `json:"attachments,omitempty"`
	HighlightedPages                   *[]*HighlightedPage                 `json:"highlighted_pages,omitempty"`
	APIDeliveryURL                     *string                             `json:"api_delivery_url,omitempty"`
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

type SignatoryFieldType string

const (
	SignatoryFieldTypeName           SignatoryFieldType = "name"
	SignatoryFieldTypeEmail          SignatoryFieldType = "email"
	SignatoryFieldTypeMobile         SignatoryFieldType = "mobile"
	SignatoryFieldTypeSignature      SignatoryFieldType = "signature"
	SignatoryFieldTypeCompanyName    SignatoryFieldType = "company"
	SignatoryFieldTypeCompanyNumber  SignatoryFieldType = "company_number"
	SignatoryFieldTypePersonalNumber SignatoryFieldType = "personal_number"
	SignatoryFieldTypeCheckbox       SignatoryFieldType = "checkbox"
	SignatoryFieldTypeRadiogroup     SignatoryFieldType = "radiogroup"
	SignatoryFieldTypeCustomText     SignatoryFieldType = "text"
)

func (s SignatoryFieldType) strp() *string {
	ss := string(s)
	return &ss
}

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

type SignatoryFieldPlacementTip string

const (
	SignatoryFieldPlacementTipLeft  SignatoryFieldPlacementTip = "left"
	SignatoryFieldPlacementTipRight SignatoryFieldPlacementTip = "right"
)

func (s SignatoryFieldPlacementTip) strp() *string {
	ss := string(s)
	return &ss
}

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
	IdleDocTimeoutPreparation *uint64 `json:"idle_doc_timeout_preparation,omitempty"`
	IdleDocTimeoutClosed      *uint64 `json:"idle_doc_timeout_closed,omitempty"`
	IdleDocTimeoutCanceled    *uint64 `json:"idle_doc_timeout_canceled,omitempty"`
	IdleDocTimeoutTimedout    *uint64 `json:"idle_doc_timeout_timedout,omitempty"`
	IdleDocTimeoutRejected    *uint64 `json:"idle_doc_timeout_rejected,omitempty"`
	IdleDocTimeoutError       *uint64 `json:"idle_doc_timeout_error,omitempty"`
	ImmediateTrash            *bool   `json:"immediate_trash,omitempty"`
}

type ContactDetails struct {
	InheritedFrom      *string         `json:"inherited_from"`
	Address            *Address        `json:"address"`
	InheritablePreview *ContactDetails `json:"inheritable_preview"`
}

type Address struct {
	CompanyNumber *string `json:"company_number,omitempty"`
	Address       *string `json:"address,omitempty"`
	Zip           *string `json:"zip,omitempty"`
	City          *string `json:"city,omitempty"`
	Country       *string `json:"country,omitempty"`
}

type HistoryStatus string

const (
	HistoryStatusInitiated       HistoryStatus = "initiated"
	HistoryStatusDraft           HistoryStatus = "draft"
	HistoryStatusCancelled       HistoryStatus = "cancelled"
	HistoryStatusRejected        HistoryStatus = "rejected"
	HistoryStatusTimeouted       HistoryStatus = "timeouted"
	HistoryStatusProblem         HistoryStatus = "problem"
	HistoryStatusDeliveryProblem HistoryStatus = "deliveryproblem"
	HistoryStatusSent            HistoryStatus = "sent"
	HistoryStatusDelivered       HistoryStatus = "delivered"
	HistoryStatusRead            HistoryStatus = "read"
	HistoryStatusOpened          HistoryStatus = "opened"
	HistoryStatusSigned          HistoryStatus = "signed"
	HistoryStatusProlonged       HistoryStatus = "prolonged"
	HistoryStatusSealed          HistoryStatus = "sealed"
	HistoryStatusExtended        HistoryStatus = "extended"
)

func (s HistoryStatus) strp() *string {
	ss := string(s)
	return &ss
}

type HistoryItem struct {
	Status HistoryStatus `json:"status"`
	Time   string        `json:"time"`
	Text   string        `json:"text"`
	Party  string        `json:"party"`
}

type DocumentHistory struct {
	Events []*HistoryItem `json:"events"`
}

type ListSortOrder string

const (
	ListSortAscending  ListSortOrder = "ascending"
	ListSortDescending ListSortOrder = "descending"
)

func (s ListSortOrder) strp() *string {
	ss := string(s)
	return &ss
}

type ListSortKey string

const (
	ListSortTitle  ListSortKey = "title"
	ListSortStatus ListSortKey = "status"
	ListSortMTime  ListSortKey = "mtime"
	ListSortAuthor ListSortKey = "author"
)

func (s ListSortKey) strp() *string {
	ss := string(s)
	return &ss
}

type ListSortParam struct {
	Order  ListSortOrder `json:"order"`
	SortBy ListSortKey   `json:"sort_by"`
}

type OAuthAuthorization struct {
	APIToken     string `json:"apitoken"`
	APISecret    string `json:"apisecret"`
	AccessToken  string `json:"accesstoken"`
	AccessSecret string `json:"accsesssecret"`
}

type Session struct {
	SessionID string `json:"session_id"`
}

type PersonalCredentialsToken struct {
	LoginToken     string `json:"login_token"`
	QRCode         string `json:"qr_code"`
	ExpirationTime string `json:"expiration_time"`
}

type Setup2FAResponse struct {
	TwofactorAlive bool    `json:"twofactor_alive"`
	QRCode         *string `json:"qr_code"`
}

type Confirm2FAResp struct {
	TwofactorAlive bool `json:"twofactor_alive"`
	TotpValid      bool `json:"totp_valid"`
}

type Disable2FAResp struct {
	TwofactorAlive bool `json:"twofactor_alive"`
}

type IsUserDeletableResp struct {
	Deletable bool    `json:"deletable"`
	Reason    *string `json:"reason"`
}

type AccessRoleSource struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type AccessRoleTarget struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type AccessRoleActions struct {
	Document          []string `json:"document"`
	FolderPolicy      []string `json:"folder_policy"`
	User              []string `json:"user"`
	UserGroup         []string `json:"user_group"`
	UserGroupPolicy   []string `json:"user_group_policy"`
	UserPersonalToken []string `json:"user_personal_token"`
	UserPolicy        []string `json:"user_policy"`
}

type AccessRole struct {
	ID             *string           `json:"id"`
	IsGenerated    bool              `json:"is_generated"`
	RoleType       string            `json:"role_type"`
	Source         AccessRoleSource  `json:"source"`
	AllowedActions AccessRoleActions `json:"allowed_actions"`
}
