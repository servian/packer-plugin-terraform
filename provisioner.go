package main

import (
	"context"

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

func (p *Provisioner) Provision(_ context.Context, ui packer.Ui, _ packer.Communicator) error {
	if p.config.SendToUi {
		ui.Say(p.config.Comment)

	}

	return nil
}
