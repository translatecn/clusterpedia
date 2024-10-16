export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
cd cmd/apiserver && go build -o ../../bin/apiserver . && cd -
cd cmd/clustersynchro-manager && go build -o ../../bin/clustersynchro-manager . && cd -
cd cmd/controller-manager && go build -o ../../bin/controller-manager . && cd -
docker build -t myself:v1 .
kind load docker-image --name clusterpedia-0 myself:v1