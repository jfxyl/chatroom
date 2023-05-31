const state = {
    currentMenu : ''
}

const mutations = {
    SET_CURRENT_MENU(state,menu) {
        state.currentMenu = menu
    },
}

const actions = {
    SET_CURRENT_MENU(context,menu) {
        context.commit('SET_CURRENT_MENU',menu)
    },
}

export default {
    state,
    mutations,
    actions
}
