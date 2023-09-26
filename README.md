# golang project with gin, gorm, wire


### run project
```
$ make generate
$ go run ./...
```

## use viper for configs
viper read configs from config.yaml file <br>
for system environment can use _ for example: datasource_redis_enabled .

## use wire as dependency injection
after create your repository,useCase and controller add new instance in the corresponding file inside each of the related folders
then run <b>make generate</b> .


