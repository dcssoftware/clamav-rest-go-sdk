# clamav-rest-go-sdk

A simple SDK to communicate with [clamav-rest-service](https://github.com/dcssoftware/clamav-rest-service)

## How to use

```go
  // address, port, isHttps
  i := NewclamAVRestInstance("127.0.0.1", 33779, false)

  // testdocument -> []byte("...")
	isInfected, err := i.ScanFile(testdocument)
	if err != nil {
		panic(err)
	}

  // true -> no virus, false -> infected
	fmt.Println("result", isInfected)
```
