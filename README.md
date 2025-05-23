# Сервис шифрования строк

Простой сервис, который шифрует текст алгоритмами MD5 или SHA256 и кеширует результаты в Redis.


```bash
docker-compose up -d
```


### Проверка работы сервиса

```bash
curl -X GET http://localhost:8080/ping
```

### Шифрование текста

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