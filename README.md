## Run docker db

```
docker run -p 3306:3306
    -e MYSQL_ROOT_PASSWORD=secret
    -e MYSQL_DATABASE=order mysql
```

## Run grpc

```
grpcurl -d '{"user_id": 123, "order_items": [{"product_code": "prod", "quantity": 4, "unit_price": 12}]}' -plaintext localhost:80 Order/Create
```

## Helm Ingress Install

```
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install nginx-ingress ingress-nginx/ingress-nginx
```
