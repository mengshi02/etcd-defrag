# etcd-defrag
An etcd data defragmentation tool

## Example

etcd defrag data by API

```go
err := etcd_defrag.Run(
    []string{
        "https://127.0.0.1:4001",
    },
    "./certificate/service.cert",
    "./certificate/service.key",
    "./certificate/ca.cert",
)
if err != nil {
    fmt.Printf("clean failure: %s\n", err.Error())
}
fmt.Println("clean success")
```

defrag data by CMD
```shell
defrag -eps="http://127.0.0.1:4001, http://127.0.0.2:4001" --trusted-ca-file="ca.cert" --cert-file="service.cert" --key-file="service.key"
```
