#!/bin/bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

function append_env_var() {
  envvar="${1}"
  entry="${2}"

  if [[ -z "${!envvar:-}" ]]; then
    export "${envvar}=${entry}"
  else
    export "${envvar}=${!envvar}, ${entry}"
  fi
}

function wait_for_resource_create() {
  type="${1}"
  name="${2}"
  ns="${3}"
  start_time=$(date '+%s')
  while ! kubectl get "${type}" "${name}" -n "${ns}"; do
    echo "Waiting for namespace creation"
    sleep 1
    if [[ $((start_time+30)) -lt $(date '+%s') ]]; then
      echo "timeout reached!"
      exit 1
    fi
  done

  kubectl rollout status --watch --timeout=240s "${type}/${name}" -n "${ns}"
}



# Create cluster
kind create cluster --config="${SCRIPT_DIR}/cluster-config.yaml"

# Set context
kubectl cluster-info --context kind-kind

# Install nginx
echo "Installing nginx ingress"
kubectl apply -f "${SCRIPT_DIR}/nginx-ingress-deploy.yaml"

# Install metrics api server
echo "Installing metrics api server"
# kubectl apply -f "${SCRIPT_DIR}/metrics-api-server.yaml"

# Wait for ingress to start
#wait_for_resource_create "deploy" "ingress-nginx-controller" "ingress-nginx"

# Wait for api server to start
#wait_for_resource_create "deploy" "metrics-server" "kube-system"