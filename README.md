# check_k8s
Nagios k8s checker

command_line    $USER1$/check_k8s -r=$ARG1$ -n=$ARG2$ $ARG3$
ARG1 = resource type (pod)
ARG2 = namespace
ARG3 = reource name

need .kube config folder in /tmp
