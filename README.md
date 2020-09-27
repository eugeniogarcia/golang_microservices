# Setup

## Instala gorilla

```ps
go get github.com/gorilla/mux
```

Instala los handlers para soportar CORS:

```ps
go get github.com/gorilla/handlers
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

Podemos ejecutar Kafka en un contenedor. Esta imagen incluye zookeeper:

```ps
docker run -d --name kafka1 -p 9092:9092 -p 2181:2181 --env ADVERTISED_HOST=172.25.240.1 --env ADVERTISED_PORT=9092 spotify/kafk
```

Para interactuar con Kafka hay varias librerias disponibles. La que vamos a usar es la de `Shopify`. Hay otras como la de `Confluent`:

```ps
go get github.com/Shopify/sarama
```

## docker-compose

Para arrancar hacemos:

```ps
docker-compose up -d
```

Para parar:

```ps
docker-compose down
```

### Comentario

#### Red

Definimos una red:

```yaml
version: "3"
networks:
  myevents:
```

Podemos ver los detalles de esta red una vez hemos iniciado los contenedores:

```ps
docker network list

NETWORK ID          NAME                DRIVER              SCOPE
2d34d8adcbd2        bridge              bridge              local
9b2d1ff4540f        host                host                local
28795e0c9cfe        none                null                local
dcbfd362f99b        src_myevents        bridge              local
```

y los detalles:

```ps
docker network inspect src_myevents

[
    {
        "Name": "src_myevents",
        "Id": "dcbfd362f99bad7e94b65fccf0f49238253312c3d5a777fb9bcc38f5f5d58085",
        "Created": "2020-09-20T12:18:36.4813258Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.20.0.0/16",
                    "Gateway": "172.20.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": true,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "00b7a423836d3aeb59f6c70c57e3573295e20b14ec65da501b625172a12fa5ac": {
                "Name": "src_rabbitmq_1",
                "EndpointID": "f2f0c4d106ca4cd262693b614c5eeb070018bf0ec7732f96daa0045e5df2d902",
                "MacAddress": "02:42:ac:14:00:02",
                "IPv4Address": "172.20.0.2/16",
                "IPv6Address": ""
            },
            "1374463af4bd0a74829ed4c6929c3df391fc648cf25d055f2ecee818c1362d80": {
                "Name": "src_events-db_1",
                "EndpointID": "5f6dd026968675563433a8c9a1ce56a7d576e0e5d0792c1d6bf95d79fd422088",
                "MacAddress": "02:42:ac:14:00:04",
                "IPv4Address": "172.20.0.4/16",
                "IPv6Address": ""
            },
            "1b1fac07e04e56a8f436b787dcbb8a8868d7f890e84f1e984e111cf4ea026dcd": {
                "Name": "src_bookings-db_1",
                "EndpointID": "e6aebc94032266de2a7219b140541c04ede45a424e0dff2d660602db0361d139",
                "MacAddress": "02:42:ac:14:00:03",
                "IPv4Address": "172.20.0.3/16",
                "IPv6Address": ""
            },
            "a5979dcda9feb5acabcd44cd06ea8d3b4855fd8eb051eccd637d615b9bf586ad": {
                "Name": "src_events_1",
                "EndpointID": "82e244422e26e270a8f64704b1456672f2c31a8a7fa7c367ac1c294920396e87",
                "MacAddress": "02:42:ac:14:00:06",
                "IPv4Address": "172.20.0.6/16",
                "IPv6Address": ""
            },
            "e945843ad87435e1c779f62debd6e84cc5e45bde1abd7631a505012064ae1ee5": {
                "Name": "src_bookings_1",
                "EndpointID": "b2313c3e540155e9095feee5a363c91ce093c264c35c1ad4df4df2fac204f8c1",
                "MacAddress": "02:42:ac:14:00:05",
                "IPv4Address": "172.20.0.5/16",
                "IPv6Address": ""
            }
        },
        "Options": {},
        "Labels": {
            "com.docker.compose.network": "myevents",
            "com.docker.compose.project": "src",
            "com.docker.compose.version": "1.27.3"
        }
    }
]
```

Podemos ver la subred que se ha creado:

```yaml
"Subnet": "172.20.0.0/16",
"Gateway": "172.20.0.1"
```

y podemos ver la ip de cada contenedor. Por ejemplo `src_rabbitmq_1`:

```yaml
            "00b7a423836d3aeb59f6c70c57e3573295e20b14ec65da501b625172a12fa5ac": {
                "Name": "src_rabbitmq_1",
                "EndpointID": "f2f0c4d106ca4cd262693b614c5eeb070018bf0ec7732f96daa0045e5df2d902",
                "MacAddress": "02:42:ac:14:00:02",
                "IPv4Address": "172.20.0.2/16",
                "IPv6Address": ""
            },
