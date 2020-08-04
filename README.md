# general-store
System that controls the transactions flow in a store, that can be applied to any one. Getting info from different sources and centralizing them.

Technologies:
- Backend Go
- Frontend Vue.js
- Database: DGraph

Getting Started

1. Clone this repo
2. Install DGraph
  1. Visit and [check][1] the steps of your os (This for linux):
    1. curl -sSf https://get.dgraph.io | bash
    2. /usr/local/bin/dgraph alpha --lru_mb 1024
    3. /usr/local/bin/dgraph zero
    4. go get github.com/dgraph-io/dgo
3. In one terminal:
  1. cd front-store
  2. npm i
  3. npm run serve
4. In other terminal:
  1. go get -u github.com/go-chi/chi
  2. got get github.com/rs/cors
  3. cd src
  4. go run main.go
5. Finish! Now test de development version


[@juanhenaoparra][2]  
<parjuag@gmail.com>  
Manizales, Caldas  

[1]: https://dgraph.io/downloads
[2]: https://twitter.com/juanhenaoparra
