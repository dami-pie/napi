FROM golang AS builder
LABEL authors="sevlak0ff"
WORKDIR /dami
COPY . .

#creating .env file
RUN echo 'PORT=:443' >> .env
RUN echo 'OTP_SECRET="damitest"' >> .env
RUN echo 'JWT_SECRET="brunoalvirubro"' >> .env
RUN echo 'DEVICE="http://device:80"' >> .env #device é o hostname do container do device (duh)
RUN echo 'CERTIFICATE="certs/localhost.crt"' >> .env
RUN echo 'PRIVATE_KEY="certs/localhost.key"' >> .env

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o dami-api .

FROM rabbitmq:3.12.0-alpine
WORKDIR /dami
COPY --from=builder /dami .

#mesmo criando um certificado, não iremos usá-lo
#TODO: resolver o bug "bad record MAC error"/"EOF" ao fazer um request https
RUN openssl req -x509 -out localhost.crt -keyout localhost.key \
      -newkey rsa:2048 -nodes -sha256 \
      -subj '/CN=dami' && \
    mv localhost.crt certs/ && \
    mv localhost.key certs/

ENTRYPOINT ["./dami-api"]