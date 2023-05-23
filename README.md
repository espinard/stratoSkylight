# stratoSkylight
Technical Assessment


This code aims at practicing deployment of simple RESTful web app using Go langage and docker. 

If you are a technical assessor, I assume you have the right environment set up :) 


Building Docker Image 

```shell
$ docker build --tag docker-stratos .
```


Running Docker container. The stratos web app will be served on port 8080

```shell
$ docker run --publish 8080:8080 docker-stratos
```
