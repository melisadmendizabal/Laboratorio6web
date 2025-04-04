# Laboratorio6web
## Instrucciones para Ejecutar el Proyecto de API
### Melisa Mendizabal 23778

¡Bienvenido al proyecto de integración con APIs! Este laboratorio tiene como objetivo aprender y trabajar con APIs y bases de datos. A continuación se detallan los pasos para ejecutar tanto el frontend, el backend, como la base de datos que usa este sistema.
![image](https://github.com/user-attachments/assets/cb38c2ec-713c-4328-976d-72633329969d)


### 1. **Ejecutando el Frontend**

Para ver la interfaz gráfica del proyecto y asegurar que todos los íconos del frontend se muestren correctamente, es necesario iniciar un servidor local. Sigue estos pasos:

1. Navega al directorio `series-tracker` (dentro de la carpeta donde se encuentra este archivo README).
2. Ejecuta el siguiente comando:
python3 -m http.server 8000
Una vez que el servidor esté corriendo, abre tu navegador y accede a la siguiente URL:
http://127.0.0.1:8000
Deberías poder ver la pantalla principal de la aplicación.

### 2. Configurando y Ejecutando la Base de Datos
Este proyecto utiliza MariaDB como sistema de gestión de base de datos. Para levantar la base de datos en un contenedor Docker, sigue los pasos a continuación:

Entra a la carpeta BD.

Construye la imagen Docker con el siguiente comando:
docker build -t Series .

Una vez que la imagen esté construida, ejecuta el contenedor:
docker run -p 3306:3306 --name lab6 Series

Esto iniciará la base de datos MariaDB en el contenedor de Docker, accesible en el puerto 3306.

### Ver el contenido de la base de datos:

Si quieres verificar los datos dentro de la base de datos, sigue estos pasos:

Entra al contenedor en ejecución:
docker exec -it lab6 bash

Luego, accede a MariaDB y ejecuta una consulta para ver las series almacenadas:
mariadb -p -e "SELECT * FROM Series.series;"
La contraseña para ingresar es rootpassword, como se indica en el archivo Dockerfile.

### 3. Iniciando el Backend
El backend está implementado en Go. Para ejecutarlo, sigue estos pasos:

Dirígete a la carpeta backend.

Compila el proyecto Go con:
go build

Luego, ejecuta el servidor con:
go run .

Esto pondrá el backend a escuchar en el puerto 8080, y se conectará a la base de datos en el puerto 3306 para manejar las peticiones.

Flujo de Comunicación
Frontend: Escucha en el puerto 8000.

Backend: Se comunica con la base de datos MariaDB en el puerto 3306 y transmite datos a través del puerto 8080 hacia el frontend.
