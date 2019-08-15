# json-response-validator

## The Problem

Assume you have sent an HTTP GET Request to an API. The API responds with a JSON payload.
Your job is to determine if that payload matches a given pattern.
The JSON payload is a map with keys of type string and values of various types.
Values can be of type string, float, array, or object.

The pattern matches the JSON payload if all of the following are true:
- for every field defined in the pattern, the payload also has that field
- The field is equal to the value in the given pattern

## Solutions

Two solutions are provided here:
1. setting up pattern types and unmarshalling the JSON into these instances to match against: [typechecking.go](./typechecking.go)
1. recursively running through a `map[string]interface` between the inputs and patterns: [mapinterface.go](./mapinterface.go)

A third solution involving the `reflect` and `mapstructure` packages is 
[outlined on StackOverflow](https://stackoverflow.com/questions/26744873/converting-map-to-struct/26746461).

For additional reference, this script was [originally attempted in Python](./json-validator.py).
