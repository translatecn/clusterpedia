ip=`ifconfig|grep en0 -A 4|grep broadcast|awk -F ' ' '{print $2}'`

cat > /tmp/subnet-rbac.yaml << EOF
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: clusterpedia-synchro
rules:
- apiGroups:
  - '*'
  resources:
  - '*'
  verbs:
  - '*'
- nonResourceURLs:
  - '*'
  verbs:
  - '*'
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: clusterpedia-synchro
  namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: clusterpedia-synchro
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: clusterpedia-synchro
subjects:
- kind: ServiceAccount
  name: clusterpedia-synchro
  namespace: default
---
EOF

for i in $(seq 0 2)
do


cat > /tmp/clusterpedia-$i.yaml << EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
nodes:
  - role: control-plane
    extraPortMappings:
      - containerPort: 6443
        hostPort: 1200$i
        protocol: TCP
kubeadmConfigPatches:
  - |
    kind: ClusterConfiguration
    apiServer:
      certSANs:
        - localhost
        - 127.0.0.1
        - $ip
EOF

kind delete clusters clusterpedia-$i

kind create cluster --config /tmp/clusterpedia-$i.yaml -n clusterpedia-$i --kubeconfig ~/.kube/clusterpedia-$i --image registry.cn-hangzhou.aliyuncs.com/acejilam/node:v1.25.0
kubectl --kubeconfig ~/.kube/clusterpedia-$i wait nodes --all --for=condition=Ready --timeout=5m

yq e ".clusters[0].cluster.server = \"https://$ip:1200$i\"" -i ~/.kube/clusterpedia-$i

kubectl --kubeconfig ~/.kube/clusterpedia-$i apply -f /tmp/subnet-rbac.yaml

if [ "$i" -eq 0 ];then
  bash +x update.sh
  cp ~/.kube/clusterpedia-$i ~/.kube/clusterpedia-$i-search
  yq e ".clusters[0].cluster.server = \"https://$ip:1200$i/apis/clusterpedia.io/v1beta1/resources\"" -i ~/.kube/clusterpedia-$i-search
  kubectl --kubeconfig ~/.kube/clusterpedia-$i apply -f ./deploy/
  kubectl --kubeconfig ~/.kube/clusterpedia-$i apply -f ./deploy/
fi

cat > /tmp/cluster-secret-$i << EOF
apiVersion: v1
kind: Secret
metadata:
  name: cluster-secret-$i
  namespace: default
  annotations:
    kubernetes.io/service-account.name: clusterpedia-synchro
type: kubernetes.io/service-account-token

EOF
kubectl --kubeconfig ~/.kube/clusterpedia-$i apply -f /tmp/cluster-secret-$i

# Get CA and Token for Service Account
SYNCHRO_CA=$(kubectl --kubeconfig ~/.kube/clusterpedia-$i -n default get secret cluster-secret-$i -o jsonpath='{.data.ca\.crt}')
SYNCHRO_TOKEN=$(kubectl --kubeconfig ~/.kube/clusterpedia-$i -n default get secret cluster-secret-$i -o jsonpath='{.data.token}')

cat > /tmp/subnet-import-$i.yaml << EOF
apiVersion: cluster.clusterpedia.io/v1alpha2
kind: PediaCluster
metadata:
  name: clusterpedia-$i
spec:
  apiserver: https://$ip:1200$i
  caData: ${SYNCHRO_CA}
  tokenData: ${SYNCHRO_TOKEN}
  syncResources
  - group: apps
    resources:
     - deployments
  - group: ""
    resources:
     - pods
     - configmaps
     - secrets



EOF
if [ "$i" -ne 0 ];then
  kubectl --kubeconfig ~/.kube/clusterpedia-0 apply -f /tmp/subnet-import-$i.yaml
fi
done


#kubectl get --raw="/apis/clusterpedia.io/v1beta1/resources/clusters/clusterpedia-1/apis/apps/v1/namespaces/kube-system/deployments/" | jq
#kubectl get --raw="/apis/clusterpedia.io/v1beta1/resources/apis/apps/v1/namespaces/kube-system/deployments/" | jq
