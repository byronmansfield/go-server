# RESTful API Built in Go

This is a simple RESTful Server written in Go. It is intended to be for my
personal learning as well as my generic boilerplate I can start from anytime I
need to whip up another REST API or even just include package(s) of it in another project.

It is a Work In Progress and work on it in my spare time, so please be patient

## Using this project

I am trying to build up a wiki for this so please check there if you do not
find what your looking for on the README.

### Building and running the server

After cloning the project you can simply build and run it by 

```
go build && ./go-server
```

Then you will have a server listening in on port 8080. Go to your web browser
and navigate to `localhost:8080`. You should see "Welcome from Go-Server!"

### Request data from it

If you want to make an async call to it to retrieve data you will need to have
a database set up and running. Lets set it up to communicate with the service
first.

Now make sure MySQL is running and has data in it. It also expects to have a
few environment variables exposed. Right now I am just using a .env file. I am
still looking for a cleaner way to do this. If you have all that in order, it 
should be ready to GET and POST data to it. 

##### GET
To get data from it like so:

```
curl -i localhost:8080/todo
```

And you should see an array of objects. (depending on whats in your database)

You can alternatively get just a single todo by name. (I plan to expand on this
later)

```
curl -i localhost:8080/todo/Code
```

Should return to you once again an array of objects. (Again plan to improve on
this) An array just in case it finds multiple matches.


##### POST
Post data to it like so:

```
curl -X POST -H "Content-type: application/json" -d '{"id":"4","name":"go to store","description":"pick up groceries","due":"2015-09-08"}' localhost:8080/todo
```

And for right now all you will get back is a 200 ok. (Like the rest, will
improve upon later)

## More later

Thats all I have for now. This is a work in progress. Please be patient.

