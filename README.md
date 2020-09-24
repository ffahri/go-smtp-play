# go smtp play

Testing go smtp server using github.com/mhale/smtpd and send mail using net/smtp

## How to start server

run receive/main.go

### How to send message
- Telnet
```
telnet localhost 2525    
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
220 webischia webischia-go-smtp ESMTP Service ready
EHLO localhost
250-webischia greets localhost
250-SIZE 0
250 ENHANCEDSTATUSCODES
MAIL FROM:<*@*.com>  
250 2.1.0 Ok
RCPT TO:<*@*.com>
250 2.1.5 Ok
DATA
354 Start mail input; end with <CR><LF>.<CR><LF>
Hey
 
.
250 2.0.0 Ok: queued
quit
221 2.0.0 webischia webischia-go-smtp ESMTP Service closing transmission channel
Connection closed by foreign host.

```

Output from log
``` 
{"level":"info","time":"2020-09-24T23:07:42+03:00","message":"Rcpt Accepted ! From : *@*.com To: *@*.com"}
```

- Go net/smtp

Run send/main.go