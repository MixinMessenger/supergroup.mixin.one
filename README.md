## Mixin Super Group

supergroup.mixin.one is the source code of "Mixin 中文群"，which is an unlimited member group base on Mixin bot.

## Prepare

1. copy `config.tpl.yaml` to `config.yaml`
2. Replace configurations in `config.yaml`, we use PostgreSQL as our database.
3. `cd web` and exec `npm install`
4. copy `env.dev.tpl.sh` to `env.dev.sh` and copy `env.prod.tpl.sh` to `env.prod.sh`
5. Replace all fields in `env.dev.sh` and `env.prod.sh`


## update if your repo is setup before 2019 May 28
```
psql -h YOUR_HOST -U YOUR_USERNAME -d YOUR_DATABASE -a -f schema.append001.sql
```
## Run

#### Server Side

1. `./supergroup.mixin.one` handle http request
2. `./supergroup.mixin.one -service message` handle messages

#### Front-end

Generate static assets `cd web && npm run build`

