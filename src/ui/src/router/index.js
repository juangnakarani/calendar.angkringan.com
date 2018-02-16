import Vue from 'vue'
import Router from 'vue-router'
import TimeList from '@/components/TimeList'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/list',
      name: 'TimeList',
      component: TimeList
    }
  ]
})
