import Cookies from 'js-cookie'

const UserInfoKey = 'naruto_user'

// 用户id
export function getUserId() {
  const user = Cookies.get(UserInfoKey)
  if (user) {
    return JSON.parse(user)['id']
  } else {
    return undefined
  }
}

// 用户名
export function getUserName() {
  const user = Cookies.get(UserInfoKey)
  if (user) {
    return JSON.parse(user)['username']
  } else {
    return undefined
  }
}

// 用户名
export function getUserCName() {
  const user = Cookies.get(UserInfoKey)
  if (user) {
    return JSON.parse(user)['cname']
  } else {
    return undefined
  }
}

// 用户全量信息
export function getUserInfo() {
  const user = Cookies.get(UserInfoKey)
  if (user) {
    return JSON.parse(user)
  } else {
    return undefined
  }
}

export function setUserInfo(UserInfo) {
  return Cookies.set(UserInfoKey, UserInfo, { expires: 1 })
}

export function removeUserInfo() {
  return Cookies.remove(UserInfoKey)
}

