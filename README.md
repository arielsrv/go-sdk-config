[![pipeline status](https://gitlab.tiendanimal.com:8088/iskaypet/digital/tools/dev/go-sdk-config/badges/main/pipeline.svg)](https://gitlab.tiendanimal.com:8088/iskaypet/digital/tools/dev/go-sdk-config/-/commits/main)
[![coverage report](https://gitlab.tiendanimal.com:8088/iskaypet/digital/tools/dev/go-sdk-config/badges/main/coverage.svg)](https://gitlab.tiendanimal.com:8088/iskaypet/digital/tools/dev/go-sdk-config/-/commits/main)
[![release](https://gitlab.tiendanimal.com:8088/iskaypet/digital/tools/dev/go-sdk-config/-/badges/release.svg)](https://gitlab.tiendanimal.com:8088/iskaypet/digital/tools/dev/go-sdk-config/-/releases)

> This SDK provides a SaaS framework to build middle-end for consumers and APIs

The intent of the project is to provide a lightweight microservice config sdk, based on Golang

The main goal is to provide a modular framework with high level abstractions to config you app which enforces best
practices

## ⚡ Table of contents

* [SDK](#sdk)
    * [Configuration](#configuration)

## SDK

### Configuration

Environment configuration is based on **Archaius Config**, you should use a similar folder
structure.
*SCOPE* env variable in remote environment is required from Kubernetes

```
└── config
    ├── config.yml (shared config)
    └── local
        └── config.yml (for local development)
    └── prod (for remote environment)
        └── config.yml (base config)
        └── {environment}.config.yml (base config)
```

The SDK provides a simple configuration hierarchy

* env variables
* resources/config/config.properties (shared config)
* resources/config/{environment}/config.properties (override shared config by environment)
* resources/config/{environment}/{scope}.config.properties (override env and shared config by scope)

example *consumers-api.uat.dp.iskaypet.com*

```
└── env variables                               (always first)
└── config
    ├── config.yml                              3th (third)
    └── local
        └── config.yml                          ignored
    └── prod
        └── config.yml (base config)            2nd (second)
        └── uat.config.yml (base config)        1st (first)
```

* 1st (first)   prod/uat.config.yml
* 2nd (second)  prod/config.yml
* 3th (third)   config.yml

example logs in production app

```
2023-03-10 16:45:03 [info] working directory: /app
2023-03-10 16:45:03 [info] loaded configuration file: /app/src/resources/config/prod/uat.config.yml
2023-03-10 16:45:03 [info] loaded configuration file: /app/src/resources/config/prod/config.yml
2023-03-10 16:45:03 [info] loaded configuration file: /app/src/resources/config/config.yml
2023-03-10 16:45:03 [info] invoke dynamic handler:FileSource
2023-03-10 16:45:03 [info] enable env source
2023-03-10 16:45:03 [info] invoke dynamic handler:EnvironmentSource
2023-03-10 16:45:03 [info] archaius init success
2023-03-10 16:45:03 [info] INFO log level
2023-03-10 16:45:03 [info] ENV: prod, SCOPE: uat
2023-03-10 16:45:03 [info] create new watcher
2023-03-10 16:45:03 [info] swagger enabled
2023-03-10 16:45:03 [info] Listening on local address 0.0.0.0:8080
2023-03-10 16:45:03 [info] Open https://consumers-api.uat.dp.iskaypet.com/ping in the browser
```