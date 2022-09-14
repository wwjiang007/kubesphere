/*
Copyright 2020 KubeSphere Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package monitoring

const (
	KubeSphereWorkspaceCount = "kubesphere_workspace_count"
	KubeSphereUserCount      = "kubesphere_user_count"
	KubeSphereClusterCount   = "kubesphere_cluser_count"
	KubeSphereAppTmplCount   = "kubesphere_app_template_count"

	WorkspaceNamespaceCount = "workspace_namespace_count"
	WorkspaceDevopsCount    = "workspace_devops_project_count"
	WorkspaceMemberCount    = "workspace_member_count"
	WorkspaceRoleCount      = "workspace_role_count"
)

var ClusterMetrics = []string{
	// multi-cluster
	"clusters_cpu_utilisation",
	"clusters_cpu_usage",
	"clusters_cpu_total",
	"clusters_memory_utilisation",
	"clusters_memory_available",
	"clusters_memory_total",
	"clusters_memory_usage_wo_cache",
	"clusters_disk_size_usage",
	"clusters_disk_size_utilisation",
	"clusters_disk_size_capacity",
	"clusters_disk_size_available",
	"clusters_pod_cpu_usage",
	"clusters_pod_cpu_requests_total",
	"clusters_pod_cpu_limits_total",
	"clusters_pod_memory_requests_total",
	"clusters_pod_memory_limits_total",
	"clusters_pod_count",
	"clusters_pod_quota",
	"clusters_pod_utilisation",
	"clusters_pod_running_count",
	"clusters_pod_succeeded_count",
	"clusters_pod_pending_count",
	"clusters_pod_failed_count",
	"clusters_pod_unknown_count",
	"clusters_pod_oomkilled_count",
	"clusters_pod_evicted_count",
	"clusters_pod_qos_guaranteed_count",
	"clusters_pod_qos_burstable_count",
	"clusters_pod_qos_besteffort_count",
	"clusters_namespace_count",
	"clusters_count",
	"clusters_node_count",
	"clusters_cronjob_count",
	"clusters_pvc_count",
	"clusters_daemonset_count",
	"clusters_deployment_count",
	"clusters_endpoint_count",
	"clusters_hpa_count",
	"clusters_job_count",
	"clusters_statefulset_count",
	"clusters_replicaset_count",
	"clusters_service_count",
	"clusters_secret_count",
	"clusters_pv_count",
	"clusters_ingresses_count",

	// cluster
	"cluster_cpu_utilisation",
	"cluster_cpu_usage",
	"cluster_cpu_total",
	"cluster_cpu_non_master_total",
	"cluster_memory_utilisation",
	"cluster_memory_available",
	"cluster_memory_total",
	"cluster_memory_non_master_total",
	"cluster_memory_usage_wo_cache",
	"cluster_net_utilisation",
	"cluster_net_bytes_transmitted",
	"cluster_net_bytes_received",
	"cluster_disk_read_iops",
	"cluster_disk_write_iops",
	"cluster_disk_read_throughput",
	"cluster_disk_write_throughput",
	"cluster_disk_size_usage",
	"cluster_disk_size_utilisation",
	"cluster_disk_size_capacity",
	"cluster_disk_size_available",
	"cluster_disk_inode_total",
	"cluster_disk_inode_usage",
	"cluster_disk_inode_utilisation",
	"cluster_namespace_count",
	"cluster_pod_count",
	"cluster_pod_quota",
	"cluster_pod_utilisation",
	"cluster_pod_running_count",
	"cluster_pod_succeeded_count",
	"cluster_pod_pending_count",
	"cluster_pod_failed_count",
	"cluster_pod_unknown_count",
	"cluster_pod_oomkilled_count",
	"cluster_pod_evicted_count",
	"cluster_pod_qos_guaranteed_count",
	"cluster_pod_qos_burstable_count",
	"cluster_pod_qos_besteffort_count",
	"cluster_pod_cpu_usage",
	"cluster_pod_cpu_non_master_usage",
	"cluster_pod_cpu_requests_total",
	"cluster_pod_cpu_requests_non_master_total",
	"cluster_pod_cpu_limits_total",
	"cluster_pod_cpu_limits_non_master_total",
	"cluster_pod_memory_usage_wo_cache",
	"cluster_pod_memory_non_master_usage_wo_cache",
	"cluster_pod_memory_requests_total",
	"cluster_pod_memory_requests_non_master_total",
	"cluster_pod_memory_limits_total",
	"cluster_pod_memory_limits_non_master_total",
	"cluster_namespace_quota_cpu_usage",
	"cluster_namespace_quota_cpu_requests_hard_total",
	"cluster_namespace_quota_cpu_limits_hard_total",
	"cluster_namespace_quota_memory_usage",
	"cluster_namespace_quota_memory_usage_wo_cache",
	"cluster_namespace_quota_memory_requests_hard_total",
	"cluster_namespace_quota_memory_limits_hard_total",
	"cluster_pod_abnormal_count",
	"cluster_node_online",
	"cluster_node_offline",
	"cluster_node_total",
	"cluster_cronjob_count",
	"cluster_pvc_count",
	"cluster_daemonset_count",
	"cluster_deployment_count",
	"cluster_endpoint_count",
	"cluster_hpa_count",
	"cluster_job_count",
	"cluster_statefulset_count",
	"cluster_replicaset_count",
	"cluster_service_count",
	"cluster_secret_count",
	"cluster_pv_count",
	"cluster_ingresses_extensions_count",
	"cluster_load1",
	"cluster_load5",
	"cluster_load15",
	"cluster_pod_abnormal_ratio",
	"cluster_node_offline_ratio",

	// gpu
	"cluster_gpu_utilization",
	"cluster_gpu_usage",
	"cluster_gpu_total",
	"cluster_gpu_memory_utilization",
	"cluster_gpu_memory_usage",
	"cluster_gpu_memory_available",
	"cluster_gpu_memory_total",
}

var NodeMetrics = []string{
	"node_cpu_utilisation",
	"node_cpu_total",
	"node_cpu_usage",
	"node_memory_utilisation",
	"node_memory_usage_wo_cache",
	"node_memory_available",
	"node_memory_total",
	"node_net_utilisation",
	"node_net_bytes_transmitted",
	"node_net_bytes_received",
	"node_disk_read_iops",
	"node_disk_write_iops",
	"node_disk_read_throughput",
	"node_disk_write_throughput",
	"node_disk_size_capacity",
	"node_disk_size_available",
	"node_disk_size_usage",
	"node_disk_size_utilisation",
	"node_disk_inode_total",
	"node_disk_inode_usage",
	"node_disk_inode_utilisation",
	"node_pod_count",
	"node_pod_quota",
	"node_pod_utilisation",
	"node_pod_running_count",
	"node_pod_succeeded_count",
	"node_pod_abnormal_count",
	"node_load1",
	"node_load5",
	"node_load15",
	"node_pod_abnormal_ratio",
	"node_pleg_quantile",

	"node_device_size_usage",
	"node_device_size_utilisation",

	// gpu
	"node_gpu_utilization",
	"node_gpu_usage",
	"node_gpu_total",
	"node_gpu_memory_utilization",
	"node_gpu_memory_usage",
	"node_gpu_memory_available",
	"node_gpu_memory_total",
	"node_gpu_temp",
	"node_gpu_power_usage",
}

var WorkspaceMetrics = []string{
	"workspace_cpu_usage",
	"workspace_memory_usage",
	"workspace_memory_usage_wo_cache",
	"workspace_net_bytes_transmitted",
	"workspace_net_bytes_received",
	"workspace_pod_count",
	"workspace_pod_running_count",
	"workspace_pod_succeeded_count",
	"workspace_pod_abnormal_count",
	"workspace_ingresses_extensions_count",
	"workspace_cronjob_count",
	"workspace_pvc_count",
	"workspace_daemonset_count",
	"workspace_deployment_count",
	"workspace_endpoint_count",
	"workspace_hpa_count",
	"workspace_job_count",
	"workspace_statefulset_count",
	"workspace_replicaset_count",
	"workspace_service_count",
	"workspace_secret_count",
	"workspace_pod_abnormal_ratio",

	// gpu
	"workspace_gpu_usage",
	"workspace_gpu_memory_usage",
}

var NamespaceMetrics = []string{
	"namespace_cpu_usage",
	"namespace_cpu_used_requests_utilisation",
	"namespace_cpu_used_limits_utilisation",
	"namespace_memory_usage",
	"namespace_memory_usage_wo_cache",
	"namespace_memory_used_requests_utilisation",
	"namespace_memory_used_limits_utilisation",
	"namespace_pvc_bytes_used",
	"namespace_pvc_count",
	"namespace_net_bytes_transmitted",
	"namespace_net_bytes_received",
	"namespace_pod_count",
	"namespace_pod_running_count",
	"namespace_pod_succeeded_count",
	"namespace_pod_abnormal_count",
	"namespace_pod_abnormal_ratio",
	"namespace_memory_limit_hard",
	"namespace_cpu_limit_hard",
	"namespace_pod_count_hard",
	"namespace_cronjob_count",
	"namespace_pvc_count",
	"namespace_daemonset_count",
	"namespace_deployment_count",
	"namespace_endpoint_count",
	"namespace_hpa_count",
	"namespace_job_count",
	"namespace_statefulset_count",
	"namespace_replicaset_count",
	"namespace_service_count",
	"namespace_secret_count",
	"namespace_configmap_count",
	"namespace_ingresses_extensions_count",
	"namespace_s2ibuilder_count",

	// gpu
	"namespace_gpu_limit_hard",
	"namespace_gpu_usage",
	"namespace_gpu_memory_usage",
}

var WorkloadMetrics = []string{
	"workload_cpu_usage",
	"workload_memory_usage",
	"workload_memory_usage_wo_cache",
	"workload_net_bytes_transmitted",
	"workload_net_bytes_received",
	"workload_deployment_replica",
	"workload_deployment_replica_available",
	"workload_statefulset_replica",
	"workload_statefulset_replica_available",
	"workload_daemonset_replica",
	"workload_daemonset_replica_available",
	"workload_deployment_unavailable_replicas_ratio",
	"workload_daemonset_unavailable_replicas_ratio",
	"workload_statefulset_unavailable_replicas_ratio",

	// gpu
	"workload_gpu_usage",
	"workload_gpu_memory_usage",
}

var PodMetrics = []string{
	"pod_cpu_usage",
	"pod_cpu_used_requests_utilisation",
	"pod_cpu_used_limits_utilisation",
	"pod_memory_usage",
	"pod_memory_usage_wo_cache",
	"pod_memory_used_requests_utilisation",
	"pod_memory_used_limits_utilisation",
	"pod_net_bytes_transmitted",
	"pod_net_bytes_received",
	"pod_pvc_bytes_used",
	"pod_pvc_bytes_utilisation",

	// gpu
	"pod_gpu_usage",
	"pod_gpu_memory_usage",
}

var ContainerMetrics = []string{
	"container_cpu_usage",
	"container_memory_usage",
	"container_memory_usage_wo_cache",
	// gpu
	"container_gpu_usage",
	"container_gpu_memory_usage",
	"container_processes_usage",
	"container_threads_usage",
}

var PVCMetrics = []string{
	"pvc_inodes_available",
	"pvc_inodes_used",
	"pvc_inodes_total",
	"pvc_inodes_utilisation",
	"pvc_bytes_available",
	"pvc_bytes_used",
	"pvc_bytes_total",
	"pvc_bytes_utilisation",
}

var IngressMetrics = []string{
	"ingress_request_count",
	"ingress_request_5xx_count",
	"ingress_request_4xx_count",
	"ingress_active_connections",
	"ingress_success_rate",
	"ingress_request_duration_average",
	"ingress_request_duration_50percentage",
	"ingress_request_duration_95percentage",
	"ingress_request_duration_99percentage",
	"ingress_request_volume",
	"ingress_request_volume_by_ingress",
	"ingress_request_network_sent",
	"ingress_request_network_received",
	"ingress_request_memory_bytes",
	"ingress_request_cpu_usage",
}

var EtcdMetrics = []string{
	"etcd_server_list",
	"etcd_server_total",
	"etcd_server_up_total",
	"etcd_server_has_leader",
	"etcd_server_is_leader",
	"etcd_server_leader_changes",
	"etcd_server_proposals_failed_rate",
	"etcd_server_proposals_applied_rate",
	"etcd_server_proposals_committed_rate",
	"etcd_server_proposals_pending_count",
	"etcd_mvcc_db_size",
	"etcd_network_client_grpc_received_bytes",
	"etcd_network_client_grpc_sent_bytes",
	"etcd_grpc_call_rate",
	"etcd_grpc_call_failed_rate",
	"etcd_grpc_server_msg_received_rate",
	"etcd_grpc_server_msg_sent_rate",
	"etcd_disk_wal_fsync_duration",
	"etcd_disk_wal_fsync_duration_quantile",
	"etcd_disk_backend_commit_duration",
	"etcd_disk_backend_commit_duration_quantile",
}

var APIServerMetrics = []string{
	"apiserver_up_sum",
	"apiserver_request_rate",
	"apiserver_request_by_verb_rate",
	"apiserver_request_latencies",
	"apiserver_request_by_verb_latencies",
}

var SchedulerMetrics = []string{
	"scheduler_up_sum",
	"scheduler_schedule_attempts",
	"scheduler_schedule_attempt_rate",
	"scheduler_e2e_scheduling_latency",
	"scheduler_e2e_scheduling_latency_quantile",
}
