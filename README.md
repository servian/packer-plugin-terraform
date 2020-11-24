# packer-provisioner-terraform

* [![Build Status](https://travis-ci.org/servian/packer-provisioner-terraform.svg?branch=main)](https://travis-ci.org/servian/packer-provisioner-terraform)
* [![license MPL-2.0](https://img.shields.io/badge/license-MPL--2.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)
* [![GoReportCard](https://goreportcard.com/badge/github.com/servian/packer-provisioner-terraform)](https://goreportcard.com/report/github.com/servian/packer-provisioner-terraform)
* [![Version](http://img.shields.io/github/release/servian/packer-provisioner-terraform/all.svg?style=flat)](https://github.com/Servian/packer-provisioner-terraform/releases)

Inspired by Megan Marsh's talk https://www.hashicorp.com/resources/extending-packer
I bit the bullet and started making my own ill advised provisioner for Terraform.

## Installation

Install the binary (you'll need `git` and `go`):

    $ go get github.com/servian/packer-provisioner-terraform

Copy the plugin into packer.d directory:

    $ mkdir $HOME/.packer.d/plugins
    $ cp $GOPATH/bin/packer-provisioner-terraform $HOME/.packer.d/plugins

## Usage

    "provisioners": [
      {
        "inline": [
          "yum install -y unzip curl"
        ],
        "type": "shell"
      },
      {
        "code_path": "./tfcode",
        "prevent_sudo": "true",
        "type": "terraform",
        "variables": {
          "consul_server_node": "false",
          "vault_alt_url": "https://example.com"
        },
        "version": "0.12.15"
      }
    ]

## parameters

 * `version`(string) - the version of Terraform to install
 * `code_path`(string) - (required) the path to the terraform code
 * `run_command`(string) - override the command to run Terraform
 * `install_command`(string) - override the command to run Terraform
 * `staging_dir`(string) - override the remote path to stage the code.
 * `variables`(map(String, String)) - set terraform variables into a terraform.auto.tfvars file

## License

The code is available as open source under the terms of the [Mozilla Public License 2.0](https://opensource.org/licenses/MPL-2.0)

