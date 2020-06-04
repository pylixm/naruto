import { login, logout } from '@/api/user'
import { setUserInfo, removeUserInfo, getUserCName, getUserName } from '@/utils/auth'
import { resetRouter } from '@/router'

const getDefaultState = () => {
  return {
    username: getUserName(),
    cname: getUserCName(),
    avatar: ''
  }
}

const state = getDefaultState()

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_USER_NAME: (state, username) => {
    state.username = username
  },
  SET_USER_CNAME: (state, cname) => {
    state.cname = cname
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  }
}

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { username, password } = userInfo
    console.log(userInfo)
    return new Promise((resolve, reject) => {
      console.log(username, password)
      login({ username: username.trim(), password: password }).then(response => {
        // 直接保存用户信息
        const { content } = response
        const username = content.name
        commit('SET_USER_NAME', username)
        commit('SET_USER_CNAME', content.cname)

        setUserInfo(content)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({ commit, state }) {
    return new Promise((resolve, reject) => {
      logout(state.token).then(() => {
        removeUserInfo()
        resetRouter()
        commit('RESET_STATE')
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

