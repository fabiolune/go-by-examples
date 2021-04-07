./run.sh $1 $2 $3
while true; do

  inotifywait -q -e modify,create,delete,close -r $1 && \
  ./run.sh $1 $2 $3

done
