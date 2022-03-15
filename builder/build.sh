#!/bin/bash
echo "docker login"
docker login -u="$USER" -p="$PASSWORD" quay.io

echo "git clone"
git clone https://github.com/openshift/windows-machine-config-operator.git --recursive

echo "$TAG"
export OPERATOR_IMAGE=quay.io/winc/wmco:"$TAG"
echo "$OPERATOR_IMAGE"
cd windows-machine-config-operator

echo "hack/olm.sh build"
hack/olm.sh build

echo "make bundle WMCO_VERSION=$TAG IMG=quay.io/winc/wmco:$TAG"
make bundle WMCO_VERSION="$TAG" IMG=quay.io/winc/wmco:"$TAG"

echo "make bundle-build BUNDLE_IMG=quay.io/winc/wmco-bundle:$TAG"
make bundle-build BUNDLE_IMG=quay.io/winc/wmco-bundle:"$TAG"

echo "docker push quay.io/winc/wmco-bundle:$TAG"
docker push quay.io/winc/wmco-bundle:"$TAG"

echo "operator-sdk bundle validate quay.io/winc/wmco-bundle:$TAG --image-builder docker"
operator-sdk bundle validate quay.io/winc/wmco-bundle:"$TAG" --image-builder docker

echo "opm index add --bundles quay.io/winc/wmco-bundle:$TAG --tag quay.io/winc/wmco-index:$TAG --container-tool docker"
opm index add --bundles quay.io/winc/wmco-bundle:"$TAG" --tag quay.io/winc/wmco-index:"$TAG" --container-tool docker

echo "docker push quay.io/winc/wmco-index:$TAG"
docker push quay.io/winc/wmco-index:"$TAG"