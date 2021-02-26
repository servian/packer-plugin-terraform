packer {
  required_plugins {
    terraform = {
      version = "= 0.0.6"
      source = "github.com/servian/terraform"
    }
  }
}

source "null" "test_server" {
  communicator = "none"
}

build {
  sources = ["source.null.test_server"]

  provisioner "terraform" {
    code_path       = "./tfcode"
    prevent_sudo    = "true"
  }
}
