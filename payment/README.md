## Call payment grpc

```
grpcurl -d '{"user_id": 123, "order_id": 12, "total_price": 32}' -plaintext localhost:3001 Payment/Create
```
