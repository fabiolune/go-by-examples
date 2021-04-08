
docker container stop go-$1-run-ex 2>/dev/null
docker container rm go-$1-run-ex 2>/dev/null;
docker build ./$1 -t go-$1-run --build-arg appname=$1

if [ ! -z "$2" ] && [ $2 = "expose" ]
then
	echo docker run -d --rm -p $3:$3 --name go-$1-run-ex go-$1-run
	docker run -d --rm -p $3:$3 --name go-$1-run-ex go-$1-run
else
	echo docker run --rm --name go-$1-run-ex go-$1-run
	docker run --rm --name go-$1-run-ex go-$1-run
fi
