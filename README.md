# Cool

Aplicacion estilo `curl` para realizar peticiones HTTP a una url

Permite enviar solicitudes GET, POST, PUT, DELETE, tambien medir el tiempo de respuesta, realizar logs para cada consulta,
especificar body y mas

Por defecto imprime el body de la consulta (a `Stdout`) y los logs a `Stderr`

Se puede usar una variable de entorno (`COOL_URL`) para especificar la URL por defecto

## Codigos de salida

- **0** sin errores
- **1** error al parsear las flags
- **2** error generico (en la conexión, al leer el body de un archivo, ...)
- **3** La consulta retorno un codigo de estado en el rango de los *300*
- **4** Idem, para los *400*
- **5** Idem, para los *500*

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

~~~
