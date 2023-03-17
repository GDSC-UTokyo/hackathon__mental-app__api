# hackathon__mental-app__api

1.mysqlサーバーを立ち上げる。
```
mysql.server start
```

2.mysqlにログインする
```
mysql -u (ユーザー名) -p(パスワード)
```

3.mysqlのrootユーザーを作成し、パスワードをrootに変更する。権限も付与する。
```
CREATE USER 'root'@'localhost' IDENTIFIED BY 'root';
```

4.データベースを作成する。
```
CREATE DATABASE `mental-app`;
```

5.Goサーバーを以下のコマンドで起動する。