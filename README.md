# check_k8s
Nagios k8s checker

~~~
command_line    $USER1$/check_k8s -r=$ARG1$ -n=$ARG2$ $ARG3$
ARG1 = resource type (pod)
ARG2 = namespace
ARG3 = reource name
~~~
~~~
status[0] == true && status[1] == true && status[2] == true        => exit code 0 "OK"
status[0] == false && status[1] == true && status[2] == true       => exit code 1 "WARNING"
status[0] == false && (status[1] == false || status[2] == false)   => exit code 1 "CRITICAL"
                             other                                 => exit code 1 "CRITICAL"
~~~



need .kube config folder in /tmp
