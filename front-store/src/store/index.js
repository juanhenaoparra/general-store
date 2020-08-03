import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    mainRoute: "http://localhost:3333",
    buyerRoute: "buyer",
  },
  mutations: {
  },
  actions: {
    getAllBuyers: function({state}, [first, page]) {
      let uri = `${state.mainRoute}/${state.buyerRoute}?first=${first}&offset=${page}`

      Vue.axios.get(uri).then(res => {
        console.log(res.data);
      })
    },
    getBuyerProfile: function ({state}, [id, first, offset]) {
      let uri = `${state.mainRoute}/${state.buyerRoute}/${id}?first=${first}&offset=${offset}`

      Vue.axios.get(uri).then(res => {
        console.log(JSON.parse(res.data).buyers[0]);
      })
    }
  },
  modules: {
  }
})
