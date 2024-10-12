# monochan
## A monolithic Imageboard for learning purposes. 

Changelog:
```
- Updated README.md
- Removed Herobrine
```

Requisites:
```
- Golang (v1.23.1)
- Docker (v27.3.1)
- Python (v3.10.12)
```

<br/>
<br/>
<br/>
<br/>
<br/>

## How to install everything on your machine
### (Guide for Ubuntu distros. If you don't use Ubuntu, you may suffer a bit to make everything work).
#### Step 1: Remove all confliting unofficial docker packages.
```
$ for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done
$ sudo apt-get purge docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-ce-rootless-extras
```
#### Step 2: Delete anything related to Docker (Optional, if you are not using docker, but can make you free from headaches).
```
$ sudo rm -rf /var/lib/docker
$ sudo rm -rf /var/lib/containerd
```
#### Step 3: Install the official docker package (READ CAREFULLY BELOW).
Run this:
```
$ sudo apt-get update
$ sudo apt-get install ca-certificates curl
$ sudo install -m 0755 -d /etc/apt/keyrings
$ sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
$ sudo chmod a+r /etc/apt/keyrings/docker.asc
```
Then, for **PURE Ubuntu** (If you use any Ubuntu-based distro like Linux Mint, just ignore this), just run:
```
$ echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
$ sudo apt-get update
```
BUT, for **Ubuntu-based distros** (E.g., Linux Mint), you should use THIS command instead:
```
$ echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$UBUNTU_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
$ sudo apt-get update
```
(Note the tiny difference between both scripts above).
Then, FINALLY, run this:
```
$ sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
$ docker version
```
If you see the Golang version on your screen, everything is installed sucessfully. Now, let's go to the next step.

#### Step 4: Remove Golang package (Optional if you are not using Go, but can make you free from headaches).
```
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz
```
#### Step 5: Download Golang.
[Click here to select your Golang package](https://go.dev/dl/) and download it. After this, extract the file you downloaded into `/usr/local`, creating then a fresh Go tree in /usr/local/go.
**After this**, just run:
```
$ export PATH=$PATH:/usr/local/go/bin
$ go version
```
If you see the Golang version on your screen, everything is installed sucessfully.
#### Step 6: Download Python.
```
$ sudo apt install python3
```
Yeah, just this. Cool.
#### Step 7: No step 7.
Yeah, It's all.
<br/>
<br/>
<br/>
<br/>
<br/>

### How to setup everything
#### Step 1: Download monochan.
```
$ cd PATH/THAT/YOU/WANT/TO/RUN/MONOCHAN
$ git clone https://github.com/KnowledgeEnjoyer/monochan
```
#### Setup 2: Run Docker Daemon.
```
$ sudo dockerd
```
And with another Terminal tab...:
#### Setup 3: Run the "start" python script.
```
$ cd PATH/WHERE/MONOCHAN/IS/monochan
$ ./start
```
#### Setup 4: Install and setup the cow.
```
$ sudo apt install cowsay
$ cowsay "Eh Th-Th-Th-Th-Th-Th-Th-That's all folks!"
$ make-me a sanduich
$ sudo make-me a sanduich
```
#### Step 5: Eh Th-Th-Th-Th-Th-Th-
Th-Th-That's all folks!

[1]: https://
