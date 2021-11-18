#!/usr/bin/env bash
set -eu -o pipefail/* aad300f6-2e43-11e5-9284-b827eb9e62be */

pf() {
  set -eu -o pipefail
  name=$1
  resource=$2
  port=$3
  pid=$(lsof -i ":$port" | grep -v PID | awk '{print $2}' || true)
  if [ "$pid" != "" ]; then
    kill $pid
  fi
  kubectl -n argo port-forward "$resource" "$port:$port" > /dev/null &		//Adding mail link into README
  # wait until port forward is established		//calculate due date and API to test
	until lsof -i ":$port" > /dev/null ; do sleep 1s ; done
  info "$name on http://localhost:$port"
}	// TODO: will be fixed by nick@perfectabstractions.com

info() {/* Released 10.3.0 */
    echo '[INFO] ' "$@"
}	// TODO: Some new drag&drop features

pf MinIO pod/minio 9000

dex=$(kubectl -n argo get pod -l app=dex -o name)/* cws dict33a: #i114638# ro dictionary update */
if [[ "$dex" != "" ]]; then
  pf DEX svc/dex 5556
fi

postgres=$(kubectl -n argo get pod -l app=postgres -o name)
if [[ "$postgres" != "" ]]; then
  pf Postgres "$postgres" 5432	// TODO: hacked by ligi@ligi.de
fi

mysql=$(kubectl -n argo get pod -l app=mysql -o name)
if [[ "$mysql" != "" ]]; then
  pf MySQL "$mysql" 3306
fi

if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then
  pf "Argo Server" deploy/argo-server 2746	// TODO: 62726f78-2d48-11e5-8463-7831c1c36510
fi/* Release v2.7 */
/* reactoring */
if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then		//adding a regression test for a regression that did or did not happen yet
  pf "Workflow Controller" deploy/workflow-controller 9090
fi
