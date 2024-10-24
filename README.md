## Agenda

[x] - O que são Virtual Machines
[x] - O que são containeres
[x] - Containeres X Virtual Machines 
[x] - Prós e contras

### Conceitos

Os repositorios remotos do docker se chamam registries

Docker hub
EKR - AWS
GKE - GCP

[x] Imagens
    [X] Imagens prontas
    [x] Imagens customizadas

[x] Volumes
    [x] Anonymous
    [x] Bind Mounts
    [x] Named Volumes
[x] Redes
[x] Compose

Comando para se criar uma imagem customizada

```zsh
docker build -t <nome_da_imagem> <caminho do dockerfile>
### exemplo
docker build -t bgw4/meuubuntu .
```

Comando para criar um container

```zsh
docker run <parameters> <imagem>
### exemplo
docker run -d --name meu-container bgw4/meuubuntu
```

Comando para acessar o terminal desse container

```zsh
docker exec -it meu-container sh
```

Comando para parar um container que está rodando

```zsh
docker stop meu-container
```

Comando para iniciar um container parado

```zsh
docker start meu-container
```

Comando para reiniciar um container que está rodando

```zsh
docker restart meu-container
```

Comando para parar um container

```zsh
docker stop meu-container
```

Comando para acessar o terminal desse container

```zsh
docker exec -it meu-container sh
```

Comando para se criar um volume nomeado

```zsh
docker volume create myvol
```

Listando containeres que estao rodando

```zsh
docker ps
```

Listando containeres que foram criados, mas estao parados

```zsh
docker ps -a
```

Visualizando os logs de containeres que estão rodando

```zsh
docker logs <container_name> -f
```

Verificando recursos de um container

```zsh
docker stats meu-container
```

Comando para iniciar containeres que estão no docker compose

```zsh
docker compose up <parameters>
# exemplo para inicar todos os containeres sem travar o terminal
docker compose up -d
# exemplo para inicar todos os containeres sem travar o terminal e forcar o build deles
docker compose up -d --build
# exemplo de como iniciar servicos expecificos (api e app) no meu docker compose
docker compose up api app -d
```

Comando para parar todos os containeres no docker compose

```zsh
docker compose down
```

