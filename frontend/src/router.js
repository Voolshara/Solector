import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/routes/marketplace.vue'
import VueAnalytics from 'vue-analytics';
import VueMeta from 'vue-meta';

Vue.use(Router)

Vue.use(VueMeta)

Vue.use(VueAnalytics, {
    id: 'G-MX568C5JY5',
    Router,
});

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
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
            path: '/instruction',
            component: () => import('./routes/instruction')
        },
        {
            path: "*",
            component: () => import('./routes/NFerr'),
        },
    ]
})