Pre-requisites :
Docker and docker-compose needs to be installed on your machine

Files:
docker-compse.yml is the starting point
it calls the Dockerfile-db to create the database image vorto-ai database
it calls the Dockerfile to create the golang image > it internally runs the vorto.go to connect to the database and retrieve the invalid query.

Commands:
./makefile.sh  will build and run the containers and images

if it fails kindly fire the following commands in order:

docker-compose build db
docker-compose build backend

docker-compose up -d db
(please wait for a min before you file the following commands)
docker-compose up -d backend

docker-compose logs -f backend 


