import layoutHeaderAside from '@/layout/header-aside'

const meta = { requiresAuth: true }

export default {
  path: '/member',
  name: 'member',
  meta,
  redirect: { name: 'member-index' },
  component: layoutHeaderAside,
  children: (pre => [
    { path: 'index', name: 'member-index', component: () => import('@/pages/member/index'), meta: { meta, title: '会员列表' } },
    { path: 'edit/:id', name: 'member-edit', component: () => import('@/pages/member/edit'), meta: { meta, title: '会员编辑' } }
  ])('member-')
}