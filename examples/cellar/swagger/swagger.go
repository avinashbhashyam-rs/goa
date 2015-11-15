//************************************************************************//
// cellar Swagger Spec
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=/home/raphael/go/src/github.com/raphael/goa/examples/cellar
// --design=github.com/raphael/goa/examples/cellar/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package swagger

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raphael/goa"
)

// MountController mounts the swagger spec controller under "/swagger.json".
func MountController(service goa.Service) {
	service.Info("mount", "ctrl", "Swagger", "action", "Show", "route", "GET /swagger.json")
	h := goa.NewHTTPRouterHandle(service, "Swagger", "Show", getSwagger)
	service.HTTPHandler().(*httprouter.Router).Handle("GET", "/swagger.json", h)
}

// getSwagger is the httprouter handle that returns the Swagger spec.
// func getSwagger(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
func getSwagger(ctx *goa.Context) error {
	ctx.Header().Set("Content-Type", "application/swagger+json")
	ctx.Header().Set("Cache-Control", "public, max-age=3600")
	return ctx.Respond(200, []byte(spec))
}

// Generated spec
const spec = `{"swagger":"2.0","info":{"title":"The virtual wine cellar","description":"A basic example of a CRUD API implemented with goa","contact":{"name":"goa team","email":"admin@goa.design","url":"http://goa.design"},"license":{"name":"MIT","url":"https://github.com/raphael/goa/blob/master/LICENSE"},"version":""},"host":"cellar.goa.design","basePath":"/cellar","schemes":["http"],"consumes":["application/json"],"produces":["application/json"],"paths":{"/accounts":{"post":{"description":"Create new account","operationId":"account#create","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/CreateAccountPayload"}}],"responses":{"201":{"description":"Resource created","headers":{"Location":{"description":"href to created resource","type":"string","pattern":"/accounts/[0-9]+"}}}},"schemes":["https"]}},"/accounts/{accountID}":{"get":{"description":"Retrieve account with given id","operationId":"account#show","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"accountID","in":"path","description":"Account ID","required":true,"type":"integer"}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/Account"}},"404":{"description":""}},"schemes":["https"]},"put":{"description":"Change account name","operationId":"account#update","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"accountID","in":"path","description":"Account ID","required":true,"type":"integer"},{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/UpdateAccountPayload"}}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]},"delete":{"operationId":"account#delete","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"accountID","in":"path","description":"Account ID","required":true,"type":"integer"}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]}},"/accounts/{accountID}/bottles":{"get":{"description":"List all bottles in account optionally filtering by year","operationId":"bottle#list","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"accountID","in":"path","required":true,"type":"string"},{"name":"years","in":"query","description":"Filter by years","required":false,"type":"array","items":{"type":"array","items":{"type":"integer"}}}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/BottleCollection"}}},"schemes":["https"]},"post":{"description":"Record new bottle","operationId":"bottle#create","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"accountID","in":"path","required":true,"type":"string"},{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/CreateBottlePayload"}}],"responses":{"201":{"description":"Resource created","headers":{"Location":{"description":"href to created resource","type":"string","pattern":"^/accounts/[0-9]+/bottles/[0-9]+$"}}}},"schemes":["https"]}},"/accounts/{accountID}/bottles/{bottleID}":{"get":{"description":"Retrieve bottle with given id","operationId":"bottle#show","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"accountID","in":"path","required":true,"type":"string"},{"name":"bottleID","in":"path","required":true,"type":"integer"}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/Bottle"}},"404":{"description":""}},"schemes":["https"]},"delete":{"operationId":"bottle#delete","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"accountID","in":"path","required":true,"type":"string"},{"name":"bottleID","in":"path","required":true,"type":"integer"}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]},"patch":{"operationId":"bottle#update","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"accountID","in":"path","required":true,"type":"string"},{"name":"bottleID","in":"path","required":true,"type":"integer"},{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/UpdateBottlePayload"}}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]}},"/accounts/{accountID}/bottles/{bottleID}/actions/rate":{"put":{"operationId":"bottle#rate","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"accountID","in":"path","required":true,"type":"string"},{"name":"bottleID","in":"path","required":true,"type":"integer"},{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/RateBottlePayload"}}],"responses":{"204":{"description":""},"404":{"description":""}},"schemes":["https"]}}},"definitions":{"Account":{"title":"Mediatype identifier: application/vnd.goa.example.account","type":"object","properties":{"created_at":{"type":"string","description":"Date of creation","format":"date-time"},"created_by":{"type":"string","description":"Email of account owner","format":"email"},"href":{"type":"string","description":"API href of account"},"id":{"type":"integer","description":"ID of account"},"name":{"type":"string","description":"Name of account"}},"description":"A tenant account"},"Bottle":{"title":"Mediatype identifier: application/vnd.goa.example.bottle","type":"object","properties":{"account":{"description":"Account that owns bottle","$ref":"#/definitions/Account"},"color":{"type":"string","enum":["red","white","rose","yellow","sparkling"]},"country":{"type":"string","minLength":2},"created_at":{"type":"string","description":"Date of creation","format":"date-time"},"href":{"type":"string","description":"API href of bottle"},"id":{"type":"integer","description":"ID of bottle"},"name":{"type":"string","minLength":2},"rating":{"type":"integer","description":"Rating of bottle between 1 and 5","minimum":1,"maximum":5},"region":{"type":"string"},"review":{"type":"string","minLength":10,"maxLength":300},"sweetness":{"type":"integer","minimum":1,"maximum":5},"updated_at":{"type":"string","description":"Date of last update","format":"date-time"},"varietal":{"type":"string","minLength":4},"vineyard":{"type":"string","minLength":2},"vintage":{"type":"integer","minimum":1900,"maximum":2020}},"description":"A bottle of wine"},"BottleCollection":{"title":"Mediatype identifier: application/vnd.goa.example.bottle; type=collection","type":"array","items":{"$ref":"#/definitions/Bottle"}},"CreateAccountPayload":{"title":"CreateAccountPayload","type":"object","properties":{"name":{"type":"string","description":"Name of account"}},"required":["name"]},"CreateBottlePayload":{"title":"CreateBottlePayload","type":"object","properties":{"color":{"type":"string","enum":["red","white","rose","yellow","sparkling"]},"country":{"type":"string","minLength":2},"name":{"type":"string","minLength":2},"region":{"type":"string"},"review":{"type":"string","minLength":10,"maxLength":300},"sweetness":{"type":"integer","minimum":1,"maximum":5},"varietal":{"type":"string","minLength":4},"vineyard":{"type":"string","minLength":2},"vintage":{"type":"integer","minimum":1900,"maximum":2020}},"required":["name","vineyard","varietal","vintage","color"]},"RateBottlePayload":{"title":"RateBottlePayload","type":"object","properties":{"rating":{"type":"integer","description":"Rating of bottle between 1 and 5","minimum":1,"maximum":5}},"required":["rating"]},"UpdateAccountPayload":{"title":"UpdateAccountPayload","type":"object","properties":{"name":{"type":"string","description":"Name of account"}},"required":["name"]},"UpdateBottlePayload":{"title":"UpdateBottlePayload","type":"object","properties":{"color":{"type":"string","enum":["red","white","rose","yellow","sparkling"]},"country":{"type":"string","minLength":2},"name":{"type":"string","minLength":2},"region":{"type":"string"},"review":{"type":"string","minLength":10,"maxLength":300},"sweetness":{"type":"integer","minimum":1,"maximum":5},"varietal":{"type":"string","minLength":4},"vineyard":{"type":"string","minLength":2},"vintage":{"type":"integer","minimum":1900,"maximum":2020}}}},"externalDocs":{"description":"goa guide","url":"http://goa.design/getting-started.html"}} `
