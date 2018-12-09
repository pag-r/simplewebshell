## Simple Go-based webshell for Docker
By default http server listens on port 9090. Commands have to be seperared by comma symbol i.e. `ls,-al,|,grep,Dockerfile`
#### Build Go binary
```sh
go build -o run run.go
```
#### Image build
```sh
docker build --tag webshell .
```
#### Docker image run
```sh
docker run -it --rm -p 9090:9090 --name webshell webshell
```
#### Usage
- Shell command output:
```sh
curl <localhost>:9090/cat,run.py
```
- Python3 script execution
```sh
curl <localhost>:9090/python3,run.py
```
- Python3 direct command execution
```sh
curl <localhost>:9090/python3,-c,'print("foo")'
```
