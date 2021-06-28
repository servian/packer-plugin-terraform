# Terraform Plugins

The Terraform Provisioner installs and runs HashiCorp Terraform on the remote instance allowing the use of modules and providers like [Local Provider](https://registry.terraform.io/providers/hashicorp/local/latest/docs).


## Installation

### Using pre-built releases

#### Using the `packer init` command

Starting from version 1.7, Packer supports a new `packer init` command allowing
automatic installation of Packer plugins. Read the
[Packer documentation](https://www.packer.io/docs/commands/init) for more information.

To install this plugin, copy and paste this code into your Packer configuration .
Then, run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    name = {
      version = ">= 0.0.7"
      source  = "github.com/servian/terraform"
    }
  }
}
```

#### Manual installation

You can find pre-built binary releases of the plugin [here](https://github.com/servian/packer-plugin-terraform/releases).
Once you have downloaded the latest archive corresponding to your target OS,
uncompress it to retrieve the plugin binary file corresponding to your platform.
To install the plugin, please follow the Packer documentation on
[installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).


#### From Source

If you prefer to build the plugin from its source code, clone the GitHub
repository locally and run the command `go build` from the root
directory. Upon successful compilation, a `packer-plugin-terraform` plugin
binary file can be found in the root directory.
To install the compiled plugin, please follow the official Packer documentation
on [installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).


## Plugin Contents

The Terraform plugin currently contains just a single provisioner:

### Provisioners

- [provisioner](/docs/provisioners/provisioner-terraform.mdx) - The Terraform provisioner is used to provisioner
  Packer builds.

