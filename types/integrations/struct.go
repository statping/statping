package integrations

import "github.com/hunterlong/statping/types/services"

type Integration struct {
	ShortName   string              `gorm:"column:name" json:"name"`
	Name        string              `gorm:"-" json:"full_name,omitempty"`
	Icon        string              `gorm:"-" json:"-"`
	Description string              `gorm:"-" json:"description,omitempty"`
	Enabled     bool                `gorm:"column:enabled;default:false" json:"enabled"`
	Fields      []*IntegrationField `gorm:"column:fields" json:"fields"`
}

type IntegrationField struct {
	Name        string      `gorm:"-" json:"name"`
	Value       interface{} `gorm:"-" json:"value"`
	Type        string      `gorm:"-" json:"type"`
	Description string      `gorm:"-" json:"description,omitempty"`
	MimeType    string      `gorm:"-" json:"mime_type,omitempty"`
}

type Integrator interface {
	Get() *Integration
	List() ([]*services.Service, error)
}
