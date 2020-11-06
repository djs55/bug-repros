# curl a service on the host in a loop

In one terminal, run a server on localhost:8080
```
PS > go run host/main.go
2020/11/06 11:17:58 Listening on 127.0.0.1:8080
2020/11/06 11:17:58 Hit enter to stop.
2020/11/06 11:18:06 Handled 100 responses
...
2020/11/06 11:18:55 Handled 10000 responses
```

In another terminal, run curl http://host.docker.internal:8080/ in a loop from a container:
```
PS > cd container
PS > docker build -t curl .
PS > docker run curl
...
Hello9998
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100     5  100     5    0     0   2500      0 --:--:-- --:--:-- --:--:--  2500
Hello9999
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100     5  100     5    0     0   1666      0 --:--:-- --:--:-- --:--:--  1666
Hello10000
```

