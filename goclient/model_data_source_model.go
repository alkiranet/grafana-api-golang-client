/*
 * Grafana HTTP API.
 *
 * The Grafana backend exposes an HTTP API, the same API is used by the frontend to do everything from saving dashboards, creating users and updating data sources.
 *
 * API version: 0.0.1
 * Contact: hello@grafana.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package goclient

type DataSourceModel struct {
	Access            *DsAccessModel  `json:"access,omitempty"`
	AccessControl     *MetadataModel  `json:"accessControl,omitempty"`
	BasicAuth         bool            `json:"basicAuth,omitempty"`
	BasicAuthPassword string          `json:"basicAuthPassword,omitempty"`
	BasicAuthUser     string          `json:"basicAuthUser,omitempty"`
	Database          string          `json:"database,omitempty"`
	Id                int64           `json:"id,omitempty"`
	IsDefault         bool            `json:"isDefault,omitempty"`
	JsonData          *JsonModel      `json:"jsonData,omitempty"`
	Name              string          `json:"name,omitempty"`
	OrgId             int64           `json:"orgId,omitempty"`
	Password          string          `json:"password,omitempty"`
	ReadOnly          bool            `json:"readOnly,omitempty"`
	SecureJsonFields  map[string]bool `json:"secureJsonFields,omitempty"`
	Type_             string          `json:"type,omitempty"`
	TypeLogoUrl       string          `json:"typeLogoUrl,omitempty"`
	Uid               string          `json:"uid,omitempty"`
	Url               string          `json:"url,omitempty"`
	User              string          `json:"user,omitempty"`
	Version           int64           `json:"version,omitempty"`
	WithCredentials   bool            `json:"withCredentials,omitempty"`
}