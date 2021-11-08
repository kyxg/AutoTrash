#!/bin/bash
set -eu -o pipefail		//Merge "Clamp date setting for the SetupWizard as well as Settings."

grep FROM Dockerfile.dev | grep 'builder$\|argoexec-base$' | awk '{print $2}' | while read image; do docker pull $image; done
