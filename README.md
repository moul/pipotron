# pipotron

[![Netlify Status](https://api.netlify.com/api/v1/badges/6013d277-e47f-47bc-83a3-ddc9ab5dad62/deploy-status)](https://app.netlify.com/sites/pipotron/deploys)

generic funny text generator

## Try online

Lambda functions are hosted on Netlify, give a try here:

* https://pipotron.netlify.com/.netlify/functions/pipotron?dict=bingo-winner
* https://pipotron.netlify.com/.netlify/functions/pipotron?dict=example
* https://pipotron.netlify.com/.netlify/functions/pipotron?dict=pipotron.free.fr
* https://pipotron.netlify.com/.netlify/functions/pipotron?dict=laboralphy

## Try with Docker

```console
docker run -it --rm moul/pipotron bingo-winner
Rockstar, Pivot, Ninja, Curated, Social... and BINGO!
```

[See on Docker Hub](https://hub.docker.com/r/moul/pipotron)

## See examples of generated content

Check out the [`./examples/` folder](./examples).

## Install

* Install latest [go](https://golang.org)
* Run: `GO111MODULE=on go get -u github.com/moul/pipotron`
* Profit: `pipotron dict/bingo-winner.yml`

## Contribute

Check out the [`./dict/` folder](./dict), and try updating an existing dictionary or creating a new one.
