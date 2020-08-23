import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import Home from '@/components/Home.vue';
import Vendors from '@/components/Vendors.vue';

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/vendors',
    name: 'VendorList',
    component: Vendors,
  },
  {
    path: '/settings',
    name: 'Settings',
    component: () => import(/* webpackChunkName: "Settings" */ '@/components/Settings.vue'),
  },
];

const router = new VueRouter({
  routes,
});

export default router;
