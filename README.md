# aws-lambda-test
Test para comunicar lambdas medite reglas de ruteo de una NAT-Gateway

Este Script puede ser funcionar en dos modos: API y HOOK

## Lambda API
Esta lambda podrá recibir POSTs directos y emulará comunicarse con un proveedor, que en este caso será un webhook de prueba del sitio https://webhook.site/
Este webhook deberá siempre mostrar la IP del NAT-Gateway y no de la lambda.
Es decir, esta lambda deberá estár detrás de una NAT-Gateway.

Para probar, se deberá hacer un POST a la lambda, al path **/api** con cualquier body, por ejemplo:
```
{
    "attr1": "value1"
}
```

El webhook deberá mostrar el body enviado y la IP del NAT-Gateway, no de la lambda.

### Configuración de la lambda en modo API

1.  Ajusta la variable de entorno WEBHOOK_URL a el webhook que creaste
2.  Ajusta la variable de entorno LAMBDA_MOOD a "api"
3.  Ajusta la variable de entorno GIN_MODE a "release"
```bash
$ WEBHOOK_URL="https://webhook.site/<your-uuid>"
$ LAMBDA_MOOD="api"
$ GIN_MODE="release"
```
4. Desplegar este script en una lambda como Zip
```bash
$ GOOS=linux GOARCH=arm64 go build -o bootstrap main.go
$ zip bootstrap.zip bootstrap
```

La NAT-Gateway deberá tener una regla de ruteo que cualquier POST saliente a "https://webhook.site/< your-uuid >"
deberá ser enmascarada con la IP de la NAT-Gateway.

## Lambda Hook
Este script servirá para poder configurar una regla de ruteo en una NAT-Gateway y que cualquier POST que
reciba la NAT-Gateway sea redirijida a la lambda hook.

### Configuración de la lambda en modo API
1.  ajusta la variable de entorno LAMBDA_MOOD a "hook"

```bash
$ LAMBDA_MOOD="hook"
$ GIN_MODE="release"
```

Para probar, se deberá hacer un GET a la lambda al path /api con cualquier body, por ejemplo:
```
{
    "attr1": "value1"
}
```

La NAT-Gateway deberá tener una regla de ruteo que cualquier GET, POST, PUT que reciba, deberá
deberá ser redireccionada a esta **Lambda hook**.