package credential

import (
	"github.com/hashicorp/boundary/internal/iam"
	"github.com/hashicorp/boundary/internal/util/template"
)

// GetOpts - iterate the inbound Options and return a struct
func GetOpts(opt ...Option) (*options, error) {
	opts := getDefaultOptions()
	for _, o := range opt {
		if o == nil {
			continue
		}
		if err := o(opts); err != nil {
			return nil, err
		}
	}
	return opts, nil
}

// Option - how Options are passed as arguments.
type Option func(*options) error

// options = how options are represented
type options struct {
	WithIamRepoFn    iam.IamRepoFactory
	WithTemplateData template.Data
}

func getDefaultOptions() *options {
	return &options{}
}

// WithIamRepoFn provides a way to pass in a repo function to use in looking up
// templated data in credential library definitions, if needed
func WithIamRepoFn(with iam.IamRepoFactory) Option {
	return func(o *options) error {
		o.WithIamRepoFn = with
		return nil
	}
}

// WithTemplateData provides a way to pass in template information
func WithTemplateData(with template.Data) Option {
	return func(o *options) error {
		o.WithTemplateData = with
		return nil
	}
}
