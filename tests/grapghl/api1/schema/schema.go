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

    // person
    oPersonCfg := graphql.ObjectConfig{}
    oPersonCfg.Name = "Person"
    oPersonCfgFields := make(graphql.Fields)
    oPersonCfg.Fields = oPersonCfgFields
    oPerson := graphql.NewObject(oPersonCfg)

    fPerson := &graphql.Field{}
    fPerson.Type = oPerson
    fPerson.Resolve = resolvePerson

    fPersonId := &graphql.Field{}
    fPersonId.Type = graphql.Int
    fPersonId.Resolve = oResolvePersonId

    fPersonFirstName := &graphql.Field{}
    fPersonFirstName.Type = graphql.String
    fPersonFirstName.Resolve = oResolvePersonFirstName

    fPersonLastName := &graphql.Field{}
    fPersonLastName.Type = graphql.String
    fPersonLastName.Resolve = oResolvePersonLastName

    oPersonCfgFields["Id"] = fPersonId
    oPersonCfgFields["FirstName"] = fPersonFirstName
    oPersonCfgFields["LastName"] = fPersonLastName

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
    queryCfgFields["person"] = fPerson
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
