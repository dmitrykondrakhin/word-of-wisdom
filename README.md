# Word of Wisdom

## Task
Design and implement “Word of Wisdom” tcp server.
* TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
* The choice of the POW algorithm should be explained.
* After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
* Docker file should be provided both for the server and for the client that solves the POW challenge

## The choice of the POW algorithm
I've chosen hashcash algorithm because, as I see, this algorithm have easy implementation but good efficiency for this task. It is good way to create some CPU-work at client side, and easy way to check it at server side. Also I've worked with email marketing automatization long time, so I've heared about hashcash algorithm for SPAM protection

## Local running
up server and client
```bash
docker compose -p word-of-wisdom --env-file ./local_deploy/.env up -d --force-recreate
```
down server and client
```bash
docker compose -p word-of-wisdom down
```

## Configuration
|env|description|default value|
|---|---|---|
|SERVER_HOST|server host|127.0.0.1|
|SERVER_PORT|server port|8080|
|CLIENT_HOST|client connection host|127.0.0.1|
|CLIENT_PORT|client connection port|8080|
|REPEATED_COUNT|request counts from client to server|60|
|HASHCASH_BITS|Number of zero bits in the hashed code|20|
