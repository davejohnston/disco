# Build

```make all```

Run the controller service
```bin/controller```

Test the service

## Get All Providers

```curl http://localhost:9999/v1/cmd/providers```
    
## Get All Resources

```curl http://localhost:9999/v1/cmd/resources```

## Add A Resource 
curl -X POST http://localhost:9999/v1/cmd/resources -d '{ "type" : "d41e2236-f147-4a10-9701-df412dd58f74", "config" : { "Host" : "vcenter.mydomain.com", "Username" : "dave", "Password" : "XYZ" } }' | jq
