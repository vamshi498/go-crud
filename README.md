# go-crud with postgres
crud application in go. It connects to postgres database

#Helpfule Docker commands
# build docker image 
docker build -t vamshi498/go-crud:yourtag . ( for example: docker build -t vamshi498/go-crud:0.2 .)

# run docker image
docker run -it -p 8080:8080 vamshi498/go-crud:yourtag . (for example: docker run -it -p 8080:8080 vamshi498/go-crud:0.2)

# list docker images
docker images

# list docker containers
docker container ls -a ( -a flag lists all the containers including the terminated ones)

# remove docker images from your local machine
docker rmi -f containerId ( -f is force . Ex: docker rmi -f ccda0e5ccbf  . (It need not be 7-10 digits. Its just a unique id. it can be 2 digits long as it is unique))