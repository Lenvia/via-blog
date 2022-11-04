
<template >
  <!-- 配置 breakpoint 属性即生效，视窗宽度小于 breakpoint 时 Sider 缩小为 collapsedWidth 宽度
  collapsed 展开-收起时的回调函数 -->
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="log">
      <span>{{ collapsed ? 'Blog' : 'My Blog' }}</span>
    </div>
    <!-- @click="goToPage" 方法将 每个菜单的key拼接到路径上跳转-->
    <a-menu theme="dark" mode="inline" @click="goToPage">
      <a-menu-item key="index">
        <a-icon type="dashboard"></a-icon>
        <span>仪表盘</span>
      </a-menu-item>

      <a-sub-menu>
        <!--  span slot 替代了直接在 a-sub-menu 后面加上 title="xxx" -->
        <span slot="title">
          <a-icon type="file" />
          <span>文章管理</span>
        </span>
        <a-menu-item key="addart">
          <a-icon type="form" />
          <span>写文章</span>
        </a-menu-item>
        <a-menu-item key="articles">
          <a-icon type="snippets" />
          <span>文章列表</span>
        </a-menu-item>
      </a-sub-menu>
      <a-menu-item key="commentlist">
        <a-icon type="message" />
        <span>评论管理</span>
      </a-menu-item>
      <a-menu-item key="categories">
        <a-icon type="book" />
        <span>分类列表</span>
      </a-menu-item>

      <a-menu-item key="users">
        <a-icon type="user" />
        <span>用户列表</span>
      </a-menu-item>

      <a-menu-item key="profile">
        <a-icon type="setting" />
        <span>个人设置</span>
      </a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data() {
    return {
      collapsed: false,
    }
  },
  methods: {
    goToPage(item) {
      this.$router.push('/' + item.key).catch((err) => err)
    },
  },
}
</script>

<style scoped>
.log {
  height: 32px;
  margin: 16px;
  background-color: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 17px;
}
</style>
