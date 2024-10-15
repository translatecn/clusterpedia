package config

import (
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	componentbaseconfig "k8s.io/component-base/config"

	crdclientset "xxxxx/pkg/generated/clientset/versioned"
	kubestatemetrics "xxxxx/pkg/kube_state_metrics"
	metricsserver "xxxxx/pkg/metrics/server"
	"xxxxx/pkg/storage"
	"xxxxx/pkg/synchromanager/clustersynchro"
)

type Config struct {
	Kubeconfig    *restclient.Config
	CRDClient     *crdclientset.Clientset
	EventRecorder record.EventRecorder

	WorkerNumber            int
	ShardingName            string
	MetricsServerConfig     metricsserver.Config
	KubeMetricsServerConfig *kubestatemetrics.ServerConfig
	StorageFactory          storage.StorageFactory
	ClusterSyncConfig       clustersynchro.ClusterSyncConfig

	LeaderElection   componentbaseconfig.LeaderElectionConfiguration
	ClientConnection componentbaseconfig.ClientConnectionConfiguration
}
