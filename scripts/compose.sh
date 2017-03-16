#!/bin/bash

# we assume we have the correct permissions to start docker containers
# start mongodb

set -x

if [ "$1" = "start" ]
then
	mongodb_image=$(docker images | grep mongodb | awk '{print $3}')
	golang_image=$(docker images | grep golang | awk '{print $3}')

	docker run -d -e "MONGODB_USER=test" -e "MONGODB_PASSWORD=test" -e "MONGODB_DATABASE=test" -e "MONGODB_ADMIN_PASSWORD=admin" --name mongodb-service $mongodb_image

	#start golang
	docker run -it  --link mongodb-service -p 80:8080 --privileged -v .:/go --name golang $golang_image bash
	
	exit 0
fi

if [ "$1" = "stop" ]
then

	# get the running containers id's
	mongodb_pid=$(docker ps -a | grep mongodb-service | awk '{print $1}')
	golang_pid=$(docker ps -a | grep golang | awk '{print $1}')

	echo "Stopping all docker containers .."

	docker stop $golang_pid $mongodb_pid

	echo "Deleting all exited docker containers .."

	docker rm $golang_pid $mongodb_pid

	exit 0
fi