```

#### Servicios

Definimos cinco servicios:

```yaml
services:
  rabbitmq:
    image: rabbitmq:3-management
    ports:
      - 15672:15672
    networks:
      - myevents

  events-db:
    image: mongo
    ports:
      - 8006:27017
    networks:
      - myevents

  bookings-db:
    image: mongo
    ports:
      - 8008:27017
    networks:
      - myevents

  events:
    build:
      context: .
      dockerfile: dockerfile.eventservice
    ports:
      - 8181:8181
      - 9100:9100
    depends_on:
      - rabbitmq
      - events-db
    environment:
      - AMQP_BROKER_URL=amqp://guest:guest@rabbitmq:5672/
      - MONGO_URL=mongodb://events-db/events
    networks:
      - myevents

  bookings:
    build:
      context: .
      dockerfile: dockerfile.bookingservice
    ports:
      - 8282:8181
      - 9101:9100
    depends_on:
      - rabbitmq
      - bookings-db
    environment:
      - AMQP_BROKER_URL=amqp://guest:guest@rabbitmq:5672/
      - MONGO_URL=mongodb://bookings-db/bookings
    networks:
      - myevents
```

Todos los servicios usan la misma red, de modo que cada container esta en la misma red local. 

```yaml
services:
  rabbitmq:
    networks:
      - myevents

  events-db:
    networks:
      - myevents

  bookings-db:
    networks:
      - myevents

  events:
    networks:
      - myevents

  bookings:
    networks:
      - myevents
```

Tenemos dos contenedores con Mongo. Cada microservicio tiene su propio almacenamiento en mongo. Los contenedores con los microservices los construimos. Indicamos:

- El dockerfile que tenemos que usar para construirlo
- Los puertos que vamos a exponer
- El contenedor del que dependen. Esto significa que nuestro contenedor utiliza otros contenedores, y de esta manera se orquestan
- Especificamos algunas variables de entorno, que luego serán usadas por el microservicio para configurarse. Estamos especificando los datos de conexión de Rabbit, y de mongo. __Notese como usamos el nombre del contenedor de rabbit y de mongo con DNS name__. En este ejemplo, por ejemplo, usamos el contenedor `events-db`:

```yaml
  events:
    build:
      context: .
      dockerfile: dockerfile.eventservice
    ports:
      - 8181:8181
      - 9100:9100
    depends_on:
      - rabbitmq
      - events-db
    environment:
      - AMQP_BROKER_URL=amqp://guest:guest@rabbitmq:5672/
      - MONGO_URL=mongodb://events-db/events
    networks:
      - myevents
```


### Error?

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
```

### Dangling images

Despues de compilar tendremos una serie de imagenes temporales que podemos eliminar:

```ps
docker image prune
```

## Glide

Glide es un gestor de paquetes para go. Utiliza un archivo llamado `glide.yaml` para guardar todas las dependencias que necesitamos. Podemos hacer que Glide inspeccione nuestro repositorio y nos cree el archivo automáticamente:

```ps
glide init
```

Tenemos que revisar el archivo `glide.yaml` para eliminar los paquetes que son locales, que hemos definido en el propio repositorio. Una vez hecho esto, podemos descargar las versiones concretas con el siguiente comando:

