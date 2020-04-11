# go-animal-rescue

go-animal-rescue is a Go client library for accessing the [Animal Rescue API v1](https://github.com/anGie44/animal-rescue/tree/v1.0)

inspired by the `go-github` client library

## Usage ##

```go
import "github.com/anGie44/go-animal-rescue/v2/animalrescue"	// with go modules enabled (GO111MODULE=on or outside GOPATH)
import "github.com/anGie44/go-animal-rescue/animalrescue" // with go modules disabled
```

Construct a new AnimalRescue client, then use the various services on the client to
access different parts of the Animal Rescue API. For example:

```go
client := animalrescue.NewClient(nil)

// get adopter in an animal rescue by ID
adopter, _, err := client.Adopters.GetAdopterByID(context.Background(), 111)
```
