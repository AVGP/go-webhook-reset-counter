# go-webhook-reset-counter
## Tiny util that accepts HTTP requests, toggles a state and counts the interactions, then resets after a configured timeout.

## What's this?

Imagine you need something that you can call via webhook and that keeps track if it has been triggered as well as how often it has been triggered.
It also allows you to have this information reset after a given time period.

## Uhm... no, really, what for?

Okay, so I'm using this to track the state of a sensor and have it reset after a while... you know, Internet of Things stuff. Yaaay, future!

## Building on other platforms

Just compile with
```shell
$ go build -o WebhookResetCounter main.go
```

## How to use this?

To start the webserver on port 8080 (default) and have it reset its state after an hour, run:

```shell
$ WebhookResetCounter --reset 3600
```

If you don't want any resetting, use `--reset 0` instead and if you don't want port 8080, use `--port` to specify any other port it should listen on instead.