```ps
glide install
```

Glide creara un directorio `vendor` y descargara todas las versiones que necesitamos. Al compilar, las versiones encontradas en el directorio `vendor` toman precedencia sobre versiones encontradas en cualquier otro sitio. De esta forma nos podemos asegurar de que estemos usando las versiones especificas de los paquetes que se han indicado en `glide.yaml`. Podemos bajar una versión actualizada con:

```ps
glide update
```

## Kubernetes

### Setup

Vamos a utilizar el `ingress`. En driver the docker no lo soporta, así que cuando arranquemos `minikube` tendremos que especificar que se use una virtual machine - `hyperv`:

```ps
minikube start --vm=true --memory=6000 --cpus=4
```

Habilitamos el `ingress`:

```ps
minikube addons enable ingress
```

La ip del ingress será:

```ps
minikube ip

192.168.1.147
```

Crearemos en hosts files las siguientes entradas:

```txt
192.168.1.147	api.myevents.example
192.168.1.147	www.myevents.example
```

Podemos usar el contenedor de docker incluido en minikube:

```ps
minikube docker-env
```

```ps
& minikube -p minikube docker-env | Invoke-Expression
```

Veamos las imagenes disponibles en el registro:

```ps
docker images

REPOSITORY                                TAG                 IMAGE ID            
k8s.gcr.io/kube-proxy                     v1.19.2             d373dd5a8593        
k8s.gcr.io/kube-controller-manager        v1.19.2             8603821e1a7a        
k8s.gcr.io/kube-apiserver                 v1.19.2             607331163122        
k8s.gcr.io/kube-scheduler                 v1.19.2             2f32d66b884f        
gcr.io/k8s-minikube/storage-provisioner   v3                  bad58561c4be        
k8s.gcr.io/etcd                           3.4.13-0            0369cf4303ff        
kubernetesui/dashboard                    v2.0.3              503bc4b7440b        
k8s.gcr.io/coredns                        1.7.0               bfe3a36ebd25        
kubernetesui/metrics-scraper              v1.0.4              86262685d9ab        
k8s.gcr.io/pause                          3.2                 80d28bedfe5d        
```

Podemos crear nuestras imagenes:

```ps
docker build -f dockerfile.eventservice  -t myevents/eventservice .
```

```ps
docker build -f dockerfile.bookingservice  -t myevents/bookingservice .
```

```ps
docker build -f Dockerfile.frontend  -t myevents/frontend .
```

### Statefulsets

Las bases de datos y rabbit necesitan un almacenamiento en disco, por este motivo creareamos unos `statefulsets`:

```ps
kubectl get statefulsets

NAME          READY   AGE
bookings-db   1/1     8m3s
events-db     1/1     4m55s
rmq           1/1     2m13s
```

Podemos ver los pods que se han creado. Observese como al estar asociados a un `statefulset`, la identidad de cada pod es fija:

```ps
kubectl get po

NAME            READY   STATUS    RESTARTS   AGE
bookings-db-0   1/1     Running   0          8m30s
events-db-0     1/1     Running   0          5m22s
rmq-0           1/1     Running   0          2m40s
```

Veamos los `pv` que se han creado:

```ps
kubectl get pv

NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                        STORAGECLASS   REASON
pvc-57bf91a3-e565-4c24-9005-434db948569b   1Gi        RWO            Delete           Bound    default/data-events-db-0     standard
pvc-639291ea-c753-4482-98ff-175039a2a844   1Gi        RWO            Delete           Bound    default/data-bookings-db-0   standard
pvc-7b61b373-7a96-4bad-a178-88d2df51372d   1Gi        RWO            Delete           Bound    default/data-rmq-0           standard
```

Y los `pvc`:

