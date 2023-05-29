#!/bin/bash

# shellcheck disable=SC1091
set -aueo pipefail

source .env

FSM_POD=$(kubectl get pods -n "$K8S_NAMESPACE" --no-headers  --selector app=fsm-controller | awk 'NR==1{print $1}')

kubectl port-forward -n "$K8S_NAMESPACE" "$FSM_POD"  9092:9092 --address 0.0.0.0
