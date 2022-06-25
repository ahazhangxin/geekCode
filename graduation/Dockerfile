FROM 69.172.74.253:8080/jenkins/golang:1.13
MAINTAINER lugj@efly.cc
EXPOSE 80
COPY ./cmd/cmd /goSrc/
COPY ./configs /goSrc/configs
RUN chmod -R 775 /goSrc
WORKDIR /goSrc
CMD /goSrc/cmd web -c /goSrc/configs/config.toml -m /goSrc/configs/model.conf