```ps
PS [EUGENIO] >kubectl get pvc
NAME                 STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS
data-bookings-db-0   Bound    pvc-639291ea-c753-4482-98ff-175039a2a844   1Gi        RWO            standard
data-events-db-0     Bound    pvc-57bf91a3-e565-4c24-9005-434db948569b   1Gi        RWO            standard
data-rmq-0           Bound    pvc-7b61b373-7a96-4bad-a178-88d2df51372d   1Gi        RWO            standard
```

Podemos observar como el pv esta bindeado a un pvc.

Hemos creado tambien unos servicios para poder acceder a la base de datos:

```ps
kubectl get svc

NAME          TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)     AGE
bookings-db   ClusterIP   None         <none>        27017/TCP   18m
events-db     ClusterIP   None         <none>        27017/TCP   9m28s
```

### Deployments

Los pods con los dos microservicios y con el frontend son stateless, así que los desplegamos con un `Deployment`. Podemos especificar el número de réplicas.

```ps
kubectl get deployment

NAME       READY   UP-TO-DATE   AVAILABLE   AGE
bookings   2/2     2            2           3m26s
events     2/2     2            2           110s
frontend   1/1     1            1           11s
```

En el servicio `bookings` y `events` hemos puesto dos instancias de pods. Los pods, a diferencia de con los `statefulsets` tienen una identidad variable:

```ps
kubectl get po

NAME                        READY   STATUS    RESTARTS   AGE
bookings-6988d6bf6f-4tlm4   1/1     Running   0          3m46s
bookings-6988d6bf6f-nm2dc   1/1     Running   0          3m46s
bookings-db-0               1/1     Running   0          27m
events-645697446c-5zj95     1/1     Running   0          2m10s
events-645697446c-nmczv     1/1     Running   0          2m10s
events-db-0                 1/1     Running   0          24m
frontend-6c987f9db-h6c6d    1/1     Running   0          31s
rmq-0                       1/1     Running   0          21m
```

Hemos creado también unos servicios:

```ps
kubectl get svc

NAME          TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
amqp-broker   ClusterIP   None            <none>        5672/TCP       9m2s
bookings      ClusterIP   10.99.174.139   <none>        80/TCP         3m32s
bookings-db   ClusterIP   None            <none>        27017/TCP      33m
events        ClusterIP   10.98.223.101   <none>        80/TCP         116s
events-db     ClusterIP   None            <none>        27017/TCP      24m
frontend      NodePort    10.109.132.34   <none>        80:31729/TCP   17s
kubernetes    ClusterIP   10.96.0.1       <none>        443/TCP        84m
```

Podemos ver que los microservicios los hemos creado como `ClusterIP`. El servicio para el frontend lo hemos creado como `NodePort`.



### Ingress

Usamos nginx como ingress. Usamos una anotación para hacer el rewrite del path. El path es un regex.

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: myevents2
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /$1
```

Lo que estamos diciendo aquí es que el recurso que se enviará al backend es `/$1`, dodne `$1` es la primera agrupación. El path es una expresión `regex`, cuando decimos primera agrupación, nos referimos por tanto a la primera agrupación del regex del path:

```yaml
spec:
  rules:
  - host: api.myevents.example
    http:
      paths:
      - path: /bookings/(.*)
        pathType: Prefix
        backend:
          service:
            name: bookings
            port:
              number: 80      
```

`$1` es lo que sigue a `/bookings`. Esto es, si hacemos una petición a `http://api.myevents.example/bookings/event/12345/bookings`, la petición que irá al backend es `http://api.myevents.example/event/12345/bookings`.

# Arquitectura de la Aplicación

## configuration

El paquete `configuration` proporciona un método llamado `ExtractConfiguration` que retorna un tipo `ServiceConfig` con los detalles de configuración de: 

- la base de datos. Tipo - ej. Mongo - y la connection string
- Endpoints del servicio (http y https)
- Tipo de broker - amqp o Kafka
- Brokers. En el caso de amqp un solo broker. En el caso de Kafka un slice de brokers

La configuración se lee de un archivo `config.json` o sino se toman valores por defecto

## msgqueue

Este paquete contiene todo lo necesario para trabajar con un broker amqp o kafka.

