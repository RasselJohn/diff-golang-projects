# Projects 1

## Image Converter

***

### For running:

Local(in `src` folder):

```
go mod tidy 
go build
./img_converter
```

or with docker:

```
docker-compose -f deploy/docker-compose.yml -p project1 up --build
```

Service will be on `http://localhost:8080/`