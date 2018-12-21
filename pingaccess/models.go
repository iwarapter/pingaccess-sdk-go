package pingaccess

//ApplicationView - An application.
type ApplicationView struct {
	AccessValidatorId  int                    `json:"accessValidatorId",omitempty`
	AgentId            int                    `json:"agentId"`
	ApplicationType    string                 `json:"applicationType",omitempty`
	CaseSensitivePath  bool                   `json:"caseSensitivePath",omitempty`
	ContextRoot        string                 `json:"contextRoot"`
	DefaultAuthType    string                 `json:"defaultAuthType"`
	Description        string                 `json:"description",omitempty`
	Destination        string                 `json:"destination",omitempty`
	Enabled            bool                   `json:"enabled",omitempty`
	Id                 int                    `json:"id",omitempty`
	IdentityMappingIds map[string]int         `json:"identityMappingIds",omitempty`
	Name               string                 `json:"name"`
	Policy             map[string]interface{} `json:"policy",omitempty`
	Realm              string                 `json:"realm",omitempty`
	RequireHTTPS       bool                   `json:"requireHTTPS",omitempty`
	SiteId             int                    `json:"siteId"`
	VirtualHostIds     []*int                 `json:"virtualHostIds"`
	WebSessionId       int                    `json:"webSessionId",omitempty`
}

//ApplicationsView - A collection of applications.
type ApplicationsView struct {
	Items []*ApplicationView `json:"items"`
}

//ConfigurationDependentFieldOption - Configuration of the dependent field option.
type ConfigurationDependentFieldOption struct {
	Options []*ConfigurationOption `json:"options"`
	Value   string                 `json:"value"`
}

//ConfigurationField - Details for configuration in the administrator console.
type ConfigurationField struct {
	Advanced     bool                     `json:"advanced"`
	ButtonGroup  string                   `json:"buttonGroup"`
	Default      string                   `json:"default"`
	Deselectable bool                     `json:"deselectable"`
	Fields       []*ConfigurationField    `json:"fields"`
	Help         Help                     `json:"help"`
	Label        string                   `json:"label"`
	Name         string                   `json:"name"`
	Options      []*ConfigurationOption   `json:"options"`
	ParentField  ConfigurationParentField `json:"parentField"`
	Required     bool                     `json:"required"`
	Type         string                   `json:"type"`
}

//ConfigurationOption - The configuration option attributes.
type ConfigurationOption struct {
	Category string `json:"category"`
	Label    string `json:"label"`
	Value    string `json:"value"`
}

//ConfigurationParentField - Configuration of the parent field.
type ConfigurationParentField struct {
	DependentFieldOptions []*ConfigurationDependentFieldOption `json:"dependentFieldOptions"`
	FieldName             string                               `json:"fieldName"`
}

//Help - Configuration of the help context of a configuration field.
type Help struct {
	Content string `json:"content"`
	Title   string `json:"title"`
	Url     string `json:"url"`
}

//IdentityMappingDescriptor
type IdentityMappingDescriptor struct {
	ClassName           string                `json:"className"`
	ConfigurationFields []*ConfigurationField `json:"configurationFields"`
	Label               string                `json:"label"`
	Type                string                `json:"type"`
}

//IdentityMappingDescriptorsView
type IdentityMappingDescriptorsView struct {
	Items []*IdentityMappingDescriptor `json:"items"`
}

