package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
)

// Config struct containing variables
type Config struct {
	Version  string `mapstructure:"version"`
	CodePath string `mapstructure:"code_path"`
	Comment  string `mapstructure:"comment"`
	SendToUI bool   `mapstructure:"ui"`
}

// Provisioner is the interface to install and run Terraform
type Provisioner struct {
	config Config
}

// Prepare parses the config and get everything ready
func (p *Provisioner) Prepare(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate: true,
	}, raws...)
	if err != nil {
		return err
	}

	return nil
}

// Provision does the work of installing Terraform and running it on the remote
func (p *Provisioner) Provision(_ context.Context, ui packer.Ui, comm packer.Communicator) error {
	if p.config.SendToUI {
		ui.Say(p.config.Comment)
	}
	stagingDir := "/tmp/packer-terraform"
	if err := p.uploadDirectory(ui, comm, stagingDir, p.config.CodePath); err != nil {
		return fmt.Errorf("Error uploading code: %s", err)
	}

	command := "cd /tmp/packer-terraform && /usr/local/bin/terraform init && /usr/local/bin/terraform apply -auto-approve"
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

	cmd.Wait()
	if cmd.ExitStatus() != 0 {
		ui.Error(out.String())
		ui.Error(outErr.String())
		return errors.New("Error bootstrapping converge")
	}

	ui.Message(strings.TrimSpace(out.String()))

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
