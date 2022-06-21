package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	radosAPI "github.com/QuentinPerez/go-radosgw/pkg/api"
)

func printRawMode(out io.Writer, data interface{}) error {
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s\n", js)
	return nil
}

func main() {
	api, err := radosAPI.New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	var user *radosAPI.User
	if os.Args[1] == "create" {
		// create a new user named JohnDoe
		user, err = api.CreateUser(radosAPI.UserConfig{
			UID:         "JohnDoe",
			DisplayName: "John Doe",
			MaxBuckets:  intPtr(-1),
		})
		if err != nil {
			log.Fatal(err)
		}
		printRawMode(os.Stdout, user)
	}

	if os.Args[1] == "get" {
		// get the user named JohnDoe
		user, err = api.GetUser("JohnDoe")
		if err != nil {
			log.Fatal(err)
		}
		printRawMode(os.Stdout, user)
	}

	if os.Args[1] == "link" {
		if os.Args[2] == "" {
			log.Fatal("missing bucket name")
		}
		// link the user named JohnDoe to the bucket named test
		err = api.LinkBucket(radosAPI.BucketConfig{
			Bucket: os.Args[2],
			UID:    "JohnDoe",
		})

	}

	if os.Args[1] == "remove" {
		// remove JohnDoe
		err = api.RemoveUser(radosAPI.UserConfig{
			UID: "JohnDoe",
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func intPtr(i int) *int {
	return &i
}