//IdentityMappingView
type IdentityMappingView struct {
	ClassName     string                 `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            int                    `json:"id",omitempty`
	Name          string                 `json:"name"`
}

//IdentityMappingsView
type IdentityMappingsView struct {
	Items []*IdentityMappingView `json:"items"`
}

//MethodView - HTTP Method
type MethodView struct {
	Name string `json:"name"`
}

//MethodsView
type MethodsView struct {
	Items []*MethodView `json:"items"`
}

//ResourceView - A resource.
type ResourceView struct {
	Anonymous               bool                   `json:"anonymous",omitempty`
	ApplicationId           int                    `json:"applicationId",omitempty`
	AuditLevel              string                 `json:"auditLevel",omitempty`
	DefaultAuthTypeOverride string                 `json:"defaultAuthTypeOverride"`
	Enabled                 bool                   `json:"enabled",omitempty`
	Id                      int                    `json:"id",omitempty`
	Methods                 []*string              `json:"methods"`
	Name                    string                 `json:"name"`
	PathPrefixes            []*string              `json:"pathPrefixes"`
	Policy                  map[string]interface{} `json:"policy",omitempty`
	RootResource            bool                   `json:"rootResource",omitempty`
}

//ResourcesView
type ResourcesView struct {
	Items []*ResourceView `json:"items"`
}

//RuleDescriptor
type RuleDescriptor struct {
	AgentCachingDisabled bool                  `json:"agentCachingDisabled"`
	Category             string                `json:"category"`
	ClassName            string                `json:"className"`
	ConfigurationFields  []*ConfigurationField `json:"configurationFields"`
	Label                string                `json:"label"`
	Modes                []string              `json:"modes"`
	Type                 string                `json:"type"`
}

//RuleDescriptorsView
type RuleDescriptorsView struct {
	Items []*RuleDescriptor `json:"items"`
}

//RuleView
type RuleView struct {
	ClassName             string                 `json:"className"`
	Configuration         map[string]interface{} `json:"configuration"`
	Id                    int                    `json:"id",omitempty`
	Name                  string                 `json:"name"`
	SupportedDestinations []*string              `json:"supportedDestinations",omitempty`
}

//RulesView
type RulesView struct {
	Items []*RuleView `json:"items"`
}

//SiteAuthenticatorDescriptor
type SiteAuthenticatorDescriptor struct {
	ClassName           string                `json:"className"`
	ConfigurationFields []*ConfigurationField `json:"configurationFields"`
	Label               string                `json:"label"`
	Type                string                `json:"type"`
}

//SiteAuthenticatorDescriptorsView
type SiteAuthenticatorDescriptorsView struct {
	Items []*SiteAuthenticatorDescriptor `json:"items"`
}

//SiteAuthenticatorView
type SiteAuthenticatorView struct {
	ClassName     string                 `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            int                    `json:"id",omitempty`
	Name          string                 `json:"name"`
}

//SiteAuthenticatorsView
type SiteAuthenticatorsView struct {
	Items []*SiteAuthenticatorView `json:"items"`
}

//SiteView - A site.
type SiteView struct {
	AvailabilityProfileId     int       `json:"availabilityProfileId",omitempty`
	ExpectedHostname          string    `json:"expectedHostname",omitempty`
	Id                        int       `json:"id",omitempty`
	KeepAliveTimeout          int       `json:"keepAliveTimeout",omitempty`
	LoadBalancingStrategyId   int       `json:"loadBalancingStrategyId",omitempty`
	MaxConnections            int       `json:"maxConnections",omitempty`
	MaxWebSocketConnections   int       `json:"maxWebSocketConnections",omitempty`
	Name                      string    `json:"name"`
	Secure                    bool      `json:"secure",omitempty`
	SendPaCookie              bool      `json:"sendPaCookie",omitempty`
	SiteAuthenticatorIds      []*int    `json:"siteAuthenticatorIds",omitempty`
	SkipHostnameVerification  bool      `json:"skipHostnameVerification",omitempty`
	Targets                   []*string `json:"targets"`
	TrustedCertificateGroupId int       `json:"trustedCertificateGroupId",omitempty`
	UseProxy                  bool      `json:"useProxy",omitempty`
	UseTargetHostHeader       bool      `json:"useTargetHostHeader",omitempty`
}

//SitesView - A collection of sites items.
type SitesView struct {
	Items []*SiteView `json:"items"`
}

//VirtualHost
type VirtualHost struct {
	AgentResourceCacheTTL     int    `json:"agentResourceCacheTTL"`
	Host                      string `json:"host"`
	Id                        int    `json:"id",omitempty`
	KeyPairId                 int    `json:"keyPairId"`
	Port                      int    `json:"port"`
	TrustedCertificateGroupId int    `json:"trustedCertificateGroupId"`
}

//VirtualHostView
type VirtualHostView struct {
	AgentResourceCacheTTL     int    `json:"agentResourceCacheTTL",omitempty`
	Host                      string `json:"host"`
	Id                        int    `json:"id",omitempty`
	KeyPairId                 int    `json:"keyPairId",omitempty`
	Port                      int    `json:"port"`
	TrustedCertificateGroupId int    `json:"trustedCertificateGroupId",omitempty`
}

//VirtualHostsView
type VirtualHostsView struct {
	Items []*VirtualHostView `json:"items"`
}
