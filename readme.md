## This repo contain avoxi test code
A restful api, which will give user the status of the IP, whether IP is `allowed` or `not allowed`

Curl i.e.
`
curl --location --request GET 'http://127.0.0.1:8052/network?ip=36.0.4.20&whitelist=USA,Nepal'
`
should return

`
{
    "IP": "Not Allowed"
}
`

Before running about command you need to run the servers.

Make sure you have `docker` and `docker-compose` install on your machine

Steps

- Clone this repo
- go inside the avoxi folder
- Run `docker-compose build && docker-compose up -d`
- After this check the log message, everything should be running properly
- Now, connect with the database, db cred info is in  `.env` file
- Go inside the rest folder you will find `.sql`, Copy the file's content and run into the database ( use `PostBird` - It's Postgre database client ) - Database cred are in `.evn` file.
- Now your table must have 3 rows 
- Now run the above curl code, you should see  
```
{
    "IP": "Allowed"
}
```

If you look info .evn file, you should see an environment variable `IP_LOCATION_URL`
if you set the URL for this, our app will hit `IP_LOCATION_URL` endpoint to find-out the country name for the given IP address

Set `IP_LOCATION_URL=https://iplocation.com/` ( Use it to test only - I am not sure whether this api service is free or not so better not to messup someone else api )

Now, rebuild your docker image
`docker-compose build && docker-compose up -d`

Now, check the log `docker-compose logs rest-avoxi`, you should see text
`IP to Country search will happen in`


##### NOTE: This code is not for production use

- TODO: Multi-stage docker image
- TODO: Improvise the data lookup
- TODO: Implement GRPC
- TODO: More error handling logic

## Happy Coding!!!

