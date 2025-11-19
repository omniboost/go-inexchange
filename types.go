package inexchange

import (
	"strings"
	"time"
)

type CommaSeparatedQueryParam []string

func (t CommaSeparatedQueryParam) MarshalSchema() string {
	return strings.Join(t, ",")
}

type PaginatedResponse[T any] struct {
	Results []T `json:"results"`
	Paging  struct {
		PageNumber int `json:"pageNumber"`
		PageSize   int `json:"pageSize"`
		TotalCount int `json:"totalCount"`
	} `json:"paging"`
}

type ConfigurationSupportedType string

type AlternateName struct {
	LanguageCode string `json:"languageCode"`
	Name         string `json:"name"`
}

type LocalizedString struct {
	LanguageCode string `json:"languageCode"`
	Content      string `json:"content"`
}

type AddressDetail struct {
	AddressLine1    string `json:"addressLine1"`
	AddressLine2    string `json:"addressLine2"`
	City            string `json:"city"`
	PostalCode      string `json:"postalCode"`
	StateCode       string `json:"stateCode"`
	CountryCode     string `json:"countryCode"`
	AddressTypeCode string `json:"addressTypeCode"`
}

type CommunicationDetail struct {
	Primary               bool   `json:"primary"`
	CommunicationTypeCode string `json:"communicationTypeCode"`
	Value                 string `json:"value"`
}

type CommunicationChannel struct {
	IsPrimary bool   `json:"isPrimary"`
	Details   string `json:"details"`
	ModeCode  string `json:"modeCode"`
}

type ContactDetail struct {
	Primary  bool   `json:"primary"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	JobTitle string `json:"jobTitle"`
}

type PropertyListItem struct {
	ID              string `json:"id"`
	Version         int64  `json:"version"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	LegalName       string `json:"legalName"`
	ParentID        string `json:"parentId"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	BusinessAddress struct {
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		City         string `json:"city"`
		PostalCode   string `json:"postalCode"`
		StateCode    string `json:"stateCode"`
		StateName    string `json:"stateName"`
		CountryCode  string `json:"countryCode"`
		CountryName  string `json:"countryName"`
	} `json:"businessAddress"`
	TaxID           string                 `json:"taxId"`
	StarRating      float64                `json:"starRating"`
	TypeID          string                 `json:"typeId"`
	StatusCode      string                 `json:"statusCode"`
	CustomFields    map[string]interface{} `json:"customFields"`
	DisplaySequence int32                  `json:"displaySequence"`
	RegionCode      string                 `json:"regionCode"`
	BrandID         string                 `json:"brandId"`
}

type Property struct {
	ID                     string                `json:"id"`
	Version                int64                 `json:"version"`
	Code                   string                `json:"code"`
	Name                   string                `json:"name"`
	ParentID               string                `json:"parentId"`
	DefaultLanguageCode    string                `json:"defaultLanguageCode"`
	TaxID                  string                `json:"taxId"`
	LegalName              string                `json:"legalName"`
	AlternateNames         []AlternateName       `json:"alternateNames"`
	SupportedLanguageCodes []string              `json:"supportedLanguageCodes"`
	AddressList            []AddressDetail       `json:"addressList"`
	CommunicationDetails   []CommunicationDetail `json:"communicationDetails"`
	Contacts               []ContactDetail       `json:"contacts"`
	LegalOwner             string                `json:"legalOwner"`
	Location               struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	CheckInTime              string                 `json:"checkInTime"`
	CheckOutTime             string                 `json:"checkOutTime"`
	FutureOperationalWindow  int32                  `json:"futureOperationalWindow"`
	AlternativePropertyCodes []string               `json:"alternativePropertyCodes"`
	AlternativePropertyIDs   []string               `json:"alternativePropertyIds"`
	CurrencyCode             string                 `json:"currencyCode"`
	CustomCurrencyFormat     string                 `json:"customCurrencyFormat"`
	ShortDataFormat          string                 `json:"shortDataFormat"`
	LongDateFormat           string                 `json:"longDateFormat"`
	TimeFormatCode           string                 `json:"timeFormatCode"`
	TimeZoneCode             string                 `json:"timeZoneCode"`
	StatusCode               string                 `json:"statusCode"`
	FiscalReportingPeriodID  string                 `json:"fiscalReportingPeriodId"`
	StarRating               float64                `json:"starRating"`
	TypeID                   string                 `json:"typeId"`
	BrandID                  string                 `json:"brandId"`
	OperatingModelCode       string                 `json:"operatingModelCode"`
	AttributeIDs             []string               `json:"attributeIds"`
	DistributionChannelIDs   []string               `json:"distributionChannelIds"`
	CustomFields             map[string]interface{} `json:"customFields"`
	ProfileStore             struct {
		PrimaryID string   `json:"primaryId"`
		OtherIDs  []string `json:"otherIds"`
	} `json:"profileStore"`

	EffectiveZoneIDs    []string `json:"effectiveZoneIds"`
	DisplaySequence     int32    `json:"displaySequence"`
	RegionCode          string   `json:"regionCode"`
	InventoryAllocation struct {
		StartTime string `json:"startTime"`
		EndTIme   string `json:"endTime"`
	} `json:"inventoryAllocation"`

	SupportedQuotationCurrencies        []string `json:"supportedQuotationCurrencies"`
	RatePlanQuotationCurrencies         []string `json:"ratePlanQuotationCurrencies"`
	PurchaseElementQuotationCurrencies  []string `json:"purchaseElementQuotationCurrencies"`
	MeetingAndEventsQuotationCurrencies []string `json:"meetingAndEventsQuotationCurrencies"`
	FirstDayOfWeek                      string   `json:"firstDayOfWeek"`
}

type TransactionCode struct {
	ID                     string            `json:"id"`
	Version                int64             `json:"version"`
	ZoneID                 string            `json:"zoneId"`
	Code                   string            `json:"code"`
	Description            []LocalizedString `json:"description"`
	IsActive               bool              `json:"isActive"`
	ForceToUse             bool              `json:"forceToUse"`
	DisplaySequence        int32             `json:"displaySequence"`
	SubCode                string            `json:"subCode"`
	CurrencyCode           string            `json:"currencyCode"`
	AcceptAsCashierPayment bool              `json:"acceptAsCashierPayment"`
	AllowToPostInAR        bool              `json:"allowToPostInAR"`
	OnlinePaymentMethod    string            `json:"onlinePaymentMethod"`
	AcceptAsDepositPayment bool              `json:"acceptAsDepositPayment"`
	AllowManualPosting     bool              `json:"allowManualPosting"`
	IsAdvancedDeposit      bool              `json:"isAdvancedDeposit"`
	IsCheckNumberMandatory bool              `json:"isCheckNumberMandatory"`
	PostingValueRange      struct {
		From float64 `json:"from"`
		To   float64 `json:"to"`
	} `json:"postingValueRange"`

	DefaultPostingValue      float64                `json:"defaultPostingValue"`
	AdjustmentID             string                 `json:"adjustmentId"`
	CustomFieldContainerJSON map[string]interface{} `json:"customFieldContainerJson"`
	IsLocked                 bool                   `json:"isLocked"`
	CreditCardSurcharge      struct {
		Percentage        float64 `json:"percentage"`
		TransactionCodeID string  `json:"transactionCodeId"`
	} `json:"creditCardSurcharge"`

	TransactionType         string `json:"transactionType"`
	IsNonRevenue            bool   `json:"isNonRevenue"`
	IsPostingAllowed        bool   `json:"isPostingAllowed"`
	TransactionSubGroupID   string `json:"transactionSubGroupId"`
	ParentTransactionCodeID string `json:"parentTransactionCodeId"`
	PaymentTypeCode         string `json:"paymentTypeCode"`
	CreditCardTypeCode      string `json:"creditCardTypeCode"`
	HasSubCodes             bool   `json:"hasSubCodes"`
	OfflinePaymentMethod    string `json:"offlinePaymentMethod"`
	IsGuestLedgerPrepayment bool   `json:"isGuestLedgerPrepayment"`
}

type TaxRule struct {
	ID                    string            `json:"id"`
	Version               int64             `json:"version"`
	ZoneID                string            `json:"zoneId"`
	Code                  string            `json:"code"`
	Description           []LocalizedString `json:"description"`
	IsActive              bool              `json:"isActive"`
	ForceToUse            bool              `json:"forceToUse"`
	DisplaySequence       int32             `json:"displaySequence"`
	IsCharge              bool              `json:"isCharge"`
	Category              string            `json:"category"`
	FlatAmount            float64           `json:"flatAmount"`
	PercentageAmount      float64           `json:"percentageAmount"`
	IsPercentage          bool              `json:"isPercentage"`
	AmountDefinition      string            `json:"amountDefinition"`
	AmountRangeDefinition struct {
		RateType         string `json:"rateType"`
		PriceType        string `json:"priceType"`
		BaseAmountType   string `json:"baseAmountType"`
		BaseAmountRanges []struct {
			From                    float64 `json:"from"`
			To                      float64 `json:"to"`
			ValueRangeType          string  `json:"valueRangeType"`
			Value                   float64 `json:"value"`
			IsIncrementationDefined bool    `json:"isIncrementationDefined"`
			IncrementationValue     float64 `json:"incrementationValue"`
			ThresholdValue          float64 `json:"thresholdValue"`
		} `json:"baseAmountRanges"`

		IsPercentage bool `json:"isPercentage"`
	} `json:"amountRangeDefinition"`

	BaseRuleID                      string `json:"baseRuleId"`
	IsBaseRule                      bool   `json:"isBaseRule"`
	IsTaxIncluded                   bool   `json:"isTaxIncluded"`
	PrintSeparately                 bool   `json:"printSeparately"`
	PostSeparately                  bool   `json:"postSeparately"`
	ChargeOnManualPosting           bool   `json:"chargeOnManualPosting"`
	ChargeMethodPerPersonDefinition struct {
		IsAdult    bool `json:"isAdult"`
		IsChildren bool `json:"isChildren"`
	} `json:"chargeMethodPerPersonDefinition"`

	ValidFromDate                Date     `json:"validFromDate"`
	ValidToDate                  Date     `json:"validToDate"`
	TransactionCodeID            string   `json:"transactionCodeId"`
	ApplicableTransactionCodeIDs []string `json:"applicableTransactionCodeIds"`
	AgeBucketIDs                 []string `json:"ageBucketIds"`
	LengthOfStay                 int32    `json:"lengthOfStay"`
	IsCommission                 bool     `json:"isCommission"`
	PurposeOfStayIDs             []string `json:"purposeOfStayIds"`
	ExcludeDayUse                bool     `json:"excludeDayUse"`
	ChargeMethod                 string   `json:"chargeMethod"`
}

type SubsidiaryListItem struct {
	ID              string `json:"id"`
	Version         int64  `json:"version"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	ParentID        string `json:"parentId"`
	ParentCode      string `json:"parentCode"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	BusinessAddress struct {
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		City         string `json:"city"`
		PostalCode   string `json:"postalCode"`
		StateCode    string `json:"stateCode"`
		StateName    string `json:"stateName"`
		CountryCode  string `json:"countryCode"`
		CountryName  string `json:"countryName"`
	} `json:"businessAddress"`

	TaxID           string `json:"taxId"`
	StatusCode      string `json:"statusCode"`
	DisplaySequence int32  `json:"displaySequence"`
}

type AccountGroup struct {
	ID                             string            `json:"id"`
	Version                        int64             `json:"version"`
	ZoneID                         string            `json:"zoneId"`
	Code                           string            `json:"code"`
	Description                    []LocalizedString `json:"description"`
	IsActive                       bool              `json:"isActive"`
	DisplaySequence                int32             `json:"displaySequence"`
	TypeCode                       string            `json:"typeCode"`
	NumberCycleID                  string            `json:"numberCycleId"`
	ExcludeFromZeroCheckOutDefault bool              `json:"excludeFromZeroCheckOutDefault"`
}

type TransactionGroup struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	ZoneID      string `json:"zoneId"`
	Version     int64  `json:"version"`
	IsActive    bool   `json:"isActive"`
	IsPayment   bool   `json:"isPayment"`
	ForceToUse  bool   `json:"forceToUse"`
	Description []struct {
		Content      string `json:"content"`
		LanguageCode string `json:"languageCode"`
	} `json:"description"`
	DisplaySequence int32  `json:"displaySequence"`
	TransactionType string `json:"transactionType"`
}

type TransactionSubGroup struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	ZoneID      string `json:"zoneId"`
	Version     int64  `json:"version"`
	IsActive    bool   `json:"isActive"`
	ForceToUse  bool   `json:"forceToUse"`
	Description []struct {
		Content      string `json:"content"`
		LanguageCode string `json:"languageCode"`
	} `json:"description"`
	DisplaySequence    int32  `json:"displaySequence"`
	TransactionGroupID string `json:"transactionGroupId"`
}

type Customer struct {
	ID              string `json:"id"`
	Version         int64  `json:"version"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	TenantName      string `json:"tenantName"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	BusinessAddress []struct {
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		City         string `json:"city"`
		PostalCode   string `json:"postalCode"`
		StateCode    string `json:"stateCode"`
		StateName    string `json:"stateName"`
		CountryCode  string `json:"countryCode"`
		CountryName  string `json:"countryName"`
	} `json:"businessAddress"`

	TaxID           string `json:"taxId"`
	StatusCode      string `json:"statusCode"`
	PropertiesCount int32  `json:"propertiesCount"`
	NonProduction   bool   `json:"nonProduction"`
	AuditData       struct {
		CreatorID      string   `json:"creatorId"`
		CreationTime   DateTime `json:"creationTime"`
		LastUpdaterID  string   `json:"lastUpdaterId"`
		LastUpdateTime DateTime `json:"lastUpdateTime"`
	} `json:"auditData"`

	CreatorName struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"creatorName"`

	RegionCodes []string `json:"regionCodes"`
}

type UserDetails struct {
	ID                        string   `json:"id"`
	AccountLoginProcedureType string   `json:"accountLoginProcedureType"`
	ExternalID                string   `json:"externalId"`
	Version                   int64    `json:"version"`
	UserName                  string   `json:"userName"`
	Email                     string   `json:"email"`
	IsSystemUser              bool     `json:"isSystemUser"`
	TenantID                  string   `json:"tenantId"`
	ActiveFrom                DateTime `json:"activeFrom"`
	ActiveTo                  DateTime `json:"activeTo"`
	DeactivationComment       string   `json:"deactivationComment"`
	BlockedTo                 DateTime `json:"blockedTo"`
	LockoutTo                 DateTime `json:"lockoutTo"`
	BlockComment              string   `json:"blockComment"`
	PasswordChangeRequired    bool     `json:"passwordChangeRequired"`
	PersonalData              struct {
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		GenderCode     string `json:"genderCode"`
		BirthDate      Date   `json:"birthDate"`
		EmployeeNumber string `json:"employeeNumber"`
		HireDate       Date   `json:"hireDate"`
	} `json:"personalData"`

	Address struct {
		CountryCode  string `json:"countryCode"`
		StateCode    string `json:"stateCode"`
		City         string `json:"city"`
		PostalCode   string `json:"postalCode"`
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
	} `json:"address"`

	CommunicationChannels         []CommunicationDetail `json:"communicationChannels"`
	PasswordExpirationDisabled    bool                  `json:"passwordExpirationDisabled"`
	AutomaticDeactivationDisabled bool                  `json:"automaticDeactivationDisabled"`
	AuthenticatorRequired         bool                  `json:"authenticatorRequired"`
}

