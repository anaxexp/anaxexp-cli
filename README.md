# AnaxExp CLI

[![Build Status](https://travis-ci.org/anaxexp/anaxexp-cli.svg?branch=master)](https://travis-ci.org/anaxexp/anaxexp-cli)
[![Docker Pulls](https://img.shields.io/docker/pulls/anaxexp/anaxexp-cli.svg)](https://hub.docker.com/r/anaxexp/anaxexp-cli)
[![Docker Stars](https://img.shields.io/docker/stars/anaxexp/anaxexp-cli.svg)](https://hub.docker.com/r/anaxexp/anaxexp-cli)
[![Docker Layers](https://images.microbadger.com/badges/image/anaxexp/anaxexp-cli.svg)](https://microbadger.com/images/anaxexp/anaxexp-cli)

This project provides a unified command line interface to [anaxexp.com](https://anaxexp.com).

## Install

Fetch the [latest release](https://github.com/anaxexp/anaxexp-cli/releases) for your platform:

#### Linux (amd64)

```bash
export ANAXEXP_CLI_LATEST_URL=$(curl -s https://api.github.com/repos/anaxexp/anaxexp-cli/releases/latest | grep linux-amd64 | grep browser_download_url | cut -d '"' -f 4)
wget -qO- "${ANAXEXP_CLI_LATEST_URL}" | sudo tar xz -C /usr/local/bin
```

#### macOS

```bash
export ANAXEXP_CLI_LATEST_URL=$(curl -s https://api.github.com/repos/anaxexp/anaxexp-cli/releases/latest | grep darwin-amd64 | grep browser_download_url | cut -d '"' -f 4)
wget -qO- "${ANAXEXP_CLI_LATEST_URL}" | tar xz -C /usr/local/bin
```

## Usage

You can run the AnaxExp CLI in your shell by typing `anaxexp`.

### Commands

The current output of `anaxexp` is as follows:

```
CLI client for AnaxExp

Usage:
    anaxexp [command]

Available Commands:
    ci
        init ANAXEXP_INSTANCE_UUID
        run COMMAND
        build SERVICE/IMAGE
        release
        deploy
    help         Help about any command
    version      Shows AnaxExp CLI version

Flags:
      --api-key string      API key
      --api-prefix string   API prefix (default "api/v2")
      --api-proto string    API protocol (default "https")
      --dind                Docker in docker mode (for init)
  -h, --help                help for anaxexp
  -v, --verbose             Verbose output

Use "anaxexp [command] --help" for more information about a command.
```
