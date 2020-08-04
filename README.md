# general-store
System that controls the transactions flow in a store, that can be applied to any one. Getting info from different sources and centralizing them.

Technologies:
- Backend Go
- Frontend Vue.js
- Database: DGraph

Getting Started

1. Clone this repo
2. Install DGraph
    - Visit and [check][1] the steps of your os (This for linux):
    - curl -sSf https://get.dgraph.io | bash
    - /usr/local/bin/dgraph alpha --lru_mb 1024
    - /usr/local/bin/dgraph zero
    - go get github.com/dgraph-io/dgo
3. In one terminal:
    - cd front-store
    - npm i
    - npm run serve
4. In other terminal:
    - go get -u github.com/go-chi/chi
    - got get github.com/rs/cors
    - cd src
    - go run main.go
5. Finish! Now test de development version


[@juanhenaoparra][2]  
<parjuag@gmail.com>  
Manizales, Caldas  

[1]: https://dgraph.io/downloads
[2]: https://twitter.com/juanhenaoparra
