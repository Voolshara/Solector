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
            path: '/test',
            component: () => import('./components/Calc_forms/form-response')
        }
    ]
})