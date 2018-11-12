// 菜单 顶栏
export default [
  { path: '/index', title: '首页', icon: 'home' },
  {
    title: '分类',
    icon: 'folder-o',
    children: [
      { path: '/category/index', title: '列表' },
      { path: '/category/create', title: '添加' }
    ]
  },
  { path: '/questions', title: '问题', icon: 'home' },
  { path: '/answers', title: '回答', icon: 'home' },
  { path: '/member', title: '会员', icon: 'home' }
]
