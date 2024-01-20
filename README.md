# Angry Customer Stress Tester

This application is a stress tester. 


Run with docker:
```bash 
docker run magrassi/stress-tester —-url http://google.com —-requests 1001 —-concurrency 10
```

Run locally:
```bash
go run main.go —-url http://google.com —-requests 1001 —-concurrency 10
```


