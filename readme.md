## Steps to get started GraphQL - gqlgen

Download glggen using 
```
go get  github.com/99designs/gqlgen
```

This will create boilerplate code for GraphQL using gqlgen
```
go run github.com/99designs/gqlgen init
```

In `graph/schema.graphqls` define schemas(structs in go), input schema, mutations(means post, put methods), queries(get methods) and do [ref](https://gqlgen.com/getting-started/)
```
make generate
```

Every time we add/modify any mutation or query or schema we need to do `make generate` so that gqlgen will update the dependencies code 


# Steps to run

```
make run
```
Currently our server is running at http://localhost:8080/#


Example mutuation `createUser` add the below thing
```
mutation CreateUser($userInput: NewUser!) {
  createUser(input: $userInput) {
    id,
    name,
    userName,
    email
  }
}

// in variables tab below

{
"userInput": {
  "name":"sergio ramos",
  "userName": "sr04",
  "email": "ramos@fmail.com"
}
}
```

Example query `getUser` 
```
query GetUser($id: String!){
  getUser(input: $id){
    id,
    name,
    email
  }
}

// in variables tab below

{
"id": "suryacr7"
}
```
