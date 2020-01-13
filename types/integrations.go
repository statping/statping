package types

type Integration struct {
	ShortName   string              `json:"name"`
	Name        string              `json:"full_name"`
	Icon        string              `json:"-"`
	Description string              `json:"description"`
	Enabled     bool                `json:"enabled"`
	Fields      []*IntegrationField `json:"fields"`
}

type IntegrationField struct {
	Name        string      `json:"name"`
	Value       interface{} `json:"value"`
	Type        string      `json:"type"`
	Description string      `json:"description,omitempty"`
	MimeType    string      `json:"mime_type,omitempty"`
}

type Integrator interface {
	Get() *Integration
	List() ([]*Service, error)
}
