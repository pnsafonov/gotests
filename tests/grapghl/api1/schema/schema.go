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

    // Type EntityA
    oEntityACfg := graphql.ObjectConfig{}
    oEntityACfg.Name = "EntityA"
    oEntityAFields := make(graphql.Fields)
    oEntityACfg.Fields = oEntityAFields
    oEntityA := graphql.NewObject(oEntityACfg)

    fEntityAId := &graphql.Field{}
    fEntityAId.Type = graphql.Int
    fEntityAId.Resolve = oResolveEntityAId

    fEntityAName := &graphql.Field{}
    fEntityAName.Type = graphql.String
    fEntityAName.Resolve = oResolveEntityAName

    fEntityAAmount := &graphql.Field{}
    fEntityAAmount.Type = graphql.Float
    fEntityAAmount.Resolve = oResolveEntityAAmount

    oEntityAFields["Id"] = fEntityAId
    oEntityAFields["Name"] = fEntityAName
    oEntityAFields["Amount"] = fEntityAAmount

    // Type EntityB
    oEntityBCfg := graphql.ObjectConfig{}
    oEntityBCfg.Name = "EntityB"
    oEntityBFields := make(graphql.Fields)
    oEntityBCfg.Fields = oEntityBFields
    oEntityB := graphql.NewObject(oEntityBCfg)

    fEntityBId := &graphql.Field{}
    fEntityBId.Type = graphql.Int
    fEntityBId.Resolve = oResolveEntityBId

    fEntityBDescription := &graphql.Field{}
    fEntityBDescription.Type = graphql.String
    fEntityBDescription.Resolve = oResolveEntityBDescription

    fEntityBEntityA := &graphql.Field{}
    fEntityBEntityA.Type = oEntityA
    fEntityBEntityA.Resolve = oResolveEntityBEntityA

    oEntityBFields["Id"] = fEntityBId
    oEntityBFields["Description"] = fEntityBDescription
    oEntityBFields["EntityA"] = fEntityBEntityA

    // Type EntityC
    oEntityCCfg := graphql.ObjectConfig{}
    oEntityCCfg.Name = "EntityC"
    oEntityCFields := make(graphql.Fields)
    oEntityCCfg.Fields = oEntityCFields
    oEntityC := graphql.NewObject(oEntityCCfg)

    fEntityCId := &graphql.Field{}
    fEntityCId.Type = graphql.Int
    fEntityCId.Resolve = oResolveEntityCId

    fEntityCTag := &graphql.Field{}
    fEntityCTag.Type = graphql.String
    fEntityCTag.Resolve = oResolveEntityCTag

    fEntityCEntityB := &graphql.Field{}
    fEntityCEntityB.Type = oEntityB
    fEntityCEntityB.Resolve = oResolveEntityCEntityB

    oEntityCFields["Id"] = fEntityCId
    oEntityCFields["Tag"] = fEntityCTag
    oEntityCFields["EntityB"] = fEntityCEntityB

    // GetEntityA
    fGetEntityA := &graphql.Field{}
    fGetEntityA.Type = oEntityA
    fGetEntityA.Resolve = resolveGetEntityA

    // GetEntityC
    fGetEntityC := &graphql.Field{}
    fGetEntityC.Type = oEntityC
    fGetEntityC.Resolve = resolveGetEntityC

    // months
    fMonths := &graphql.Field{}
    fMonths.Type = graphql.NewList(graphql.String)
    fMonths.Resolve = resolveMonths

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

    //
    queryCfgFields := make(graphql.Fields)
    queryCfgFields["resolveMonths"] = fMonths
    queryCfgFields["person"] = fPerson
    queryCfgFields["helloworld"] = fHelloWorld
    queryCfgFields["sayhello"] = fSayHello
    queryCfgFields["GetEntityC"] = fGetEntityC
    queryCfgFields["GetEntityA"] = fGetEntityA

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
