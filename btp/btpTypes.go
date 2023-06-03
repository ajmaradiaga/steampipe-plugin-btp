package btp

type CustomProperty struct {
	AccountGUID string `json:"accountGUID"`
	Key         string `json:"key"`
	Value       string `json:"value"`
}

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

type Directory struct {
	GUID              string           `json:"guid"`
	ParentType        string           `json:"parentType"`
	GlobalAccountGUID string           `json:"globalAccountGUID"`
	DisplayName       string           `json:"displayName"`
	CreatedDate       int64            `json:"createdDate"`
	CreatedBy         string           `json:"createdBy"`
	ModifiedDate      int64            `json:"modifiedDate"`
	EntityState       string           `json:"entityState"`
	StateMessage      string           `json:"stateMessage"`
	DirectoryType     string           `json:"directoryType"`
	DirectoryFeatures []string         `json:"directoryFeatures"`
	CustomProperties  []CustomProperty `json:"customProperties"`
	ContractStatus    string           `json:"contractStatus"`
	ConsumptionBased  bool             `json:"consumptionBased"`
	ParentGUID0       string           `json:"parentGuid"`
	ParentGUID1       string           `json:"parentGUID"`
}

type DataCenter struct {
	Name                   string `json:"name"`
	DisplayName            string `json:"displayName"`
	Region                 string `json:"region"`
	Environment            string `json:"environment"`
	IaasProvider           string `json:"iaasProvider"`
	SupportsTrial          bool   `json:"supportsTrial"`
	ProvisioningServiceURL string `json:"provisioningServiceUrl"`
	SaasRegistryServiceURL string `json:"saasRegistryServiceUrl"`
	Domain                 string `json:"domain"`
	GeoAccess              string `json:"geoAccess"`
}

type Service struct {
	Name             string `json:"name"`
	DisplayName      string `json:"displayName"`
	Description      string `json:"description"`
	BusinessCategory struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
	} `json:"businessCategory"`
	OwnerType     string `json:"ownerType"`
	TermsOfUseURL string `json:"termsOfUseUrl"`
	ServicePlans  []struct {
		Name                      string  `json:"name"`
		DisplayName               string  `json:"displayName"`
		Description               string  `json:"description"`
		UniqueIdentifier          string  `json:"uniqueIdentifier"`
		ProvisioningMethod        string  `json:"provisioningMethod"`
		Amount                    float64 `json:"amount"`
		RemainingAmount           float64 `json:"remainingAmount"`
		ProvidedBy                string  `json:"providedBy"`
		Beta                      bool    `json:"beta"`
		AvailableForInternal      bool    `json:"availableForInternal"`
		InternalQuotaLimit        any     `json:"internalQuotaLimit"`
		AutoAssign                bool    `json:"autoAssign"`
		AutoDistributeAmount      any     `json:"autoDistributeAmount"`
		MaxAllowedSubaccountQuota any     `json:"maxAllowedSubaccountQuota"`
		Category                  string  `json:"category"`
		SourceEntitlements        []struct {
			EntitlementName    string  `json:"entitlementName"`
			Amount             float64 `json:"amount"`
			ProductID          string  `json:"productId"`
			Restricted         bool    `json:"restricted"`
			AllowedSubaccounts any     `json:"allowedSubaccounts"`
			CommercialModel    struct {
				Name             string `json:"name"`
				DisplayName      string `json:"displayName"`
				ConsumptionBased bool   `json:"consumptionBased"`
				Description      string `json:"description"`
			} `json:"commercialModel"`
			AutoAssign bool `json:"autoAssign"`
		} `json:"sourceEntitlements"`
		DataCenters []struct {
			Name                   string `json:"name"`
			DisplayName            string `json:"displayName"`
			Region                 string `json:"region"`
			Environment            string `json:"environment"`
			IaasProvider           string `json:"iaasProvider"`
			SupportsTrial          bool   `json:"supportsTrial"`
			ProvisioningServiceURL string `json:"provisioningServiceUrl"`
			SaasRegistryServiceURL string `json:"saasRegistryServiceUrl"`
			Domain                 string `json:"domain"`
			IsMainDataCenter       any    `json:"isMainDataCenter"`
			GeoAccess              string `json:"geoAccess"`
		} `json:"dataCenters"`
		Resources              any   `json:"resources"`
		DependsOn              any   `json:"dependsOn"`
		AdditionalPlanFeatures []any `json:"additionalPlanFeatures"`
		AssignedResources      any   `json:"assignedResources"`
		Unlimited              bool  `json:"unlimited"`
	} `json:"servicePlans"`
	IconBase64             string `json:"iconBase64"`
	ApplicationCoordinates struct {
		IconFormat   string `json:"iconFormat"`
		InventoryIds []struct {
			Key string `json:"key"`
		} `json:"inventoryIds"`
		Visibility         string `json:"visibility"`
		ServiceDescription []struct {
			LinkCategory        string `json:"linkCategory"`
			Title               string `json:"title"`
			PropagateTheme      string `json:"propagateTheme"`
			DescriptionCategory string `json:"descriptionCategory"`
			LinkURL             string `json:"linkURL"`
		} `json:"serviceDescription"`
		ServiceCategories []struct {
			Name string `json:"name"`
		} `json:"serviceCategories"`
		RegionInformation []struct {
			Key string `json:"key"`
		} `json:"regionInformation"`
		CFService struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Metadata    struct {
				DisplayName string `json:"displayName"`
			} `json:"metadata"`
			Plans []struct {
				TechnicalName string `json:"technicalName"`
				Name          string `json:"name"`
				Description   string `json:"description"`
				Metadata      struct {
				} `json:"metadata"`
			} `json:"plans"`
		} `json:"CFService"`
	} `json:"applicationCoordinates"`
}
