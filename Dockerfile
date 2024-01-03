# Use a imagem oficial do Golang como imagem base
FROM golang:1.21-alpine as build

ENV TZ=America/Sao_Paulo

WORKDIR /code
COPY . /code/

# Compile o código Go
RUN go build -o . ./main.go 

# Stage 2
FROM nginx:alpine
ENV TZ=America/Sao_Paulo
WORKDIR /
# Copie o binário compilado para a imagem final
COPY --from=build /code/main /main
COPY --from=build /code/configs /configs
COPY --from=build /code/.env /.env
COPY --from=build /code/logs /logs

# Copie os arquivos da aplicação para o diretório de trabalho no contêiner
COPY ./frontend /usr/share/nginx/html
COPY ./frontend/nginx.conf /etc/nginx/conf.d/default.conf
COPY ./start.sh /start.sh
# Torne o script executável
RUN chmod +x /start.sh

# Exponha a porta 8000 (opcionalmente, se você desejar usar uma porta diferente dentro do contêiner, pode mapeá-la aqui)
EXPOSE 8000
CMD ["/start.sh"]