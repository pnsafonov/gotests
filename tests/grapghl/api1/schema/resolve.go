package schema

import "github.com/graphql-go/graphql"

var (
    months = []string{
        "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",
    }
)

type EntityA struct {
    Id      int
    Name    string
    Amount  float32
}

type EntityB struct {
    Id          int
    Description string
    EntityA     EntityA
}

type EntityC struct {
    Id      int
    Tag     string
    EntityB EntityB
}

func resolveGetEntityA(p graphql.ResolveParams) (interface{}, error) {
    entA := EntityA{}
    entA.Id = 3354
    entA.Name = "Name of EntityA"

    return entA, nil
}

// query{
//  GetEntityC{
//    EntityB {
//      Id
//      Description
//      EntityA {
//        Amount
//        Id
//        Name
//      }
//    }
//    Id
//    Tag
//  }
//}
func resolveGetEntityC(p graphql.ResolveParams) (interface{}, error) {
    entA := EntityA{}
    entA.Id = 7782
    entA.Name = "Some name of EntityA"

    entB := EntityB{}
    entB.Id = 37
    entB.Description = "This is description of EntityB"
    entB.EntityA = entA

    entC := EntityC{}
    entC.Id = 111099
    entC.EntityB = entB

    return entC, nil
}

func oResolveEntityAId(p graphql.ResolveParams) (interface{}, error) {
    if ent, ok := p.Source.(EntityA); ok {
        return ent.Id, nil
    }
    return nil, nil
}

func oResolveEntityAName(p graphql.ResolveParams) (interface{}, error) {
    if ent, ok := p.Source.(EntityA); ok {
        return ent.Name, nil
    }
    return nil, nil
}

func oResolveEntityAAmount(p graphql.ResolveParams) (interface{}, error) {
    if ent, ok := p.Source.(EntityA); ok {
        return ent.Amount, nil
    }
    return nil, nil
}

func oResolveEntityBId(p graphql.ResolveParams) (interface{}, error) {
    if ent, ok := p.Source.(EntityB); ok {
        return ent.Id, nil
    }
    return nil, nil
}

func oResolveEntityBDescription(p graphql.ResolveParams) (interface{}, error) {
    if ent, ok := p.Source.(EntityB); ok {
        return ent.Description, nil
    }
    return nil, nil
}

func oResolveEntityBEntityA(p graphql.ResolveParams) (interface{}, error) {
    if ent, ok := p.Source.(EntityB); ok {
        return ent.EntityA, nil
    }
    return nil, nil
}

func oResolveEntityCId(p graphql.ResolveParams) (interface{}, error) {
    if ent, ok := p.Source.(EntityC); ok {
        return ent.Id, nil
    }
    return nil, nil
}

func oResolveEntityCTag(p graphql.ResolveParams) (interface{}, error) {
    if ent, ok := p.Source.(EntityC); ok {
        return ent.Tag, nil
    }
    return nil, nil
}

func oResolveEntityCEntityB(p graphql.ResolveParams) (interface{}, error) {
    if ent, ok := p.Source.(EntityC); ok {
        return ent.EntityB, nil
    }
    return nil, nil
}

//query {
//    resolveMonths
//}
func resolveMonths(p graphql.ResolveParams) (interface{}, error) {
    return months, nil
}

type Person struct {
    Id        int32
    FirstName string
    LastName  string
}

//query {
//    person{
//
// }
//}
func resolvePerson(p graphql.ResolveParams) (interface{}, error) {
    person := Person{}
    person.Id = 7
    person.FirstName = "Vicotor"
    person.LastName = "Tchernomardin"
    return person, nil
}

func oResolvePersonId(p graphql.ResolveParams) (interface{}, error) {
    if person, ok := p.Source.(Person); ok {
        return person.Id, nil
    }
    return nil, nil
}

func oResolvePersonFirstName(p graphql.ResolveParams) (interface{}, error) {
    if person, ok := p.Source.(Person); ok {
        return person.FirstName, nil
    }
    return nil, nil
}

func oResolvePersonLastName(p graphql.ResolveParams) (interface{}, error) {
    if person, ok := p.Source.(Person); ok {
        return person.LastName, nil
    }
    return nil, nil
}

//query {
//    sayhello(msg: "Hello Again!")
//}
func resolveSayHello(p graphql.ResolveParams) (interface{}, error) {
    arg1 := p.Args["msg"]
    return arg1, nil
}

//query {
//    helloworld
//}
func resolveHelloWorld(p graphql.ResolveParams) (interface{}, error) {
    return "Hello World!!", nil
}