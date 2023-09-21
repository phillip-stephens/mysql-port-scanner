# MySQL Port Scanner

A port scanner to determine if an instance of MySQL server is running at a given IP address/port.
If there is an instance running, a sample of information included from the handshake will be printed to the console.

## Build

From the `mysql-port-scanner/` directory, run

```shell
go build .
```

## Run

```shell
Usage of ./mysql-port-scanner:
  -ip string
        the IP address this scanner should connect to (default "0.0.0.0")
  -port uint
        the port this scanner should connect to
  -help
        prints out this usage information
```

### Example

```shell
./mysql-port-scanner -ip 127.0.0.1 -port 3306
```

## Sample Output

```shell
go build .  && ./mysql-port-scanner --ip 127.0.0.1 --port 3306
MySQL instance is running on:
IP: 127.0.0.1
Port: 3306

Protocol: 10
Version: 8.1.0
Thread ID: 12
Authentication Plugin: mysql_native_password
```

## Testing

To test functionality, you can use the included `docker-compose.yml` file to start a MySQL docker container running on the default port (`3306`)

[Source of MySQL docker-compose.yml](https://hub.docker.com/_/mysql)

To start the MySQL docker container, from the `mysql-port-scanner/` directory run

```shell
docker-compose -f docker-compose.yml up
```
