# go-crud
crud application in go

#Helpfule Docker commands
# build docker image 
docker build -t vamshi498/go-crud:yourtag . ( for example: docker build -t vamshi498/go-crud:0.2 .)

# run docker image
docker run -it -p 8080:8080 vamshi498/go-crud:yourtag . (for example: docker run -it -p 8080:8080 vamshi498/go-crud:0.2)

# list docker images
docker images

# list docker containers
docker container ls -a ( -a flag lists all the containers including the terminated ones)