FROM golang:1.8.0

MAINTAINER Luis Mayta <@slovacus>

COPY ./glide.yaml /go/src/course-go/
COPY ./glide.lock /go/src/course-go/

# set working directory to /app/
WORKDIR /go/src/course-go/

# install golang dependencies
RUN curl https://glide.sh/get | sh \
  && glide install
