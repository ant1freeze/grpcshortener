<h4>GRPCShortener is a link shortener that can take a long url as input and convert it to a short one.
 
And this service can take a short url as input and convert it to a long one, if it is in the database.</h4>

The project is built on the base of gRPC framework, use PostgreSQL and include gRPC server, gRPC client and goose (for initial migration) in docker cintainers.

**Requirements**:

    Docker version 19.03.5, build 633a0ea838

    docker-compose version 1.25.3, build d4d1b42b

**Usage**:

    All docker containers based on golang:1.16-alpine 

    Database: psql (PostgreSQL) 11.12 (Debian 11.12-0+deb10u1)

    For migrations: github.com/pressly/goose/v3

    For get env args: github.com/spf13/viper

**Before install** grpcshotener you need **install postgreSQL** and **create password for default user postgresql if not exists:

```linux
user:~$ sudo -i -u postgres

postgres@user:~$ psql

postgres=# ALTER USER postgres PASSWORD 'mynewpassword';
```

And then add this credentials to ./configs/app.env

**OR**

You can **create user/password and database for it**, for example:

```postgresql
postgres=> create user test with password 'testpass' createdb;

postgres=> create database test;
```

And then add this credentials to ./configs/app.env

<h5>Install:</h5>

```linux
user:~$ git clone https://github.com/ant1freeze/grpcshortener.git

user:~$ cd grpcshortener

user:~$ docker build . -t=shorter_client -f=Dockerfile_client

user:~$ docker-compose up --build
```

There are 2 docker containers in docker-compose.yml (shortener_server and goose).
Server starts on the port 50051 by default (you can change it in ./configs/app.env)
Then goose makes migration (create table urls in your database).

After 2 docker containers were up you can run client with args:

<h6>For CREATE short url:</h6>

 ```linux
user:~$ docker run --net=host shorter_client create google.com
 ```
 
 return:
 
 ```linux
 2021/09/08 12:50:59 localhost/gvahJggOzY
 ```
 
 <h6>For GET long url from db:</h6>
 
   ```linux
 user:~$ docker run --net=host shorter_client get gvahJggOzY
  ```
  
  return:
  
  ```linux
  2021/09/08 12:51:08 google.com
  ```
  
  If long url doesn't exist in database - return "Didn't find anything."
  
  ```linux
  user:~$ docker run --net=host shorter_client get aaaaaa
  2021/09/08 13:09:14 Didn't find anything.
  ```
  
  If you write only get/create without url or nothing, docker will return help:
   
  ```linux
  user:~$ docker run --net=host shorter_client get
  2021/09/08 13:10:11 Need type 'get <short URL>' or 'create <long URL>'
  ```
