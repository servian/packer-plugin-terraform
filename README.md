# packer-provisioner-terraform

* [![Build Status](https://www.travis-ci.com/servian/packer-plugin-terraform.svg?branch=main)](https://travis-ci.com/servian/packer-plugin-terraform)
* [![license MPL-2.0](https://img.shields.io/badge/license-MPL--2.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)
* [![GoReportCard](https://goreportcard.com/badge/github.com/servian/packer-plugin-terraform)](https://goreportcard.com/report/github.com/servian/packer-plugin-terraform)
* [![Version](http://img.shields.io/github/release/servian/packer-plugin-terraform/all.svg?style=flat)](https://github.com/Servian/packer-plugin-terraform/releases)

Inspired by Megan Marsh's talk https://www.hashicorp.com/resources/extending-packer
I bit the bullet and started making my own ill advised provisioner for Terraform.

## Usage

    packer {
      required_plugins {
        terraform = {
          version = "= 0.0.6"
          source = "github.com/servian/terraform"
        }
      }
    }

    source "docker" "test_server" {
      commit = true
      image  = "amazonlinux:2"
    }

    build {
      sources = ["source.docker.test_server"]

      provisioner "terraform" {
        code_path       = "./tfcode"
        prevent_sudo    = "true"
        variables = {
          consul_server_node = false
        }
      }
    }

## parameters

 * `version`(string) - the version of Terraform to install
 * `code_path`(string) - (required) the path to the terraform code
 * `run_command`(string) - override the command to run Terraform
 * `install_command`(string) - override the command to run Terraform
 * `staging_dir`(string) - override the remote path to stage the code.
 * `variables`(map(String, String)) - set terraform variables into a terraform.auto.tfvars file

## License

The code is available as open source under the terms of the [Mozilla Public License 2.0](https://opensource.org/licenses/MPL-2.0)

