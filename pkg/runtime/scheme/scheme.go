package scheme

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/kubernetes/pkg/api/legacyscheme"

	unstructuredscheme "xxxxx/pkg/runtime/scheme/unstructured"
)

var (
	LegacyResourceScheme         = legacyscheme.Scheme
	LegacyResourceCodecs         = legacyscheme.Codecs
	LegacyResourceParameterCodec = legacyscheme.ParameterCodec

	UnstructuredScheme = unstructuredscheme.NewScheme()
	UnstructuredCodecs = unstructured.UnstructuredJSONScheme
)