- Define una interface `Event` que deben implementar todos los mensajes. Esta interface especifica un método que nos devolvera el topico del mensaje `EventName()` 
- Define una interface `EventEmitter` que debe implementar quien quiera publicar un mensaje
- Define una interface `EventListener` que debe implementar quien quiera escuchar mensajes:
    - El método `Listen(events ...string)` devuelve un canal por el que llegarán los mensajes. El método admite un número variable de argumentos en los que poder indicar cuales son los topicos que nos interesan
    - El método `Mapper()` devuelve el `EventMapper`. El `EventMapper` nos permite convertir el mensaje recibido en el tipo/clase correspondiente
- Define el interface `EventMapper`. EL interface tiene un solo método `MapEvent(string, interface{}) (Event, error)`. El método tiene como argumentos el topico del interface, y el payload - del tipo interface{}-

Además de todos estos tipos, se incluyen en este paquete dos implementaciones del interface `EventMapper`:
- StaticEventMapper. 
- DynamicEventMapper

### builder

Este subpaquete incluye un helper que nos permite construir un `msgqueue.EventListener`. Si la variable de entorno `AMQP_URL` esta presente crea un listener AMQP. En caso de no estar presente, comprueba si tenemos la variable de entorno `KAFKA_BROKERS` y si la tenemos, crea un listener de Kafka. Sino tenemos ninguna de las dos, devuelve un errror.

### amqp

Incluye los helpers para construir un emiter o un listener de AMQP. Tanto el listener como el emiter tendrán los metodos, helpers necesarios para trabajar con AMQP, así como los datos que definen un listener o un emiter.

Un listener tiene la siguiente estructura:

```go
type amqpEventListener struct {
	connection *amqp.Connection
	exchange   string
	queue      string
	mapper     msgqueue.EventMapper
}
```

Guarda la conexión con AMQP, el nombre del exchange y de la cola, así como el mapper que hay que utilizar para mapear el evento a su tipo. El Emiter tiene la siguiente estructura:

```go
type amqpEventEmitter struct {
	connection *amqp.Connection
	exchange   string
	//El canal de eventos
	events chan *emittedEvent
}
```

Guarda la conexión con AMQP, y el nombre del exchange.

#### Emiter

Para crear un emiter se toma una conexión, se obtiene un Channel, se crea el exchange en el Channel, y luego se cierra el Channel. En el objeto guardamos la conexión, y el nombre del exchange. 

```go
channel, err := a.connection.Channel()
```

```go
err = channel.ExchangeDeclare(a.exchange,
		"topic", //exchange de tipo topic
		true,    //persistent - No se elimina cuando el broker se cierra
		false,   // autodelete - no se elimina cuando el canal se cierre
		false,   //internal - no es interno. Se podrá publicar mensajes
		false,   //noWait - esperamos a que el exchange nos confirme que el mensaje se haya publicado
        nil)
```

Cuando queramos publicar un mensaje, tomamos la conexión de la estructura de datos, obtenemos un Channel, y publicamos el mensaje en el exchange que esta indicado en la estructura de datos.

```go
channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()
```

```go
err = channel.Publish(a.exchange,
    event.EventName(), //routing
    false,             //comprobamos que el mesanje haya llegado a una cola por lo menos
    false,             //espera hasta que el mensaje haya sido consumido por al
    msg)               //Mensaje
```

Es decir, que el Channel se obtiene de la conexión cada vez que necesitamos publicar un mensaje, no se reutiliza - de echo no lo guardamos en la estructura. La razón por la que no se reutiliza es que no es threadsafe.

El mensaje AMQP tiene la siguiente pita:

```go
msg := amqp.Publishing{
    Headers:     amqp.Table{"x-event-name": event.EventName()},
    ContentType: "application/json",
    Body:        jsonBody,
}
```

El Evento que publicamos se utiliza para "rellenar" dos propiedades del mensaje. Por un lado el EventName - recordemos que hemos definido que el evento implemente el interface `Event` - lo usamos en el `Headers`, y por otro lado el `Body` lleva el mensaje en bruto.

