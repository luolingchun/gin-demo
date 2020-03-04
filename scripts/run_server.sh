swag init --dir=src --output=src/docs

go build -o ./bin/gin-demo.exe gin-demo/src

./bin/gin-demo.exe