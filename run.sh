
docker container stop go-$1-run-ex 2>/dev/null
docker container rm go-$1-run-ex 2>/dev/null;
docker build ./$1 -t go-$1-run 
docker run -d --rm -p 8080:8080 --name go-$1-run-ex go-$1-run
