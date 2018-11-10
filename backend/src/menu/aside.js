// 菜单 侧边栏
export default [
  { path: '/index', title: '首页', icon: 'home' },
  {
    title: '分类',
    icon: 'folder-o',
    children: [
      { path: '/demo/page1', title: '列表' },
      { path: '/demo/page2', title: '添加' }
    ]
  },
  { path: '/questions', title: '问题', icon: 'home' },
  { path: '/answers', title: '回答', icon: 'home' },
  { path: '/members', title: '会员', icon: 'home' }
]
