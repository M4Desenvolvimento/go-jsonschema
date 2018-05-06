package p

// AWSCodeCommitConnection
type AWSCodeCommitConnection struct {
	AccessKeyID                 string `json:"accessKeyID"`
	InitialRepositoryEnablement bool   `json:"initialRepositoryEnablement,omitempty"`
	Region                      string `json:"region"`
	RepositoryPathPattern       string `json:"repositoryPathPattern,omitempty"`
	SecretAccessKey             string `json:"secretAccessKey"`
}

// BitbucketServerConnection
type BitbucketServerConnection struct {
	Certificate                 string `json:"certificate,omitempty"`
	GitURLType                  string `json:"gitURLType,omitempty"`
	InitialRepositoryEnablement bool   `json:"initialRepositoryEnablement,omitempty"`
	Password                    string `json:"password,omitempty"`
	RepositoryPathPattern       string `json:"repositoryPathPattern,omitempty"`
	Token                       string `json:"token,omitempty"`
	Url                         string `json:"url"`
	Username                    string `json:"username,omitempty"`
}

// ExperimentalFeatures Experimental features to enable or disable. Features that are now enabled by default are marked as deprecated.
type ExperimentalFeatures struct {
	HostSurveysLocally     string `json:"hostSurveysLocally,omitempty"`
	JumpToDefOSSIndex      string `json:"jumpToDefOSSIndex,omitempty"`
	SearchTimeoutParameter string `json:"searchTimeoutParameter,omitempty"`
	ShowMissingRepos       string `json:"showMissingRepos,omitempty"`
}

// GitHubConnection
type GitHubConnection struct {
	Certificate                 string   `json:"certificate,omitempty"`
	GitURLType                  string   `json:"gitURLType,omitempty"`
	InitialRepositoryEnablement bool     `json:"initialRepositoryEnablement,omitempty"`
	PreemptivelyClone           bool     `json:"preemptivelyClone,omitempty"`
	Repos                       []string `json:"repos,omitempty"`
	RepositoryPathPattern       string   `json:"repositoryPathPattern,omitempty"`
	RepositoryQuery             []string `json:"repositoryQuery,omitempty"`
	Token                       string   `json:"token"`
	Url                         string   `json:"url,omitempty"`
}

// GitLabConnection
type GitLabConnection struct {
	Certificate                 string   `json:"certificate,omitempty"`
	GitURLType                  string   `json:"gitURLType,omitempty"`
	InitialRepositoryEnablement bool     `json:"initialRepositoryEnablement,omitempty"`
	ProjectQuery                []string `json:"projectQuery,omitempty"`
	RepositoryPathPattern       string   `json:"repositoryPathPattern,omitempty"`
	Token                       string   `json:"token"`
	Url                         string   `json:"url"`
}

// GitoliteConnection
type GitoliteConnection struct {
	Blacklist                  string `json:"blacklist,omitempty"`
	Host                       string `json:"host"`
	PhabricatorMetadataCommand string `json:"phabricatorMetadataCommand,omitempty"`
	Prefix                     string `json:"prefix"`
}

// Langservers
type Langservers struct {
	Address               string                 `json:"address"`
	Disabled              bool                   `json:"disabled,omitempty"`
	InitializationOptions map[string]interface{} `json:"initializationOptions,omitempty"`
	Language              string                 `json:"language"`
	Metadata              *Metadata              `json:"metadata,omitempty"`
}

// Links
type Links struct {
	Blob       string `json:"blob,omitempty"`
	Commit     string `json:"commit,omitempty"`
	Repository string `json:"repository,omitempty"`
	Tree       string `json:"tree,omitempty"`
}

// Metadata Language server metadata. Used to populate various UI elements.
type Metadata struct {
	DocsURL      string `json:"docsURL,omitempty"`
	Experimental bool   `json:"experimental,omitempty"`
	HomepageURL  string `json:"homepageURL,omitempty"`
	IssuesURL    string `json:"issuesURL,omitempty"`
}

// OpenIDConnectAuthProvider Configures the OpenID Connect authentication provider for SSO.
type OpenIDConnectAuthProvider struct {
	ClientID           string `json:"clientID"`
	ClientSecret       string `json:"clientSecret"`
	Issuer             string `json:"issuer"`
	OverrideToken      string `json:"overrideToken,omitempty"`
	RequireEmailDomain string `json:"requireEmailDomain,omitempty"`
}

