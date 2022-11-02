
<template>
  <a-config-provier :locale="locale">
    <div id="app">
      <router-view />
    </div>
  </a-config-provier>
</template>

<script>
import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
export default {
  data() {
    return {
      locale: zhCN,
    }
  },
}
</script>
