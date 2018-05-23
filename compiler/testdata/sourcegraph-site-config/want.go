package p

import (
	"encoding/json"
	"errors"
	"fmt"
)

type AWSCodeCommitConnection struct {
	AccessKeyID                 string `json:"accessKeyID"`
	InitialRepositoryEnablement bool   `json:"initialRepositoryEnablement,omitempty"`
	Region                      string `json:"region"`
	RepositoryPathPattern       string `json:"repositoryPathPattern,omitempty"`
	SecretAccessKey             string `json:"secretAccessKey"`
}
type AuthProviders struct {
	Builtin       *BuiltinAuthProvider
	Saml          *SAMLAuthProvider
	Openidconnect *OpenIDConnectAuthProvider
	HttpHeader    *HTTPHeaderAuthProvider
}

func (v AuthProviders) MarshalJSON() ([]byte, error) {
	if v.Builtin != nil {
		return json.Marshal(v.Builtin)
	}
	if v.Saml != nil {
		return json.Marshal(v.Saml)
	}
	if v.Openidconnect != nil {
		return json.Marshal(v.Openidconnect)
	}
	if v.HttpHeader != nil {
		return json.Marshal(v.HttpHeader)
	}
	return nil, errors.New("tagged union type must have exactly 1 non-nil field value")
}
func (v *AuthProviders) UnmarshalJSON(data []byte) error {
	var d struct {
		DiscriminantProperty string `json:"type"`
	}
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	switch d.DiscriminantProperty {
	case "builtin":
		return json.Unmarshal(data, &v.Builtin)
	case "http-header":
		return json.Unmarshal(data, &v.HttpHeader)
	case "openidconnect":
		return json.Unmarshal(data, &v.Openidconnect)
	case "saml":
		return json.Unmarshal(data, &v.Saml)
	}
	return fmt.Errorf("tagged union type must have a %q property whose value is one of %s", "type", []string{"builtin", "saml", "openidconnect", "http-header"})
}

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

// BuiltinAuthProvider description: Configures the builtin username-password authentication provider.
type BuiltinAuthProvider struct {
	AllowSignup bool   `json:"allowSignup,omitempty"`
	Type        string `json:"type"`
}

// ExperimentalFeatures description: Experimental features to enable or disable. Features that are now enabled by default are marked as deprecated.
type ExperimentalFeatures struct {
	EnhancedSAML           string `json:"enhancedSAML,omitempty"`
	HostSurveysLocally     string `json:"hostSurveysLocally,omitempty"`
	JumpToDefOSSIndex      string `json:"jumpToDefOSSIndex,omitempty"`
	MultipleAuthProviders  string `json:"multipleAuthProviders,omitempty"`
	SearchTimeoutParameter string `json:"searchTimeoutParameter,omitempty"`
}
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
type GitLabConnection struct {
	Certificate                 string   `json:"certificate,omitempty"`
	GitURLType                  string   `json:"gitURLType,omitempty"`
	InitialRepositoryEnablement bool     `json:"initialRepositoryEnablement,omitempty"`
	ProjectQuery                []string `json:"projectQuery,omitempty"`
	RepositoryPathPattern       string   `json:"repositoryPathPattern,omitempty"`
	Token                       string   `json:"token"`
	Url                         string   `json:"url"`
}
type GitoliteConnection struct {
	Blacklist                  string `json:"blacklist,omitempty"`
	Host                       string `json:"host"`
	PhabricatorMetadataCommand string `json:"phabricatorMetadataCommand,omitempty"`
	Prefix                     string `json:"prefix"`
}

// HTTPHeaderAuthProvider description: Configures the HTTP header authentication provider (which authenticates users by consulting an HTTP request header set by an authentication proxy such as https://github.com/bitly/oauth2_proxy).
type HTTPHeaderAuthProvider struct {
	Type           string `json:"type"`
	UsernameHeader string `json:"usernameHeader"`
}
type Langservers struct {
	Address               string                 `json:"address,omitempty"`
	Disabled              bool                   `json:"disabled,omitempty"`
	InitializationOptions map[string]interface{} `json:"initializationOptions,omitempty"`
	Language              string                 `json:"language"`
	Metadata              *Metadata              `json:"metadata,omitempty"`
}
type Links struct {
	Blob       string `json:"blob,omitempty"`
	Commit     string `json:"commit,omitempty"`
	Repository string `json:"repository,omitempty"`
	Tree       string `json:"tree,omitempty"`
}

// Metadata description: Language server metadata. Used to populate various UI elements.
type Metadata struct {
	DocsURL      string `json:"docsURL,omitempty"`
	Experimental bool   `json:"experimental,omitempty"`
	HomepageURL  string `json:"homepageURL,omitempty"`
	IssuesURL    string `json:"issuesURL,omitempty"`
}

