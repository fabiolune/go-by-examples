if [ -z "$1" ]
then
	echo "need to specify a project name:"
	echo "./create.sh <project>"
	exit 1
fi

if [ -d $1 ]
then
	echo "The project $1 already exixts"
	exit 1
fi

echo "creating project $1..."

mkdir $1
cd $1
mkdir src

cat << EOF > Dockerfile
## BUILD stage
FROM golang:1.15-alpine as build

WORKDIR /app
COPY ./src /app/
RUN go build


## RUN stage
## REMARK: when using scratch, no certificate authorities are actually available and http equests will fail
## an aproach is described here: https://titanwolf.org/Network/Articles/Article?AID=674fa38a-dfd5-4f67-bfd9-3b3197ac3ad4#gsc.tab=0
FROM scratch

ARG appname
WORKDIR /app
COPY --from=build /app/\${appname} /app/app

CMD ["./app"]
EOF

cd src
go mod init $1

cat << EOF > main.go
package main

import "fmt"

func main(){
	fmt.Println("Hello world from $1")
}

EOF
