# RESTful Server Built in Go

This is a simple RESTful Server written in Go. It is intended to be for my
personal learning as well as my generic boilerplate I can start from anytime I
need to whip up another REST API or even just include it in my go project.

## Using this project

### Building and running the server

After cloning the project you can simply build and run it by 

```
go build && ./go_server
```

Then you will have a server listening in on port 8080. Go to your web browser
and navigate to `localhost:8080`. You should see "Hello from Go HTTP!"

### Request data from it

If you want to make an async call to it to retrieve data you will need to have
a database set up and running. Lets set it up to communicate with the service
first.

I will have to come back to this part

Now make sure MySQL is running and has data in it. Now it should be ready to
GET and POST data to it. 

To get data from it like so:

```
curl -i localhost:8080/todos
```

Post data to it like so:

```
curl -X POST -H "Content-type: application/json" -d '{"id":"4","name":"go to store","completed":"false","due":"2015-09-08"}' localhost:8080/todos
```

## More later

Thats all I have for now. This is a work in progress. Please be patient.
