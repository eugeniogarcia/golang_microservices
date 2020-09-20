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

Instalar la librería para go de AMQP - servira para Rabbit o para cualquier otro producto que implemente el protocolo AMQP:

```ps
go get -u github.com/streadway/amqp
```

```ps
github.com/mitchellh/mapstructure
```
### Usar RabbitMQ
Ejecutar una imagen con RabbitMQ, incluyendo la consola de gesión:

```ps
docker run --detach --name rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

La consola de gestión se puede abrir en `http:\\localhost:15672`. El usuario y contraseña son `guest`.

## Kafka

Hay varias librerias para trabajar con Kafka. La que vamos a usar es la de `Shopify`. Hay otras como la de `Confluent`:

```ps
go get github.com/Shopify/sarama
```

## docker-compose

Al hacer si se arroja este error:

```ps
docker-compose up -d

docker.credentials.errors.InitializationError: docker-credential-gcloud not installed or not available in PATH
```

Hay que eliminar el bloque `credHelpers` del docker config `C:\Users\Eugenio\.docker\config.json`. Estaba así:

```json
{
	"auths": {
		"https://index.docker.io/v1/": {}
	},
	"HttpHeaders": {
		"User-Agent": "Docker-Client/19.03.12 (windows)"
	},
	"credsStore": "desktop",
	"credHelpers": {
		"asia.gcr.io": "gcloud",
		"eu.gcr.io": "gcloud",
		"gcr.io": "gcloud",
		"marketplace.gcr.io": "gcloud",
		"staging-k8s.gcr.io": "gcloud",
		"us.gcr.io": "gcloud"
	},
	"stackOrchestrator": "swarm"
}
```

y lo deje así:

```json
{
	"auths": {
		"https://index.docker.io/v1/": {}
	},
	"HttpHeaders": {
		"User-Agent": "Docker-Client/19.03.12 (windows)"
	},
	"credsStore": "desktop",
	"stackOrchestrator": "swarm"
}
```json