# go-animal-rescue

go-animal-rescue is a Go client library for accessing the [Animal Rescue API v1.1](https://github.com/anGie44/animal-rescue/tree/v1.1)

inspired by the `go-github` client library

## Usage ##

```go
import "github.com/anGie44/go-animal-rescue" // with go modules disabled
```

Construct a new AnimalRescue client, then use the various services on the client to
access different parts of the Animal Rescue API. For example:

```go
client := animalrescue.NewClient(nil)

// get adopter in an animal rescue by ID
adopter, _, err := client.Adopters.GetAdopterByID(context.Background(), 111)
```
