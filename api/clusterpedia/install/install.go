package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	internal "xxxxx/api/clusterpedia"
	"xxxxx/api/clusterpedia/v1beta1"
)

func Install(scheme *runtime.Scheme) {
	utilruntime.Must(internal.Install(scheme))
	utilruntime.Must(v1beta1.Install(scheme))
	utilruntime.Must(scheme.SetVersionPriority(v1beta1.SchemeGroupVersion))
}
