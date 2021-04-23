CREATE TABLE `host` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(64) PRIMARY KEY NOT NULL,
  `name` varchar(255) UNIQUE NOT NULL,
  `ip` varchar(128) UNIQUE NOT NULL,
  `os` varchar(64) DEFAULT null,
  `os_version` varchar(64) DEFAULT null,
  `cpu_core` integer(64) DEFAULT null,
  `memory` integer(64) DEFAULT null,
  `has_gpu` tinyint(1) DEFAULT null,
  `gpu_num` integer(64) DEFAULT null,
  `gpu_info` varchar(128) DEFAULT null,
  `port` varchar(64) DEFAULT null,
  `status` varchar(64) DEFAULT null,
  `tags` varchar(255) DEFAULT null,
  `credential_id` varchar(64) DEFAULT null,
  `zone_id` varchar(255) DEFAULT null,
  `message` mediumtext,
  `architectures` varchar(255) DEFAULT null
);

CREATE TABLE `volume` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(64) PRIMARY KEY NOT NULL,
  `host_id` varchar(64) DEFAULT null,
  `size` integer(64) DEFAULT null,
  `name` integer(255) DEFAULT null
);

CREATE TABLE `credential` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(255) PRIMARY KEY NOT NULL,
  `name` varchar(255) UNIQUE NOT NULL,
  `username` varchar(255) DEFAULT null,
  `password` varchar(255) DEFAULT null,
  `private_key` mediumtext DEFAULT null,
  `type` varchar(255) DEFAULT null
);

CREATE TABLE `zone` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(64) PRIMARY KEY NOT NULL,
  `name` varchar(255) UNIQUE NOT NULL,
  `vars` mediumtext,
  `status` varchar(64) DEFAULT null,
  `region_id` varchar(64) DEFAULT null
);

CREATE TABLE `region` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(64) PRIMARY KEY NOT NULL,
  `name` varchar(255) UNIQUE NOT NULL,
  `datecenter` varchar(64) DEFAULT null,
  `provider` varchar(64) DEFAULT null,
  `vars` mediumtext
);

CREATE TABLE `cluster` (
  `id` varchar(255) PRIMARY KEY NOT NULL,
  `name` varchar(255) UNIQUE NOT NULL,
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `spec_id` varchar(255) DEFAULT null,
  `secrect_id` varchar(255) DEFAULT null,
  `status_id` varchar(255) DEFAULT null
);

CREATE TABLE `cluster_nodes` (
  `id` varchar(255) PRIMARY KEY NOT NULL,
  `name` varchar(255) DEFAULT null,
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `host_id` varchar(255),
  `cluster_id` varchar(255),
  `role` varchar(255) DEFAULT null,
  `status` varchar(255) DEFAULT null,
  `message` mediumtext
);

CREATE TABLE `cluster_secret` (
  `id` varchar(255) PRIMARY KEY NOT NULL,
  `kubeadm_token` mediumtext,
  `kubernetes_token` mediumtext
);

CREATE TABLE `cluster_spec` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(255) PRIMARY KEY NOT NULL,
  `version` varchar(255) DEFAULT null,
  `network_type` varchar(255) DEFAULT null,
  `runtime_type` varchar(255) DEFAULT null,
  `docker_storage_dir` varchar(255) DEFAULT null,
  `containerd_storage_dir` varchar(255) DEFAULT null,
  `lb_kube_apiserver_ip` varchar(255) DEFAULT null,
  `kube_api_server_port` varchar(255) DEFAULT null,
  `kube_pod_subnet` varchar(255) DEFAULT null,
  `kube_service_subnet` varchar(255) DEFAULT null,
  `worker_amount` integer(11) DEFAULT null,
  `kube_max_pods` integer(11) DEFAULT null,
  `kube_proxy_mode` varchar(255) DEFAULT null,
  `architectures` varchar(255) DEFAULT null,
  `support_gpu` varchar(255) DEFAULT null
);

CREATE TABLE `cluster_status` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(255) PRIMARY KEY NOT NULL,
  `message` mediumtext,
  `phase` varchar(255) DEFAULT null
);

CREATE TABLE `cluster_log` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(255) PRIMARY KEY NOT NULL,
  `cluster_id` varchar(255) DEFAULT null,
  `type` varchar(255) DEFAULT null,
  `message` mediumtext
);

