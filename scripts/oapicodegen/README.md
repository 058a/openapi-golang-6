~/go/bin/oapi-codegen -generate types,server,spec -package hello ./../../api/hello.yaml > ./../../internal/infrastructure/oapicodegen/hello/hello.go && go mod tidy

~/go/bin/oapi-codegen -generate types,server,spec -package stock ./../../api/stock.yaml > ./../../internal/infrastructure/oapicodegen/stock/stock.go && go mod tidy