GRPCShortener is a link shortener that can take a long url as input and convert it to a short one.
And thos service can take a short url as input and convert it to a long one, if it is in the database.

The project is built on the base of gRPC framework, use PostgreSQL and include gRPC server, gRPC client and goose (for initial migration) in docker cintainers.

Before install grpcshotener you need install postgreSQL and create password for default user postgresql if not exists:

user:~$ sudo -i -u postgres
postgres@user:~$ psql
postgres=# ALTER USER postgres PASSWORD 'mynewpassword';

OR

You can create user/password and database for it, for example:

postgres=> create user test with password 'testpass' createdb;
postgres=> create database test;

And then add this credentials to ./configs/app.env

Install:

git clone https://github.com/ant1freeze/grpcshortener.git
cd grpcshortener
docker-compose up --build

There are 2 docker containers in docker-compose.yml (shortener_server and goose).
Server starts on the port 50051 by default (you can change it in ./configs/app.env)
Then goose makes migration (create table urls in your database).

After 2 docker containers were up you can run client with args:

For CREATE short url:
 docker run --net=host shorter_client create <url>
 
 For GET long url from db:
  docker run --net=host shorter_client get <url>
  
  If long url doesn't exist in database - return "Didn't find anything."
  
  
