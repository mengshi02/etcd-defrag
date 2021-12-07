# etcd-defrag
An etcd data defragmentation tool

## Example

clean etcd data

```go
err := etcd_defrag.Clean(
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
