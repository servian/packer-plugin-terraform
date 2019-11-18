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
	Version    string `mapstructure:"version"`
	CodePath   string `mapstructure:"code_path"`
	RunCommand string `mapstructure:"run_command"`
	StagingDir string `mapstructure:"staging_dir"`

	ctx interpolate.Context
}

// Provisioner is the interface to install and run Terraform
type Provisioner struct {
	config        Config
	guestCommands *provisioner.GuestCommands
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

	if p.config.RunCommand == "" {
		p.config.RunCommand = "cd {{.StagingDir}} \u0026\u0026 /usr/local/bin/terraform init \u0026\u0026 /usr/local/bin/terraform apply -auto-approve"
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

	ui.Message("Running Terraform")
	p.config.ctx.Data = &Config{
		StagingDir: p.config.StagingDir,
	}
	command, err := interpolate.Render(p.config.RunCommand, &p.config.ctx)
	if err != nil {
		return err
	}

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
