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

### Running udp.go

- Start server

```shell
go run udp.go
```

- Test connection

```shell
nc -u localhost 3000
```

| TCP                                         | UDP                                        |
| ------------------------------------------- | ------------------------------------------ |
| Connection oriented                         | Connection less                            |
| Slower, reliable, packet loss doesn't occur | Faster, unreliable, chances of packet loss |
| Header sizes are bigger                     | Header sizes are smaller                   |
| Segment sequencing and Acknowledge segments | No sequencing and No acknowledgment        |
| Emails, File Transfers, and Web Browsing    | Live Streaming, Gaming                     |

### Reference

- [TCP and UDP in GO](https://github.com/jeroendk/go-tcp-udp)
- [TCP and UDP in NodeJS along with explanation](https://www.youtube.com/watch?v=qqRYkcta6IE&list=PLQnljOFTspQUNnO4p00ua_C5mKTfldiYT&index=6)
