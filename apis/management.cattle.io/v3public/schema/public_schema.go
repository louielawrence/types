package schema

import (
	"net/http"

	"github.com/rancher/norman/types"
	"github.com/rancher/types/apis/management.cattle.io/v3"
	"github.com/rancher/types/apis/management.cattle.io/v3public"
	"github.com/rancher/types/factory"
)

var (
	PublicVersion = types.APIVersion{
		Version: "v3public",
		Group:   "management.cattle.io",
		Path:    "/v3-public",
	}

	PublicSchemas = factory.Schemas(&PublicVersion).
			Init(authProvidersTypes)
)

func authProvidersTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImportAndCustomize(&PublicVersion, v3.Token{}, func(schema *types.Schema) {
			// No collection methods causes the store to not create a CRD for it
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{}
		}).
		MustImportAndCustomize(&PublicVersion, v3public.AuthProvider{}, func(schema *types.Schema) {
			schema.CollectionMethods = []string{http.MethodGet}
		}).
		// Local provider
		MustImportAndCustomize(&PublicVersion, v3public.LocalProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "basicLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImport(&PublicVersion, v3public.BasicLogin{}).
		// Github provider
		MustImportAndCustomize(&PublicVersion, v3public.GithubProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "githubLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImport(&PublicVersion, v3public.GithubLogin{}).
		// Phabricator provider
		MustImportAndCustomize(&PublicVersion, v3public.PhabricatorProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "phabricatorLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImport(&PublicVersion, v3public.PhabricatorLogin{}).
		// Active Directory provider
		MustImportAndCustomize(&PublicVersion, v3public.ActiveDirectoryProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "basicLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		// Azure AD provider
		MustImportAndCustomize(&PublicVersion, v3public.AzureADProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "azureADLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImport(&PublicVersion, v3public.AzureADLogin{}).
		// Saml provider
		MustImportAndCustomize(&PublicVersion, v3public.PingProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "samlLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImport(&PublicVersion, v3public.SamlLogin{}).
		// OpenLdap provider
		MustImportAndCustomize(&PublicVersion, v3public.OpenLdapProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "basicLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		// FreeIpa provider
		MustImportAndCustomize(&PublicVersion, v3public.FreeIpaProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "basicLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		})
}
