variable "consul_server_node" {
  description = "Enable server mode for Consul"
  default     = true
  type        = bool
}

variable "nomad_alt_url" {
  description = "URL to retrieve Nomad Binary from"
  default     = "https://releases.hashicorp.com/nomad/0.10.1/nomad_0.10.1_linux_amd64.zip"
  type        = string
}

resource "local_file" "consul_service" {
  sensitive_content = templatefile("${path.module}/service.tpl",
    {
      consul_server_node = var.consul_server_node
      nomad_alt_url      = var.nomad_alt_url
    }
  )
  filename        = "/etc/systemd/system/consul.service"
  file_permission = "0644"

  provisioner "local-exec" {
    command = "echo systemctl daemon-reload"
  }
}


