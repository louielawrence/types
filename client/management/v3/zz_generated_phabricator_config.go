package client

const (
	PhabricatorConfigType                     = "phabricatorConfig"
	PhabricatorConfigFieldAccessMode          = "accessMode"
	PhabricatorConfigFieldAllowedPrincipalIDs = "allowedPrincipalIds"
	PhabricatorConfigFieldAnnotations         = "annotations"
	PhabricatorConfigFieldClientID            = "clientId"
	PhabricatorConfigFieldClientSecret        = "clientSecret"
	PhabricatorConfigFieldCreated             = "created"
	PhabricatorConfigFieldCreatorID           = "creatorId"
	PhabricatorConfigFieldEnabled             = "enabled"
	PhabricatorConfigFieldHostname            = "hostname"
	PhabricatorConfigFieldLabels              = "labels"
	PhabricatorConfigFieldName                = "name"
	PhabricatorConfigFieldOwnerReferences     = "ownerReferences"
	PhabricatorConfigFieldRemoved             = "removed"
	PhabricatorConfigFieldTLS                 = "tls"
	PhabricatorConfigFieldType                = "type"
	PhabricatorConfigFieldUuid                = "uuid"
)

type PhabricatorConfig struct {
	AccessMode          string            `json:"accessMode,omitempty" yaml:"accessMode,omitempty"`
	AllowedPrincipalIDs []string          `json:"allowedPrincipalIds,omitempty" yaml:"allowedPrincipalIds,omitempty"`
	Annotations         map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClientID            string            `json:"clientId,omitempty" yaml:"clientId,omitempty"`
	ClientSecret        string            `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
	Created             string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID           string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Enabled             bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Hostname            string            `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	Labels              map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences     []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed             string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	TLS                 bool              `json:"tls,omitempty" yaml:"tls,omitempty"`
	Type                string            `json:"type,omitempty" yaml:"type,omitempty"`
	Uuid                string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
