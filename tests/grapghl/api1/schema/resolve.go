package schema

import "github.com/graphql-go/graphql"

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