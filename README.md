## Credit

#### Dependencies
    golang, nodejs, postgresql

#### API server

    cp config.sample.yaml config.yaml

    go run main.go

#### Web assets

    (cd web && npm install)
    
    npm --prefix web run dev

#### Open

    localhost:8001