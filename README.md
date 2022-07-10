## Setup
```
go mod tidy
```

## Usage
```
go run server.go
```

## generate schema.graphqls
```
go run github.com/99designs/gqlgen generate
```

##  *.graphqls extension
```
From this we get that files with the *.graphqls extension are interpreted as schema files and used to generate the bindings and types under graph/.

graph/generated/generated.go is where all the “glue” between the graphql specification and go is generated; you’ll hopefully never need to look inside
graph/model/models_gen.go is where the unmanaged types and fields are generated
graph/resolver.go is where we will add our dependencies, which will be shared between the different types of resolvers
graph/schema.resolvers.go contains the go receivers for the resolvers which correspond to the queries and mutations defined in schema.graphqls
```

## structure
```
├── go.mod
 
├── go.sum
 
├── gqlgen.yml
 
├── graph
 
│ ├── generated
 
│ │ └── generated.go
 
│ ├── model
 
│ │ └── models_gen.go
 
│ ├── resolver.go
 
│ ├── schema.graphqls
 
│ └── schema.resolvers.go
 
└── server.go
```

##Tutorial
```
    to see the tutorial visit [here](https://www.inovex.de/de/blog/graphql-application-golang-tutorial/)
```