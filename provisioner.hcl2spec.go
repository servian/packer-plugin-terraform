package main

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	Version        *string           `mapstructure:"version" cty:"version"`
	CodePath       *string           `mapstructure:"code_path" cty:"code_path"`
	ExecuteCommand *string           `mapstructure:"run_command" cty:"run_command"`
	InstallCommand *string           `mapstructure:"install_command" cty:"install_command"`
	StagingDir     *string           `mapstructure:"staging_dir" cty:"staging_dir"`
	PreventSudo    *bool             `mapstructure:"prevent_sudo" cty:"prevent_sudo"`
	Variables      map[string]string `mapstructure:"variables" cty:"variables"`
	GuestOSType    *string           `mapstructure:"guest_os_type" cty:"guest_os_type"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":   &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type": &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},

		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"version":                    &hcldec.AttrSpec{Name: "version", Type: cty.String, Required: false},
		"code_path":                  &hcldec.AttrSpec{Name: "cookbook_path", Type: cty.List(cty.String), Required: false},
		"run_command":                &hcldec.AttrSpec{Name: "run_command", Type: cty.String, Required: false},
		"install_command":            &hcldec.AttrSpec{Name: "install_command", Type: cty.String, Required: false},
		"staging_dir":                &hcldec.AttrSpec{Name: "staging_dir", Type: cty.String, Required: false},
		"prevent_sudo":               &hcldec.AttrSpec{Name: "prevent_sudo", Type: cty.Bool, Required: false},
		"guest_os_type":              &hcldec.AttrSpec{Name: "guest_os_type", Type: cty.String, Required: false},
	}
	return s
}
