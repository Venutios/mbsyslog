mbSyslog
========
[![Build Status](https://travis-ci.com/Venutios/mbsyslog.svg?branch=master)](https://travis-ci.com/Venutios/mbsyslog)
<a href='https://coveralls.io/github/Venutios/mbsyslog?branch=master'><img src='https://coveralls.io/repos/github/Venutios/mbsyslog/badge.svg?branch=master' alt='Coverage Status' /></a>
[![Go Report Card](https://goreportcard.com/badge/github.com/Venutios/mbsyslog)](https://goreportcard.com/report/github.com/Venutios/mbsyslog)

An implementation of a Syslog server (RFC3164 and RFC5424) in Go.

## Install
```go get github.com/Venutios/mbsyslog```

## Usage
Starting a Syslog server and receiving messages.
```
messages := make(chan mbsyslog.Message, 5)
server := mbsyslog.NewServer(messages)

go func() {
	if err := server.Listen(); err != nil {
		panic(fmt.Sprintf("Server failed to start listening: %s", err.Error()))
	}
}()

for {
	select {
	case m := <-messages:
		fmt.Println(m)
	default:
		time.Sleep(100 * time.Millisecond)
	}
}
```

Stopping a Syslog server.
```
s.Stop()
for s.Running() {
    select {
    case m := <-messages:
        fmt.Println(m)
    default:
        time.Sleep(100 * time.Millisecond)
    }
}
```
