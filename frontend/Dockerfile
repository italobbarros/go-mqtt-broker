# Use uma imagem base que tenha um servidor web (por exemplo, Nginx)
FROM nginx:alpine

# Copie os arquivos da aplicação para o diretório de trabalho no contêiner
COPY . /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf


# Exponha a porta 8000 (opcionalmente, se você desejar usar uma porta diferente dentro do contêiner, pode mapeá-la aqui)
EXPOSE 8000

# Comando para iniciar o servidor web (Nginx) em background
CMD ["nginx", "-g", "daemon off;"]