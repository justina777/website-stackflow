FROM golang:1.11-alpine3.9

MAINTAINER "Justina <chiahuei.lin@gmail.com>"

RUN apk add --no-cache bash curl tar git 

RUN go get gopkg.in/resty.v1 
    # && 
RUN mkdir -p ${GOPATH}/src/github.com/justina777/website-stackflow
ADD . ${GOPATH}/src/github.com/justina777/website-stackflow
RUN cd ${GOPATH}/src/github.com/justina777/website-stackflow \
    && go build 
WORKDIR ${GOPATH}/src/github.com/justina777/website-stackflow
EXPOSE 8080

CMD ["./website-stackflow"]