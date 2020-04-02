# Service
__Stefan Arentz, April 2020__

This is a very simple package with just a Run() function. It takes a function that you want to run as a service, waiting for it to finish, or receiving a SIGINT or SIGTERM.

```
func service(ctx context.Context) {
    ticker := time.NewTicker(500 * time.Millisecond)
    for {
        select {
	case <-ctx.Done():
            // We are done, just exit
            return
        case t := <-ticker.C:
            fmt.Println("Ticking at", t)
        }
    }
}

func main() {
    log.Println("Starting service")
    service.Run(context.Background(), service)
}
```