// Phabricator
type Phabricator struct {
	Repos []*Repos `json:"repos,omitempty"`
	Token string   `json:"token,omitempty"`
	Url   string   `json:"url,omitempty"`
}

// Repos
type Repos struct {
	Callsign string `json:"callsign"`
	Path     string `json:"path"`
}

// Repository
type Repository struct {
	Links *Links `json:"links,omitempty"`
	Path  string `json:"path"`
	Type  string `json:"type,omitempty"`
	Url   string `json:"url"`
}

// SAMLAuthProvider Configures the SAML authentication provider for SSO.
type SAMLAuthProvider struct {
	IdentityProviderMetadata    string `json:"identityProviderMetadata,omitempty"`
	IdentityProviderMetadataURL string `json:"identityProviderMetadataURL,omitempty"`
	ServiceProviderCertificate  string `json:"serviceProviderCertificate"`
	ServiceProviderPrivateKey   string `json:"serviceProviderPrivateKey"`
}

// SMTPServerConfig The SMTP server used to send transactional emails (such as email verifications, reset-password emails, and notifications).
type SMTPServerConfig struct {
	Authentication string `json:"authentication"`
	Domain         string `json:"domain,omitempty"`
	Host           string `json:"host"`
	Password       string `json:"password,omitempty"`
	Port           int    `json:"port"`
	Username       string `json:"username,omitempty"`
}

// SearchSavedQueries
type SearchSavedQueries struct {
	Description    string `json:"description"`
	Key            string `json:"key"`
	Notify         bool   `json:"notify,omitempty"`
	NotifySlack    bool   `json:"notifySlack,omitempty"`
	Query          string `json:"query"`
	ShowOnHomepage bool   `json:"showOnHomepage,omitempty"`
}

