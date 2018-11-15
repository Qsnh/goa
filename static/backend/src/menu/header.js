// 菜单 顶栏
export default [
  { path: '/index', title: '首页', icon: 'home' },
  {
    title: '分类',
    icon: 'navicon',
    children: [
      { path: '/category/index', title: '列表' },
      { path: '/category/create', title: '添加' }
    ]
  },
  { path: '/questions', title: '问题', icon: 'envelope' },
  { path: '/answers', title: '回答', icon: 'comment' },
  { path: '/member/index', title: '会员', icon: 'users' },
  { path: '/settings', title: '系统配置', icon: 'cogs' }
]
