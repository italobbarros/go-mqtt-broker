# go-mqtt-broker

# TO DO 
## Frontend
- [x] Tree topics viewer
- [x] Info topic by topic-info endpoint
- [ ] General information about total of messages,sessions,subscripts...
- [ ] login
- [ ] save and edit username and password
- [ ] save and edit rules permission and security
  
## Backend
- [x] Connected MQTT
- [x] Publish Qos 0,1,2
- [x] Subscribe Qos 0,1
- [x] Ping
- [x] UnSubscribe
- [x] Endpoint tree mqtt topics
- [x] Endpoint info mqtt topic
- [ ] Save Username and Password on json and validate
- [ ] Save Username and Password on redis or DB
- [ ] Specific rules permission and security
- [ ] TLS authenticator
- [ ] Save payload on specific Database



# How to Run
If you prefer, you can run it using docker-compose. It is a better way to execute the code since it specifies all the parameters in a single file. The correct command is:
## Docker Compose

```bash
docker-compose up -d --build
```
But is possible to run using a docker command, like a bellow:
## Docker
### Frontend
```bash
docker build -t broker-management  .
```
```bash
docker run -d --name broker-management -p 3000:8000 --network=broker broker-management 
```
### Backend
```bash
docker build -t broker-mqtt  .
```
```bash
docker run -d --name broker-mqtt-1 -p 8080:8080 -p 1883:1883 --network=broker broker-mqtt 
```

# Testing

Using 500 clients publishing 100 messages, each message on a different topic. It was executed 12 times and averaged in groups of 4 for an average.

| Broker Name    | Highest (msg/s) | Average (msg/s) | Lowest (msg/s) |
|----------------|:---------------:|:---------------:|:--------------:|
| go-mqtt-broker | -           | -            | -           |
| Mosquitto      | _           | -            | -           |
| EMQX           | -           | -            | -           |