type UnitAssignmentListItem struct {
	UnitID          string `json:"unitId"`
	ParentID        string `json:"parentId"`
	Type            string `json:"type"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	IsPublished     bool   `json:"isPublished"`
	RegionCode      string `json:"regionCode"`
	DisplaySequence int32  `json:"displaySequence"`
}

type ContactProfileResponse struct {
	ContactProfileID                  string   `json:"contactProfileId"`
	ContactProfileRole                string   `json:"contactProfileRole"`
	IsPrimary                         bool     `json:"isPrimary"`
	Roles                             []string `json:"roles"`
	ContactProfileAvailabilityDetails struct {
		ProfileStoreID                   string   `json:"profileStoreId"`
		RegionCode                       string   `json:"regionCode"`
		IsAvailableOutsideOfSourceRegion bool     `json:"isAvailableOutsideOfSourceRegion"`
		AvailableInRegions               []string `json:"availableInRegions"`
	} `json:"contactProfileAvailabilityDetails"`
}

type CodeIDValue struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type LocalizedIDCodeValue struct {
	ID           string `json:"id"`
	Code         string `json:"code"`
	LanguageCode string `json:"languageCode"`
	Description  string `json:"description"`
}

type ProfileMetadata struct {
	CreatorID            string   `json:"creatorId"`
	CreationTime         DateTime `json:"creationTime"`
	LastUpdaterID        string   `json:"lastUpdaterId"`
	LastUpdateTime       DateTime `json:"lastUpdateTime"`
	CreationPropertyID   string   `json:"creationPropertyId"`
	LastUpdatePropertyID string   `json:"lastUpdatePropertyId"`
}

type ProfileAddress struct {
	AddressTypeCode struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"addressTypeCode"`

	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City         string `json:"city"`
	PostalCode   string `json:"postalCode"`
	StateCode    struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"stateCode"`

	CountryCode struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"countryCode"`

	LanguageCode struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"languageCode"`

	IsPrimary   bool `json:"isPrimary"`
	IsArPrimary bool `json:"isArPrimary"`
	DistrictID  struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"districtId"`

	ID string `json:"id"`
}

type ProfileCommunicationChannel struct {
	ID       string `json:"id"`
	TypeCode struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"typeCode"`

	ModeCode struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"modeCode"`

	Details   string `json:"details"`
	IsPrimary bool   `json:"isPrimary"`
}

type ProfilesExternalIdentifier struct {
	ExternalSystemID string `json:"externalSystemId"`
	Identifier       string `json:"identifier"`
}

type CompanyProfile struct {
	ID       string `json:"id"`
	Version  string `json:"version"`
	TypeCode struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"typeCode"`

	LinkedProfileCount int32 `json:"linkedProfileCount"`
	LinkedProfileIDs   []struct {
		ProfileRoleCode string `json:"profileRoleCode"`
		ProfileID       string `json:"profileId"`
	} `json:"linkedProfileIds"`

	Status struct {
		IsActive               bool `json:"isActive"`
		DeactivationReasonCode struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"deactivationReasonCode"`

		IsRestricted          bool `json:"isRestricted"`
		RestrictionReasonCode struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"restrictionReasonCode"`

		RestrictionComment string `json:"restrictionComment"`
		IsLocked           bool   `json:"isLocked"`
		LockedReason       string `json:"lockedReason"`
		ForReview          bool   `json:"forReview"`
	} `json:"status"`

	Details struct {
		FullName      string `json:"fullName"`
		Abbreviation  string `json:"abbreviation"`
		CorporateID   string `json:"corporateId"`
		TaxID         string `json:"taxId"`
		IndustryCodes []struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"industryCodes"`

		BusinessSegmentCodes []struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"businessSegmentCodes"`

		PreferredLanguageProfileDetails struct {
			LanguageCode struct {
				Code          string `json:"code"`
				LanguageCodee string `json:"languageCode"`
				Description   string `json:"description"`
			} `json:"languageCode"`

			FullName string `json:"fullName"`
		} `json:"preferredLanguageProfileDetails"`

		PreferredCommunicationLanguageCode struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"preferredCommunicationLanguageCode"`
	} `json:"details"`

	Addresses             []ProfileAddress              `json:"addresses"`
	CommunicationChannels []ProfileCommunicationChannel `json:"communicationChannels"`

	Metadata ProfileMetadata `json:"metadata"`

	NotDuplicateProfiles struct {
		Collection []struct {
			NotDuplicateProfileID string `json:"notDuplicateProfileId"`
			ProfileID             string `json:"profileId"`
			ProfileRoleCode       string `json:"profileRoleCode"`
		}

		Count int32 `json:"count"`
	} `json:"notDuplicateProfiles"`

	FreezeStatus struct {
		IsFrozen   bool   `json:"isFrozen"`
		Reason     string `json:"reason"`
		PropertyID string `json:"propertyId"`
		Source     string `json:"source"`
	} `json:"freezeStatus"`

	CustomFieldContainerJSON  string                       `json:"customFieldContainerJson"`
	ExternalSystemIdentifiers []ProfilesExternalIdentifier `json:"externalSystemIdentifiers"`
	NotesIdentifiers          []string                     `json:"notesIdentifiers"`
	ContactProfiles           []ContactProfileResponse     `json:"contactProfiles"`
	ContactFor                []ContactProfileResponse     `json:"contactFor"`
	ProfileOwners             struct {
		Primary struct {
			ID string `json:"id"`
		} `json:"primary"`

		Others []struct {
			ID string `json:"id"`
		} `json:"others"`
	} `json:"profileOwners"`

	RelationshipDetails struct {
		IsRoot bool `json:"isRoot"`
	} `json:"relationshipDetails"`

	ExternalAccountReceivable struct {
		ExternalAccountReceivableNumber string `json:"externalAccountReceivableNumber"`
	} `json:"externalAccountReceivable"`

	HasImage bool `json:"hasImage"`
}

type CompanyProfileV0 struct {
	ID           string `json:"id"`
	IsActive     bool   `json:"isActive"`
	IsRestricted bool   `json:"isRestricted"`
	Version      int    `json:"version"`
	TypeCode     string `json:"typeCode"`
	TypeID       string `json:"typeId"`

	Details struct {
		FullName                           string   `json:"fullName"`
		TaxID                              string   `json:"taxId"`
		BusinessSegmentCodes               []string `json:"businessSegmentCodes"`
		PreferredCommunicationLanguageCode string   `json:"preferredCommunicationLanguageCode"`
		BusinessSegmentIds                 []string `json:"businessSegmentIds"`
	} `json:"details"`

	CommunicationDetails []struct {
		ID       string `json:"id"`
		TypeID   string `json:"typeId"`
		TypeCode string `json:"typeCode"`
		Mode     string `json:"mode"`
		Details  string `json:"details"`
		Primary  bool   `json:"primary"`
	} `json:"communicationDetails"`

	Addresses []struct {
		ID           string `json:"id"`
		CountryCode  string `json:"countryCode"`
		City         string `json:"city"`
		PostCode     string `json:"postCode"`
		AddressLine1 string `json:"addressLine1"`
		TypeCode     string `json:"typeCode"`
		TypeID       string `json:"typeId"`
		Primary      bool   `json:"primary"`
		LanguageCode string `json:"languageCode"`
		ArPrimary    bool   `json:"arPrimary"`

		CountrySpecificFields []struct {
			CountrySpecificField LocalizedIDCodeValue `json:"countrySpecificField"`

			Value struct {
				Type      string `json:"type"`
				String    string `json:"string"`
				LocalDate Date   `json:"localDate"`
				Array     []any  `json:"array"`
			}
		} `json:"countrySpecificFields"`
	} `json:"addresses"`

	IsLocked  bool `json:"isLocked"`
	ForReview bool `json:"forReview"`

	NotDuplicateProfiles struct {
		Collection []struct {
			NotDuplicateProfileID string `json:"notDuplicateProfileId"`
			ProfileID             string `json:"profileId"`
			ProfileRoleCode       string `json:"profileRoleCode"`
		} `json:"collection"`

		NotDuplicateProfilesCount int `json:"notDuplicateProfilesCount"`
	} `json:"notDuplicateProfiles"`

	FreezeStatus struct {
		IsFrozen   bool   `json:"isFrozen"`
		Reason     string `json:"reason"`
		PropertyID string `json:"propertyId"`
		Source     string `json:"source"`
	} `json:"freezeStatus"`

	ExternalSystemIdentifiers []ProfilesExternalIdentifier `json:"externalSystemIdentifiers"`
	NotesIdentifiers          []string                     `json:"notesIdentifiers"`
	ContactProfiles           []ContactProfileResponse     `json:"contactProfiles"`
	ContactFor                []ContactProfileResponse     `json:"contactFor"`
	HasImage                  bool                         `json:"hasImage"`
	ProfileStoreID            string                       `json:"profileStoreId"`

	Metadata ProfileMetadata `json:"metadata"`

	Embedded struct {
		LocalExternalAccountReceivable struct {
			Number string `json:"number"`
		} `json:"localExternalAccountReceivable"`
	} `json:"_embedded"`

	ProfileAvailabilityDetails struct {
		ProfileStoreID                   string   `json:"profileStoreId"`
		RegionCode                       string   `json:"regionCode"`
		IsAvailableOutsideOfSourceRegion bool     `json:"isAvailableOutsideOfSourceRegion"`
		AvailableInRegions               []string `json:"availableInRegions"`
	} `json:"profileAvailabilityDetails"`

	ExternalAccountReceivableNumber string `json:"externalAccountReceivableNumber"`
}