#### Listener

Similar al caso anterior. Con la conexión obtenemos un Channel, que emplearemos para declarar el Exchange y la Cola:

```go
channel, err := a.connection.Channel()
defer channel.Close()
```

```go
//Crea el exchange
err = channel.ExchangeDeclare(a.exchange, "topic", true, false, false, false, nil)
if err != nil {
    return err
}

//Crea la cola
_, err = channel.QueueDeclare(a.queue, true, false, false, false, nil)
if err != nil {
    return fmt.Errorf("could not declare queue %s: %s", a.queue, err)
}
```

La conexión, exchange y la cola se guardan en la estructura del Listener. No hemos hecho ningun binding entre el exchange y la cola. 

En `func (l *amqpEventListener) Listen(eventNames ...string) (<-chan msgqueue.Event, <-chan error, error)` nos encargamos de escuchar mensajes. El método retorna un channerl con los mensajes que llegan en la cola. El argumento `eventNames` lo usamos para crear el binding de la cola con el exchange:

```go
	for _, event := range eventNames {
		if err := channel.QueueBind(l.queue, event, l.exchange, false, nil); err != nil {
			return nil, nil, fmt.Errorf("could not bind event %s to queue %s: %s", event, l.queue, err)
		}
	}
```

Una vez establecido el binding de la cola con el exchange, con `Consume` obtenemos el channel: 

```go
msgs, err := channel.Consume(l.queue, //Nombre de la cola
    "",    //el identificador del cliente. Si no indicamos nada, se autogenera uno
    false, //No hacemos auto acknowledge
    false, //No es exclusivo, así que puede haber varios consumidores en la misma cola
    false,
    false,
    nil)
```

### kafka

En este paquete implementamos un emiter y un listener con Kafka, usando la librería de `Shopify`.

#### Emiter

La estructura del emiter es:

```go
type kafkaEventEmitter struct {
	producer sarama.SyncProducer
}
```

Guardamos un `SyncProducer` - la otra posibilidad sería guardar el productor asíncrono. 

Creamos un canal. Enviaremos la conexión con Kafka por el canal:

```go
result := make(chan sarama.Client)
```

Creamos la conexión/cliente - en Kafka usamos un cliente para publicar mensajes. Especificamos los brokers y una configuración. En este caso no cambiamos los valores por defecto de la configuración:

```go
config := sarama.NewConfig()
conn, err := sarama.NewClient(brokers, config)
if err == nil {
    log.Println("connection successfully established")
    result <- conn
    return
}
```

Los brokers que usamos los tomamos de la variable de entorno:

```go
brokers := []string{"localhost:9092"}

if brokerList := os.Getenv("KAFKA_BROKERS"); brokerList != "" {
    brokers = strings.Split(brokerList, ",")
}
```

utilizaremos la conexión/cliente para crear el producer - en este caso síncrono. Guardaremos el producer en la estructura:

```go
producer, err := sarama.NewSyncProducerFromClient(client)
if err != nil {
    return nil, err
}

emitter := kafkaEventEmitter{
    producer: producer,
}
```

Cuando queramos publicar un mensaje usaremos `Emit`. Preparamos el mensaje. Observemos que tiene la misma estructura que el mensaje AMQP que definimos antes:

```go
jsonBody, err := json.Marshal(messageEnvelope{
    evt.EventName(),
    evt,
})
```

Finalmente publicamos el mensaje:

```go
msg := &sarama.ProducerMessage{
    Topic: "events",
    Value: sarama.ByteEncoder(jsonBody),
}
```

#### Listener

La estructura del listener es como sigue:

```go
type kafkaEventListener struct {
	consumer   sarama.Consumer
	partitions []int32
	mapper     msgqueue.EventMapper
}
```

Tenemos por un lado el consumidor, y por otro lado la partición a la que nos vamos a conectar. Destacar que en la librería de `Confluent` no hace falta especificar la partición, pero en la de `Shopify` nos tenemos que encargar, que le vamos ha hacer!!.

