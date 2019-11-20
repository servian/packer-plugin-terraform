package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/provisioner"
	"github.com/hashicorp/packer/template/interpolate"
)

// Config struct containing variables
type Config struct {
	Version        string `mapstructure:"version"`
	CodePath       string `mapstructure:"code_path"`
	RunCommand     string `mapstructure:"run_command"`
	InstallCommand string `mapstructure:"install_command"`
	StagingDir     string `mapstructure:"staging_dir"`
	PreventSudo    bool   `mapstructure:"prevent_sudo"`

	ctx interpolate.Context
}

// Provisioner is the interface to install and run Terraform
type Provisioner struct {
	config        Config
	guestCommands *provisioner.GuestCommands
}

// RunTemplate for temp storage of interpolation vars
type RunTemplate struct {
	StagingDir string
	Sudo       bool
	Version    string
}

// Prepare parses the config and get everything ready
func (p *Provisioner) Prepare(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
	}, raws...)
	if err != nil {
		return err
	}

	if p.config.StagingDir == "" {
		p.config.StagingDir = "/tmp/packer-terraform"
	}

	if p.config.Version == "" {
		p.config.Version = "0.12.15"
	}

	if p.config.InstallCommand == "" {
		p.config.InstallCommand = "curl https://releases.hashicorp.com/terraform/{{.Version}}/terraform_{{.Version}}_linux_amd64.zip -so /tmp/terraform.zip \u0026\u0026 {{if .Sudo}}sudo {{end}}unzip -d /usr/local/bin/ /tmp/terraform.zip"
	}

	if p.config.RunCommand == "" {
		p.config.RunCommand = "cd {{.StagingDir}} \u0026\u0026 {{if .Sudo}}sudo {{end}}/usr/local/bin/terraform init \u0026\u0026 {{if .Sudo}}sudo {{end}}/usr/local/bin/terraform apply -auto-approve"
	}

	return nil
}

// Provision does the work of installing Terraform and running it on the remote
func (p *Provisioner) Provision(_ context.Context, ui packer.Ui, comm packer.Communicator) error {
	ui.Say("Provisioning with Terraform...")

	ui.Message("Uploading Code")
	if err := p.uploadDirectory(ui, comm, p.config.StagingDir, p.config.CodePath); err != nil {
		return fmt.Errorf("Error uploading code: %s", err)
	}

	ui.Message("Installing Terraform")
	p.config.ctx.Data = &RunTemplate{
		StagingDir: p.config.StagingDir,
		Version:    p.config.Version,
		Sudo:       !p.config.PreventSudo,
	}
	command, err := interpolate.Render(p.config.InstallCommand, &p.config.ctx)
	if err != nil {
		return err
	}
	if err := p.runCommand(ui, comm, command); err != nil {
		return fmt.Errorf("Error running Terraform: %s", err)
	}

	ui.Message("Running Terraform")
	command, err = interpolate.Render(p.config.RunCommand, &p.config.ctx)
	if err != nil {
		return err
	}
	if err := p.runCommand(ui, comm, command); err != nil {
		return fmt.Errorf("Error installing Terraform: %s", err)
	}

	return nil
}

func (p *Provisioner) runCommand(ui packer.Ui, comm packer.Communicator, command string) error {
	var out, outErr bytes.Buffer
	cmd := &packer.RemoteCmd{
		Command: command,
		Stdin:   nil,
		Stdout:  &out,
		Stderr:  &outErr,
	}

	ctx := context.TODO()
	if err := cmd.RunWithUi(ctx, comm, ui); err != nil {
		return err
	}
	if cmd.ExitStatus() != 0 {
		return fmt.Errorf("non-zero exit status")
	}
	return nil
}

func (p *Provisioner) uploadDirectory(ui packer.Ui, comm packer.Communicator, dst string, src string) error {
	// Make sure there is a trailing "/" so that the directory isn't
	// created on the other side.
	if src[len(src)-1] != '/' {
		src = src + "/"
	}

	return comm.UploadDir(dst, src, nil)
}
