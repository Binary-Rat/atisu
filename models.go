package atisu

type City struct {
	CityID    int    `json:"city_id"`
	IsSuccess bool   `json:"is_success"`
	Street    string `json:"street"`
}

type Cities map[string]City

type Response struct {
	TotalCount  int                     `json:"total_count"`
	HiddenCount int                     `json:"hidden_count"`
	Accounts    map[string]Account      `json:"accounts"`
	Trucks      []Truck                 `json:"trucks"`
	Responses   map[string]ResponseData `json:"responses"`
}

type Account struct {
	AtiID                   string    `json:"ati_id"`
	FirmName                string    `json:"firm_name"`
	FullFirmName            string    `json:"full_firm_name"`
	Brand                   string    `json:"brand"`
	Ownership               string    `json:"ownership"`
	OwnershipID             int       `json:"ownership_id"`
	FirmTypeID              int       `json:"firm_type_id"`
	FirmType                string    `json:"firm_type"`
	ClaimsSum               float64   `json:"claims_sum"`
	ClaimsCount             int       `json:"claims_count"`
	BadPartnerMentionsCount int       `json:"bad_partner_mentions_count"`
	BadPartnerFirmsCount    int       `json:"bad_partner_firms_count"`
	RecommendationCount     int       `json:"recommendation_count"`
	RecommendationsCount    int       `json:"recommendations_count"`
	IsAtiPartner            bool      `json:"is_ati_partner"`
	IsOdksMember            bool      `json:"is_odks_member"`
	CityID                  int       `json:"city_id"`
	Passport                Passport  `json:"passport"`
	Contacts                []Contact `json:"contacts"`
	UserLists               UserLists `json:"user_lists"`
	Comment                 string    `json:"comment"`
}

type Passport struct {
	AtiDataMatchPoint      float64 `json:"ati_data_match_point"`
	AccountLifetimePoint   float64 `json:"account_lifetime_point"`
	BusinessActivityPoint  float64 `json:"business_activity_point"`
	RoundTablePoint        float64 `json:"round_table_point"`
	ClaimPoint             float64 `json:"claim_point"`
	ProfActivityPoint      float64 `json:"prof_activity_point"`
	AtiAdministrationPoint float64 `json:"ati_administration_point"`
	ClonesPoint            float64 `json:"clones_point"`
	EgrPoint               float64 `json:"egr_point"`
	MassRegistrationPoint  float64 `json:"mass_registration_point"`
	MassFounderPoint       float64 `json:"mass_founder_point"`
	FirmLifetimePoint      float64 `json:"firm_lifetime_point"`
	NegativePointsSum      float64 `json:"negative_points_sum"`
	TotalScore             float64 `json:"total_score"`
	Status                 int     `json:"status"`
	StatusDescription      string  `json:"status_description"`
}

type Contact struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Telephone       string `json:"telephone"`
	TelephoneBrand  string `json:"telephone_brand"`
	TelephoneRegion string `json:"telephone_region"`
	Email           bool   `json:"email"`
	ICQ             string `json:"icq"`
	Mobile          string `json:"mobile"`
	MobileBrand     string `json:"mobile_brand"`
	MobileRegion    string `json:"mobile_region"`
	SkypeName       string `json:"skype_name"`
	Fax             string `json:"fax"`
	MobileOperator  string `json:"mobile_operator"`
	CityID          int    `json:"city_id"`
}

type UserLists struct {
	White          bool        `json:"white"`
	Neutral        bool        `json:"neutral"`
	Black          bool        `json:"black"`
	EmojiFirmLists []EmojiList `json:"emoji_firm_lists"`
}

type EmojiList struct {
	Emoji           string     `json:"emoji"`
	TotalListsCount int        `json:"total_lists_count"`
	HeadLists       []HeadList `json:"head_lists"`
}

type HeadList struct {
	ListName string `json:"list_name"`
	ListID   string `json:"list_id"`
	Emoji    string `json:"emoji"`
	Type     int    `json:"type"`
	Fixed    bool   `json:"fixed"`
}

type Truck struct {
	AtiID          string        `json:"ati_id"`
	TruckID        string        `json:"truck_id"`
	AddedAt        string        `json:"added_at"`
	UpdatedAt      string        `json:"updated_at"`
	ContactID1     int           `json:"contact_id1"`
	ContactID2     int           `json:"contact_id2"`
	DepartmentID   int           `json:"department_id"`
	Note           string        `json:"note"`
	FirstDate      string        `json:"first_date"`
	LastDate       string        `json:"last_date"`
	PeriodicityID  int           `json:"periodicity_id"`
	DateType       int           `json:"date_type"`
	TruePrice      float64       `json:"true_price"`
	TrueCurrencyID int           `json:"true_currency_id"`
	Transport      Transport     `json:"transport"`
	Loading        Loading       `json:"loading"`
	UnloadingList  []Unloading   `json:"unloading_list"`
	PaymentOptions PaymentOption `json:"payment_options"`
	ShowContacts   bool          `json:"show_contacts"`
	Comments       []Comment     `json:"comments"`
}

