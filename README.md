# Introduction

This repo contains some initial experiments with the Go programming language

# Project structure

Every folder contains an example, and all the examples share the same structure:

```shell
<folder>
  |
  ├─ Dockerfile
  └─ <src folder>
         |
         ├─ <go module definition>
         ├─ <go entrypoint file>
         └─ ...

```

All the dockerfiles have a common multistaged build structure nad accept an input parameter (which by definition is the name of the folder)

# Build and run the project

The project can simply be built using docker; to simplify this (And based on the conventions applied to the dockerfile), from the root of the repo simply run:

```shell
./run <name of the single project folder without '/'>
```

To continuously watch for changes in a project folder and react building and running it, execute:

```shell
./watchrun <name of the single project folder without '/'>
```

When the project ot run needs to be exposed on a specific port, the above command can be changed into:

```shell
./watchrun <name of the single project folder without '/'> expose <port to expose>
```

The port (let's assume 8080) in this case will be exposed with the docker run option `-p 8080:8080`
