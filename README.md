# Сервис шифрования строк

запуск: собрать, запустить докер, и можно пробовать:

```bash
docker-compose up -d
```

### проверка работы сервиса

```bash
curl -X GET http://localhost:8080/ping
```

### шифрование текста
можно через curl, можно через postman

MD5:
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"input":"текст","algorithm":"md5"}' \
  http://localhost:8080/encrypt
```

SHA256:
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"input":"текст","algorithm":"sha256"}' \
  http://localhost:8080/encrypt
``` 
