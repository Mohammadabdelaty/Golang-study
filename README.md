# Here We *GO*

This is a Golang study.

### Installation
For ubuntu go to https://go.dev/doc/install
copy the download link 

```bash
wget https://go.dev/dl/go1.25.1.linux-amd64.tar.gz
sudo tar -C /usr/local/ -xzf go1.25.1.linux-amd64.tar.gz

vim .bashrc
export GOROOT=/path/to/your/go/installation
export PATH=$PATH:$GOROOT/bin

# For wsl 
source .bashrc # for ubuntu server .profile
go version
go version go1.25.1 linux/amd64
```

In packages we are trying to understand the packages in Go.

to start, you need to initialize go in this dir

```bash
go mod init firstGo
```

To include packages in `packages` dir you need to import it in `main.go` as follows:
```go
import "firstGo/packages"
```
 