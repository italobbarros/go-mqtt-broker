# go-mqtt-broker

# TO DO
- [x] Connected MQTT
- [x] Publish Qos 0,1,2
- [x] Subscribe Qos 0,1
- [x] Ping
- [x] UnSubscribe
- [ ] Save Username and Password on redis or DB
- [ ] TLS authenticator

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
