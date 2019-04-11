# website-stackflow
The website for fetching questions regarding to Android from [Stack Overflow](https://stackoverflow.com/).


## Prerequisites
1. Docker 
- [Install Docker Desktop for Windows](https://docs.docker.com/docker-for-windows/install/)
- [Install Docker Desktop for Mac](https://docs.docker.com/docker-for-mac/install/)

2. Git 
- [Install Git on Your Machine](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Running the Web Application

1. Clone source codes from GitHub
```
$ git clone https://github.com/justina777/website-stackflow.git .
```

2. Build the web application into Docker image
```
$ docker build --rm -t website-stackflow -f app.Dockerfile .
```

3. Run the web application through docker
```
$ docker run --rm -it -p 8080:8080 website-stackflow
```

4. Browse the website on local machine through Chrome
```
http://localhost:8080
```