Obtenemos los brokers, de la misma forma que hicimos con el emiter:

```go
brokers := []string{"localhost:9092"}

if brokerList := os.Getenv("KAFKA_BROKERS"); brokerList != "" {
    brokers = strings.Split(brokerList, ",")
}
```

Prepara la lista de particiones. Sino especificamos ninguna consumira de cualquier partición:

```go
partitions := []int32{}

if partitionList := os.Getenv("KAFKA_PARTITIONS"); partitionList != "" {
    partitionStrings := strings.Split(partitionList, ",")
    partitions = make([]int32, len(partitionStrings))

    for i := range partitionStrings {
        partition, err := strconv.Atoi(partitionStrings[i])
        if err != nil {
            return nil, err
        }
        partitions[i] = int32(partition)
    }
}
```

Guardamos el consumer y la relación de particiones en la que estamos interesados en la estructura del listener:

```go
consumer, err := sarama.NewConsumerFromClient(client)

listener := &kafkaEventListener{
    consumer:   consumer,
    partitions: partitions,
    mapper:     msgqueue.NewEventMapper(),
}
```

Finalmente implementamos el método `Listen(events ...string) (<-chan msgqueue.Event, <-chan error, error)` que usaremos para consumir los mensajes. Los mensajes los publicaremos en un canal. El topico en este caso esta hardcodeado:


```go
topic := "events"
results := make(chan msgqueue.Event)
```

Si no tenemos particiones consumimos sin especificarlas:

```go
partitions := k.partitions
if len(partitions) == 0 {
    partitions, err = k.consumer.Partitions(topic)
    if err != nil {
        return nil, nil, err
    }
}
```

Pero si tenemos particiones lanzamos una go rutina por cada una de las particiones:

```go
for _, partition := range partitions {

    pConsumer, err := k.consumer.ConsumePartition(topic, partition, 0)
    if err != nil {
        return nil, nil, err
    }

    go func() {
        for msg := range pConsumer.Messages() {
            body := messageEnvelope{}
            err := json.Unmarshal(msg.Value, &body)
            if err != nil {
                errors <- fmt.Errorf("could not JSON-decode message: %v", err)
                continue
            }

            event, err := k.mapper.MapEvent(body.EventName, body.Payload)
            if err != nil {
                errors <- fmt.Errorf("could not map message: %v", err)
                continue
            }

            results <- event
        }
    }()
```

## contracts

En este paquete definimos el modelo de datos. Tanto el cliente como el consumidor compartiran este modelo.

En el contrato implementamos la estructura y los interfaces. Por ejemplo, `Event`:

```go
type EventCreatedEvent struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	LocationID string    `json:"location_id"`
	Start      time.Time `json:"start_date"`
	End        time.Time `json:"end_date"`
}

// EventName returns the event's name
func (c *EventCreatedEvent) EventName() string {
	return "eventCreated"
}
```

`EventName()` devuelbe el routing key que usaremos cuando se publique el mensaje.

## eventservice

Lee la configuración:

```go
confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
flag.Parse()
//extract configuration
config, _ := configuration.ExtractConfiguration(*confPath)
```

Comprueba que broker de mensajería esta configurado:

```go
switch config.MessageBrokerType {
case "amqp":
  conn, err := amqp.Dial(config.AMQPMessageBroker)
  if err != nil {
    panic(err)
  }

  eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
  if err != nil {
    panic(err)
  }
case "kafka":
  conf := sarama.NewConfig()
  conf.Producer.Return.Successes = true
  conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
  if err != nil {
    panic(err)
  }

  eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
  if err != nil {
    panic(err)
  }
```

Si AMQP se conecta con el broker y configura el emiter:

```go
conn, err := amqp.Dial(config.AMQPMessageBroker)

eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
```

Configura la capa de persistencia:

```go
dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
```

Con la capa de persistencia y el emiter configurados, se arranca el servidor http:

