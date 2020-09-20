# Setup

## Instala gorilla

```ps
go get github.com/gorilla/mux
```

## Instala el driver de Mongo

```ps
go get gopkg.in/mgo.v2
```

## Crea un certificado

```ps
go run C:\go\src\crypto\tls\generate_cert.go --host localhost

2020/09/19 01:12:21 wrote cert.pem
2020/09/19 01:12:21 wrote key.pem
```

Nos crea el `cert.pem` y `key.pem` en el directorio actual. Otros parametros que se pueden pasar a `generate_cert.go` son:

- `--start-date`
- `--duration`
- `--rsa-bits`
- `--help`

## RabbitMQ

Ejecutar una imagen con RabbitMQ, incluyendo la consola de gesión:

```ps
docker run --detach --name rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

La consola de gestión se puede abrir en `http:\\localhost:15672`. El usuario y contraseña son `guest`.

Instalar la librería para go de AMQP:

```ps
go get -u github.com/streadway/amqp
```

```ps
github.com/mitchellh/mapstructure
```