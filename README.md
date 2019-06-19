# pipotron

[![CircleCI](https://circleci.com/gh/moul/pipotron.svg?style=svg)](https://circleci.com/gh/moul/pipotron)
[![Netlify Status](https://api.netlify.com/api/v1/badges/da26e7a7-179f-49b1-89b4-3103d382ee10/deploy-status)](https://app.netlify.com/sites/pipotron/deploys)
[![](https://images.microbadger.com/badges/image/moul/pipotron.svg)](https://microbadger.com/images/moul/pipotron "Get your own image badge on microbadger.com")

generic funny text generator

## Try online

Basic UI: https://pipotron.moul.io/

Lambda functions are hosted on Netlify, give a try here:

* https://pipotron.moul.io/.netlify/functions/pipotron?dict=bingo-winner
* https://pipotron.moul.io/.netlify/functions/pipotron?dict=example
* https://pipotron.moul.io/.netlify/functions/pipotron?dict=laboralphy
* https://pipotron.moul.io/.netlify/functions/pipotron?dict=moijaime
* https://pipotron.moul.io/.netlify/functions/pipotron?dict=pipotron.free.fr

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
