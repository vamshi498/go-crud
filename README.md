# go-crud with postgres
crud application in go. It connects to postgres database
We are using golang-migrate https://github.com/golang-migrate/migrate to perform db migrations

# start the container and run the application 
```docker-compose up```  

# insert data using /addUser endpoint
Ex:  
``` curl -XPOST -H "Content-type: application/json" -d '{"firstname":"Vamshi","lastname":"Muthyapu"}' 'http://localhost:8080/addUser' ```  

# retreive data using /users/{id} endpoint 
Ex:  
```curl http://localhost:8080/users/1 ```  

Response for above curl would look like  
``` {"Firstname":"Muthyapu","Lastname":"","Address":{"City":"","State":"","Country":""}} ```  

# db migration
To perform db migrations using golang-migrate , we need to have two files
1. 01_first_migration.up.sql
2. 01_first_migration.down.sql ( roll back in case migration is not successfull)

#Helpful Docker commands
# build docker image 
docker build -t vamshi498/go-crud:yourtag . ( for example: docker build -t vamshi498/go-crud:0.2 .)

# run docker image
docker run -it -p 8080:8080 vamshi498/go-crud:yourtag . (for example: docker run -it -p 8080:8080 vamshi498/go-crud:0.2)

# list docker images
docker images

# list docker containers
docker container ls -a ( -a flag lists all the containers including the terminated ones)

# remove docker images from your local machine
docker rmi -f imageId ( -f is force . Ex: docker rmi -f ccda0e5ccbf  . (It need not be 7-10 digits. Its just a unique id. it can be 2 digits long as it is unique))

# remove container 
docker rm containerId

# remove all images 
docker system prune -a 

# If any process is running and you want to kill it. 
For ex you want to kill PID of a service which is using 8080
ps -aux | grep 8080 --> this will give you PID.  After this do kill -9 PID

# docker compose 
1.docker-compose up  (builds and runs the stack)  
2.docker-compose down ( stops the services)  
3.docker-compose up --build ( rebuilds the stack and then starts the services)  

# get details about docker container like IPAddress, config etc
docker inspect containerId
