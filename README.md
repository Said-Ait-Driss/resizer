# Resizer
## Light wight package to resize images from a given source

### request body :
```go
    {
        imageUrl string 
        width    int    
        height   int    
    }
```

### service listening on 3000 :


### runing the script :
```go
    cd resizer
    go mod tidy
    go run cmd/main.go
```