CREATE TABLE `tenant` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(64) PRIMARY KEY NOT NULL,
  `name` varchar(255) UNIQUE NOT NULL,
  `description` varchar(128) DEFAULT null
);

CREATE TABLE `tenant_member` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(64) PRIMARY KEY NOT NULL,
  `tenant_id` varchar(64) DEFAULT null,
  `user_id` varchar(64) DEFAULT null,
  `role` varchar(64) DEFAULT null
);

CREATE TABLE `tenant_resource` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(64) PRIMARY KEY NOT NULL,
  `resource_type` varchar(128) DEFAULT null,
  `resource_id` varchar(64) DEFAULT null,
  `tenant_id` varchar(64) DEFAULT null
);

CREATE TABLE `user` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(64) PRIMARY KEY NOT NULL,
  `name` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255) DEFAULT null,
  `email` varchar(255) UNIQUE NOT NULL,
  `is_active` tinyint(1) DEFAULT 1,
  `language` varchar(64) DEFAULT null,
  `is_admin` tinyint(1) DEFAULT 0
);

CREATE TABLE `user_cluster_resource_limit` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(64) PRIMARY KEY NOT NULL,
  `user_id` varchar(64) DEFAULT null,
  `cluster_id` varchar(255) DEFAULT null,
  `max_cpu_core` integer(64) DEFAULT null,
  `max_memory` integer(64) DEFAULT null,
  `max_pod` integer(64) DEFAULT null
);

CREATE TABLE `system_setting` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(64) PRIMARY KEY NOT NULL,
  `key` varchar(255) UNIQUE NOT NULL,
  `value` varchar(255) UNIQUE NOT NULL,
  `tab` varchar(64) DEFAULT null
);

CREATE TABLE `apptemplate` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(64) PRIMARY KEY NOT NULL,
  `name` varchar(255) UNIQUE NOT NULL,
  `version` varchar(255) DEFAULT null,
  `describe` varchar(255) DEFAULT null,
  `status` varchar(255) DEFAULT null,
  `vars` mediumtext,
  `logo` varchar(255) DEFAULT null
);

CREATE TABLE `cluster_app` (
  `created_at` datetime DEFAULT null,
  `updated_at` datetime DEFAULT null,
  `id` varchar(64) PRIMARY KEY NOT NULL,
  `name` varchar(255) UNIQUE NOT NULL,
  `cluster_id` varchar(255) DEFAULT null,
  `tenant_id` varchar(255) DEFAULT null,
  `apptemplate_id` varchar(64) DEFAULT null,
  `status` varchar(255) DEFAULT null,
  `vars` mediumtext
);

ALTER TABLE `host` ADD FOREIGN KEY (`credential_id`) REFERENCES `credential` (`id`);

ALTER TABLE `volume` ADD FOREIGN KEY (`host_id`) REFERENCES `host` (`id`);

ALTER TABLE `zone` ADD FOREIGN KEY (`region_id`) REFERENCES `region` (`id`);

ALTER TABLE `cluster` ADD FOREIGN KEY (`spec_id`) REFERENCES `cluster_spec` (`id`);

ALTER TABLE `cluster` ADD FOREIGN KEY (`secrect_id`) REFERENCES `cluster_secret` (`id`);

ALTER TABLE `cluster` ADD FOREIGN KEY (`status_id`) REFERENCES `cluster_status` (`id`);

ALTER TABLE `cluster_nodes` ADD FOREIGN KEY (`host_id`) REFERENCES `host` (`id`);

ALTER TABLE `cluster_nodes` ADD FOREIGN KEY (`cluster_id`) REFERENCES `cluster` (`id`);

ALTER TABLE `cluster_log` ADD FOREIGN KEY (`cluster_id`) REFERENCES `cluster` (`id`);

ALTER TABLE `tenant_member` ADD FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`);

ALTER TABLE `tenant_member` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `tenant_resource` ADD FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`);

ALTER TABLE `user_cluster_resource_limit` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `user_cluster_resource_limit` ADD FOREIGN KEY (`cluster_id`) REFERENCES `cluster` (`id`);

ALTER TABLE `cluster_app` ADD FOREIGN KEY (`cluster_id`) REFERENCES `cluster` (`id`);

ALTER TABLE `cluster_app` ADD FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`);

ALTER TABLE `cluster_app` ADD FOREIGN KEY (`apptemplate_id`) REFERENCES `apptemplate` (`id`);