// OpenIDConnectAuthProvider description: Configures the OpenID Connect authentication provider for SSO.
type OpenIDConnectAuthProvider struct {
	ClientID           string `json:"clientID"`
	ClientSecret       string `json:"clientSecret"`
	Issuer             string `json:"issuer"`
	OverrideToken      string `json:"overrideToken,omitempty"`
	RequireEmailDomain string `json:"requireEmailDomain,omitempty"`
	Type               string `json:"type"`
}
type Phabricator struct {
	Repos []*Repos `json:"repos,omitempty"`
	Token string   `json:"token,omitempty"`
	Url   string   `json:"url,omitempty"`
}
type Repos struct {
	Callsign string `json:"callsign"`
	Path     string `json:"path"`
}
type Repository struct {
	Links *Links `json:"links,omitempty"`
	Path  string `json:"path"`
	Type  string `json:"type,omitempty"`
	Url   string `json:"url"`
}

// SAMLAuthProvider description: Configures the SAML authentication provider for SSO.
type SAMLAuthProvider struct {
	IdentityProviderMetadata    string `json:"identityProviderMetadata,omitempty"`
	IdentityProviderMetadataURL string `json:"identityProviderMetadataURL,omitempty"`
	ServiceProviderCertificate  string `json:"serviceProviderCertificate"`
	ServiceProviderPrivateKey   string `json:"serviceProviderPrivateKey"`
	Type                        string `json:"type"`
}

// SMTPServerConfig description: The SMTP server used to send transactional emails (such as email verifications, reset-password emails, and notifications).
type SMTPServerConfig struct {
	Authentication string `json:"authentication"`
	Domain         string `json:"domain,omitempty"`
	Host           string `json:"host"`
	Password       string `json:"password,omitempty"`
	Port           int    `json:"port"`
	Username       string `json:"username,omitempty"`
}
type SearchSavedQueries struct {
	Description    string `json:"description"`
	Key            string `json:"key"`
	Notify         bool   `json:"notify,omitempty"`
	NotifySlack    bool   `json:"notifySlack,omitempty"`
	Query          string `json:"query"`
	ShowOnHomepage bool   `json:"showOnHomepage,omitempty"`
}
type SearchScope struct {
	Description string `json:"description,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}

// Settings description: Configuration settings for users and organizations on Sourcegraph.
type Settings struct {
	Motd                   []string                  `json:"motd,omitempty"`
	NotificationsSlack     *SlackNotificationsConfig `json:"notifications.slack,omitempty"`
	SearchRepositoryGroups map[string][]string       `json:"search.repositoryGroups,omitempty"`
	SearchSavedQueries     []*SearchSavedQueries     `json:"search.savedQueries,omitempty"`
	SearchScopes           []*SearchScope            `json:"search.scopes,omitempty"`
}

// SiteConfiguration description: Configuration for a Sourcegraph site.
type SiteConfiguration struct {
	AppURL                            string                       `json:"appURL,omitempty"`
	AuthAllowSignup                   bool                         `json:"auth.allowSignup,omitempty"`
	AuthDisableAccessTokens           bool                         `json:"auth.disableAccessTokens,omitempty"`
	AuthOpenIDConnect                 *OpenIDConnectAuthProvider   `json:"auth.openIDConnect,omitempty"`
	AuthProvider                      string                       `json:"auth.provider,omitempty"`
	AuthProviders                     []AuthProviders              `json:"auth.providers,omitempty"`
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
	HttpStrictTransportSecurity       string                       `json:"httpStrictTransportSecurity,omitempty"`
	HttpToHttpsRedirect               interface{}                  `json:"httpToHttpsRedirect,omitempty"`
	Langservers                       []*Langservers               `json:"langservers,omitempty"`
	LightstepAccessToken              string                       `json:"lightstepAccessToken,omitempty"`
	LightstepProject                  string                       `json:"lightstepProject,omitempty"`
	MaxReposToSearch                  int                          `json:"maxReposToSearch,omitempty"`
	NoGoGetDomains                    string                       `json:"noGoGetDomains,omitempty"`
	OidcClientID                      string                       `json:"oidcClientID,omitempty"`
	OidcClientSecret                  string                       `json:"oidcClientSecret,omitempty"`
	OidcEmailDomain                   string                       `json:"oidcEmailDomain,omitempty"`
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
	TlsLetsencrypt                    string                       `json:"tls.letsencrypt,omitempty"`
	TlsCert                           string                       `json:"tlsCert,omitempty"`
	TlsKey                            string                       `json:"tlsKey,omitempty"`
	UpdateChannel                     string                       `json:"update.channel,omitempty"`
	UseJaeger                         bool                         `json:"useJaeger,omitempty"`
}

// SlackNotificationsConfig description: Configuration for sending notifications to Slack.
type SlackNotificationsConfig struct {
	WebhookURL string `json:"webhookURL"`
}
