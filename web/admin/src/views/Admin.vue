<template>
  <a-layout class="container">
    <Nav></Nav>
    <a-layout>
      <a-layout-header class="headerBtn">
        <Header></Header>
      </a-layout-header>
      <a-layout-content>
        <!-- 通过绑定一个fullPath，可以识别当前页面路由的完整地址，
        当地址发生改变(包括参数改变)则重新渲染页面(例如动态路由参数的变化） -->
        <router-view :key="$route.fullPath"></router-view>
      </a-layout-content>
      <a-layout-ooter class="footer">
        <Footer></Footer>
      </a-layout-ooter>
    </a-layout>
  </a-layout>
</template>

<script>
import Nav from '../components/admin/Nav'
import Footer from '../components/admin/Footer'
import Header from '../components/admin/Header'

export default {
  components: { Nav, Footer, Header },
}
</script>

<style scoped>
.container {
  height: 100%;
}
.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
