FROM golang:1.19

WORKDIR /usr/src/app

COPY src .

RUN go mod download && go mod verify

RUN go build -v -o start

ENV PATH=/usr/src/app
EXPOSE 8080

CMD ["start"]


#You can then build and run the Docker image:

#$ docker build -t users-image:1 .

#$ docker run -it -p 8080:8080 -e mysql_users_host=172.17.0.1 --rm --name users-container users-image:1 

#$ docker run -it -p 8080:8080 -e mysql_users_host=172.18.0.2 --rm --name users-container users-image:1 

#$ docker run -it -p 8080:8080 -e mysql_users_host=db_moova --network=mysql_default --rm --name users-container users-image:1 