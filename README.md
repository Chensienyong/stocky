# Stocky

## Description

Stocky is the microservice for the Stock information.

## SLO and SLI
- Availability: 100%
- Rata-rata response time: <100ms

## Architecture Diagram

![architecture_diagram](/doc/architecture.png)

## Onboarding and Development Guide
### Prerequisite
- Git
- Go 1.9 or later

### Setup

- Install Git

  See [Git Installation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

- Install Go (Golang)

  See [Golang Installation](https://golang.org/doc/install)

- Clone this repo in your local

  If you have not set your GOPATH, set it using [this](https://golang.org/doc/code.html#GOPATH) guide.

  ```sh
  git clone git@github.com:chensienyong/stocky.git
  ```

- Go to Arcade directory, then sync the vendor file

  ```sh
  cd $GOPATH/src/github.com/chensienyong/stocky
  make mod
  ```

- Copy env.sample and if necessary, modify the env value(s)

  ```sh
  cp env.sample .env
  ```

- Run Stocky

  ```sh
  make start
  ```

- Check whether it is ran correctly. It should return `OK` message

  ```sh
  curl -X GET "http://localhost:1234/healthz"
  ```

### Development

- Make a new branch with descriptive name about the change(s) and checkout to the new branch

  ```sh
  git checkout -b branch-name
  ```

- Make your change(s) and make the test(s)

- Run the test in your local environment

  ```sh
  make test
  ```

- Commit and push your change to upstream repository

  ```sh
  git commit -m "a meaningful commit message"
  git push origin branch-name
  ```

- Open Pull Request in Repository

- Pull request should only be merged if review phase is passed

## FAQ

<details>
<summary>Can i contribute to this repo</summary>

Of course, why not?

</details>
