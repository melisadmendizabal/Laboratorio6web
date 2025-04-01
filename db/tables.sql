--Melisa Mendizabal 23778

-- Crear la base de datos 'simple_login' si no existe
-- Por un momento no entendí y dejé el mismo nombre de tabla de la base de datos de Dennis ;(
CREATE DATABASE IF NOT EXISTS simple_login;
USE simple_login;

-- Crear la tabla 'series' si no existe
CREATE TABLE IF NOT EXISTS series(
        id INT AUTO_INCREMENT PRIMARY KEY,
        title varchar(250) not null,
        status varchar(50) not null,
        last_episode_watched int not null,
        total_episodes int not null,
        ranking int
);

-- Crear un usuario 'app_user' si no existe, con contraseña 'app_'
create user if not exists 'app_user'@'%' identified by 'app_password';
grant all privileges on simple_login.* to 'app_user'@'%';
-- Recargar los privilegios para asegurarse de que se apliquen los cambios
flush privileges;
