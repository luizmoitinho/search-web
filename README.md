# search-web
### Configurando o ambiente
1. Clone o projeto
2. Execute os seguintes comandos em sua máquina
```shell
go mod init
go mod tidy
```

### Testando
1. Busca de IPs na Web
```shell
go run main.go ip --host amazon.com.br
```
> Nota 1: No lugar de "amazon.com.br" informe qualquer host valido na web para saber o seu IP
> Nota 2: Caso não seja informado o trecho "--host amazon.com.br" a busca a ser feita será para 'google.com"

2. Buscar por servidores
```shell
 go run main.go servers --host amazon.com.br
 ```

3. Execute com o build do projeto
```shell
 ./search-web ip --host amazon.com.br
```
