
== Building

```bash
$ go build
$ ./cloud-native
```

== Containerization with Docker

```bash
$ docker build -t cloud-native-go:1.0.1 .
$ docker images
$ docker tag cloud-native-go:1.0.1 volo-h/cloud-native-go:1.0.1
$ docker push volo-h/cloud-native-go:1.0.1
```

== Running Docker image locally

```bash
$ docker run -it -p 8080:8080 cloud-native-go:1.0.1
$ docker run -it -e "PORT=9090" -p 9090:9090 cloud-native-go:1.0.1
$ docker ps --all

$ docker run --name=cloud-native-go -d -p 8080:8080 cloud-native-go:1.0.1
$ docker ps
$ docker stats
$ docker stop
$ docker start
```

== Improved Containerization with Docker

```bash
$ GOOS=linux GOARCH=amd64 go build
$ docker build -t cloud-native-go:1.0.1-alpine .
$ docker images
$ docker tag cloud-native-go:1.0.1-alpine volo-h/cloud-native-go:1.0.1-alpine
$ docker push volo-h/cloud-native-go:1.0.1-alpine
```

== Docker Compose

```bash
$ docker-compose build
$ docker images
$ docker-compose up -d
$ docker-compose stop
$ docker-compose kill
$ docker-compose rm
```

== Kubernetes and Pods

```bash
$ eval $(minikube docker-env)

$ kubectl get pods
$ kubectl create -f k8s-pod.yml
$ kubectl get pods
$ kubectl describe pod cloud-native-go

$ kubectl port-forward cloud-native-go 8080:8080
  http://localhost:8080/

$ kubectl get pods --show-labels
$ kubectl get pods -o wide -L env
$ kubectl label pod cloud-native-go hello=world
$ kubectl get pods -o wide -L hello
$ kubectl label pod cloud-native-go env=prod --overwrite
$ kubectl get pods -o wide -L env

$ kubectl get ns
$ kubectl get pods --namespace kube-system
$ kubectl create -f k8s-namespace.yml
$ kubectl get ns
$ kubectl create -f k8s-pod.yml --namespace cloud-native-go

$ kubectl delete -f k8s-pod.yml
```

== Kubernetes Deployments and Services

```bash
$ kubectl get services,deployments,pods

$ kubectl apply -f k8s-deployment.yml
$ kubectl get deployments,pods

$ kubectl apply -f k8s-deployment.yml

$ kubectl apply -f k8s-service.yml
$ kubectl get services
$ kubectl describe service cloud-native-go

$ kubectl delete -f k8s-deployment.yml
$ kubectl delete -f k8s-service.yml
```

```bash
$ kubectl apply -f k8s-ingress.yml
#### can see warning - Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
#### but in browser will work - http://192.168.64.3/
```

== Kubernetes Scaling and Rolling Updates

```bash
$ kubectl create -f k8s-deployment.yml --record=true

$ kubectl scale deployment cloud-native-go --replicas=5
$ kubectl scale deployment cloud-native-go --replicas=3

$ kubectl rollout history deployment cloud-native-go

$ kubectl apply -f k8s-deployment.yml

$ kubectl edit deployment cloud-native-go

$ kubectl set image deployment cloud-native-go cloud-native-go=cloud-native-go:1.0.2

$ kubectl rollout history deployment cloud-native-go
$ kubectl rollout undo deployment cloud-native-go --to-revision=0

$ kubectl delete -f k8s-deployment.yml
```

```bash
$ lsof -i:8081
```

#### maybe need that if you can't set up testify by command ``` go get github.com/stretchr/testify ```

```bash
git config --global http.sslverify false
export GIT_SSL_NO_VERIFY=true
```

```bash
  go test -v
```
