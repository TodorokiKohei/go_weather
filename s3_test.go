package main

import (
	"context"
	"strings"
	"testing"
)

func TestObjectClinet_Connect(t *testing.T) {
	endpoint := "localhost:9000"
	accessKeyId := "root"
	secreteAccessKey := "password"

	obj := &ObjectClinet{}
	err := obj.Connect(endpoint, accessKeyId, secreteAccessKey)
	if err != nil {
		t.Fatal(err)
	}
}

func TestObjectClinet_Post(t *testing.T) {
	endpoint := "localhost:9000"
	accessKeyId := "root"
	secreteAccessKey := "password"

	body := `{"message": "text"}`
	obj := &ObjectClinet{}
	obj.Connect(endpoint, accessKeyId, secreteAccessKey)

	ctx := context.Background()
	_, err := obj.Post(ctx, "test-bucket", "test-message", strings.NewReader(body), int64(len(body)))
	if err != nil {
		t.Fatal(err)
	}
}
