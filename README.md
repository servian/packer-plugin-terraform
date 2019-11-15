# packer-provisioner-terraform

Inspired by Megan Marsh's talk https://www.hashicorp.com/resources/extending-packer
I bit the bullet and started making my own ill advised provisioner for Terraform.

## Usage

    "provisioners": [
      {
        "type": "terraform",
        "version": "0.12.14",
        "code_path": "./terraform"
      }
    ]
