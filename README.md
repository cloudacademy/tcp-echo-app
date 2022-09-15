![Build Status](https://github.com/cloudacademy/tcp-echo-app/actions/workflows/go.yml/badge.svg)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/cloudacademy/tcp-echo-app)

# TCP Echo App
A simple TCP echoing application.

The TCP echoing application been developed using the Go programming language. It is designed to simply echo back any data that it recieves, together with the client's source IP address. Clients connect to this application and communicate using TCP (layer-4) connections.

## Usage
To start the TCP echo application, configure the `HOSTPORT` environment variable. `HOSTPORT` represents the listening address and port that the TCP echo application listens on.

Startup:
```
HOSTPORT=0.0.0.0:9091 tcpapp
```

## Docker
The TCP echoing application has been packaged into a Docker image. The Docker image can be pulled with the following command:

```
docker pull cloudacademydevops/tcpapp:v1
```

Use the following command to launch the TCP echoing application within Docker:
```
docker run --name tcpapp -p 9091:9091 --detach cloudacademydevops/tcpapp:v1
```

## Kubernetes
Use the following command to launch the TCP echoing application as a Deployment resource within a cluster:

```
cat << EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: tcpapp
  namespace: tcpapp
spec:
  selector:
    matchLabels:
      app.kubernetes.io/name: tcpapp
  replicas: 2
  template:
    metadata:
      labels:
        app.kubernetes.io/name: tcpapp
    spec:
      containers:
      - image: cloudacademydevops/tcpapp:v1
        imagePullPolicy: Always
        name: tcpapp
        ports:
        - containerPort: 9091
EOF
```

## Build
The following commands can be used to build and package the source code:

Current operating system:
```
go build .
```

Linux operating system:
```
CGO_ENABLED=0 GOOS=linux go build -o tcpapp .
```

Docker:
```
docker buildx build --platform=linux/amd64 -t cloudacademydevops/tcpapp:v1 .
```