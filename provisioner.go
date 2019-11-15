package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
)

type Config struct {
	Version  string `mapstructure:"version"`
	CodePath string `mapstructure:"code_path"`
	Comment  string `mapstructure:"comment"`
	SendToUi bool   `mapstructure:"ui"`
}

type Provisioner struct {
	config Config
}

func (p *Provisioner) Prepare(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate: true,
	}, raws...)
	if err != nil {
		return err
	}

	return nil
}

func (p *Provisioner) Provision(_ context.Context, ui packer.Ui, comm packer.Communicator) error {
	if p.config.SendToUi {
		ui.Say(p.config.Comment)
	}
	stagingDir := "/tmp/packer-terraform"
	if err := p.uploadDirectory(ui, comm, stagingDir, p.config.CodePath); err != nil {
		return fmt.Errorf("Error uploading code: %s", err)
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
