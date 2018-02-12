import Vue from 'vue'
import Router from 'vue-router'
import TimeTable from '@/components/TimeTable'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'TimeTable',
      component: TimeTable
    }
  ]
})
