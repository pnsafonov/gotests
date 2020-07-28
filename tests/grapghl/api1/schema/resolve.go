package schema

import "github.com/graphql-go/graphql"

//query {
//    HelloWorld
//}
func resolveHelloWorld(p graphql.ResolveParams) (interface{}, error) {
    return "Hello World!!", nil
}