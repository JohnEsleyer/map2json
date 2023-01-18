# map2json
A small Go library that converts a hashmap into a JSON file and vice versa.

#### Installation
```
go get github.com/johnesleyer/map2json
```

### Usage
Converting map to json file
```go
if err := map2json.ToJson(m, "example.json"); err != nil{
  fmt.Println("Error while converting to json:", err)  
}
```

Converting json file to map
```go
result, err := map2json.ToMap("example.json")
if err != nil{
  fmt.Println("Error while converting to map:", err)
}
fmt.Println(result)
```
