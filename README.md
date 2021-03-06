# go-http-client

#### How to new a client
```
u, _ := url.Parse("http://www.google.com")
builder := (*client.CBuilder)(nil).NewBuilder()
builder.BaseUrl = u
client := builder.Build().NewClient()

resposne, _ := client.Get("/").Execute()
fmt.Println("reponse body", resposne.Payload)
```

#### Http Client, support following method
Get
Post
Delete
Put
....
#### Easy to add path param, header
```
client.Post("/subPath", "/thirdPath").Header("User-Agent", "My-Client").AddParam("Param", "Value").Execute()
```
#### Support Json Entity
```
client.Post().JsonEntity(<struct{}>).Execute()
```
#### Easy to debug
Request&Response are easy to check
```
request url >>>>> https://www.google.com?key=valu
request header >>>>> map[User-Agent:[my-client]]
response status code <<<<< 200
response header <<<<< map[Cache-Control:[private, max-age=0] Content-Type:[text/html; charset=ISO-8859-1] P3p:[CP="This is not a P3P policy! See https://www.google.com/support/accounts/answer/151657?hl=en for more info."] Server:[gws] X-Frame-Options:[SAMEORIGIN] Set-Cookie:[NID=89=gyjRzw_bfEUEPAgnFKqFsr2oUydsLrJwnhLdIvl-Nxbq4rMa5wIQmWeWQpiR1lgr0tPLQFdD1ZtU8Uj745iBMvc8erFSVH9GFdTKXHeAOWpRRVigQ4MU7HAp5XGyvdIfwQkWl377tRujQA; expires=Fri, 28-Apr-2017 08:34:51 GMT; path=/; domain=.google.com; HttpOnly] Date:[Thu, 27 Oct 2016 08:34:51 GMT] Expires:[-1] X-Xss-Protection:[1; mode=block] Alt-Svc:[quic=":443"; ma=2592000; v="36,35,34"]]
response body <<<<< <!
```
#### Get reponse body
```golang
respone, error := client.GET().Execute()
if error != nil {
	fmt.Println("reponse body", respone.Payload)
}
```
#### Response to Entity
```golang
entity  := <struct{}>
client.Post().JsonEntity(<struct{}>).ExecuteForEntity(&entity)
```
#### Support client authorizaiton header provider
there is simpleAuthroziaitonProvider for eample
