FROM golang AS build-stage
LABEL authors="sevlak0ff"
WORKDIR /dami

#copia e baixa as dependências primeiro, fazendo com que o docker crie a cache dessa layer
COPY go.mod go.sum ./
RUN go mod download

#copia o resto do código
COPY . .

#cria .env
RUN echo 'PORT=:443' >> .env
RUN echo 'OTP_SECRET="damitest"' >> .env
RUN echo 'JWT_SECRET="brunoalvirubro"' >> .env
RUN echo 'DEVICE="http://device:80"' >> .env #device é o hostname do container do device (duh)
RUN echo 'CERTIFICATE="certs/localhost.crt"' >> .env
RUN echo 'PRIVATE_KEY="certs/localhost.key"' >> .env


RUN CGO_ENABLED=0 GOOS=linux go build -a -o dami-api .

FROM rabbitmq:3.12.0-alpine
WORKDIR /dami
COPY --from=build-stage /dami .

#mesmo criando um certificado, não iremos usá-lo
#TODO: resolver o bug "bad record MAC error"/"EOF" ao fazer um request https
RUN openssl req -x509 -out localhost.crt -keyout localhost.key \
      -newkey rsa:2048 -nodes -sha256 \
      -subj '/CN=dami' && \
    mv localhost.crt certs/ && \
    mv localhost.key certs/