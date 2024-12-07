package main

import (
	"errors"
	"fmt"
)

// Function that returns an error
func doError() (string, error) {
	return "", errors.New("this is my error")
}

// Function that does not return an error
func doNoError() (string, error) {
	return "My response", nil
}

// Function that returns an fmt error (advanced)
func doFmtError() error {
	errCode := 401
	return fmt.Errorf("this is my error code: %d", errCode)
}

/*
The error type is an interface type. An error variable represents any value that can describe itself as string.
The zero value of an interface is nil. Here's the interface declaration. Here the error interface declares a
Error() method that returns a string. We can create our own error type just by using the Error() string method.

type error interface {
	Error() string
}

The the most commonly used error emplementation is the errors package's unexported errorString.
When you call errors.New(), all it's doing is returning the unexported errorString struct that has
the message you passed in assignned to "s string" property.

type errorString struct {
	s string
}

And the Error() method defined on the errorString returns that string (e.s) that's attached to the struct.

func (e *errorString) Error() string {
	return e.s
}

*/

/*
 To make your own errors, implement Error() method that returns a string on your struct type.
*/

// Creating my own error types to contain within them some addtitonal information that can be used for debugging.

type ErrAccessTokenInvalid struct {
	Url     string // url used during request
	Code    int    // error code received back
	Message string // error message
}

func (eati ErrAccessTokenInvalid) Error() string {
	return fmt.Sprintf(
		"invalid access token - url: %s, response code: %d, response %s",
		eati.Url,
		eati.Code,
		eati.Message,
	)
}

type ErrRefreshTokenInvalid struct {
	Url     string
	Code    int
	Message string
}

func (erti ErrRefreshTokenInvalid) Error() string {
	return fmt.Sprintf(
		"invalid refresh token - url: %s, response code: %d, response: %s",
		erti.Url,
		erti.Code,
		erti.Message,
	)
}

type ErrAuthCodeInvalid struct {
	Url     string
	Code    int
	Message string
}

func (eaci ErrAuthCodeInvalid) Error() string {
	return fmt.Sprintf(
		"invalid auth code - url: %s, response code: %d, response %s",
		eaci.Url,
		eaci.Code,
		eaci.Message,
	)
}

func runAccessToken(url string, succeed bool) (string, error) {
	if succeed {
		return "Success!", nil
	}
	return "", ErrAccessTokenInvalid{url, 401, "access token expired"}
}

func runRefreshToken(url string, succeed bool) (string, error) {
	if succeed {
		return "Success!", nil
	}
	return "", ErrRefreshTokenInvalid{url, 400, "access token invalid"}
}

/*
The error type is an interface type. An error variable represents any value that can describe itself as string.
The zero value of an interface is nil. Here's the interface declaration. Here the error interface declares a
Error() method that returns a string. We can create our own error type just by using the Error() string method.

	type error interface {
		Error() string
	}

The the most commonly used error emplementation is the errors package's unexported errorString.
When you call errors.New(), all it's doing is returning the unexported errorString struct that has
the message you passed in assignned to "s string" property.

	type errorString struct {
		s string
	}

And the Error() method defined on the errorString returns that string (e.s) that's attached to the struct.

	func (e *errorString) Error() string {
		return e.s
	}
*/
func runAuthCode(url string, succeed bool) (string, error) {
	if succeed {
		return "Success!", nil
	}
	return "", ErrAuthCodeInvalid{url, 400, "authorization code expired"}
}

// --------ERROR WRAPPING --------------
// Advanced Error types - Wrap, Unwrap, Is & As.
// Three error types defined here. TOP, MIDDLE & BOTTOM

// ------TOP ERROR--------

type ErrTop struct {
	msg string
	mid error
}

func (et ErrTop) Error() string {
	return et.msg
}

func (et ErrTop) Unwrap() error {
	return et.mid
}

//---- MIDDLE ERROR --------

type ErrMiddle struct {
	msg    string
	bottom error
}

func (em ErrMiddle) Error() string {
	return em.msg
}

func (em ErrMiddle) Unwrap() error {
	return em.bottom
}

// -----BOTTOM ERROR-----Mimicking a bottow level error in a call stack ---

type ErrBottom struct {
	msg string
}

func (eb ErrBottom) Error() string {
	return eb.msg
}

// ------With ERROR WRAPPING, you get additonal functionality with Printf from fmt package. There's an extra verb with "%w" ----
// What "%w" does is, if you pass in an error, it will expand out the entire error chain.
// View main.go for demo
