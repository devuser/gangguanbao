import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/orders',
    },
    {
      path: '/statistics',
      name: 'Statistics',
      component: () => import('@/views/Statistics.vue'),
      meta: { title: '统计分析', breadcrumb: ['首页', '统计分析'] },
    },
    {
      path: '/orders',
      name: 'Orders',
      component: () => import('@/views/OrderList.vue'),
      meta: { title: '订单查询', breadcrumb: ['首页', '订单查询'] },
    },
    {
      path: '/materials',
      name: 'Materials',
      component: () => import('@/views/MaterialList.vue'),
      meta: { title: '材质管理', breadcrumb: ['首页', '材质管理'] },
    },
    {
      path: '/specs',
      name: 'Specs',
      component: () => import('@/views/SpecList.vue'),
      meta: { title: '规格管理', breadcrumb: ['首页', '规格管理'] },
    },
    {
      path: '/vendors',
      name: 'Vendors',
      component: () => import('@/views/VendorList.vue'),
      meta: { title: '厂家管理', breadcrumb: ['首页', '厂家管理'] },
    },
  ],
});

export default router;
