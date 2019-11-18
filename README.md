# packer-provisioner-terraform

Inspired by Megan Marsh's talk https://www.hashicorp.com/resources/extending-packer
I bit the bullet and started making my own ill advised provisioner for Terraform.

## Usage

    "provisioners": [
      {
        "type": "terraform",
        "code_path": "./tfcode"
      }
    ]

## parameters

 * `version`(string) - the version of Terraform to install
 * `code_path`(string) - (required) the path to the terraform code
 * `run_command`(string) - override the command to run Terraform
 * `install_command`(string) - override the command to run Terraform
 * `staging_dir`(string) - override the remote path to stage the code.  
