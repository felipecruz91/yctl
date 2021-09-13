# yctl

[![build](https://github.com/felipecruz91/yctl/actions/workflows/build.yml/badge.svg)](https://github.com/felipecruz91/yctl/actions/workflows/build.yml)

Control your Yeelight devices from your terminal.

## Download

Target a specific version and platform for the binary. For version `v0.1.0` and linux, you would:

```cli
  curl -L https://github.com/felipecruz91/yctl/releases/download/v0.1.0/yctl -o yctl
  sudo mv yctl /usr/local/bin
```

For other platforms like Windows or MacOS choose the binary from the links [here](https://github.com/felipecruz91/yctl/releases).

## Usage

### Switch on

```cli
yctl --ip 192.168.1.56 on
```

### Switch off

```cli
yctl --ip 192.168.1.56 off
```
