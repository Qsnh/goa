import layoutHeaderAside from '@/layout/header-aside'

const meta = { requiresAuth: true }

export default {
  path: '/category',
  name: 'category',
  meta,
  redirect: { name: 'category-index' },
  component: layoutHeaderAside,
  children: (pre => [
    { path: 'index', name: 'category-index', component: () => import('@/pages/category/index'), meta: { meta, title: '分类列表' } },
    { path: 'create', name: 'category-create', component: () => import('@/pages/category/create'), meta: { meta, title: '添加分类' } },
    { path: 'edit/:id', name: 'category-edit', component: () => import('@/pages/category/edit'), meta: { meta, title: '编辑分类' } }
  ])('category-')
}
