#!/bin/bash

set -aueo pipefail

# shellcheck disable=SC1091
source .env

POD="$(kubectl get pods -n "$K8S_NAMESPACE" --selector app=fsm-controller --no-headers | awk '{print $1}' | head -n1)"

kubectl logs "${POD}" -c fsm-controller -n "$K8S_NAMESPACE" -f
