# Running for example: docker build -f deploy/Dockerfile -t project1 .

FROM  golang:1.17.11-buster

WORKDIR /home/app/
COPY . /home/app/
RUN cd /home/app/src && go mod init img_convert && go mod tidy && go build

EXPOSE 8080



