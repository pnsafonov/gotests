package schema

import (
    "github.com/graphql-go/graphql"
    "log"
)

var (
    Schema graphql.Schema
)

func init() {
    var (
        err error
    )

    // sayhello
    fSayHelloArg1 := &graphql.ArgumentConfig{}
    fSayHelloArg1.Type = graphql.String

    fSayHello := &graphql.Field{}
    fSayHello.Type = graphql.String
    fSayHello.Args = make(graphql.FieldConfigArgument)
    fSayHello.Args["msg"] = fSayHelloArg1
    fSayHello.Resolve = resolveSayHello

    // helloworld
    fHelloWorld := &graphql.Field{}
    fHelloWorld.Type = graphql.String
    fHelloWorld.Resolve = resolveHelloWorld

    queryCfgFields := make(graphql.Fields)
    queryCfgFields["helloworld"] = fHelloWorld
    queryCfgFields["sayhello"] = fSayHello

    queryCfg := graphql.ObjectConfig{}
    queryCfg.Name = "Query"
    queryCfg.Fields = queryCfgFields

    query := graphql.NewObject(queryCfg)

    schemaConfig := graphql.SchemaConfig{}
    schemaConfig.Query = query

    Schema, err = graphql.NewSchema(schemaConfig)
    if err != nil {
        log.Fatalf("err = %v", err)
    }
}
