package btp

type GlobalAccount struct {
	GUID             string `json:"guid,omitempty"`
	DisplayName      string `json:"displayName,omitempty"`
	CreatedDate      int64  `json:"createdDate,omitempty"`
	ModifiedDate     int64  `json:"modifiedDate,omitempty"`
	EntityState      string `json:"entityState,omitempty"`
	StateMessage     string `json:"stateMessage,omitempty"`
	Subdomain        string `json:"subdomain,omitempty"`
	ContractStatus   string `json:"contractStatus,omitempty"`
	CommercialModel  string `json:"commercialModel,omitempty"`
	ConsumptionBased bool   `json:"consumptionBased,omitempty"`
	LicenseType      string `json:"licenseType,omitempty"`
	GeoAccess        string `json:"geoAccess,omitempty"`
	RenewalDate      int64  `json:"renewalDate,omitempty"`
}

type Subaccount struct {
	GUID              string `json:"guid"`
	TechnicalName     string `json:"technicalName"`
	DisplayName       string `json:"displayName"`
	GlobalAccountGUID string `json:"globalAccountGUID"`
	ParentGUID        string `json:"parentGUID"`
	ParentType        string `json:"parentType"`
	Region            string `json:"region"`
	Subdomain         string `json:"subdomain"`
	BetaEnabled       bool   `json:"betaEnabled"`
	UsedForProduction string `json:"usedForProduction"`
	Description       string `json:"description"`
	State             string `json:"state"`
	StateMessage      string `json:"stateMessage"`
	CreatedDate       int64  `json:"createdDate"`
	CreatedBy         string `json:"createdBy"`
	ModifiedDate      int64  `json:"modifiedDate"`
}
