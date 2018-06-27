FROM bitnami/minideb-extras:jessie-r14-buildpack as buildenv

RUN  apt-get update; apt-get install -y git wget 

ENV  GOPATH=/gopath
ENV  PATH=$GOPATH/bin:$PATH
RUN  wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
RUN  tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz
ENV  PATH=$PATH:/usr/local/go/bin
RUN  mkdir -p /gopath/src/github.com/vtuson
RUN  git clone https://github.com/vtuson/cmsidecar.git /gopath/src/github.com/vtuson/cmsidecar
RUN  mkdir -p /gopath/bin
RUN  go install github.com/vtuson/cmsidecar
RUN wget https://storage.googleapis.com/kubernetes-helm/helm-v2.9.1-linux-amd64.tar.gz
RUN tar xvf helm-v2.9.1-linux-amd64.tar.gz -C /

FROM bitnami/minideb:latest

COPY --from=buildenv /gopath/bin/cmsidecar /
COPY --from=buildenv /linux-amd64/helm /usr/local/bin/helm
RUN  apt-get update; apt-get install -y git 
RUN helm init --client-only  
CMD ["/cmsidecar"]