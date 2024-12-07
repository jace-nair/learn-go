package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// ------------ EXAMPLE 1: context is for controlling timouts --------------

func helloHandler(w http.ResponseWriter, r *http.Request) {

	// *http.Request has context built into it
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("API Response! from Example 1")
	case <-ctx.Done():
		fmt.Println("Oh no the context in Example 1 expired")
		http.Error(w, "Request context timeout in Example 1", http.StatusRequestTimeout)
		return
	}
}

// ------------ EXAMPLE 2: context is for cancelling go routines --------------

func exampleTimeout(ctx context.Context) {

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second) // create a Timeout by passing in a parent context and time to cancel

	defer cancel()

	done := make(chan struct{}) // make a channel named "done" of struct type

	// Create an anonymous go function sub-routine that mimics an API that takes 3 seconds to complete, 1 second greater than context.
	// We can control or hanldle this go subroutine more elegently with "context"
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	// Case 1 is listening for values coming out of our "done" channel
	case <-done:
		fmt.Println("Called the API in Example 2")
	// Listening for signals coming out of context with timeout
	case <-ctxWithTimeout.Done():
		fmt.Println("Oh no my timeout expired! in Example 2", ctxWithTimeout.Err())
		// do some logic to handle this
	}
}

// -------------- EXAMPLE 2 (another one) ------------------------

func doSomethingCool(ctx context.Context) {

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second) // create a Timeout by passing in a parent context and time to cancel

	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Oh no, I've exceeded the deadline in Example 2 (another one)")
		}
		time.Sleep(2 * time.Second)
	}()

	// This code simulates a long running process
	for {
		select {
		case <-ctxWithTimeout.Done():
			fmt.Println("timed out in Example 2 (another one)")
			return
		default:
			fmt.Println("doing something cool in Example 2 (another one)")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// ------------- EXAMPLE 3: context is for passing metadata, information or state across go applications or process ----------------
// Carry values for authentication, tokens or user information
// Have abilities to pass state across go applications

func exampleWithValues() {

	// Create a key of int type
	type key int
	// Declare and assign Userkey of key type
	const UserKey key = 0

	ctx := context.Background()

	// Create a context with value by passing in the parent context and key value pair
	ctxWithValue := context.WithValue(ctx, UserKey, "123")

	if userID, ok := ctxWithValue.Value(UserKey).(string); ok {
		fmt.Println("this is the userID in Example 3", userID)
	} else {
		fmt.Println("this is a protected route in Example 3 - no userID found")
	}
}

// -------------- EXAMPLE 3 (another one) ------------------------

// Create a rIDType3 of int type
type rIDType3 int

// Declare and assign requestID3 of rIDType3
const requestID3 rIDType3 = 0

func enrichContext3(ctx context.Context) context.Context {
	return context.WithValue(ctx, requestID3, "12345")
}

func doSomethingCool3(ctx context.Context) {
	rID3 := ctx.Value(requestID3)
	fmt.Println("This is the requestID3 in Example 3 (another one): ", rID3)
}
