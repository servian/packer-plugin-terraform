packer {
  required_plugins {
    terraform = {
      version = "= 0.0.7"
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
    variable_string = jsonencode({
        consul_server_node = false
        nomad_alt_url = "https://example.com"
    })
    version = "1.0.0"
  }
}
