import Vue from 'vue'
import Router from 'vue-router'
import TimeList from '@/components/TimeList'
import MyAccount from '@/components/MyAccount'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'TimeList',
      component: TimeList
    },
    {
      path: '/myaccount',
      name: 'MyAccount',
      component: MyAccount
    }
  ]
})
