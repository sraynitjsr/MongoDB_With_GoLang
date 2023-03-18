package main

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func TestInsertOne(t *testing.T) {
	// Set up a MongoDB client
	client, ctx, cancel, err := connect("mongodb://localhost:27017", 30)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer close(client, ctx, cancel)

	// Create a test document
	document := bson.D{
		{Key: "Roll", Value: 10},
		{Key: "Mathematics", Value: 70},
		{Key: "Science", Value: 85},
	}

	// Insert the test document into the test database
	result, err := insertOne(client, ctx, "test-db", "test-col", document)
	if err != nil {
		t.Fatalf("Failed to insert document: %v", err)
	}

	// Check that the document was inserted successfully
	if result.InsertedID == nil {
		t.Fatalf("Failed to get the ID of the inserted document")
	}
}

func TestConnect(t *testing.T) {
	// Connect to a non-existent MongoDB server to test the error handling
	_, _, _, err := connect("mongodb://localhost:12345", 30)
	if err == nil {
		t.Fatalf("Expected an error when connecting to a non-existent MongoDB server")
	}

	// Connect to a valid MongoDB server to test the successful connection
	client, ctx, cancel, err := connect("mongodb://localhost:27017", 30)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer close(client, ctx, cancel)
}

func TestClose(t *testing.T) {
	// Set up a MongoDB client
	client, ctx, cancel, err := connect("mongodb://localhost:27017", 30)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Close the MongoDB client and check that it was disconnected successfully
	close(client, ctx, cancel)
	select {
	case <-ctx.Done():
		// The context was cancelled, so the client should have been disconnected
	case <-time.After(5 * time.Second):
		// The client was not disconnected within 5 seconds, so something went wrong
		t.Fatalf("Failed to disconnect the MongoDB client within 5 seconds")
	}
}
