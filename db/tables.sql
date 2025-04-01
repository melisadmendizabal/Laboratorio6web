CREATE DATABASE IF NOT EXISTS simple_login;
USE simple_login;

CREATE TABLE IF NOT EXISTS series(
        id INT AUTO_INCREMENT PRIMARY KEY,
        title varchar(250) not null,
        status varchar(50) not null,
        last_episode_watched int not null,
        total_episodes int not null,
        ranking int
);

create user if not exists 'app_user'@'%' identified by 'app_password';
grant all privileges on simple_login.* to 'app_user'@'%';
flush privileges;
