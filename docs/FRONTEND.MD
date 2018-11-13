# 前端开发

### 安装依赖

```
npm install
```

### 运行

```
npm run watch
```

有下面这些命令：

| 命令 | 说明 |
| --- | --- |
| `npm run watch` | 编译前端文件，并监听8080端口，本地文件发生变化会自动热更新 |
| `npm run dev` | 开发环境下的编译 |
| `npm run build` | 生产环境下的编译 |


### 自定义Vue组件

> Vue的相关资源都存放在 `views/assets` 目录下面。

首先，在 `views/assets/components` 目录下面创建一个 `.vue` 文件，比如说这里叫做：`demo.vue`，内容如下：

```vue
<template>
    <div><a>Hello World.</a></div>
</template>
<script>
export default {}
</script>
```

然后我们需要在 `views/assets/index.js` 文件中注册这个组件，需要添加下面的内容：

```vue
import DemoComponent from './components/demo.vue'

Vue.component('demo-dmeo', DemoComponent)
```

这里需要重新编译下，如果开启了 `serve` 就直接忽略。

然后我们就可以在模板中这样使用自定义的组件了：

```
<demo-demo></demo-demo>
```

### 前端的CSS样式

GOA的样式都写在 `views/assets/scss/app.scss` 文件中，使用的是 SCSS 语法。