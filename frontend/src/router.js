import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/routes/home.vue'

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [
        {
            path:'/',
            component: Home
        },
        {
            path: '/selection',
            component: () => import('./routes/selection.vue')
        },
        {
            path: '/feedback',
            component: () => import('./components/Fill/GoogleForm')
        },
        {
            path: "*",
            component: () => import('./components/NFerr'),
        }
    ]
})