type Transport struct {
	CarType       int     `json:"car_type"`
	BodyType      int     `json:"body_type"`
	LoadingType   int     `json:"loading_type"`
	Weight        float64 `json:"weight"`
	Volume        float64 `json:"volume"`
	TruckLength   float64 `json:"truck_length"`
	TruckWidth    float64 `json:"truck_width"`
	TruckHeight   float64 `json:"truck_height"`
	TrailerLength float64 `json:"trailer_length"`
	TrailerWidth  float64 `json:"trailer_width"`
	TrailerHeight float64 `json:"trailer_height"`
	Hydrolift     bool    `json:"hydrolift"`
	PartialLoad   bool    `json:"partial_load"`
	Koniki        bool    `json:"koniki"`
	TIR           bool    `json:"tir"`
	EKMT          bool    `json:"ekmt"`
	ADRTypes      int     `json:"adr_types"`
}

type Loading struct {
	CityID   int `json:"city_id"`
	Radius   int `json:"radius"`
	Distance int `json:"distance"`
}

type Unloading struct {
	Main          bool    `json:"main"`
	Visible       bool    `json:"visible"`
	CashSum       float64 `json:"cash_sum"`
	SumWithNDS    float64 `json:"sum_with_nds"`
	SumWithoutNDS float64 `json:"sum_without_nds"`
	CurrencyID    int     `json:"currency_id"`
	PointType     int     `json:"point_type"`
	PointID       int     `json:"point_id"`
	Radius        int     `json:"radius"`
	Distance      int     `json:"distance"`
}

type PaymentOption struct {
	Card bool `json:"card"`
	Torg bool `json:"torg"`
}

type Comment struct {
	CommentID       string `json:"comment_id"`
	Message         string `json:"message"`
	ContactID       int    `json:"contact_id"`
	ContactName     string `json:"contact_name"`
	CommentDateTime string `json:"comment_date_time"`
}

type ResponseData struct {
	Offers     []Offer     `json:"offers"`
	Absents    []Absent    `json:"absents"`
	Complaints []Complaint `json:"complaints"`
}

type Offer struct {
	TruckID          string         `json:"truck_id"`
	AdditionDate     string         `json:"addition_date"`
	ChangeDate       string         `json:"change_date"`
	CashSum          float64        `json:"cash_sum"`
	CurrencyID       int            `json:"currency_id"`
	SumWithNDS       float64        `json:"sum_with_nds"`
	NDSCurrencyID    int            `json:"nds_currency_id"`
	SumWithoutNDS    float64        `json:"sum_without_nds"`
	NotNDSCurrencyID int            `json:"not_nds_currency_id"`
	Note             string         `json:"note"`
	PrepayPercent    int            `json:"prepay_percent"`
	DelayPaymentDays int            `json:"delay_payment_days"`
	LoadingCityID    int            `json:"loading_city_id"`
	UnloadingCityID  int            `json:"unloading_city_id"`
	Weight           float64        `json:"weight"`
	Volume           float64        `json:"volume"`
	CarDeliveryDate  string         `json:"car_delivery_date"`
	UnloadPayment    bool           `json:"unload_payment"`
	Coloading        bool           `json:"coloading"`
	SenderFirmInfo   SenderFirmInfo `json:"sender_firm_info"`
}

type Absent struct {
	TruckID        string         `json:"truck_id"`
	AddDate        string         `json:"add_date"`
	SenderFirmInfo SenderFirmInfo `json:"sender_firm_info"`
}

type Complaint struct {
	TruckID        string         `json:"truck_id"`
	AddDate        string         `json:"add_date"`
	ChangeDate     string         `json:"change_date"`
	Note           string         `json:"note"`
	SenderFirmInfo SenderFirmInfo `json:"sender_firm_info"`
}

type SenderFirmInfo struct {
	AtiID     string      `json:"ati_id"`
	ContactID int         `json:"contact_id"`
	Contact   ContactInfo `json:"contact"`
	Name      string      `json:"name"`
	Score     float64     `json:"score"`
	Status    int         `json:"status"`
}

type ContactInfo struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	MobilePhone string `json:"mobile_phone"`
	City        string `json:"city"`
	ICQ         string `json:"icq"`
	SkypeName   string `json:"skype_name"`
	Fax         string `json:"fax"`
	Email       string `json:"email"`
}

type DateOption struct {
	DateOption string `json:"date_option"`
}

type CityFilter struct {
	ID   int `json:"id"`
	Type int `json:"type"`
}

type Weight struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type Volume struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type Filter struct {
	Dates       DateOption `json:"dates"`
	From        CityFilter `json:"from"`
	To          CityFilter `json:"to"`
	Weight      Weight     `json:"weight"`
	Volume      Volume     `json:"volume"`
	TruckType   int        `json:"truck_type"`
	LoadingType int        `json:"loading_type"`
	SortingType int        `json:"sorting_type"`
}

type requestCars struct {
	Page         int    `json:"page"`
	ItemsPerPage int    `json:"items_per_page"`
	Filter       Filter `json:"filter"`
}
