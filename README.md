# Portal en Mantenimiento
Este microServicio permitirá levantar un servidor http con la capacidad de servir una pagina
estática que muestre un mensaje de mantenimiento.
El servicio estara corriendo en: 

> domain:port

Descripción
-----------
Descripción de los directorios:
  * **config:** Directorio que contiene un archivo ENV con la configuración del servicio,
             parámetros importantes a configurar son (Host, Port, Template_name).
  * **statics:** Directorio donde se almacenaran los contenidos estáticos como css, js, img.
  * **templates:** Directorio donde se almacenaran los templates.
 
Archivo de Configuración config.env
-----------------------------------

>host = dominio del portal que entrara en mantenimiento.
>
>port = puerto en el cual estará escuchando el servidor.
>
>template_dir = directorio donde se deben almacenar los templates.
>
>template_name = nombre del template que se mostrara.
>
>static_route = ruta estática utilizada para invocar los css, js e img.
>
>static_dir = directorio donde deben almacenar los contenidos estaticos.

Testing
-------

Ejemplo de Configuración para testing

~~~ENV
  MAINTENANCE_PORT=9090
  RETRY_TIME=3600
  HOST=testmaintenance.dev
  TEMPLATE_DIR=Testing/templates/
  TEMPLATE_NAME=Mantenimiento
  STATIC_ROUTE=/assets/
  STATIC_DIR=./Testing/assets/
~~~

**MAINTENANCE_PORT:**
Es el puerto donde estara escuchando la aplicación

**RETRY_TIME:**
Tiempo de reintento para los robots

**HOST:**
En el nombre de host puedes poner cualquier dominio,
pero en tu archivo etc/host debes hacer que redireccione a la ip de tu docker-machine (por default es: 192.168.99.100) para fines de prueba.
~~~
  > 192.168.99.100       testmaintenance.dev
~~~

**TEMPLATE_DIR:**
Directorio donde se encuentran almacenado el template, dentro del proyecto hay un directorio llamado Public donde debe estar todo el contenido.

**TEMPLATE_NAME:**
Es el nombre de la plantilla que se utilizara, este nombre se define con el formato de plantillas golang.
> **Nombre de Template para Testing:** Mantenimiento
 
**STATIC_ROUTE:**
Esta es el nombre de la ruta que se configurara en el servidor para servir todos los contenidos estaticos.
> El tempate html debe llamar a los contenidos estaticos teniendo en cuenta el nombre de la ruta.

**STATIC_DIR:**
Esta es la ruta donde esta el directorio que contiene los contenidos estaticos.

Requerimientos del servicio
---------------------------

  * Docker
  * Docker-compose
  * Docker-machine (para fines de prueba)

Ejecución del servicio
----------------------

  * **Primer paso:** clonar dentro de la carpeta "src/" el proyecto:
~~~
  > git clone https://git.yachay.pe/osanchezj/ms-portals-maintenance.git
~~~~

  * **Segundo paso:** 
    Crear el contenedor de Docker
 ~~~
  > cd maintenance_service/
  > PORT=<port> docker-compose up -d
 ~~~ 
  > El comando "docker-compose up -d" descargara la imagen de golang, creara el contenedor y lo iniciara.
 > Se le pasa una variable de entorno llamada PORT, esta variable debe tener el mismo numero de puerto que se configuro en el archivo .ENV
 
  * **Tercer paso:** Ingresar desde tu navegador al servicio indicando el puerto 
~~~
  http://testmaintenance.dev:port
~~~~
  
  **Nota:** Para fines de producción este contenedor debe estar atras de un nginx o algun otro servidor web que haga el proxeo entre tu portla y el servicio de mantenimiento.