// SearchScope
type SearchScope struct {
	Description string `json:"description,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}

// Settings Configuration settings for users and organizations on Sourcegraph.
type Settings struct {
	Motd                   []string                  `json:"motd,omitempty"`
	NotificationsSlack     *SlackNotificationsConfig `json:"notifications.slack,omitempty"`
	SearchRepositoryGroups map[string][]string       `json:"search.repositoryGroups,omitempty"`
	SearchSavedQueries     []*SearchSavedQueries     `json:"search.savedQueries,omitempty"`
	SearchScopes           []*SearchScope            `json:"search.scopes,omitempty"`
}

// SiteConfiguration Configuration for a Sourcegraph site.
type SiteConfiguration struct {
	AdminUsernames                    string                       `json:"adminUsernames,omitempty"`
	AppURL                            string                       `json:"appURL,omitempty"`
	AuthAllowSignup                   bool                         `json:"auth.allowSignup,omitempty"`
	AuthDisableAccessTokens           bool                         `json:"auth.disableAccessTokens,omitempty"`
	AuthOpenIDConnect                 *OpenIDConnectAuthProvider   `json:"auth.openIDConnect,omitempty"`
	AuthProvider                      string                       `json:"auth.provider,omitempty"`
	AuthPublic                        bool                         `json:"auth.public,omitempty"`
	AuthSaml                          *SAMLAuthProvider            `json:"auth.saml,omitempty"`
	AuthUserIdentityHTTPHeader        string                       `json:"auth.userIdentityHTTPHeader,omitempty"`
	AuthUserOrgMap                    map[string][]string          `json:"auth.userOrgMap,omitempty"`
	AwsCodeCommit                     []*AWSCodeCommitConnection   `json:"awsCodeCommit,omitempty"`
	BitbucketServer                   []*BitbucketServerConnection `json:"bitbucketServer,omitempty"`
	BlacklistGoGet                    []string                     `json:"blacklistGoGet,omitempty"`
	CorsOrigin                        string                       `json:"corsOrigin,omitempty"`
	DisableAutoGitUpdates             bool                         `json:"disableAutoGitUpdates,omitempty"`
	DisableBrowserExtension           bool                         `json:"disableBrowserExtension,omitempty"`
	DisableBuiltInSearches            bool                         `json:"disableBuiltInSearches,omitempty"`
	DisableExampleSearches            bool                         `json:"disableExampleSearches,omitempty"`
	DisablePublicRepoRedirects        bool                         `json:"disablePublicRepoRedirects,omitempty"`
	DisableTelemetry                  bool                         `json:"disableTelemetry,omitempty"`
	DontIncludeSymbolResultsByDefault bool                         `json:"dontIncludeSymbolResultsByDefault,omitempty"`
	EmailAddress                      string                       `json:"email.address,omitempty"`
	EmailSmtp                         *SMTPServerConfig            `json:"email.smtp,omitempty"`
	ExecuteGradleOriginalRootPaths    string                       `json:"executeGradleOriginalRootPaths,omitempty"`
	ExperimentalFeatures              *ExperimentalFeatures        `json:"experimentalFeatures,omitempty"`
	GitMaxConcurrentClones            int                          `json:"gitMaxConcurrentClones,omitempty"`
	GitOriginMap                      string                       `json:"gitOriginMap,omitempty"`
	Github                            []*GitHubConnection          `json:"github,omitempty"`
	GithubClientID                    string                       `json:"githubClientID,omitempty"`
	GithubClientSecret                string                       `json:"githubClientSecret,omitempty"`
	GithubEnterpriseAccessToken       string                       `json:"githubEnterpriseAccessToken,omitempty"`
	GithubEnterpriseCert              string                       `json:"githubEnterpriseCert,omitempty"`
	GithubEnterpriseURL               string                       `json:"githubEnterpriseURL,omitempty"`
	GithubPersonalAccessToken         string                       `json:"githubPersonalAccessToken,omitempty"`
	Gitlab                            []*GitLabConnection          `json:"gitlab,omitempty"`
	Gitolite                          []*GitoliteConnection        `json:"gitolite,omitempty"`
	HtmlBodyBottom                    string                       `json:"htmlBodyBottom,omitempty"`
	HtmlBodyTop                       string                       `json:"htmlBodyTop,omitempty"`
	HtmlHeadBottom                    string                       `json:"htmlHeadBottom,omitempty"`
	HtmlHeadTop                       string                       `json:"htmlHeadTop,omitempty"`
	HttpToHttpsRedirect               bool                         `json:"httpToHttpsRedirect,omitempty"`
	Langservers                       []*Langservers               `json:"langservers,omitempty"`
	LightstepAccessToken              string                       `json:"lightstepAccessToken,omitempty"`
	LightstepProject                  string                       `json:"lightstepProject,omitempty"`
	MaxReposToSearch                  int                          `json:"maxReposToSearch,omitempty"`
	NoGoGetDomains                    string                       `json:"noGoGetDomains,omitempty"`
	OidcClientID                      string                       `json:"oidcClientID,omitempty"`
	OidcClientSecret                  string                       `json:"oidcClientSecret,omitempty"`
	OidcEmailDomain                   string                       `json:"oidcEmailDomain,omitempty"`
	OidcOverrideToken                 string                       `json:"oidcOverrideToken,omitempty"`
	OidcProvider                      string                       `json:"oidcProvider,omitempty"`
	Phabricator                       []*Phabricator               `json:"phabricator,omitempty"`
	PhabricatorURL                    string                       `json:"phabricatorURL,omitempty"`
	PrivateArtifactRepoID             string                       `json:"privateArtifactRepoID,omitempty"`
	PrivateArtifactRepoPassword       string                       `json:"privateArtifactRepoPassword,omitempty"`
	PrivateArtifactRepoURL            string                       `json:"privateArtifactRepoURL,omitempty"`
	PrivateArtifactRepoUsername       string                       `json:"privateArtifactRepoUsername,omitempty"`
	RepoListUpdateInterval            int                          `json:"repoListUpdateInterval,omitempty"`
	ReposList                         []*Repository                `json:"repos.list,omitempty"`
	SamlIDProviderMetadataURL         string                       `json:"samlIDProviderMetadataURL,omitempty"`
	SamlSPCert                        string                       `json:"samlSPCert,omitempty"`
	SamlSPKey                         string                       `json:"samlSPKey,omitempty"`
	SearchScopes                      []*SearchScope               `json:"searchScopes,omitempty"`
	SecretKey                         string                       `json:"secretKey,omitempty"`
	Settings                          *Settings                    `json:"settings,omitempty"`
	SiteID                            string                       `json:"siteID,omitempty"`
	SsoUserHeader                     string                       `json:"ssoUserHeader,omitempty"`
	TlsLetsencrypt                    string                       `json:"tls.letsencrypt,omitempty"`
	TlsCert                           string                       `json:"tlsCert,omitempty"`
	TlsKey                            string                       `json:"tlsKey,omitempty"`
	UpdateChannel                     string                       `json:"update.channel,omitempty"`
	UseJaeger                         bool                         `json:"useJaeger,omitempty"`
}

// SlackNotificationsConfig Configuration for sending notifications to Slack.
type SlackNotificationsConfig struct {
	WebhookURL string `json:"webhookURL"`
}
