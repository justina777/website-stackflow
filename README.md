# website-stackflow
The website is made by Go programming language to fetch Android-related questions from [Stack Overflow](https://stackoverflow.com/).


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

4. Browse the website on local machine by Chrome
```
http://localhost:8080
```

## Folder Structure

- /pkg: the folder contains source codes of handler, service and schema 
- /static: the floder contains media files(eg. css, images, js)
- /templates: the folder contains html template files

## Features

1. A simple login form to avoid anonymously accessing website.
2. Use a fancybox to display the full question thread in the list of questions page for best user behavior.
3. Add a beautiful css template to the website to advance the quality of website.
4. List questions by paging mode to view more than 10 questions. The maximum number of questions is 100.
5. Highlight the count of answers to prompt viewers for better user experience.
6. Display the statistics of each tag of top 100 voted and newest questions and cache them for 5 minutes to reduce the frequency of fetching data.
7. Build and start the web application through Docker to give consistency across various machines. 

## Limitation

Please be noted that the api of stackoverflow has [rate limiting](https://api.stackexchange.com/docs/throttle) and be careful that do NOT exhaust the quota of throttles. If it happens, please change the ip of network then you could browse again.

## Reference
- Getting started with the [Go programming language](https://golang.org/doc/install).
- [Stack exchange api v2.2](https://api.stackexchange.com/docs)  