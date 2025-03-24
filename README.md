# Cool

Aplicacion estilo `curl` para realizar peticiones HTTP o HTTPS a una url

## Características

- Permite enviar solicitudes GET, POST, PUT, DELETE, o cualquier otra, midiendo el tiempo de respuesta e imprimiendo los headers de la respuesta
- Por defecto imprime el body de la consulta a `Stdout` y los logs a `Stderr`, esto permite guardar por separado los logs del body
- Al especificar la flag `-b` se cambia el método automáticamente a `POST` y se establece el header `Content-Type: application/json`, facilitando hacer consultas al no tener que especificar cada opción por separado

## Codigos de salida

- **0** sin errores
- **1** error al parsear las flags
- **2** error generico (en la conexión, al leer el body de un archivo, ...)

## Ejemplos

~~~bash

cool                                                # Consulta tipo GET a la dirección por defecto (http://localhost:8080)

cool URL                                            # Indica una url

cool -m METHOD                                      # Metodo a utilizar (por defecto GET)

cool -ct CONT_TYPE                                  # Especifica el Content-Type (por defecto 'application/json')

cool -b '{"key":"val", ...}'                        # Especifica un body

cool -b @endpoint1.json                             # Lee el body desde un archivo

cool -Q                                             # No imprime el body

cool -q                                             # No imprime los logs

cool -H "Authorization: Bearer $TOKEN"              # Especifica headers (se puede usar varias veces para cada header)

~~~
