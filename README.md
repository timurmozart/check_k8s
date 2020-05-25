# check_k8s
Nagios k8s checker

~~~
command_line    $USER1$/check_k8s -r=$ARG1$ -n=$ARG2$ $ARG3$
ARG1 = resource type (pod)
ARG2 = namespace
ARG3 = reource name
~~~
kubectl get service -n jenkins

|NAME                         | READY | STATUS  | RESTARTS | AGE|
| --- | --- | --- | --- | --- |
|jenkins-7d84d4dbcd-kpdph      | 0/1   | Running | 0        | 56s|
|nginx-backend-b4cddd8b8-gh6hm | 1/1   | Running | 0        | 3d5h|

kubectl get pod -n jenkins

|NAME |READY|STATUS|RESTARTS|AGE|
| --- | --- | --- | --- | --- |
|jenkins-.......|1/1|Running|0|56s|
|......-..... | 1/1   | Running | 0        | 3d5h|
|     |     |     |     |     |

status[0] = curl jenkins_ip:30001/login == 200 
status[1] = STATUS
status[2] =   

~~~
status[0] == true && status[1] == true && status[2] == true        => exit code 0 "OK"
status[0] == false && status[1] == true && status[2] == true       => exit code 1 "WARNING"
status[0] == false && (status[1] == false || status[2] == false)   => exit code 1 "CRITICAL"
                             other                                 => exit code 1 "CRITICAL"
~~~




need .kube config folder in /tmp
