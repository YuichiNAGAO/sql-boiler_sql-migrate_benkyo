
[SQLBoiler](https://github.com/volatiletech/sqlboiler)と[sql-migrate](https://github.com/rubenv/sql-migrate)の勉強用のレポジトリです。


sql-migrate
- migrationファイルの作成
```bash
sql-migrate new create_uesrs
```
- migrationの実行
```bash
sql-migrate up
```

SQLBoiler
```bash
sqlboiler mysql --config "config/sqlboiler.toml" --pkgname "models"  --output "./db/models" --no-tests --wipe
```