type IndividualProfile struct {
	ID       string `json:"id"`
	Version  string `json:"version"`
	TypeCode struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"typeCode"`

	Status struct {
		IsActive               bool `json:"isActive"`
		DeactivationReasonCode struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"deactivationReasonCode"`

		IsRestricted          bool `json:"isRestricted"`
		RestrictionReasonCode struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"restrictionReasonCode"`

		RestrictionComment string `json:"restrictionComment"`
		IsLocked           bool   `json:"isLocked"`
		LockedReason       string `json:"lockedReason"`
		ForReview          bool   `json:"forReview"`
	} `json:"status"`

	Details struct {
		TitleCode       LocalizedIDCodeValue `json:"titleCode"`
		FirstName       string               `json:"firstName"`
		LastName        string               `json:"lastName"`
		NationalityCode struct {
			Code         string `json:"code"`
			LanguageCode string `json:"languageCode"`
			Description  string `json:"description"`
		} `json:"nationalityCode"`

		Greeting string `json:"greeting"`
		VIPCode  struct {
			Color        string `json:"color"`
			ID           string `json:"id"`
			Code         string `json:"code"`
			LanguageCode string `json:"languageCode"`
			Description  string `json:"description"`
		} `json:"vipCode"`

		GenderCode struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"genderCode"`

		BirthdayAnniversaryDate *struct {
			Month int32 `json:"month"`
			Day   int32 `json:"day"`
			Year  int32 `json:"year"`
		} `json:"birthdayAnniversaryDate"`

		PreferredLanguageProfileDetails *struct {
			LanguageCode struct {
				Code         string `json:"code"`
				LanguageCode string `json:"languageCode"`
				Description  string `json:"description"`
			} `json:"languageCode"`

			LocalizedIndividualProfileDetails *struct {
				Title     string `json:"title"`
				FirstName string `json:"firstName"`
				LastName  string `json:"lastName"`
				Greeting  string `json:"greeting"`
			} `json:"localizedIndividualProfileDetails"`
		} `json:"preferredLanguageProfileDetails"`

		Anniversaries []struct {
			Date struct {
				Month int32 `json:"month"`
				Day   int32 `json:"day"`
				Year  int32 `json:"year"`
			} `json:"date"`

			Description string `json:"description"`
		} `json:"anniversaries"`

		CompanyName                        string               `json:"companyName"`
		PositionName                       string               `json:"positionName"`
		MiddleName                         string               `json:"middleName"`
		Suffix                             LocalizedIDCodeValue `json:"suffix"`
		PreferredCommunicationLanguageCode struct {
			Code         string `json:"code"`
			LanguageCode string `json:"languageCode"`
			Description  string `json:"description"`
		} `json:"preferredCommunicationLanguageCode"`

		HonoraryTitle LocalizedIDCodeValue `json:"honoraryTitle"`

		TaxID           string `json:"taxId"`
		AdditionalNames []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"additionalNames"`
	} `json:"details"`

	Addresses             []ProfileAddress              `json:"addresses"`
	CommunicationChannels []ProfileCommunicationChannel `json:"communicationChannels"`
	DisabilityCodes       []LocalizedIDCodeValue        `json:"disabilityCodes"`
	TaxExemption          struct {
		Code    LocalizedIDCodeValue `json:"code"`
		Comment string               `json:"comment"`
	} `json:"taxExemption"`

	Memberships []struct {
		ID         string               `json:"id"`
		IsActive   bool                 `json:"isActive"`
		LevelCode  LocalizedIDCodeValue `json:"levelCode"`
		TypeCode   LocalizedIDCodeValue `json:"typeCode"`
		Number     string               `json:"number"`
		SinceDate  Date                 `json:"sinceDate"`
		ExpiryDate Date                 `json:"expiryDate"`
	} `json:"memberships"`

	Documents []struct {
		ID               string      `json:"id"`
		ParentVersion    string      `json:"parentVersion"`
		TypeCode         CodeIDValue `json:"typeCode"`
		Number           string      `json:"number"`
		ExpiryDate       Date        `json:"expiryDate"`
		IssueDate        Date        `json:"issueDate"`
		IssueCountryCode struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"issueCountryCode"`

		BirthCountryCode struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"birthCountryCode"`

		BirthDate       Date   `json:"birthDate"`
		BirthPlace      string `json:"birthPlace"`
		NationalityCode struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"nationalityCode"`

		AdditionalInfo string `json:"additionalInfo"`
	} `json:"documents"`

	Metadata             ProfileMetadata `json:"metadata"`
	NotDuplicateProfiles struct {
		Collection []struct {
			NotDuplicateProfileID string `json:"notDuplicateProfileId"`
			ProfileID             string `json:"profileId"`
			ProfileRoleCode       string `json:"profileRoleCode"`
		} `json:"collection"`

		Count int32 `json:"count"`
	} `json:"notDuplicateProfiles"`

	FreezeStatus struct {
		IsFrozen   bool   `json:"isFrozen"`
		Reason     string `json:"reason"`
		PropertyID string `json:"propertyId"`
		Source     string `json:"source"`
	} `json:"freezeStatus"`

	Consents []struct {
		ConsentID   string          `json:"consentId"`
		IsGranted   bool            `json:"isGranted"`
		GrantedAt   DateTime        `json:"grantedAt"`
		ConsentType string          `json:"consentType"`
		Source      string          `json:"source"`
		Metadata    ProfileMetadata `json:"metadata"`
		ExpiryDate  Date            `json:"expiryDate"`
	} `json:"consents"`

	CustomFieldContainerJSON    string                       `json:"customFieldContainerJson"`
	CustomPiiFieldContainerJSON string                       `json:"customPiiFieldContainerJson"`
	ExternalSystemIdentifiers   []ProfilesExternalIdentifier `json:"externalSystemIdentifiers"`
	NotesIdentifiers            []string                     `json:"notesIdentifiers"`
	LinkedProfileCount          int32                        `json:"linkedProfileCount"`
	LinkedProfileIDs            []struct {
		ProfileRoleCode string `json:"profileRoleCode"`
		ProfileID       string `json:"profileId"`
	} `json:"linkedProfileIds"`

	GroupStatus struct {
		IsGrouped  bool   `json:"isGrouped"`
		GroupID    string `json:"groupId"`
		PropertyID string `json:"propertyId"`
		Code       string `json:"code"`
		Name       string `json:"name"`
	} `json:"groupStatus"`

	ContactProfiles  []ContactProfileResponse `json:"contactProfiles"`
	ContactFor       []ContactProfileResponse `json:"contactFor"`
	IsIncognito      bool                     `json:"isIncognito"`
	IncognitoDetails struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"incognitoDetails"`

	ProfileOwners struct {
		Primary struct {
			ID string `json:"id"`
		} `json:"primary"`

		Others []struct {
			ID string `json:"id"`
		} `json:"others"`
	} `json:"profileOwners"`

	Preferences []struct {
		ID               string               `json:"id"`
		PreferenceCode   LocalizedIDCodeValue `json:"preferenceCode"`
		IsPreferred      bool                 `json:"isPreferred"`
		IsGlobal         bool                 `json:"isGlobal"`
		OriginPropertyID string               `json:"originPropertyId"`
		IsRequired       bool                 `json:"isRequired"`
		Weight           int32                `json:"weight"`
	} `json:"preferences"`

	ExternalAccountReceivable struct {
		ExternalAccountReceivableNumber string `json:"externalAccountReceivableNumber"`
	} `json:"externalAccountReceivable"`
}

type IndividualProfileV0 struct {
	ID           string `json:"id"`
	IsActive     bool   `json:"isActive"`
	IsRestricted bool   `json:"isRestricted"`
	Version      int    `json:"version"`
	TypeCode     string `json:"typeCode"`
	TypeID       string `json:"typeId"`

	Details struct {
		FirstName                          string `json:"firstName"`
		LastName                           string `json:"lastName"`
		TitleCode                          string `json:"titleCode"`
		TitleID                            string `json:"titleId"`
		Greeting                           string `json:"greeting"`
		PreferredCommunicationLanguageCode string `json:"preferredCommunicationLanguageCode"`
	} `json:"details"`

	CommunicationDetails []CommunicationDetail `json:"communicationDetails"`
	Addresses            []struct {
		ID                    string        `json:"id"`
		CountryCode           string        `json:"countryCode"`
		City                  string        `json:"city"`
		PostCode              string        `json:"postCode"`
		AddressLine1          string        `json:"addressLine1"`
		TypeCode              string        `json:"typeCode"`
		TypeID                string        `json:"typeId"`
		Primary               bool          `json:"primary"`
		LanguageCode          string        `json:"languageCode"`
		ArPrimary             bool          `json:"arPrimary"`
		CountrySpecificFields []interface{} `json:"countrySpecificFields"`
	} `json:"addresses"`

	Preferences []struct {
		ID               string               `json:"id"`
		PreferenceCode   LocalizedIDCodeValue `json:"preferenceCode"`
		IsPreferred      bool                 `json:"isPreferred"`
		IsGlobal         bool                 `json:"isGlobal"`
		OriginPropertyID string               `json:"originPropertyId"`
		IsRequired       bool                 `json:"isRequired"`
		Weight           int32                `json:"weight"`
	} `json:"preferences"`

	IdentificationDocuments []struct {
		TypeCode         string   `json:"typeCode,omitempty"`
		Number           string   `json:"number,omitempty"`
		ExpiryDate       Date     `json:"expiryDate,omitempty"`
		IssueDate        Date     `json:"issueDate,omitempty"`
		IssueCountryCode string   `json:"issueCountryCode,omitempty"`
		IssuePlace       string   `json:"issuePlace,omitempty"`
		IssuedBy         string   `json:"issuedBy,omitempty"`
		BirthCountryCode string   `json:"birthCountryCode,omitempty"`
		BirthPlace       string   `json:"birthPlace,omitempty"`
		BirthDate        Date     `json:"birthDate,omitempty"`
		NationailtyCode  string   `json:"nationailtyCode,omitempty"`
		AdditionalInfo   string   `json:"additionalInfo,omitempty"`
		ImageRefs        []string `json:"imageRefs,omitempty"`
	} `json:"identificationDocuments"`

	IsLocked             bool `json:"isLocked"`
	ForReview            bool `json:"forReview"`
	NotDuplicateProfiles struct {
		Collection []struct {
			NotDuplicateProfileID string `json:"notDuplicateProfileId"`
			ProfileID             string `json:"profileId"`
			ProfileRoleCode       string `json:"profileRoleCode"`
		} `json:"collection"`

		NotDuplicateProfilesCount int `json:"notDuplicateProfilesCount"`
	} `json:"notDuplicateProfiles"`

	FreezeStatus struct {
		IsFrozen   bool   `json:"isFrozen"`
		Reason     string `json:"reason"`
		PropertyID string `json:"propertyId"`
		Source     string `json:"source"`
	} `json:"freezeStatus"`

	Consents []struct {
		ConsentID   string          `json:"consentId"`
		IsGranted   bool            `json:"isGranted"`
		GrantedAt   DateTime        `json:"grantedAt"`
		ConsentType string          `json:"consentType"`
		Source      string          `json:"source"`
		Metadata    ProfileMetadata `json:"metadata"`
		ExpiryDate  Date            `json:"expiryDate"`
	} `json:"consents"`

	ExternalSystemIdentifiers []ProfilesExternalIdentifier `json:"externalSystemIdentifiers"`
	ContactProfiles           []ContactProfileResponse     `json:"contactProfiles"`
	ContactFor                []ContactProfileResponse     `json:"contactFor"`
	HasImage                  bool                         `json:"hasImage"`
	ProfileStoreID            string                       `json:"profileStoreId"`
	ProfileLastActivity       struct {
		PerformedBy string    `json:"performedBy"`
		PerformedAt time.Time `json:"performedAt"`
	} `json:"profileLastActivity"`

	Metadata struct {
		CreatorID               string    `json:"creatorId"`
		CreationTime            time.Time `json:"creationTime"`
		LastUpdaterID           string    `json:"lastUpdaterId"`
		LastUpdateTime          time.Time `json:"lastUpdateTime"`
		CreatedInPropertyID     string    `json:"createdInPropertyId"`
		LastUpdatedInPropertyID string    `json:"lastUpdatedInPropertyId"`
	} `json:"metadata"`

	Embedded struct {
		LocalExternalAccountReceivable struct {
			Number string `json:"number"`
		} `json:"localExternalAccountReceivable"`
	} `json:"_embedded"`

	ProfileAvailabilityDetails struct {
		ProfileStoreID                   string   `json:"profileStoreId"`
		RegionCode                       string   `json:"regionCode"`
		IsAvailableOutsideOfSourceRegion bool     `json:"isAvailableOutsideOfSourceRegion"`
		AvailableInRegions               []string `json:"availableInRegions"`
	} `json:"profileAvailabilityDetails"`

	ExternalAccountReceivableNumber string `json:"externalAccountReceivableNumber"`
}

type IndividualProfileV2 struct {
	ID       string               `json:"id"`
	TypeCode LocalizedIDCodeValue `json:"typeCode"`

	Status struct {
		IsActive               bool `json:"isActive"`
		DeactivationReasonCode struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"deactivationReasonCode"`

		IsRestricted          bool `json:"isRestricted"`
		RestrictionReasonCode struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"restrictionReasonCode"`

		RestrictionComment string `json:"restrictionComment"`
		IsLocked           bool   `json:"isLocked"`
		LockedReason       string `json:"lockedReason"`
		ForReview          bool   `json:"forReview"`
	} `json:"status"`

	Details struct {
		TitleCode       LocalizedIDCodeValue `json:"titleCode"`
		FirstName       string               `json:"firstName"`
		LastName        string               `json:"lastName"`
		NationalityCode struct {
			Code         string `json:"code"`
			LanguageCode string `json:"languageCode"`
			Description  string `json:"description"`
		} `json:"nationalityCode"`

		Greeting string `json:"greeting"`
		VIPCode  struct {
			Color        string `json:"color"`
			ID           string `json:"id"`
			Code         string `json:"code"`
			LanguageCode string `json:"languageCode"`
			Description  string `json:"description"`
		} `json:"vipCode"`

		GenderDetails LocalizedIDCodeValue `json:"genderDetails"`

		BirthdayAnniversaryDate struct {
			Month int32 `json:"month"`
			Day   int32 `json:"day"`
			Year  int32 `json:"year"`
		} `json:"birthdayAnniversaryDate"`

		PreferredLanguageProfileDetails struct {
			LanguageCode struct {
				Code         string `json:"code"`
				LanguageCode string `json:"languageCode"`
				Description  string `json:"description"`
			} `json:"languageCode"`

			LocalizedIndividualProfileDetails struct {
				Title     string `json:"title"`
				FirstName string `json:"firstName"`
				LastName  string `json:"lastName"`
				Greeting  string `json:"greeting"`
			} `json:"localizedIndividualProfileDetails"`
		} `json:"preferredLanguageProfileDetails"`

		Anniversaries []struct {
			Date struct {
				Month int32 `json:"month"`
				Day   int32 `json:"day"`
				Year  int32 `json:"year"`
			} `json:"date"`

			Description string `json:"description"`
		} `json:"anniversaries"`

		CompanyName                        string               `json:"companyName"`
		PositionName                       string               `json:"positionName"`
		MiddleName                         string               `json:"middleName"`
		Suffix                             LocalizedIDCodeValue `json:"suffix"`
		PreferredCommunicationLanguageCode struct {
			Code         string `json:"code"`
			LanguageCode string `json:"languageCode"`
			Description  string `json:"description"`
		} `json:"preferredCommunicationLanguageCode"`

		HonoraryTitle LocalizedIDCodeValue `json:"honoraryTitle"`

		TaxID           string `json:"taxId"`
		AdditionalNames []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"additionalNames"`

		CountrySpecificFields []struct {
			CountrySpecificField LocalizedIDCodeValue `json:"countrySpecificField"`

			Value struct {
				Type      string `json:"type"`
				String    string `json:"string"`
				LocalDate Date   `json:"localDate"`
				Array     []any  `json:"array"`
			}
		} `json:"countrySpecificFields"`
	} `json:"details"`

	Addresses             []ProfileAddress              `json:"addresses"`
	CommunicationChannels []ProfileCommunicationChannel `json:"communicationChannels"`
	DisabilityCodes       []LocalizedIDCodeValue        `json:"disabilityCodes"`
	NotesIdentifiers      []string                      `json:"notesIdentifiers"`
	Memberships           []struct {
		ID         string               `json:"id"`
		IsActive   bool                 `json:"isActive"`
		LevelCode  LocalizedIDCodeValue `json:"levelCode"`
		TypeCode   LocalizedIDCodeValue `json:"typeCode"`
		Number     string               `json:"number"`
		SinceDate  Date                 `json:"sinceDate"`
		ExpiryDate Date                 `json:"expiryDate"`
	} `json:"memberships"`

	Documents []struct {
		ID               string               `json:"id"`
		ParentVersion    int64                `json:"parentVersion"`
		TypeCode         LocalizedIDCodeValue `json:"typeCode"`
		Number           string               `json:"number"`
		ExpiryDate       Date                 `json:"expiryDate"`
		IssueDate        Date                 `json:"issueDate"`
		IssueCountryCode struct {
			Code         string `json:"code"`
			LanguageCode string `json:"languageCode"`
			Description  string `json:"description"`
		} `json:"issueCountryCode"`

		IssuePlace       string `json:"issuePlace"`
		IssuedBy         string `json:"issuedBy"`
		BirthCountryCode struct {
			Code         string `json:"code"`
			LanguageCode string `json:"languageCode"`
			Description  string `json:"description"`
		} `json:"birthCountryCode"`

		BirthDate       Date   `json:"birthDate"`
		BirthPlace      string `json:"birthPlace"`
		NationalityCode struct {
			Code         string `json:"code"`
			LanguageCode string `json:"languageCode"`
			Description  string `json:"description"`
		} `json:"nationalityCode"`

		AdditionalInfo        string `json:"additionalInfo"`
		CountrySpecificFields []struct {
			CountrySpecificField LocalizedIDCodeValue `json:"countrySpecificField"`
			Value                struct {
				Type      string `json:"type"`
				String    string `json:"string"`
				LocalDate Date   `json:"localDate"`
				Array     []any  `json:"array"`
			} `json:"value"`
		} `json:"countrySpecificFields"`
	} `json:"documents"`

	Metadata ProfileMetadata `json:"metadata"`

	NotDuplicateProfiles struct {
		Collection []struct {
			NotDuplicateProfileID string `json:"notDuplicateProfileId"`
			ProfileID             string `json:"profileId"`
			ProfileRoleCode       string `json:"profileRoleCode"`
		} `json:"collection"`

		Count int32 `json:"count"`
	} `json:"notDuplicateProfiles"`

	FreezeStatus struct {
		IsFrozen   bool   `json:"isFrozen"`
		Reason     string `json:"reason"`
		PropertyID string `json:"propertyId"`
		Source     string `json:"source"`
	} `json:"freezeStatus"`

	Consents []struct {
		ConsentID   string          `json:"consentId"`
		IsGranted   bool            `json:"isGranted"`
		GrantedAt   DateTime        `json:"grantedAt"`
		ConsentType string          `json:"consentType"`
		Source      string          `json:"source"`
		Metadata    ProfileMetadata `json:"metadata"`
		ExpiryDate  Date            `json:"expiryDate"`
	} `json:"consents"`

	CustomFieldContainerJSON    string                       `json:"customFieldContainerJson"`
	CustomPiiFieldContainerJSON string                       `json:"customPiiFieldContainerJson"`
	ExternalSystemIdentifiers   []ProfilesExternalIdentifier `json:"externalSystemIdentifiers"`
	GroupStatus                 struct {
		IsGrouped  bool   `json:"isGrouped"`
		GroupID    string `json:"groupId"`
		PropertyID string `json:"propertyId"`
		Code       string `json:"code"`
		Name       string `json:"name"`
	} `json:"groupStatus"`

	ContactProfiles  []ContactProfileResponse `json:"contactProfiles"`
	ContactFor       []ContactProfileResponse `json:"contactFor"`
	IsIncognito      bool                     `json:"isIncognito"`
	IncognitoDetails struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"incognitoDetails"`

	ProfileOwners struct {
		Primary struct {
			ID string `json:"id"`
		} `json:"primary"`

		Others []struct {
			ID string `json:"id"`
		} `json:"others"`
	} `json:"profileOwners"`

	ProfileLastActivity struct {
		PerformedAt DateTime `json:"performedAt"`
		DaysAgo     int32    `json:"daysAgo"`
		PerformedBy string   `json:"performedBy"`
	} `json:"profileLastActivity"`

	TaxExemption struct {
		Code    LocalizedIDCodeValue `json:"code"`
		Comment string               `json:"comment"`
	} `json:"taxExemption"`

	Preferences []struct {
		ID               string               `json:"id"`
		PreferenceCode   LocalizedIDCodeValue `json:"preferenceCode"`
		IsPreferred      bool                 `json:"isPreferred"`
		IsGlobal         bool                 `json:"isGlobal"`
		OriginPropertyID string               `json:"originPropertyId"`
		IsRequired       bool                 `json:"isRequired"`
		Weight           int32                `json:"weight"`
	} `json:"preferences"`

	Version           int64 `json:"version"`
	MembershipDetails struct {
		Manual []struct {
			MembershipID string `json:"membershipId"`
			Number       string `json:"number"`
			Status       struct {
				IsActive bool `json:"isActive"`
			} `json:"status"`

			MembershipType  LocalizedIDCodeValue `json:"membershipType"`
			MembershipLevel struct {
				LocalizedIDCodeValue
				Color string `json:"color"`
			} `json:"membershipLevel"`

			SinceDate  Date `json:"sinceDate"`
			ExpiryDate Date `json:"expiryDate"`
		} `json:"manual"`

		Internal []struct {
			MembershipID string `json:"membershipId"`
			Number       string `json:"number"`
			Status       struct {
				IsActive bool `json:"isActive"`
			} `json:"status"`

			MembershipType  LocalizedIDCodeValue `json:"membershipType"`
			MembershipLevel struct {
				LocalizedIDCodeValue
				Color string `json:"color"`
			} `json:"membershipLevel"`

			SinceDate  Date `json:"sinceDate"`
			ExpiryDate Date `json:"expiryDate"`
		} `json:"internal"`
	} `json:"membershipDetails"`

	ExternalAccountReceivable struct {
		ExternalAccountReceivableNumber string `json:"externalAccountReceivableNumber"`
	} `json:"externalAccountReceivable"`

	HasImage bool `json:"hasImage"`
}

type TravelAgentProfile struct {
	ID       string `json:"id"`
	Version  string `json:"version"`
	TypeCode struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"typeCode"`

	Status struct {
		IsActive               bool `json:"isActive"`
		DeactivationReasonCode struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"deactivationReasonCode"`

		IsRestricted          bool `json:"isRestricted"`
		RestrictionReasonCode struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"restrictionReasonCode"`

		RestrictionComment string `json:"restrictionComment"`
		IsLocked           bool   `json:"isLocked"`
		LockedReason       string `json:"lockedReason"`
		ForReview          bool   `json:"forReview"`
	} `json:"status"`

	LinkedProfileCount int32 `json:"linkedProfileCount"`
	LinkedProfileIDs   []struct {
		ProfileRoleCode string `json:"profileRoleCode"`
		ProfileID       string `json:"profileId"`
	} `json:"linkedProfileIds"`

	Details struct {
		Description                     string        `json:"description"`
		FullName                        string        `json:"fullName"`
		Abbreviation                    string        `json:"abbreviation"`
		IATA                            string        `json:"iata"`
		TaxID                           string        `json:"taxId"`
		BusinessSegmentation            []CodeIDValue `json:"businessSegmentation"`
		PreferredLanguageProfileDetails struct {
			LanguageCode struct {
				Code        string `json:"code"`
				Description string `json:"description"`
			} `json:"languageCode"`

			FullName string `json:"fullName"`
		} `json:"preferredLanguageProfileDetails"`

		PreferredCommunicationLanguageCode struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"preferredCommunicationLanguageCode"`
	} `json:"details"`

	Addresses             []ProfileAddress              `json:"addresses"`
	CommunicationChannels []ProfileCommunicationChannel `json:"communicationChannels"`
	Metadata              ProfileMetadata               `json:"metadata"`
	NotDuplicateProfiles  struct {
		Collection []struct {
			NotDuplicateProfileID string `json:"notDuplicateProfileId"`
			ProfileID             string `json:"profileId"`
			ProfileRoleCode       string `json:"profileRoleCode"`
		}

		Count int32 `json:"count"`
	} `json:"notDuplicateProfiles"`

	FreezeStatus struct {
		IsFrozen   bool   `json:"isFrozen"`
		Reason     string `json:"reason"`
		PropertyID string `json:"propertyId"`
		Source     string `json:"source"`
	} `json:"freezeStatus"`

	CustomFieldContainerJSON  string                       `json:"customFieldContainerJson"`
	ExternalSystemIdentifiers []ProfilesExternalIdentifier `json:"externalSystemIdentifiers"`
	NotesIdentifiers          []string                     `json:"notesIdentifiers"`
	ContactProfiles           []ContactProfileResponse     `json:"contactProfiles"`
	ContactFor                []ContactProfileResponse     `json:"contactFor"`
	ProfileOwners             struct {
		Primary struct {
			ID string `json:"id"`
		} `json:"primary"`

		Others []struct {
			ID string `json:"id"`
		} `json:"others"`
	} `json:"profileOwners"`

	RelationshipDetails struct {
		IsRoot bool `json:"isRoot"`
	} `json:"relationshipDetails"`

	ExternalAccountReceivable struct {
		ExternalAccountReceivableNumber string `json:"externalAccountReceivableNumber"`
	} `json:"externalAccountReceivable"`

	HasImage bool `json:"hasImage"`
}

type TravelAgentProfileV0 struct {
	ID           string `json:"id"`
	IsActive     bool   `json:"isActive"`
	IsRestricted bool   `json:"isRestricted"`
	Version      int    `json:"version"`
	TypeCode     string `json:"typeCode"`
	TypeID       string `json:"typeId"`

	Details struct {
		FullName                           string        `json:"fullName"`
		BusinessSegmentCodes               []interface{} `json:"businessSegmentCodes"`
		PreferredCommunicationLanguageCode string        `json:"preferredCommunicationLanguageCode"`
		BusinessSegmentIds                 []interface{} `json:"businessSegmentIds"`
	} `json:"details"`

	CommunicationDetails []CommunicationDetail `json:"communicationDetails"`
	Addresses            []struct {
		ID                    string        `json:"id"`
		CountryCode           string        `json:"countryCode"`
		StateCode             string        `json:"stateCode"`
		City                  string        `json:"city"`
		PostCode              string        `json:"postCode"`
		AddressLine1          string        `json:"addressLine1"`
		AddressLine2          string        `json:"addressLine2"`
		TypeCode              string        `json:"typeCode"`
		TypeID                string        `json:"typeId"`
		Primary               bool          `json:"primary"`
		LanguageCode          string        `json:"languageCode"`
		ArPrimary             bool          `json:"arPrimary"`
		CountrySpecificFields []interface{} `json:"countrySpecificFields"`
	} `json:"addresses"`

	IsLocked  bool `json:"isLocked"`
	ForReview bool `json:"forReview"`

	FreezeStatus struct {
		IsFrozen bool   `json:"isFrozen"`
		Reason   string `json:"reason"`
	} `json:"freezeStatus"`

	ExternalSystemIdentifiers []struct {
		ExternalSystemID string `json:"externalSystemId,omitempty"`
		Identifier       string `json:"identifier,omitempty"`
	} `json:"externalSystemIdentifiers,omitempty"`

	HasImage       bool   `json:"hasImage"`
	ProfileStoreID string `json:"profileStoreId"`

	Metadata struct {
		CreatorID               string    `json:"creatorId"`
		CreationTime            time.Time `json:"creationTime"`
		LastUpdaterID           string    `json:"lastUpdaterId"`
		LastUpdateTime          time.Time `json:"lastUpdateTime"`
		CreatedInPropertyID     string    `json:"createdInPropertyId"`
		LastUpdatedInPropertyID string    `json:"lastUpdatedInPropertyId"`
	} `json:"metadata"`

	Embedded struct {
		LocalExternalAccountReceivable struct {
			Number string `json:"number"`
		} `json:"localExternalAccountReceivable"`
	} `json:"_embedded"`

	ProfileAvailabilityDetails struct {
		ProfileStoreID                   string   `json:"profileStoreId"`
		RegionCode                       string   `json:"regionCode"`
		IsAvailableOutsideOfSourceRegion bool     `json:"isAvailableOutsideOfSourceRegion"`
		AvailableInRegions               []string `json:"availableInRegions"`
	} `json:"profileAvailabilityDetails"`

	ExternalAccountReceivableNumber string `json:"externalAccountReceivableNumber"`
}

type Transaction struct {
	ID      string `json:"id"`
	Account struct {
		ID     string `json:"id"`
		Number string `json:"number"`
	} `json:"account"`

	OriginAccount struct {
		ID     string `json:"id"`
		Number string `json:"number"`
	} `json:"originAccount"`

	Version             int64    `json:"version"`
	TransactionDate     Date     `json:"transactionDate"`
	TransactionDateTime DateTime `json:"transactionDateTime"`
	BusinessDate        Date     `json:"businessDate"`
	RevenueDate         Date     `json:"revenueDate"`
	UpdatedAt           DateTime `json:"updatedAt"`
	User                struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`

	Cashier struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"cashier"`

	Comment      string `json:"comment"`
	PostingGroup struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"postingGroup"`

	PaymentMethod string `json:"paymentMethod"`
	FolioWindow   struct {
		ID        string `json:"id"`
		Number    int    `json:"number"`
		ProfileID string `json:"profileId"`
	} `json:"folioWindow"`

	InvoiceNumber   string `json:"invoiceNumber"`
	TransactionType string `json:"transactionType"`
	LedgerType      string `json:"ledgerType"`
	TransactionCode struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"transactionCode"`

	TransactionGroup struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"transactionGroup"`

	TransactionSubGroup struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"transactionSubGroup"`

	Reservation struct {
		ID                 string `json:"id"`
		RoomNumber         string `json:"roomNumber"`
		ConfirmationNumber string `json:"confirmationNumber"`
		RoomType           struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"roomType"`

		PrimaryGuest struct{} `json:"primaryGuest"`

		MarketSegment struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"marketSegment"`

		BookingSource struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"bookingSource"`

		DistributionChannel struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"distributionChannel"`
	} `json:"reservation"`

	TaxTransactions []struct {
		ID                   string  `json:"id"`
		PropertyID           string  `json:"propertyId"`
		TenantID             string  `json:"tenantId"`
		IsIncluded           bool    `json:"isIncluded"`
		PricePerUnit         float64 `json:"pricePerUnit"`
		Quantity             float64 `json:"quantity"`
		TotalAmount          float64 `json:"totalAmount"`
		TotalAmountUnrounded float64 `json:"totalAmountUnrounded"`
		TaxAmountCurrency    string  `json:"taxAmountCurrency"`
		TaxTransactionCodeID string  `json:"taxTransactionCodeId"`
		TaxTransactionCode   string  `json:"taxTransactionCode"`
		TransactionID        string  `json:"transactionId"`
		PostingGroupID       string  `json:"postingGroupId"`
	} `json:"taxTransactions"`

	NetAmount      float64 `json:"netAmount"`
	GrossAmount    float64 `json:"grossAmount"`
	TaxAmount      float64 `json:"taxAmount"`
	PaymentAmount  float64 `json:"paymentAmount"`
	NetUnitPrice   float64 `json:"netUnitPrice"`
	GrossUnitPrice float64 `json:"grossUnitPrice"`
	Quantity       float64 `json:"quantity"`
	Currency       string  `json:"currency"`
	Metadata       struct {
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"metaData"`
}

