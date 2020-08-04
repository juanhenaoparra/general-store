import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    mainRoute: `http://${location.hostname}:3333`,
    buyerRoute: "buyer",
    syncRoute: "sync",
    allBuyers: [],
    syncResponse : undefined,
    currentProfile: undefined,
  },
  getters: {
    getIpsByCurrent: (state) => {
      if(state.currentProfile){
        let listOthers = state.currentProfile['~by_buyer'].map(transaction => {
          return {
            address: transaction.since_ip.address,
            buyers: transaction.since_ip['~since_ip'].map(postTransaction => {
              return postTransaction.by_buyer;
            })
        }
        });

        return listOthers;
      }
    },
  },
  mutations: {
    fillBuyers(state, value) {
      state.allBuyers.push(...value);
    },
    fillSync(state, value) {
      state.syncResponse = value;
    },
    fillCurrentProfile(state, value) {
      state.currentProfile = value;
    },
  },
  actions: {
    getAllBuyers: function({state, commit}, [first, page]) {
      let uri = `${state.mainRoute}/${state.buyerRoute}?first=${first}&offset=${page}`

      Vue.axios.get(uri).then(res => {
        let data = res.data.buyers;
        commit('fillBuyers', data)
      })
    },
    seeMyProfile: function ({state, commit}, [id, first, offset]) {
      let uri = `${state.mainRoute}/${state.buyerRoute}/${id}?first=${first}&offset=${offset}`

      Vue.axios.get(uri).then(res => {
        let parsedProfile = JSON.parse(res.data).buyers[0];
        commit('fillCurrentProfile', parsedProfile);
      })
    },
    syncData: function ({ state, commit }, date) {
      let dateNow = "";

      if (date){
        dateNow = date;
      }else{
        dateNow = new Date();
      }

      let withoutHours = new Date(dateNow.getFullYear(), dateNow.getMonth(),dateNow.getDate());
      let dateParsed = parseInt(Date.parse(withoutHours)/1000);


      let uri = `${state.mainRoute}/${state.syncRoute}?date=${dateParsed}`

      Vue.axios.get(uri).then(res => {
        commit('fillSync', res.data)
      })
    }
  },
  modules: {
  }
})