```go
httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPint, dbhandler, eventEmitter)
```

### Servidor http

Para implementar el servidor http usamos la librería `github.com/gorilla/mux`. Gorilla necesita que se definan handlers que se asociarán a los recursos http:

```go
handler := newEventHandler(dbHandler, eventEmitter)
```

La librería usa routers. Creamos un router:

```go
r := mux.NewRouter()
```  

El router se configura definiendo los recursos, verbos y handlers:

- recurso `/events`, con verbo `GET` y dos path parameters, es gestionando por `handler.findEventHandler`:

```go
eventsrouter := r.PathPrefix("/events").Subrouter()
eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
```

```go
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
```

```go
eventsrouter.Methods("GET").Path("/{eventID}").HandlerFunc(handler.oneEventHandler)
```

```go
eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
```

Una vez que se han definido todas las rutas, procedemos a iniciar la escucha. Como queremos en este caso escuchar por http y https, lanzamos dos go rutinas:

```go
httpErrChan := make(chan error)
httptlsErrChan := make(chan error)

server := handlers.CORS()(r)

go func() {
  httptlsErrChan <- http.ListenAndServeTLS(tlsendpoint, "cert.pem", "key.pem", server)
}()

go func() {
  httpErrChan <- http.ListenAndServe(endpoint, server)
}()
```

### Handlers

Tipicamente tomaremos los parametros y headers del request

```go
func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		fmt.Fprint(w, `No search criteria found, you can either search by id via /id/4
						to search by name via /name/coldplayconcert`)
		return
	}
```

Con los parametros ya procedemos a ejecutar la lógica, publicar mensajes si procede o actuar sobre la base de datos:

```go
		event, err = eh.dbhandler.FindEventByName(searchkey)
```

La respuesta de la api se tiene que construir:

```go
if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error occured %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)
}
```

## bookingservice

Como en el eventservice lee la configuración y comprueba que broker de mensajería esta configurado. En este servicio vamos a escuchar - listener - y publicar mensajes - emiter.

```go
switch config.MessageBrokerType {
case "amqp":
  conn, err := amqp.Dial(config.AMQPMessageBroker)
  panicIfErr(err)

  //Escuchamos por eventos events, y booking
  eventListener, err = msgqueue_amqp.NewAMQPEventListener(conn, "events", "booking")
  panicIfErr(err)

  //A su vez publicamos events
  eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
  panicIfErr(err)
case "kafka":
  conf := sarama.NewConfig()
  conf.Producer.Return.Successes = true
  conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
  panicIfErr(err)

  eventListener, err = kafka.NewKafkaEventListener(conn, []int32{})
  panicIfErr(err)

  eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
  panicIfErr(err)
default:
  panic("Bad message broker type: " + config.MessageBrokerType)
}
```

Notese como hemos configurado también los listeners. En el caso de AMPQ configura los bindings `events` y `booking`:

```go
eventListener, err = msgqueue_amqp.NewAMQPEventListener(conn, "events", "booking")
panicIfErr(err)
```

En el caso de Kafka escucha de todas las particiones:

```go
eventListener, err = kafka.NewKafkaEventListener(conn, []int32{})
panicIfErr(err)
```

Como este servicio consume mensajes, lanza una go rutina para consumirlos:

```ps
//Procesa los eventos en una go rutina
processor := listener.EventProcessor{eventListener, dbhandler}
go processor.ProcessEvents()
```

Por lo demás el servicio es análigo al eventservice, se configura la capa de persistencia, y se inicia el servidor.

### Servidor http

Otro estilo, pero mis cosa:

```go
r := mux.NewRouter()
r.Methods("post").Path("/events/{eventID}/bookings").Handler(&CreateBookingHandler{eventEmitter, database})

srv := http.Server{
  Handler:      handlers.CORS()(r),
  Addr:         listenAddr,
  WriteTimeout: 2 * time.Second,
  ReadTimeout:  1 * time.Second,
}

srv.ListenAndServe()
```

