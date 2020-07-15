![Go](https://github.com/jparrill/pada-apps/workflows/Go/badge.svg)

# pada-apps for your i3 Bar

Let's say that you wanna create your i3 bar applications using Golang...this is my take here, Applications in Golang for your i3Bar, w/e is fine meanwhile could execute custom scripts.

## Pada Apps

All of the apps have the same behaviour un common, `Red == bad`, `Green == good` and at some point `Yellow == not found` this last one usually comes with a description message

- **Date:** It will return the date as fancy prompt
```
λ pada-apps/bin git:(master) [⇡] via 🐹 v1.14.3 ./date       
 Tue 14 Jul 2020 |  18:23
```

- **Eth:** It will return the status of the interface of the system, if it has an IP assigned, will be considered it as Up, if not, will be as Down
```
λ pada-apps/bin git:(master) [⇡] via 🐹 v1.14.3 ./eth
Usage : ./eth <interface name>
Interfaces: lo wlp2s0 ens4 virbr0 virbr0-nic virbr1 virbr1-nic tun0

## In Red because is disconnected
λ pada-apps/bin git:(master) [⇡] via 🐹 v1.14.3 ./eth ens4


## In Green because is connected
λ pada-apps/bin git:(master) [⇡] via 🐹 v1.14.3 ./eth wlp2s0


## In Yellow because does not exist
λ pada-apps/bin git:(master) [⇡] via 🐹 v1.14.3 ./eth wlp2s 
No Int: wlp2s %
```

- **Connection:** It will return the icon with the status of the internet connectivity
```
## In Green because is connected
λ pada-apps/bin git:(master) [⇡] via 🐹 v1.14.3 ./connection


## In Red when it's disconnected
λ pada-apps/bin git:(master) [⇡] via 🐹 v1.14.3 ./connection

```

## Build the apps

- Use the `Makefile` to generate the binaries

```
λ pada-apps git:(master) [⇡] via 🐹 v1.14.3 make
mkdir -p "/home/jparrill/go/src/github.com/jparrill/pada-apps/bin"
building: cmd/connection/main.go
building: cmd/date/main.go
building: cmd/eth/main.go
building: cmd/hello/main.go

λ pada-apps git:(master) [⇡!] via 🐹 v1.14.3 ls bin
connection  date  eth  hello
```

- To clean them up just
```
λ pada-apps git:(master) [⇡!] via 🐹 v1.14.3 make clean
go clean
rm -rf "/home/jparrill/go/src/github.com/jparrill/pada-apps/bin"

λ pada-apps git:(master) [⇡!] via 🐹 v1.14.3 ls bin
ls: cannot access 'bin': No such file or directory
```


## Testing

- To launch the tests just execute this make target
```
make test
```
