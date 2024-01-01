# Use a imagem oficial do Golang como imagem base
FROM golang:1.21-alpine as build

ENV TZ=America/Sao_Paulo

WORKDIR /code
COPY . /code/

# Compile o código Go
RUN go build -o . ./main.go 

# Stage 2
FROM scratch
ENV TZ=America/Sao_Paulo
WORKDIR /
# Copie o binário compilado para a imagem final
COPY --from=build /code/main /main
COPY --from=build /code/configs /configs
COPY --from=build /code/.env /.env
COPY --from=build /code/logs /logs

