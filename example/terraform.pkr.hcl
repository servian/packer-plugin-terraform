packer {
  required_plugins {
    terraform = {
      version = ">= 0.0.7"
      source  = "github.com/servian/terraform"
    }
  }
}

source "docker" "amazon" {
  commit = true
  image  = "amazonlinux:2"
}

build {
  sources = ["source.docker.amazon"]

  provisioner "shell" {
    inline = [
      "yum install -y unzip"
    ]
  }

  provisioner "terraform" {
    code_path       = "./tfcode"
    prevent_sudo    = "true"
    variable_string = jsonencode({
        consul_server_node = false
        nomad_alt_url = "https://example.com"
    })
  }

  post-processor "docker-tag" {
    repository = "tristanmorgan/packer-tf-test"
    tags       = ["latest"]
  }
}
