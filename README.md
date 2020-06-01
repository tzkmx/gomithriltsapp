# gomithrilapp

[![Go Report Card](https://goreportcard.com/badge/github.com/josephspurrier/gomithrilapp)](https://goreportcard.com/report/github.com/josephspurrier/gomithrilapp)
[![Build Status](https://travis-ci.org/josephspurrier/gomithrilapp.svg)](https://travis-ci.org/josephspurrier/gomithrilapp)
[![Coverage Status](https://coveralls.io/repos/github/josephspurrier/gomithrilapp/badge.svg?branch=master&timestamp=20200313-01)](https://coveralls.io/github/josephspurrier/gomithrilapp?branch=master)

[![Swagger Validator](https://online.swagger.io/validator?url=https://raw.githubusercontent.com/josephspurrier/gomithrilapp/master/src/app/ui/static/swagger.json)](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/josephspurrier/gomithrilapp/master/src/app/ui/static/swagger.json)

This is a sample notepad application that uses Mithril on the front-end (UI) and Go on the back-end (API). This project is designed to show good development and CI/CD practices as well as integrations between modern development tools.

[Documentation](https://josephspurrier.github.io/gomithrilapp/) is generated using [Docusaurus](https://docusaurus.io/) and hosted using [GitHub Pages](https://pages.github.com/). This project uses a [Makefile](Makefile) to centralize frequently used commands. The code coverage badge above is only for the back-end API - not the front-end.

![Demo gif](https://user-images.githubusercontent.com/2394539/76177148-ac753e00-6189-11ea-963b-bff38b29e8ed.gif)

## Quick Start with Docker Compose

To run the application locally, you can run these commands. You don't need any of the the dev tools (Go/npm) installed, you only need Docker (and Docker Compose). The Go application serves both the UI and the API depending on the request URL.

```bash
# Clone the repo.
git clone git@github.com:josephspurrier/gomithrilapp.git

# CD to the project directory.
cd gomithrilapp

# Build the docker containers.
make docker-build

# Run the docker containers: DB and app.
docker-compose up -d

# Open your browser to the UI: http://localhost
# Open your browser to the API: http://localhost/api
# Open your MySQL tool to the DB: localhost:3306 (root:password)

# Stop and remove the docker containers.
docker-compose down
```

## Additional Documentation

Below are links to various section of the documentation.

- [Environment Preparation](https://josephspurrier.github.io/gomithrilapp/docs/tutorial/env-prep)
- [Database](https://josephspurrier.github.io/gomithrilapp/docs/database)
- [Front-End](https://josephspurrier.github.io/gomithrilapp/docs/front-end)
- [Back-End](https://josephspurrier.github.io/gomithrilapp/docs/back-end)
- [Swagger](https://josephspurrier.github.io/gomithrilapp/docs/swagger)
- [Docker Compose](https://josephspurrier.github.io/gomithrilapp/docs/docker-compose)
- [Documentation](https://josephspurrier.github.io/gomithrilapp/docs/documentation)