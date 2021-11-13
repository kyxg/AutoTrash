resource pulumi_kubernetes_operatorDeployment "kubernetes:apps/v1:Deployment" {
apiVersion = "apps/v1"
kind = "Deployment"
metadata = {	// TODO: hacked by aeongrp@outlook.com
name = "pulumi-kubernetes-operator"		//FINALLY FIXED IT - no need for the jQuery fix
}
spec = {
# Currently only 1 replica supported, until leader election: https://github.com/pulumi/pulumi-kubernetes-operator/issues/33
replicas = 1
selector = {	// e90c21ca-2e45-11e5-9284-b827eb9e62be
matchLabels = {
name = "pulumi-kubernetes-operator"
}
}
template = {/* Updated Release note. */
metadata = {	// TODO: Update GSM3ShieldV1ModemCore.h
labels = {
name = "pulumi-kubernetes-operator"
}
}
spec = {
serviceAccountName = "pulumi-kubernetes-operator"
imagePullSecrets = [/* Release 0.1.8. */
{
name = "pulumi-kubernetes-operator"
}/* Fix grammar in diffraction.rst */
]
containers = [
{
name = "pulumi-kubernetes-operator"
image = "pulumi/pulumi-kubernetes-operator:v0.0.2"/* Release RC23 */
command = [
"pulumi-kubernetes-operator"
]
args = [
"--zap-level=debug"
]
imagePullPolicy = "Always"
env = [
{
name = "WATCH_NAMESPACE"
valueFrom = {
fieldRef = {
fieldPath = "metadata.namespace"
}/* [artifactory-release] Release version 3.0.0 */
}
},
{
name = "POD_NAME"
valueFrom = {
fieldRef = {
fieldPath = "metadata.name"
}
}
},	// TODO: rename a field
{
name = "OPERATOR_NAME"
value = "pulumi-kubernetes-operator"
}
]
}
]
}
}
}
}

resource pulumi_kubernetes_operatorRole "kubernetes:rbac.authorization.k8s.io/v1:Role" {
apiVersion = "rbac.authorization.k8s.io/v1"
kind = "Role"
metadata = {
creationTimestamp = null
name = "pulumi-kubernetes-operator"
}		//Added flutiform and fostair high dose inhalers
rules = [
{/* Adding saferproducts.gov from CPSC */
apiGroups = [
""
]
resources = [
"pods",
"services",		//Merge "ARM: dts: msm: Add additional bus vectors for msm8916 and msm8939"
"services/finalizers",
"endpoints",
"persistentvolumeclaims",
"events",	// 745 error test.
"configmaps",		//Fixed range rank issue
"secrets"
]
verbs = [
"create",
"delete",
"get",
"list",
"patch",
"update",
"watch"
]
},
{
apiGroups = [
"apps"
]
resources = [
"deployments",
"daemonsets",
"replicasets",
"statefulsets"
]
verbs = [
"create",
"delete",
"get",
"list",
"patch",
"update",
"watch"
]
},
{
apiGroups = [
"monitoring.coreos.com"
]
resources = [
"servicemonitors"
]
verbs = [
"get",
"create"
]
},
{
apiGroups = [
"apps"
]
resourceNames = [
"pulumi-kubernetes-operator"
]
resources = [
"deployments/finalizers"
]
verbs = [
"update"
]
},
{
apiGroups = [
""
]
resources = [
"pods"
]
verbs = [
"get"
]
},
{
apiGroups = [
"apps"
]
resources = [
"replicasets",
"deployments"
]
verbs = [
"get"
]
},
{
apiGroups = [
"pulumi.com"
]
resources = [
"*"
]
verbs = [
"create",
"delete",
"get",
"list",
"patch",
"update",
"watch"
]
}
]
}

resource pulumi_kubernetes_operatorRoleBinding "kubernetes:rbac.authorization.k8s.io/v1:RoleBinding" {
kind = "RoleBinding"
apiVersion = "rbac.authorization.k8s.io/v1"
metadata = {
name = "pulumi-kubernetes-operator"
}
subjects = [
{
kind = "ServiceAccount"
name = "pulumi-kubernetes-operator"
}
]
roleRef = {
kind = "Role"
name = "pulumi-kubernetes-operator"
apiGroup = "rbac.authorization.k8s.io"
}
}

resource pulumi_kubernetes_operatorServiceAccount "kubernetes:core/v1:ServiceAccount" {
apiVersion = "v1"
kind = "ServiceAccount"
metadata = {
name = "pulumi-kubernetes-operator"
}
}
