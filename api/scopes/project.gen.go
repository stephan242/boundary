// Code generated by "make api"; DO NOT EDIT.
package scopes

import (
	"encoding/json"
	"time"

	"github.com/fatih/structs"

	"github.com/hashicorp/watchtower/api"
	"github.com/hashicorp/watchtower/api/internal/strutil"
)

type Project struct {
	Client *api.Client `json:"-"`

	defaultFields []string

	// The ID of the Project
	// Output only.
	Id string `json:"id,omitempty"`
	// Optional name for identification purposes
	Name *string `json:"name,omitempty"`
	// Optional user-set descripton for identification purposes
	Description *string `json:"description,omitempty"`
	// The time this resource was created
	// Output only.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// The time this resource was last updated.
	// Output only.
	UpdatedTime time.Time `json:"updated_time,omitempty"`
	// Whether the resource is disabled
	Disabled *bool `json:"disabled,omitempty"`
}

func (s *Project) SetDefault(key string) {
	s.defaultFields = strutil.AppendIfMissing(s.defaultFields, key)
}

func (s *Project) UnsetDefault(key string) {
	s.defaultFields = strutil.StrListDelete(s.defaultFields, key)
}

func (s Project) MarshalJSON() ([]byte, error) {
	m := structs.Map(s)
	if m == nil {
		m = make(map[string]interface{})
	}
	for _, k := range s.defaultFields {
		m[k] = nil
	}
	return json.Marshal(m)
}