type Transactions []Transaction

type PDFResponseBody struct {
	Content     []byte
	ContentType string
	FileName    string
}

type CashieringAccount struct {
	ID                    string `json:"id"`
	Version               int64  `json:"version"`
	AccountNumber         string `json:"accountNumber"`
	ExternalAccountNumber string `json:"externalAccountNumber"`
	AccountGroupCode      struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"accountGroupCode"`

	ProfileID         string   `json:"profileId"`
	ProfileRoleCode   string   `json:"profileRoleCode"`
	ValidFromDate     DateTime `json:"validFromDate"`
	ValidTODate       string   `json:"validToDate"`
	AccountStatusCode struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"accountStatusCode"`

	ReservationStatusCode struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"reservationStatusCode"`

	CreditLimit float64 `json:"creditLimit"`
	Nickname    string  `json:"nickname"`
	NoPost      bool    `json:"noPost"`
	NetBalance  struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"netBalance"`

	GrossBalance struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"grossBalance"`

	OriginatorRole string   `json:"originatorRole"`
	OriginatorID   string   `json:"originatorId"`
	MembershipIDs  []string `json:"membershipIds"`
	ChannelCode    struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"channelCode"`

	SourceCode struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"sourceCode"`

	MarketSegmentCode struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"marketSegmentCode"`

	LinkedProfilesV2 []struct {
		ProfileID         string `json:"profileId"`
		ProfileRole       string `json:"profileRole"`
		IsPrimary         bool   `json:"isPrimary"`
		TaxExemptReasonID string `json:"taxExemptReasonId"`
	} `json:"linkedProfilesV2"`

	OutstandingPrepaidCommission struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"outstandingPrepaidCommission"`

	ContactProfiles map[string][]struct {
		ID        string   `json:"id"`
		Roles     []string `json:"roles"`
		IsPrimary bool     `json:"isPrimary"`
	} `json:"contactProfiles"`

	IsExcludedFromZeroCheckOut bool `json:"isExcludedFromZeroCheckOut"`
	AllowSkipInvoiceGeneration bool `json:"allowSkipInvoiceGeneration"`

	ReservationGroup struct {
		ID   string `json:"id"`
		Code string `json:"code"`
	} `json:"reservationGroup"`

	MultiReservationLinks []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Identifier string `json:"identifier"`
	} `json:"multiReservationLinks"`

	MeetingAndEvents struct {
		BookingCode        string `json:"bookingCode"`
		BookingChannelCode string `json:"bookingChannelCode"`
	} `json:"meetingAndEvents"`

	ItineraryID string `json:"itineraryId"`
}

type ProfileFlowControl struct {
	MatchConfiguration struct {
		MatchEngine string `json:"matchEngine"`
		MergeRules  struct {
			TargetProfile string `json:"targetProfile"`
			MergeRules    []struct {
				Field     string `json:"field"`
				MergeRule string `json:"mergeRule"`
			} `json:"mergeRules,omitempty"`
		} `json:"mergeRules,omitempty"`
	} `json:"matchConfiguration,omitempty"`

	ForceUpsertWhenUsableProfileIsAvailable bool `json:"forceUpsertWhenUsableProfileIsAvailable,omitempty"`
	UseRecentTargetProfile                  bool `json:"useRecentTargetProfile,omitempty"`
}

type ProfileUpsert struct {
	FlowControl ProfileFlowControl `json:"_flowControl,omitempty"`
}

type CreateIndividualV1 struct {
	ProfileUpsert

	TypeCode string `json:"typeCode,omitempty"`
	Details  struct {
		TitleCode       LocalizedIDCodeValue `json:"titleCode,omitempty"`
		FirstName       string               `json:"firstName,omitempty"`
		LastName        string               `json:"lastName,omitempty"`
		NationalityCode string               `json:"nationalityCode,omitempty"`
		Greeting        string               `json:"greeting,omitempty"`
		VIPCode         string               `json:"vipCode,omitempty"`
		GenderCode      string               `json:"genderCode,omitempty"`
		GenderID        string               `json:"genderId,omitempty"`

		BirthdayAnniversaryDate struct {
			Month int32 `json:"month,omitempty"`
			Day   int32 `json:"day,omitempty"`
			Year  int32 `json:"year,omitempty"`
		} `json:"birthdayAnniversaryDate,omitempty"`

		PreferredLanguageProfileDetails struct {
			LanguageCode string `json:"languageCode,omitempty"`

			LocalizedIndividualProfileDetails struct {
				Title     string `json:"title,omitempty"`
				FirstName string `json:"firstName,omitempty"`
				LastName  string `json:"lastName,omitempty"`
				Greeting  string `json:"greeting,omitempty"`
			} `json:"localizedIndividualProfileDetails,omitempty"`
		} `json:"preferredLanguageProfileDetails,omitempty"`

		Anniversaries []struct {
			Date struct {
				Month int32 `json:"month,omitempty"`
				Day   int32 `json:"day,omitempty"`
				Year  int32 `json:"year,omitempty"`
			} `json:"date,omitempty"`

			Description string `json:"description,omitempty"`
		} `json:"anniversaries,omitempty"`

		CompanyName                        string `json:"companyName,omitempty"`
		PositionName                       string `json:"positionName,omitempty"`
		MiddleName                         string `json:"middleName,omitempty"`
		SuffixID                           string `json:"suffixId,omitempty"`
		PreferredCommunicationLanguageCode string `json:"preferredCommunicationLanguageCode,omitempty"`
		HonoraryTitleID                    string `json:"honoraryTitleId,omitempty"`

		TaxID           string `json:"taxId,omitempty"`
		AdditionalNames []struct {
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"additionalNames,omitempty"`
	} `json:"details,omitempty"`

	Addresses []struct {
		AddressLine1    string `json:"addressLine1,omitempty"`
		AddressLine2    string `json:"addressLine2,omitempty"`
		City            string `json:"city,omitempty"`
		CountryCode     string `json:"countryCode,omitempty"`
		IsArPrimary     bool   `json:"isArPrimary,omitempty"`
		IsPrimary       bool   `json:"isPrimary,omitempty"`
		LanguageCode    string `json:"languageCode,omitempty"`
		PostalCode      string `json:"postalCode,omitempty"`
		StateCode       string `json:"stateCode,omitempty"`
		AddressTypeCode string `json:"addressTypeCode,omitempty"`
		DistrictID      string `json:"districtId,omitempty"`
	} `json:"addresses,omitempty"`

	CommunicationChannels []struct {
		TypeCode  string `json:"typeCode,omitempty"`
		ModeCode  string `json:"modeCode,omitempty"`
		Details   string `json:"details,omitempty"`
		IsPrimary bool   `json:"isPrimary,omitempty"`
	} `json:"communicationChannels,omitempty"`

	DisabilityCodes []string `json:"disabilityCodes,omitempty"`
	TaxExemption    struct {
		Code    string `json:"code,omitempty"`
		Comment string `json:"comment,omitempty"`
	} `json:"taxExemption,omitempty"`

	Consents []struct {
		ConsentID   string `json:"consentId,omitempty"`
		Source      string `json:"source,omitempty"`
		Description string `json:"description,omitempty"`
		ExpiryDate  Date   `json:"expiryDate,omitempty"`
		IsGranted   bool   `json:"isGranted,omitempty"`
	} `json:"consents,omitempty"`

	ExternalSystemIdentifiers []struct {
		ExternalSystemID string `json:"externalSystemId,omitempty"`
		Identifier       string `json:"identifier,omitempty"`
	} `json:"externalSystemIdentifiers,omitempty"`

	GroupStatus struct {
		IsGrouped bool   `json:"isGrouped,omitempty"`
		GroupID   string `json:"groupId,omitempty"`
	} `json:"groupStatus,omitempty"`

	IdentificationDocuments []struct {
		TypeCode         string   `json:"typeCode,omitempty"`
		Number           string   `json:"number,omitempty"`
		ExpiryDate       Date     `json:"expiryDate,omitempty"`
		IssueDate        Date     `json:"issueDate,omitempty"`
		IssueCountryCode string   `json:"issueCountryCode,omitempty"`
		IssuePlace       string   `json:"issuePlace,omitempty"`
		IssuedBy         string   `json:"issuedBy,omitempty"`
		BirthCountryCode string   `json:"birthCountryCode,omitempty"`
		BirthPlace       string   `json:"birthPlace,omitempty"`
		BirthDate        Date     `json:"birthDate,omitempty"`
		NationailtyCode  string   `json:"nationailtyCode,omitempty"`
		AdditionalInfo   string   `json:"additionalInfo,omitempty"`
		ImageRefs        []string `json:"imageRefs,omitempty"`
	} `json:"identificationDocuments,omitempty"`

	Preferences []struct {
		PreferenceCode string  `json:"preferenceCode,omitempty"`
		IsGlobal       bool    `json:"isGlobal,omitempty"`
		IsPreferred    bool    `json:"isPreferred,omitempty"`
		IsRequired     bool    `json:"isRequired,omitempty"`
		Weight         float64 `json:"weight,omitempty"`
	} `json:"preferences,omitempty"`

	Memberships []struct {
		LevelCode  string `json:"levelCode,omitempty"`
		TypeCode   string `json:"typeCode,omitempty"`
		Number     string `json:"number,omitempty"`
		SinceDate  Date   `json:"sinceDate,omitempty"`
		ExpiryDate Date   `json:"expiryDate,omitempty"`
		IsActive   bool   `json:"isActive,omitempty"`
	} `json:"memberships,omitempty"`

	LinkedProfiles []struct {
		ProfileID       string `json:"profileId,omitempty"`
		ProfileRoleCode string `json:"profileRoleCode,omitempty"`
	} `json:"linkedProfiles,omitempty"`

	NotDuplicateProfiles []string `json:"notDuplicateProfiles,omitempty"`
	Notes                []struct {
		Content  string `json:"content,omitempty"`
		IsGlobal bool   `json:"isGlobal,omitempty"`
		TypeCode string `json:"typeCode,omitempty"`
	} `json:"notes,omitempty"`

	Attachments []struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		IsGlobal    bool   `json:"isGlobal,omitempty"`
	} `json:"attachments,omitempty"`

	Alerts []struct {
		Content  string   `json:"content,omitempty"`
		Title    string   `json:"title,omitempty"`
		IsGlobal bool     `json:"isGlobal,omitempty"`
		Areas    []string `json:"areas,omitempty"`
	} `json:"alerts,omitempty"`

	CreditCardsStandalone []struct {
		TransactionCodeID string `json:"transactionCodeId,omitempty"`
		CreditCardToken   string `json:"creditCardToken,omitempty"`
		Prefix6           string `json:"prefix6,omitempty"`
		Suffix4           string `json:"suffix4,omitempty"`
		ExpiryMonth       int32  `json:"expiryMonth,omitempty"`
		ExpiryYear        int32  `json:"expiryYear,omitempty"`
	} `json:"creditCardsStandalone,omitempty"`

	CreditCardsOffline []struct {
		TransactionCodeID string `json:"transactionCodeId,omitempty"`
		CardHolder        string `json:"cardHolder,omitempty"`
		CardNumber        string `json:"cardNumber,omitempty"`
		ExpiryYear        int32  `json:"expiryYear,omitempty"`
		ExpiryMonth       int32  `json:"expiryMonth,omitempty"`
		CardToken         string `json:"cardToken,omitempty"`
	} `json:"creditCardsOffline,omitempty"`

	OtherPaymentMethods []struct {
		TransactionCodeID string `json:"transactionCodeId,omitempty"`
	} `json:"otherPaymentMethods,omitempty"`

	ProfileOwners struct {
		Primary struct {
			ID string `json:"id,omitempty"`
		} `json:"primary,omitempty"`

		Others []struct {
			ID string `json:"id,omitempty"`
		} `json:"others,omitempty"`
	} `json:"profileOwners,omitempty"`

	ProfileStoreID           string `json:"profileStoreId,omitempty"`
	WaiveOffSurchargeDetails []struct {
		ExemptFromFees bool   `json:"exemptFromFees,omitempty"`
		PropertyID     string `json:"propertyId,omitempty"`
	} `json:"waiveOffSurchargeDetails,omitempty"`

	GlobalSalesMarketing []struct {
		SalesMarketingFieldID string `json:"salesMarketingFieldId,omitempty"`
		Items                 []struct {
			SalesMarketingValueID string `json:"salesMarketingValueId,omitempty"`
		} `json:"items,omitempty"`
	} `json:"globalSalesMarketing,omitempty"`

	PropertySalesMarketing []struct {
		SalesMarketingFieldID string `json:"salesMarketingFieldId,omitempty"`
		Items                 []struct {
			SalesMarketingValueID string `json:"salesMarketingValueId,omitempty"`
		} `json:"items,omitempty"`
	} `json:"propertySalesMarketing,omitempty"`

	RestrictionDetails struct {
		IsRestricted        bool   `json:"isRestricted,omitempty"`
		RestrictionReasonID string `json:"restrictionReasonId,omitempty"`
	} `json:"restrictionDetails,omitempty"`

	TaxExemptDetails struct {
		Comment string `json:"comment,omitempty"`
		Reasons []struct {
			TaxExemptReasonId string `json:"taxExemptReasonId,omitempty"`
		} `json:"reasons,omitempty"`
	} `json:"taxExemptDetails,omitempty"`
}

