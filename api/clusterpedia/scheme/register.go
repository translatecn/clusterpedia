package scheme

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"xxxxx/api/clusterpedia/install"
)

var scheme = runtime.NewScheme()

var Codecs = serializer.NewCodecFactory(scheme)

var ParameterCodec = runtime.NewParameterCodec(scheme)

func init() {
	install.Install(scheme)
}
