package schema

import "github.com/graphql-go/graphql"

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

func iResolvePersonId(p graphql.ResolveParams) (interface{}, error) {
    if person, ok := p.Source.(Person); ok {
        return person.Id, nil
    }
    return nil, nil
}

func iResolvePersonFirstName(p graphql.ResolveParams) (interface{}, error) {
    if person, ok := p.Source.(Person); ok {
        return person.FirstName, nil
    }
    return nil, nil
}

func iResolvePersonLastName(p graphql.ResolveParams) (interface{}, error) {
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