type PatchIndividual struct {
	Details *struct {
		TitleID         string `json:"titleId,omitempty"`
		FirstName       string `json:"firstName,omitempty"`
		LastName        string `json:"lastName,omitempty"`
		NationalityCode string `json:"nationalityCode,omitempty"`
		Greeting        string `json:"greeting,omitempty"`
		VIPID           string `json:"vipId,omitempty"`
		GenderCode      string `json:"genderCode,omitempty"`
		GenderID        string `json:"genderId,omitempty"`

		BirthdayAnniversaryDate *struct {
			Month int32 `json:"month,omitempty"`
			Day   int32 `json:"day,omitempty"`
			Year  int32 `json:"year,omitempty"`
		} `json:"birthdayAnniversaryDate,omitempty"`

		PreferredLanguageProfileDetails *struct {
			LanguageCode string `json:"languageCode,omitempty"`

			LocalizedIndividualProfileDetails *struct {
				Title     string `json:"title,omitempty"`
				FirstName string `json:"firstName,omitempty"`
				LastName  string `json:"lastName,omitempty"`
				Greeting  string `json:"greeting,omitempty"`
			} `json:"localizedIndividualProfileDetails,omitempty"`
		} `json:"preferredLanguageProfileDetails,omitempty"`

		CompanyName                        string `json:"companyName,omitempty"`
		PositionName                       string `json:"positionName,omitempty"`
		MiddleName                         string `json:"middleName,omitempty"`
		SuffixID                           string `json:"suffixId,omitempty"`
		PreferredCommunicationLanguageCode string `json:"preferredCommunicationLanguageCode,omitempty"`
		HonoraryTitleID                    string `json:"honoraryTitleId,omitempty"`

		Anniversaries []struct {
			Date struct {
				Month int32 `json:"month,omitempty"`
				Day   int32 `json:"day,omitempty"`
				Year  int32 `json:"year,omitempty"`
			} `json:"date,omitempty"`

			Description string `json:"description,omitempty"`
		} `json:"anniversaries,omitempty"`

		TaxID           string `json:"taxId,omitempty"`
		AdditionalNames []struct {
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"additionalNames,omitempty"`
	} `json:"details,omitempty"`

	IncognitoDetails *struct {
		IsIncognito bool   `json:"isIncognito,omitempty"`
		FirstName   string `json:"firstName,omitempty"`
		LastName    string `json:"lastName,omitempty"`
	} `json:"incognitoDetails,omitempty"`

	CommunicationChannels []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TypeID  string `json:"typeId,omitempty"`
			Mode    string `json:"mode,omitempty"`
			Details string `json:"details,omitempty"`
			Primary bool   `json:"primary,omitempty"`
		} `json:"value,omitempty"`
	} `json:"communicationChannels,omitempty"`

	AddressDetails []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			AddressLine1    string `json:"addressLine1,omitempty"`
			AddressLine2    string `json:"addressLine2,omitempty"`
			City            string `json:"city,omitempty"`
			CountryCode     string `json:"countryCode,omitempty"`
			IsArPrimary     bool   `json:"isArPrimary,omitempty"`
			IsPrimary       bool   `json:"isPrimary,omitempty"`
			LanguageCode    string `json:"languageCode,omitempty"`
			PostalCode      string `json:"postalCode,omitempty"`
			StateCode       string `json:"stateCode,omitempty"`
			AddressTypeCode string `json:"addressTypeCode,omitempty"`
			DistrictID      string `json:"districtId,omitempty"`
		} `json:"value,omitempty"`
	} `json:"addressDetails,omitempty"`

	IdentificationDocuments []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TypeCode         string   `json:"typeCode,omitempty"`
			Number           string   `json:"number,omitempty"`
			ExpiryDate       Date     `json:"expiryDate,omitempty"`
			IssueDate        Date     `json:"issueDate,omitempty"`
			IssueCountryCode string   `json:"issueCountryCode,omitempty"`
			IssuePlace       string   `json:"issuePlace,omitempty"`
			IssuedBy         string   `json:"issuedBy,omitempty"`
			BirthCountryCode string   `json:"birthCountryCode,omitempty"`
			BirthPlace       string   `json:"birthPlace,omitempty"`
			BirthDate        Date     `json:"birthDate,omitempty"`
			NationailtyCode  string   `json:"nationailtyCode,omitempty"`
			AdditionalInfo   string   `json:"additionalInfo,omitempty"`
			ImageRefs        []string `json:"imageRefs,omitempty"`
		} `json:"value,omitempty"`
	} `json:"identificationDocuments,omitempty"`

	Consents []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			ConsentID   string `json:"consentId,omitempty"`
			Source      string `json:"source,omitempty"`
			Description string `json:"description,omitempty"`
			ExpiryDate  Date   `json:"expiryDate,omitempty"`
			IsGranted   bool   `json:"isGranted,omitempty"`
		} `json:"value,omitempty"`
	} `json:"consents,omitempty"`

	DisabilityStatuses []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			ID string `json:"id,omitempty"`
		} `json:"value,omitempty"`
	} `json:"disabilityStatuses,omitempty"`

	AttachedEntitiesPaths     []string `json:"attachedEntitiesPaths,omitempty"`
	ExternalSystemIdentifiers []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     *struct {
			ExternalSystemID string `json:"externalSystemId,omitempty"`
			Identifier       string `json:"identifier,omitempty"`
		} `json:"value,omitempty"`
	} `json:"externalSystemIdentifiers,omitempty"`

	ProfileOwners []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			ID        string `json:"id,omitempty"`
			IsPrimary bool   `json:"isPrimary,omitempty"`
		} `json:"value,omitempty"`
	} `json:"profileOwners,omitempty"`

	Attachments []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			Name        string `json:"name,omitempty"`
			Description string `json:"description,omitempty"`
			IsGlobal    bool   `json:"isGlobal,omitempty"`
		} `json:"value,omitempty"`
	} `json:"attachments,omitempty"`

	AttachmentFiles []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			Name string `json:"name,omitempty"`
		} `json:"value,omitempty"`
	} `json:"attachmentFiles,omitempty"`

	OtherPaymentMethod []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TransactionCodeID string `json:"transactionCodeId,omitempty"`
		} `json:"value,omitempty"`
	} `json:"otherPaymentMethod,omitempty"`

	CreditCardStandalone []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TransactionCodeID string `json:"transactionCodeId,omitempty"`
			CreditCardToken   string `json:"creditCardToken,omitempty"`
			Prefix6           string `json:"prefix6,omitempty"`
			Suffix4           string `json:"suffix4,omitempty"`
			ExpiryMonth       int32  `json:"expiryMonth,omitempty"`
			ExpiryYear        int32  `json:"expiryYear,omitempty"`
		} `json:"value,omitempty"`
	} `json:"creditCardStandalone,omitempty"`

	CreditCardOffline []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TransactionCodeID string `json:"transactionCodeId,omitempty"`
			CardHolder        string `json:"cardHolder,omitempty"`
			CardNumber        string `json:"cardNumber,omitempty"`
			ExpiryYear        int32  `json:"expiryYear,omitempty"`
			ExpiryMonth       int32  `json:"expiryMonth,omitempty"`
			CardToken         string `json:"cardToken,omitempty"`
		} `json:"value,omitempty"`
	} `json:"creditCardOffline,omitempty"`

	Memberships []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			LevelCode  string `json:"levelCode,omitempty"`
			TypeCode   string `json:"typeCode,omitempty"`
			Number     string `json:"number,omitempty"`
			SinceDate  Date   `json:"sinceDate,omitempty"`
			ExpiryDate Date   `json:"expiryDate,omitempty"`
			IsActive   bool   `json:"isActive,omitempty"`
		} `json:"value,omitempty"`
	} `json:"memberships,omitempty"`

	TaxExemptDetails *struct {
		Comment string `json:"comment,omitempty"`
		Reasons []struct {
			TaxExemptReasonId string `json:"taxExemptReasonId,omitempty"`
		} `json:"reasons,omitempty"`
	} `json:"taxExemptDetails,omitempty"`
}

type PatchCompany struct {
	Details *struct {
		FullName                        string   `json:"fullName,omitempty"`
		CorporateID                     string   `json:"corporateId,omitempty"`
		Abbreviation                    string   `json:"abbreviation,omitempty"`
		TaxID                           string   `json:"taxId,omitempty"`
		IndustryCodes                   []string `json:"industryCodes,omitempty"`
		BusinessSegmentCodes            []string `json:"businessSegmentCodes,omitempty"`
		PreferredLanguageProfileDetails struct {
			LanguageCode string `json:"languageCode,omitempty"`
			FullName     string `json:"fullName,omitempty"`
		} `json:"preferredLanguageProfileDetails,omitempty"`

		PreferredCommunicationLanguageCode string `json:"preferredCommunicationLanguageCode,omitempty"`
	} `json:"details,omitempty"`

	CommunicationChannels []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TypeID  string `json:"typeId,omitempty"`
			Mode    string `json:"mode,omitempty"`
			Details string `json:"details,omitempty"`
			Primary bool   `json:"primary,omitempty"`
		} `json:"value,omitempty"`
	} `json:"communicationChannels,omitempty"`

	AddressDetails []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			AddressLine1    string `json:"addressLine1,omitempty"`
			AddressLine2    string `json:"addressLine2,omitempty"`
			City            string `json:"city,omitempty"`
			CountryCode     string `json:"countryCode,omitempty"`
			IsArPrimary     bool   `json:"isArPrimary,omitempty"`
			IsPrimary       bool   `json:"isPrimary,omitempty"`
			LanguageCode    string `json:"languageCode,omitempty"`
			PostalCode      string `json:"postalCode,omitempty"`
			StateCode       string `json:"stateCode,omitempty"`
			AddressTypeCode string `json:"addressTypeCode,omitempty"`
			DistrictID      string `json:"districtId,omitempty"`
		} `json:"value,omitempty"`
	} `json:"addressDetails,omitempty"`

	IdentificationDocuments []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TypeCode         string   `json:"typeCode,omitempty"`
			Number           string   `json:"number,omitempty"`
			ExpiryDate       Date     `json:"expiryDate,omitempty"`
			IssueDate        Date     `json:"issueDate,omitempty"`
			IssueCountryCode string   `json:"issueCountryCode,omitempty"`
			IssuePlace       string   `json:"issuePlace,omitempty"`
			IssuedBy         string   `json:"issuedBy,omitempty"`
			BirthCountryCode string   `json:"birthCountryCode,omitempty"`
			BirthPlace       string   `json:"birthPlace,omitempty"`
			BirthDate        Date     `json:"birthDate,omitempty"`
			NationailtyCode  string   `json:"nationailtyCode,omitempty"`
			AdditionalInfo   string   `json:"additionalInfo,omitempty"`
			ImageRefs        []string `json:"imageRefs,omitempty"`
		} `json:"value,omitempty"`
	} `json:"identificationDocuments,omitempty"`

	Consents []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			ConsentID   string `json:"consentId,omitempty"`
			Source      string `json:"source,omitempty"`
			Description string `json:"description,omitempty"`
			ExpiryDate  Date   `json:"expiryDate,omitempty"`
			IsGranted   bool   `json:"isGranted,omitempty"`
		} `json:"value,omitempty"`
	} `json:"consents,omitempty"`

	DisabilityStatuses []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			ID string `json:"id,omitempty"`
		} `json:"value,omitempty"`
	} `json:"disabilityStatuses,omitempty"`

	AttachedEntitiesPaths     []string `json:"attachedEntitiesPaths,omitempty"`
	ExternalSystemIdentifiers []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     *struct {
			ExternalSystemID string `json:"externalSystemId,omitempty"`
			Identifier       string `json:"identifier,omitempty"`
		} `json:"value,omitempty"`
	} `json:"externalSystemIdentifiers,omitempty"`

	ProfileOwners []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			ID        string `json:"id,omitempty"`
			IsPrimary bool   `json:"isPrimary,omitempty"`
		} `json:"value,omitempty"`
	} `json:"profileOwners,omitempty"`

	Attachments []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			Name        string `json:"name,omitempty"`
			Description string `json:"description,omitempty"`
			IsGlobal    bool   `json:"isGlobal,omitempty"`
		} `json:"value,omitempty"`
	} `json:"attachments,omitempty"`

	AttachmentFiles []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			Name string `json:"name,omitempty"`
		} `json:"value,omitempty"`
	} `json:"attachmentFiles,omitempty"`

	OtherPaymentMethod []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TransactionCodeID string `json:"transactionCodeId,omitempty"`
		} `json:"value,omitempty"`
	} `json:"otherPaymentMethod,omitempty"`

	CreditCardStandalone []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TransactionCodeID string `json:"transactionCodeId,omitempty"`
			CreditCardToken   string `json:"creditCardToken,omitempty"`
			Prefix6           string `json:"prefix6,omitempty"`
			Suffix4           string `json:"suffix4,omitempty"`
			ExpiryMonth       int32  `json:"expiryMonth,omitempty"`
			ExpiryYear        int32  `json:"expiryYear,omitempty"`
		} `json:"value,omitempty"`
	} `json:"creditCardStandalone,omitempty"`

	CreditCardOffline []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TransactionCodeID string `json:"transactionCodeId,omitempty"`
			CardHolder        string `json:"cardHolder,omitempty"`
			CardNumber        string `json:"cardNumber,omitempty"`
			ExpiryYear        int32  `json:"expiryYear,omitempty"`
			ExpiryMonth       int32  `json:"expiryMonth,omitempty"`
			CardToken         string `json:"cardToken,omitempty"`
		} `json:"value,omitempty"`
	} `json:"creditCardOffline,omitempty"`

	Memberships []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			LevelCode  string `json:"levelCode,omitempty"`
			TypeCode   string `json:"typeCode,omitempty"`
			Number     string `json:"number,omitempty"`
			SinceDate  Date   `json:"sinceDate,omitempty"`
			ExpiryDate Date   `json:"expiryDate,omitempty"`
			IsActive   bool   `json:"isActive,omitempty"`
		} `json:"value,omitempty"`
	} `json:"memberships,omitempty"`

	TaxExemptDetails *struct {
		Comment string `json:"comment,omitempty"`
		Reasons []struct {
			TaxExemptReasonId string `json:"taxExemptReasonId,omitempty"`
		} `json:"reasons,omitempty"`
	} `json:"taxExemptDetails,omitempty"`
}

type PatchTravelAgent struct {
	Details *struct {
		FullName                        string   `json:"fullName,omitempty"`
		CorporateID                     string   `json:"corporateId,omitempty"`
		Abbreviation                    string   `json:"abbreviation,omitempty"`
		TaxID                           string   `json:"taxId,omitempty"`
		IndustryCodes                   []string `json:"industryCodes,omitempty"`
		BusinessSegmentCodes            []string `json:"businessSegmentCodes,omitempty"`
		PreferredLanguageProfileDetails struct {
			LanguageCode string `json:"languageCode,omitempty"`
			FullName     string `json:"fullName,omitempty"`
		} `json:"preferredLanguageProfileDetails,omitempty"`

		PreferredCommunicationLanguageCode string `json:"preferredCommunicationLanguageCode,omitempty"`
	} `json:"details,omitempty"`

	CommunicationChannels []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TypeID  string `json:"typeId,omitempty"`
			Mode    string `json:"mode,omitempty"`
			Details string `json:"details,omitempty"`
			Primary bool   `json:"primary,omitempty"`
		} `json:"value,omitempty"`
	} `json:"communicationChannels,omitempty"`

	AddressDetails []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			AddressLine1    string `json:"addressLine1,omitempty"`
			AddressLine2    string `json:"addressLine2,omitempty"`
			City            string `json:"city,omitempty"`
			CountryCode     string `json:"countryCode,omitempty"`
			IsArPrimary     bool   `json:"isArPrimary,omitempty"`
			IsPrimary       bool   `json:"isPrimary,omitempty"`
			LanguageCode    string `json:"languageCode,omitempty"`
			PostalCode      string `json:"postalCode,omitempty"`
			StateCode       string `json:"stateCode,omitempty"`
			AddressTypeCode string `json:"addressTypeCode,omitempty"`
			DistrictID      string `json:"districtId,omitempty"`
		} `json:"value,omitempty"`
	} `json:"addressDetails,omitempty"`

	IdentificationDocuments []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TypeCode         string   `json:"typeCode,omitempty"`
			Number           string   `json:"number,omitempty"`
			ExpiryDate       Date     `json:"expiryDate,omitempty"`
			IssueDate        Date     `json:"issueDate,omitempty"`
			IssueCountryCode string   `json:"issueCountryCode,omitempty"`
			IssuePlace       string   `json:"issuePlace,omitempty"`
			IssuedBy         string   `json:"issuedBy,omitempty"`
			BirthCountryCode string   `json:"birthCountryCode,omitempty"`
			BirthPlace       string   `json:"birthPlace,omitempty"`
			BirthDate        Date     `json:"birthDate,omitempty"`
			NationailtyCode  string   `json:"nationailtyCode,omitempty"`
			AdditionalInfo   string   `json:"additionalInfo,omitempty"`
			ImageRefs        []string `json:"imageRefs,omitempty"`
		} `json:"value,omitempty"`
	} `json:"identificationDocuments,omitempty"`

	Consents []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			ConsentID   string `json:"consentId,omitempty"`
			Source      string `json:"source,omitempty"`
			Description string `json:"description,omitempty"`
			ExpiryDate  Date   `json:"expiryDate,omitempty"`
			IsGranted   bool   `json:"isGranted,omitempty"`
		} `json:"value,omitempty"`
	} `json:"consents,omitempty"`

	DisabilityStatuses []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			ID string `json:"id,omitempty"`
		} `json:"value,omitempty"`
	} `json:"disabilityStatuses,omitempty"`

	AttachedEntitiesPaths     []string `json:"attachedEntitiesPaths,omitempty"`
	ExternalSystemIdentifiers []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     *struct {
			ExternalSystemID string `json:"externalSystemId,omitempty"`
			Identifier       string `json:"identifier,omitempty"`
		} `json:"value,omitempty"`
	} `json:"externalSystemIdentifiers,omitempty"`

	ProfileOwners []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			ID        string `json:"id,omitempty"`
			IsPrimary bool   `json:"isPrimary,omitempty"`
		} `json:"value,omitempty"`
	} `json:"profileOwners,omitempty"`

	Attachments []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			Name        string `json:"name,omitempty"`
			Description string `json:"description,omitempty"`
			IsGlobal    bool   `json:"isGlobal,omitempty"`
		} `json:"value,omitempty"`
	} `json:"attachments,omitempty"`

	AttachmentFiles []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			Name string `json:"name,omitempty"`
		} `json:"value,omitempty"`
	} `json:"attachmentFiles,omitempty"`

	OtherPaymentMethod []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TransactionCodeID string `json:"transactionCodeId,omitempty"`
		} `json:"value,omitempty"`
	} `json:"otherPaymentMethod,omitempty"`

	CreditCardStandalone []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TransactionCodeID string `json:"transactionCodeId,omitempty"`
			CreditCardToken   string `json:"creditCardToken,omitempty"`
			Prefix6           string `json:"prefix6,omitempty"`
			Suffix4           string `json:"suffix4,omitempty"`
			ExpiryMonth       int32  `json:"expiryMonth,omitempty"`
			ExpiryYear        int32  `json:"expiryYear,omitempty"`
		} `json:"value,omitempty"`
	} `json:"creditCardStandalone,omitempty"`

	CreditCardOffline []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			TransactionCodeID string `json:"transactionCodeId,omitempty"`
			CardHolder        string `json:"cardHolder,omitempty"`
			CardNumber        string `json:"cardNumber,omitempty"`
			ExpiryYear        int32  `json:"expiryYear,omitempty"`
			ExpiryMonth       int32  `json:"expiryMonth,omitempty"`
			CardToken         string `json:"cardToken,omitempty"`
		} `json:"value,omitempty"`
	} `json:"creditCardOffline,omitempty"`

	Memberships []struct {
		Operation string `json:"operation,omitempty"`
		ID        string `json:"id,omitempty"`
		Value     struct {
			LevelCode  string `json:"levelCode,omitempty"`
			TypeCode   string `json:"typeCode,omitempty"`
			Number     string `json:"number,omitempty"`
			SinceDate  Date   `json:"sinceDate,omitempty"`
			ExpiryDate Date   `json:"expiryDate,omitempty"`
			IsActive   bool   `json:"isActive,omitempty"`
		} `json:"value,omitempty"`
	} `json:"memberships,omitempty"`

	TaxExemptDetails *struct {
		Comment string `json:"comment,omitempty"`
		Reasons []struct {
			TaxExemptReasonId string `json:"taxExemptReasonId,omitempty"`
		} `json:"reasons,omitempty"`
	} `json:"taxExemptDetails,omitempty"`
}

type ExternalSystem struct {
	ID   string `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Name []struct {
		Content      string `json:"content,omitempty"`
		LanguageCode string `json:"languageCode,omitempty"`
	} `json:"name,omitempty"`
	ZoneID      string `json:"zoneId,omitempty"`
	Version     int64  `json:"version,omitempty"`
	IsActive    bool   `json:"isActive"`
	ForceToUse  bool   `json:"forceToUse"`
	Description []struct {
		Content      string `json:"content,omitempty"`
		LanguageCode string `json:"languageCode,omitempty"`
	} `json:"description,omitempty"`
	DisplaySequence int32 `json:"displaySequence,omitempty"`
}

type Zone struct {
	ID                 string   `json:"id"`
	Version            int64    `json:"version"`
	OwningPropertyID   string   `json:"owningPropertyId"`
	OwningSubsidiaryID string   `json:"owningSubsidiaryId"`
	TypeCode           string   `json:"typeCode"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	StatusCode         string   `json:"statusCode"`
	EntityTypeCodes    []string `json:"entityTypeCodes"`
}

type ProfileResponse struct {
	ID       string `json:"id,omitempty"`
	RoleCode struct {
		Code         string `json:"code,omitempty"`
		LanguageCode string `json:"languageCode,omitempty"`
		Description  string `json:"description,omitempty"`
	} `json:"roleCode,omitempty"`

	TypeCode LocalizedIDCodeValue `json:"typeCode,omitempty"`
	Status   struct {
		IsActive               bool                 `json:"isActive,omitempty"`
		DeactivationReasonCode LocalizedIDCodeValue `json:"deactivationReasonCode,omitempty"`
		IsRestricted           bool                 `json:"isRestricted,omitempty"`
		RestrictionReasonCode  LocalizedIDCodeValue `json:"restrictionReasonCode,omitempty"`
		RestrictionComment     string               `json:"restrictionComment,omitempty"`
		IsLocked               bool                 `json:"isLocked,omitempty"`
		LockedReason           string               `json:"lockedReason,omitempty"`
		ForReview              bool                 `json:"forReview,omitempty"`
	} `json:"status,omitempty"`

	IndividualDetails struct {
		TitleCode       LocalizedIDCodeValue `json:"titleCode,omitempty"`
		FirstName       string               `json:"firstName,omitempty"`
		LastName        string               `json:"lastName,omitempty"`
		NationalityCode struct {
			Code         string `json:"code,omitempty"`
			LanguageCode string `json:"languageCode,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"nationalityCode,omitempty"`

		Greeting string `json:"greeting,omitempty"`
		VIPCode  struct {
			Color        string `json:"color,omitempty"`
			ID           string `json:"id,omitempty"`
			Code         string `json:"code,omitempty"`
			LanguageCode string `json:"languageCode,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"vipCode,omitempty"`

		GenderDetails           LocalizedIDCodeValue `json:"genderDetails,omitempty"`
		BirthdayAnniversaryDate struct {
			Month int32 `json:"month,omitempty"`
			Day   int32 `json:"day,omitempty"`
			Year  int32 `json:"year,omitempty"`
		} `json:"birthdayAnniversaryDate,omitempty"`

		PreferredLanguageProfileDetails struct {
			LanguageCode struct {
				Code         string `json:"code,omitempty"`
				LanguageCode string `json:"languageCode,omitempty"`
				Description  string `json:"description,omitempty"`
			} `json:"languageCode,omitempty"`

			LocalizedIndividualProfileDetails struct {
				Title     string `json:"title,omitempty"`
				FirstName string `json:"firstName,omitempty"`
				LastName  string `json:"lastName,omitempty"`
				Greeting  string `json:"greeting,omitempty"`
			} `json:"localizedIndividualProfileDetails,omitempty"`
		} `json:"preferredLanguageProfileDetails,omitempty"`

		Anniversaries []struct {
			Date struct {
				Month int32 `json:"month,omitempty"`
				Day   int32 `json:"day,omitempty"`
				Year  int32 `json:"year,omitempty"`
			} `json:"date,omitempty"`

			Description string `json:"description,omitempty"`
		} `json:"anniversaries,omitempty"`

		CompanyName                        string               `json:"companyName,omitempty"`
		PositionName                       string               `json:"positionName,omitempty"`
		MiddleName                         string               `json:"middleName,omitempty"`
		Suffix                             LocalizedIDCodeValue `json:"suffix,omitempty"`
		PreferredCommunicationLanguageCode struct {
			Code         string `json:"code,omitempty"`
			LanguageCode string `json:"languageCode,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"preferredCommunicationLanguageCode,omitempty"`

		HonoraryTitle   LocalizedIDCodeValue `json:"honoraryTitle,omitempty"`
		TaxID           string               `json:"taxId,omitempty"`
		AdditionalNames []struct {
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"additionalNames,omitempty"`
	} `json:"individualDetails,omitempty"`

	CompanyDetails struct {
		FullName                        string                 `json:"fullName,omitempty"`
		Abbreviation                    string                 `json:"abbreviation,omitempty"`
		CorporateID                     string                 `json:"corporateId,omitempty"`
		TaxID                           string                 `json:"taxId,omitempty"`
		IndustryCodes                   []LocalizedIDCodeValue `json:"industryCodes,omitempty"`
		BusinessSegmentCodes            []LocalizedIDCodeValue `json:"bUsinessSegmentCodes,omitempty"`
		PreferredLanguageProfileDetails struct {
			LanguageCode struct {
				Code         string `json:"code,omitempty"`
				LanguageCode string `json:"languageCode,omitempty"`
				Description  string `json:"description,omitempty"`
			} `json:"languageCode,omitempty"`

			FullName string `json:"fullName,omitempty"`
		} `json:"preferredLanguageProfileDetails,omitempty"`

		PreferredCommunicationLanguageCode struct {
			Code         string `json:"code,omitempty"`
			LanguageCode string `json:"languageCode,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"preferredCommunicationLanguageCode,omitempty"`
	} `json:"companyDetails,omitempty"`

	TravelAgentDetails struct {
		FullName                        string                 `json:"fullName,omitempty"`
		Abbreviation                    string                 `json:"abbreviation,omitempty"`
		IATA                            string                 `json:"iata,omitempty"`
		TaxID                           string                 `json:"taxId,omitempty"`
		BusinessSegmentCode             []LocalizedIDCodeValue `json:"businessSegmentCode,omitempty"`
		PreferredLanguageProfileDetails struct {
			LanguageCode struct {
				Code         string `json:"code,omitempty"`
				LanguageCode string `json:"languageCode,omitempty"`
				Description  string `json:"description,omitempty"`
			} `json:"languageCode,omitempty"`

			FullName string `json:"fullName,omitempty"`
		} `json:"preferredLanguageProfileDetails,omitempty"`

		PreferredCommunicationLanguageCode struct {
			Code         string `json:"code,omitempty"`
			LanguageCode string `json:"languageCode,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"preferredCommunicationLanguageCode,omitempty"`

		IndustryCodes []LocalizedIDCodeValue `json:"industryCodes,omitempty"`
	} `json:"travelAgentDetails,omitempty"`

	Addresses []struct {
		AddressTypeCode LocalizedIDCodeValue `json:"addressTypeCode,omitempty"`
		AddressLine1    string               `json:"addressLine1,omitempty"`
		AddressLine2    string               `json:"addressLine2,omitempty"`
		City            string               `json:"city,omitempty"`
		PostalCode      string               `json:"postalCode,omitempty"`
		StateCode       struct {
			Code         string `json:"code,omitempty"`
			LanguageCode string `json:"languageCode,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"stateCode,omitempty"`

		CountryCode struct {
			Code         string `json:"code,omitempty"`
			LanguageCode string `json:"languageCode,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"countryCode,omitempty"`

		LanguageCode struct {
			Code         string `json:"code,omitempty"`
			LanguageCode string `json:"languageCode,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"languageCode,omitempty"`

		IsPrimary   bool                 `json:"isPrimary,omitempty"`
		IsArPrimary bool                 `json:"isArPrimary,omitempty"`
		DistrictID  LocalizedIDCodeValue `json:"districtId,omitempty"`
		ID          string               `json:"id,omitempty"`
	} `json:"addresses,omitempty"`

	CommunicationChannels []struct {
		ID       string               `json:"id,omitempty"`
		TypeCode LocalizedIDCodeValue `json:"typeCode,omitempty"`
		ModeCode struct {
			Code         string `json:"code,omitempty"`
			LanguageCode string `json:"languageCode,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"modeCode,omitempty"`

		Details   string `json:"details,omitempty"`
		IsPrimary bool   `json:"isPrimary,omitempty"`
	} `json:"communicationChannels,omitempty"`

	DisabilityCodes  []LocalizedIDCodeValue `json:"disabilityCodes,omitempty"`
	NotesIdentifiers []string               `json:"notesIdentifiers,omitempty"`
	Memberships      []struct {
		ID         string               `json:"id,omitempty"`
		IsActive   bool                 `json:"isActive,omitempty"`
		LevelCode  LocalizedIDCodeValue `json:"levelCode,omitempty"`
		TypeCode   LocalizedIDCodeValue `json:"typeCode,omitempty"`
		Number     string               `json:"number,omitempty"`
		SinceDate  Date                 `json:"sinceDate,omitempty"`
		ExpiryDate Date                 `json:"expiryDate,omitempty"`
	} `json:"memberships,omitempty"`

	Documents []struct {
		ID               string               `json:"id,omitempty"`
		ParentVersion    int64                `json:"parentVersion,omitempty"`
		TypeCode         LocalizedIDCodeValue `json:"typeCode,omitempty"`
		Number           string               `json:"number,omitempty"`
		ExpiryDate       Date                 `json:"expiryDate,omitempty"`
		IssueDate        Date                 `json:"issueDate,omitempty"`
		IssueCountryCode struct {
			Code         string `json:"code,omitempty"`
			LanguageCode string `json:"languageCode,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"issueCountryCode,omitempty"`

		IssuePlace       string `json:"issuePlace,omitempty"`
		IssuedBy         string `json:"issuedBy,omitempty"`
		BirthCountryCode struct {
			Code         string `json:"code,omitempty"`
			LanguageCode string `json:"languageCode,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"birthCountryCode,omitempty"`

		BirthDate       Date   `json:"birthDate,omitempty"`
		BirthPlace      string `json:"birthPlace,omitempty"`
		NationalityCode struct {
			Code         string `json:"code,omitempty"`
			LanguageCode string `json:"languageCode,omitempty"`
			Description  string `json:"description,omitempty"`
		} `json:"nationalityCode,omitempty"`

		AdditionalInfo string `json:"additionalInfo,omitempty"`
	} `json:"documents,omitempty"`

	Metadata ProfileMetadata `json:"metadata,omitempty"`

	NotDuplicateProfiles struct {
		Collection []struct {
			NotDuplicateProfileID string `json:"notDuplicateProfileId,omitempty"`
			ProfileID             string `json:"profileId,omitempty"`
			ProfileRoleCode       string `json:"profileRoleCode,omitempty"`
		} `json:"collection,omitempty"`

		Count int32 `json:"count,omitempty"`
	} `json:"notDuplicateProfiles,omitempty"`

	FreezeStatus struct {
		IsFrozen   bool   `json:"isFrozen,omitempty"`
		Reason     string `json:"reason,omitempty"`
		PropertyID string `json:"propertyId,omitempty"`
		Source     string `json:"source,omitempty"`
	} `json:"freezeStatus,omitempty"`

	Consents []struct {
		ConsentID   string          `json:"consentId,omitempty"`
		IsGranted   bool            `json:"isGranted,omitempty"`
		GrantedAt   DateTime        `json:"grantedAt,omitempty"`
		ConsentType string          `json:"consentType,omitempty"`
		Source      string          `json:"source,omitempty"`
		Description string          `json:"description,omitempty"`
		Metadata    ProfileMetadata `json:"metadata,omitempty"`
		ExpiryDate  Date            `json:"expiryDate,omitempty"`
	} `json:"consents,omitempty"`

	ExternalSystemIdentifiers []struct {
		ExternalSystemID string `json:"externalSystemId,omitempty"`
		Identifier       string `json:"identifier,omitempty"`
	} `json:"externalSystemIdentifiers,omitempty"`

	GroupStatus struct {
		IsGrouped  bool   `json:"isGrouped,omitempty"`
		GroupID    string `json:"groupId,omitempty"`
		PropertyID string `json:"propertyId,omitempty"`
		Code       string `json:"code,omitempty"`
		Name       string `json:"name,omitempty"`
	} `json:"groupStatus,omitempty"`

	ContactProfiles []struct {
		ContactProfileID                  string   `json:"contactProfileId,omitempty"`
		ContactProfileRole                string   `json:"contactProfileRole,omitempty"`
		IsPrimary                         bool     `json:"isPrimary,omitempty"`
		Roles                             []string `json:"roles,omitempty"`
		ContactProfileAvailabilityDetails struct {
			ProfileStoreID                   string   `json:"profileStoreId,omitempty"`
			RegionCode                       string   `json:"regionCode,omitempty"`
			IsAvailableOutsideOfSourceRegion bool     `json:"isAvailableOutsideOfSourceRegion,omitempty"`
			AvailableInRegions               []string `json:"availableInRegions,omitempty"`
		} `json:"contactProfileAvailabilityDetails,omitempty"`
	} `json:"contactProfiles,omitempty"`

	ContactFor []struct {
		ContactProfileID                  string   `json:"contactProfileId,omitempty"`
		ContactProfileRole                string   `json:"contactProfileRole,omitempty"`
		IsPrimary                         bool     `json:"isPrimary,omitempty"`
		Roles                             []string `json:"roles,omitempty"`
		ContactProfileAvailabilityDetails struct {
			ProfileStoreID                   string   `json:"profileStoreId,omitempty"`
			RegionCode                       string   `json:"regionCode,omitempty"`
			IsAvailableOutsideOfSourceRegion bool     `json:"isAvailableOutsideOfSourceRegion,omitempty"`
			AvailableInRegions               []string `json:"availableInRegions,omitempty"`
		} `json:"contactProfileAvailabilityDetails,omitempty"`
	} `json:"contactFor,omitempty"`

	IsIncognito      bool `json:"isIncognito,omitempty"`
	IncognitoDetails struct {
		FirstName string `json:"firstName,omitempty"`
		LastName  string `json:"lastName,omitempty"`
	} `json:"incognitoDetails,omitempty"`

	ProfileOwners struct {
		Primary struct {
			ID string `json:"id,omitempty"`
		} `json:"primary,omitempty"`

		Others []struct {
			ID string `json:"id,omitempty"`
		} `json:"others,omitempty"`
	} `json:"profileOwners,omitempty"`

	ProfileLastActivity struct {
		PerformedAt DateTime `json:"performedAt,omitempty"`
		DaysAgo     int32    `json:"daysAgo,omitempty"`
		PerformedBy string   `json:"performedBy,omitempty"`
	} `json:"profileLastActivity,omitempty"`

	TaxExemption struct {
		Code    LocalizedIDCodeValue `json:"code,omitempty"`
		Comment string               `json:"comment,omitempty"`
	} `json:"taxExemption,omitempty"`

	Preferences []struct {
		ID               string               `json:"id,omitempty"`
		PreferenceCode   LocalizedIDCodeValue `json:"preferenceCode,omitempty"`
		IsPreferred      bool                 `json:"isPreferred,omitempty"`
		IsGlobal         bool                 `json:"isGlobal,omitempty"`
		OriginPropertyID string               `json:"originPropertyId,omitempty"`
		IsRequired       bool                 `json:"isRequired,omitempty"`
		Weight           int32                `json:"weight,omitempty"`
	} `json:"preferences,omitempty"`

	Version int64 `json:"version,omitempty"`

	MembershipDetails struct {
		Manual []struct {
			MembershipID string `json:"membershipId,omitempty"`
			Number       string `json:"number,omitempty"`
			Status       struct {
				IsActive bool `json:"isActive,omitempty"`
			} `json:"status,omitempty"`

			MembershipType  LocalizedIDCodeValue `json:"membershipType,omitempty"`
			MembershipLevel struct {
				Color        string `json:"color,omitempty"`
				ID           string `json:"id,omitempty"`
				Code         string `json:"code,omitempty"`
				LanguageCode string `json:"languageCode,omitempty"`
				Description  string `json:"description,omitempty"`
			} `json:"membershipLevel,omitempty"`

			SinceDate  Date `json:"sinceDate,omitempty"`
			ExpiryDate Date `json:"expiryDate,omitempty"`
		} `json:"manual,omitempty"`

		Internal []struct {
			MembershipID string `json:"membershipId,omitempty"`
			Number       string `json:"number,omitempty"`
			Status       struct {
				IsActive bool `json:"isActive,omitempty"`
			} `json:"status,omitempty"`

			MembershipType  LocalizedIDCodeValue `json:"membershipType,omitempty"`
			MembershipLevel struct {
				Color        string `json:"color,omitempty"`
				ID           string `json:"id,omitempty"`
				Code         string `json:"code,omitempty"`
				LanguageCode string `json:"languageCode,omitempty"`
				Description  string `json:"description,omitempty"`
			} `json:"membershipLevel,omitempty"`

			SinceDate  Date `json:"sinceDate,omitempty"`
			ExpiryDate Date `json:"expiryDate,omitempty"`
		} `json:"internal,omitempty"`
	} `json:"membershipDetails,omitempty"`

	RelationshipDetails struct {
		IsRoot bool `json:"isRoot,omitempty"`
	} `json:"relationshipDetails,omitempty"`

	ExternalAccountReceivable struct {
		ExternalAccountReceivableNumber string `json:"externalAccountReceivableNumber,omitempty"`
	} `json:"externalAccountReceivable,omitempty"`

	HasImage       bool   `json:"hasImage,omitempty"`
	ProfileStoreID string `json:"profileStoreId,omitempty"`
}

type FolioItem struct {
	ID              string `json:"id,omitempty"`
	Version         int64  `json:"version,omitempty"`
	ProfileID       string `json:"profileId,omitempty"`
	ProfileRoleCode struct {
		Code        string `json:"code,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"profileRoleCode,omitempty"`
	Number        string `json:"number,omitempty"`
	FolioTypeCode struct {
		Code        string `json:"code,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"folioTypeCode,omitempty"`

	FolioStyleCode struct {
		Code        string `json:"code,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"folioStyleCode,omitempty"`

	AdditionalText  string `json:"additionalText,omitempty"`
	FolioStatusCode struct {
		Code        string `json:"code,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"folioStatusCode,omitempty"`

	BillingStatusCode struct {
		Code        string `json:"code,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"billingStatusCode,omitempty"`

	BillViewAllowed       bool `json:"billViewAllowed,omitempty"`
	RemoteCheckOutAllowed bool `json:"remoteCheckOutAllowed,omitempty"`
	GrossBalance          struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"grossBalance,omitempty"`

	InvoiceNumber                                  string   `json:"invoiceNumber,omitempty"`
	IsAutomaticallyCreated                         bool     `json:"isAutomaticallyCreated,omitempty"`
	VoidedFolioWindowID                            string   `json:"voidedFolioWindowId,omitempty"`
	VoidReason                                     string   `json:"voidReason,omitempty"`
	ContactProfileID                               string   `json:"contactProfileId,omitempty"`
	ContactProfileRoles                            []string `json:"contactProfileRoles,omitempty"`
	OriginFolioWindowID                            string   `json:"originFolioWindowId,omitempty"`
	IsLocked                                       bool     `json:"isLocked,omitempty"`
	IsInvoiceGenerated                             bool     `json:"isInvoiceGenerated,omitempty"`
	IsInvoiceCorrectionFinishedSuccessfully        bool     `json:"isInvoiceCorrectionFinishedSuccessfully,omitempty"`
	IsInvoiceGenerationFinishedSuccessfully        bool     `json:"isInvoiceGenerationFinishedSuccessfully,omitempty"`
	IsAdvanceDepositCorrectionFinishedSuccessfully bool     `json:"isAdvanceDepositCorrectionFinishedSuccessfully,omitempty"`
	IsPostDirectlyInArFinishedSuccessfully         bool     `json:"isPostDirectlyInArFinishedSuccessfully,omitempty"`
	IsOriginalInvoiceProvided                      bool     `json:"isOriginalInvoiceProvided,omitempty"`
	IsInternalCorrection                           bool     `json:"isInternalCorrection,omitempty"`
	FolioConstraintsItem                           struct {
		ID           string `json:"id,omitempty"`
		BalanceLimit struct {
			Amount   float64 `json:"amount,omitempty"`
			Currency string  `json:"currency,omitempty"`
		} `json:"balanceLimit,omitempty"`

		AllowedTransactionCodesIDs []string `json:"allowedTransactionCodesIds,omitempty"`
		AllowedPaymentTypes        []string `json:"allowedPaymentTypes,omitempty"`
	} `json:"folioConstrainsItem,omitempty"`

	IsTaxExemptionApplied                                 bool `json:"isTaxExemptionApplied,omitempty"`
	IsCheckedOutWithoutInvoiceGeneration                  bool `json:"isCheckedOutWithoutInvoiceGeneration,omitempty"`
	IsGuestLedgerPrepaymentCreationFinishedSuccessfully   bool `json:"isGuestLedgerPrepaymentCreationFinishedSuccessfully,omitempty"`
	IsGuestLedgerPrepaymentCorrectionFinishedSuccessfully bool `json:"isGuestLedgerPrepaymentCorrectionFinishedSuccessfully,omitempty"`
}

type FolioInternalTransaction struct {
	ID           string `json:"id,omitempty"`
	NetUnitPrice struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"netUnitPrice,omitempty"`

	GrossUnitPrice struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"grossUnitPrice,omitempty"`

	Quantity  float64 `json:"quantity,omitempty"`
	NetAmount struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"netAmount,omitempty"`

	GrossAmount struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"grossAmount,omitempty"`

	TransactionCode struct {
		Code        string `json:"code,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"transactionCode,omitempty"`

	Comment           string `json:"comment,omitempty"`
	CashierNumber     int64  `json:"cashierNumber,omitempty"`
	TransactionAmount struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"transactionAmount,omitempty"`

	IsMainTransaction bool `json:"isMainTransaction,omitempty"`
	Taxes             []struct {
		TransactionCode struct {
			Code        string `json:"code,omitempty"`
			Description string `json:"description,omitempty"`
		} `json:"transactionCode,omitempty"`

		UnitPrice struct {
			Amount   float64 `json:"amount,omitempty"`
			Currency string  `json:"currency,omitempty"`
		} `json:"unitPrice,omitempty"`

		Quantity float64 `json:"quantity,omitempty"`

		Amount struct {
			Amount   float64 `json:"amount,omitempty"`
			Currency string  `json:"currency,omitempty"`
		} `json:"amount,omitempty"`

		IsIncluded bool   `json:"isIncluded,omitempty"`
		TaxRuleID  string `json:"taxRuleId,omitempty"`
	} `json:"taxes,omitempty"`

	RevenueDate string `json:"revenueDate,omitempty"`
}

type FolioTransaction struct {
	ID            string   `json:"id,omitempty"`
	Version       int64    `json:"version,omitempty"`
	CreationDate  DateTime `json:"creationDate,omitempty"`
	BusinessDate  DateTime `json:"businessDate,omitempty"`
	ExternalDate  string   `json:"externalDate,omitempty"`
	WorkstationID string   `json:"workstationId,omitempty"`
	NetUnitPrice  struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"netUnitPrice,omitempty"`

	GrossUnitPrice struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"grossUnitPrice,omitempty"`

	BaseUnitPrice struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"baseUnitPrice,omitempty"`

	Quantity float64 `json:"quantity,omitempty"`

	NetAmount struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"netAmount,omitempty"`

	GrossAmount struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"grossAmount,omitempty"`

	TransactionCode struct {
		Code        string `json:"code,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"transactionCode,omitempty"`

	Comment                      string  `json:"comment,omitempty"`
	Remarks                      string  `json:"remarks,omitempty"`
	Description                  string  `json:"description,omitempty"`
	CheckNumber                  string  `json:"checkNumber,omitempty"`
	HasGuestCheckDetails         bool    `json:"hasGuestCheckDetails,omitempty"`
	Reference                    string  `json:"reference,omitempty"`
	NumberOfCovers               []int32 `json:"numberOfCovers,omitempty"`
	TransactionGroupTypeCode     string  `json:"transactionGroupTypeCode,omitempty"`
	CashierID                    int64   `json:"cashierId,omitempty"`
	IsVoided                     bool    `json:"isVoided,omitempty"`
	IsVoidedCreditCardOperation  bool    `json:"isVoidedCreditCardOperation,omitempty"`
	IsVoidingCreditCardOperation bool    `json:"isVoidingCreditCardOperation,omitempty"`
	IsRefunded                   bool    `json:"isRefunded,omitempty"`
	RefundedAmount               struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"refundedAmount,omitempty"`

	InternalTransactions []FolioInternalTransaction `json:"internalTransactions,omitempty"`
	IsPackage            bool                       `json:"isPackage,omitempty"`
	TransactionAmount    struct {
		Amount   float64 `json:"amount,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"transactionAmount,omitempty"`
	VoidedFolioTransactionBusinessDate string `json:"voidedFolioTransactionBusinessDate,omitempty"`
	PostingMethod                      string `json:"postingMethod,omitempty"`
	PostingOriginator                  string `json:"postingOriginator,omitempty"`
	UserDetails                        struct {
		ID        string `json:"id,omitempty"`
		FirstName string `json:"firstName,omitempty"`
		LastName  string `json:"lastName,omitempty"`
		FullName  string `json:"fullName,omitempty"`
	} `json:"userDetails,omitempty"`

	IsRounding     bool `json:"isRounding,omitempty"`
	OriginatorInfo struct {
		ProfileID     string `json:"profileId,omitempty"`
		AccountID     string `json:"accountId,omitempty"`
		RoomID        string `json:"roomId,omitempty"`
		AccountNumber string `json:"accountNumber,omitempty"`
		RoomNumber    string `json:"roomNumber,omitempty"`
	} `json:"originatorInfo,omitempty"`

	PerformedActions              []string `json:"performedActions,omitempty"`
	IsAdvancedDepositCredit       bool     `json:"isAdvancedDepositCredit,omitempty"`
	IsTaxTransaction              bool     `json:"isTaxTransaction,omitempty"`
	IsTaxIncluded                 bool     `json:"isTaxIncluded,omitempty"`
	OriginatorID                  string   `json:"originatorId,omitempty"`
	ShowALlInternalTransactions   bool     `json:"showAllInternalTransactions,omitempty"`
	IsAdvancedDeposit             bool     `json:"isAdvancedDeposit,omitempty"`
	IsOnSitePayment               bool     `json:"isOnSitePayment,omitempty"`
	FixedChargeID                 string   `json:"fixedChargeId,omitempty"`
	CreditCardOperationID         string   `json:"creditCardOperationId,omitempty"`
	IsCreditCardOperation         bool     `json:"isCreditCardOperation,omitempty"`
	IsGuestLedgerPrepaymentCredit bool     `json:"isGuestLedgerPrepaymentCredit,omitempty"`
	IsGuestLedgerPrepayment       bool     `json:"isGuestLedgerPrepayment,omitempty"`
}

type AggregatedExtractResult struct {
	TaxRule              []TaxRule             `json:"taxRule"`
	TransactionCodes     []TransactionCode     `json:"transactionCode"`
	TransactionGroups    []TransactionGroup    `json:"transactionGroup"`
	TransactionSubGroups []TransactionSubGroup `json:"TransactionSubGroup"`
}

type AggregatedPaginatedResponse struct {
	Results AggregatedExtractResult `json:"results"`
	Paging  struct {
		PageNumber int `json:"pageNumber"`
		PageSize   int `json:"pageSize"`
		TotalCount int `json:"totalCount"`
	} `json:"paging"`
}

func (r AggregatedPaginatedResponse) TotalCount() int {
	return len(r.Results.TaxRule) + len(r.Results.TransactionCodes) + len(r.Results.TransactionGroups) + len(r.Results.TransactionSubGroups)
}
