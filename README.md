<div align="center">
  <a href="https://github.com/cosasdepuma/Masterchef">
    <img alt="Logo" src="https://github.com/cosasdepuma/Masterchef/raw/main/.github/README/gophers/masterchef.svg" width="300"/>
  </a>
  <h1>Masterchef Framework</h1>
  <h3>👨‍🍳 The most delicious pentesting tool 👩‍🍳</h3>

  <h4>Under development...</h4>
</div>
<br/>

<a href="#">
  <img alt="Cooker" src="https://github.com/cosasdepuma/Masterchef/raw/main/.github/README/gophers/cooker.svg" height="120" align="left"/>
</a>

**Masterchef** is a graphical vulnerability exploitation and validation tool that helps you break down your penetration testing workflow by creating tasty recipes.

It is the ideal tool for getting started in the world of offensive cybersecurity. Visual, intuitive, customizable and with a lot of delicious ingredients! _So yummy!_

<br/>

## 🍱 Table of contents

- [Disclaimer](#-disclaimer)
- [Download and install](#-download-and-install)
  - [Compiled version](#-compiled-version)
  - [Go](#-go)
  - [Docker](#-docker)
  - [_Sauce_ code](#-sauce-code)
- [Getting started](#-getting-started)
- [Examples](#-examples)

## 🍅 Disclaimer

The use of the application may be a criminal act, depending on the regulations of each country.

**The author of this document is not responsible in case criminal charges are brought against any individual or corporation using the tool against the stipulated laws, as well as for damages caused by a misuse of the tool. It is the responsibility of the end user to obey all applicable laws.**

It is recommended that use be limited to controlled environments and/or penetration testing with prior approval.

## 🥤 Download and install

Download and install **Masterchef Framework** quickly with the steps described here.

### 🥡 Compiled version

Click the button below to download the **Masterchef** binary.

<div>
&nbsp;&nbsp;&nbsp;
<a href="https://github.com/cosasdepuma/Masterchef/releases/download/v0.1.0-tfg/masterchef-0.1.0.exe"><img alt="Windows" src="https://github.com/cosasdepuma/Masterchef/raw/main/.github/README/download/windows.png" width="100" /></a>
<a href="https://github.com/cosasdepuma/Masterchef/releases/download/v0.1.0-tfg/masterchef-0.1.0.dmg"><img alt="Mac" src="https://github.com/cosasdepuma/Masterchef/raw/main/.github/README/download/macos.png" width="100" /></a>
<a href="https://github.com/cosasdepuma/Masterchef/releases/download/v0.1.0-tfg/masterchef-0.1.0.elf"><img alt="Linux" src="https://github.com/cosasdepuma/Masterchef/raw/main/.github/README/download/linux.png" width="100" /></a>
<br><br>
</div>

Don't see your operating system here? Try one of the [other downloads](https://github.com/cosasdepuma/Masterchef/releases).

**Masterchef binary requires no installation or dependencies**.

### 🥣 Go

> Go 1.16+ is required

Install the application using **go get**:

```sh
go install github.com/cosasdepuma/masterchef@latest
```

### 🐟 Docker

It is possible to compile and run **Masterchef** in a container using **docker**:

```sh
# Download the container from Docker Hub
docker pull cosasdepuma/masterchef

# Run the container
docker run --name masterchef -d -p 7767:7767 cosasdepuma/masterchef
```

Alternatively, you can pull up the **Masterchef** service, as well as its documentation, using **docker-compose**:

```sh
cd masterchef/
docker-compose up
```

### 🍲 _Sauce_ code

If you are a ~~paranoid~~ _sybarite_, you can download and install the program yourself from the source code.

[Download](https://github.com/cosasdepuma/Masterchef/archive/refs/heads/main.zip) or clone the repository using **git**:

```sh
git clone https://github.com/cosasdepuma/Masterchef
```

Compile the code:

```sh
cd masterchef/

# Manually
yarn --cwd frontend install
yarn --cwd frontend run compile
cd backend/
export GOOS=linux     # windows, darwin...
export GOARCH=amd64   # 386, arm, arm64...
go build -a -ldflags="-s -w -extldflags \"-static\"" -o ../dist/masterchef main.go
cd ..
upx -9 --ultra-brute dist/masterchef # optional

# Automatically
make clean
make

# Using Docker
docker build -t cosasdepuma:masterchef .
```

## 🍴 Getting started

Run the application by double clicking or using the console:

```sh
./masterchef
```

It is possible to modify the address and the number of threads per request using **environment variables**:

```sh
export MC_ADDR=0.0.0.0:7767 # Address
export MC_THREADS=200       # Threads
```

Once the application is started, access the displayed web address. By default, it is [http://127.0.0.1:7767/](http://127.0.0.1:7767/).

If everything works correctly, the following interface will be shown:

![Interface](https://github.com/cosasdepuma/Masterchef/raw/main/.github/README/screenshot/interface.png)

## 📚 Examples

You can find recipes with examples of functionality in the [cookbook](https://github.com/MasterchefFramework/Masterchef/tree/main/cookbook) folder.
