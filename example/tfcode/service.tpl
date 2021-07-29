[Unit]
Description="HashiCorp Consul - A service mesh solution - ${consul_server_node}"
Documentation=${nomad_alt_url}
Requires=network-online.target
After=network-online.target
ConditionFileNotEmpty=/opt/consul/conf/consul.hcl

[Service]
User=consul
Group=consul
Capabilities=CAP_NET_BIND_SERVICE
EnvironmentFile=-/etc/sysconfig/consul
ExecStart=/opt/consul/bin/consul agent -config-dir=/opt/consul/conf
ExecReload=/opt/consul/bin/consul reload
KillMode=process
Restart=on-failure
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
