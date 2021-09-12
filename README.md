# Go-Python-Docker-App

A mock application that puts together some code written in Python and Golang using docker.

In Go there is a simple application that sends a random message as a POST request to a simpe FastAPI server.

The server then reads the JSON request and writes the message to a file.

The points of this is just to test docker networks and how I could get two programs to talk to each other.

To run this app:
```
docker-compose build
docker-compose up
```
