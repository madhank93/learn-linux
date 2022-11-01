### Running tcp.js

- Start server

```shell
node tcp.js
```

- Test connection

```shell
telnet localhost 8080
```

### Running udp.js

- Start server

```shell
node udp.js
```

- Test connection

```shell
echo "hi" | nc -w1 -u localhost 8081
```
