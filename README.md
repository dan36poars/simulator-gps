# Simulator-gps

fake gps data produced to apache kafka. to get in real time a json latitude and longitude coordinates when has been pass a clientId and routeId parameters from specific client to system, then all the time that a new drive will want started a new move we can consume this data in our front-end.

## DevOps
infra folder

- apache kafka

## Docker commands

if you don't have a docker user group in the environment, always start typing `sudo` in each command

- start container simulator

```shell
    $ docker-compose up -d
```

- Check if working
```bash
    $ docker ps
```

```bash
    $ docker-compose ps
```

- get container shell command

simulator is the name of the your container

```bash
    $ docker exec -it simulator bash
```

- state listen a topic in kafka

readTest is the topic in this sample

```bash
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=readTest
```

> add line in **/etc/hosts** file if you receive a error like **don't resolve localhost names**
> ```
> 172.17.0.1    kafka
> ```