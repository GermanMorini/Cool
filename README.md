# Cool

Aplicacion estilo `curl` para realizar peticiones HTTP o HTTPS a una url

## Características

- Permite enviar solicitudes GET, POST, PUT, DELETE, tambien medir el tiempo de respuesta, realizar logs para cada consulta, especificar body y mas
- Por defecto imprime el body de la consulta a `Stdout` y los logs a `Stderr`, esto permite guardar por separado los logs/body
- Permite usar logs en formato **JSON** o simplemente imprimir el tiempo de respuesta
- Se puede usar una variable de entorno (`COOL_URL`) para especificar la URL por defecto, así no se debe especificar con cada consulta
- Al especificar la flag `-b` se cambia el método automáticamente a `POST`, facilitando hacer consultas al no tener que especificar el método con cada consulta

## Codigos de salida

- **0** sin errores
- **1** error al parsear las flags
- **2** error generico (en la conexión, al leer el body de un archivo, ...)

## Flags

~~~bash

cool                                                # Consulta tipo GET a la dirección por defecto (http://localhost:8080)

cool -u URL                                         # Indica una url

cool -u URL -p PATH                                 # Indica un path

cool -m METHOD                                      # Metodo a utilizar (por defecto GET)

cool -ct CONT_TYPE                                  # Especifica el Content-Type (por defecto 'application/json')

cool -b '{"key":"val", ...}'                        # Especifica un body

cool -b @endpoint1.json                             # Lee el body desde un archivo

cool -Q                                             # No imprime el body

cool -q                                             # No imprime los logs

cool -j                                             # Logs en formato json

cool -rt                                            # Solo imprime el tiempo que tardó la consulta

cool -H "Authorization: Bearer $TK"                 # Especifica headers (se puede usar varias veces para cada header)

~~~
