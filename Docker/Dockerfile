FROM bitnami/minideb-extras:jessie-r14-buildpack as buildenv

RUN bitnami-pkg install go-1.8.3-0 --checksum 557d43c4099bd852c702094b6789293aed678b253b80c34c764010a9449ff136

ENV GOPATH=/gopath
ENV PATH=$GOPATH/bin:/opt/bitnami/go/bin:$PATH

RUN  apt-get update &  apt-get install git 
RUN  git clone https://github.com/vtuson/cmsidecar.git 
RUN  go build ./cmsidecar

FROM bitnami/minideb:latest

COPY --from=buildenv /cmsidecar/cmsidecar /
CMD ["/cmsidecar"]