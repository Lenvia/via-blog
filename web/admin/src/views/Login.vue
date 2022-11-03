<template>
  <div class="container">
    <div class="loginBox">
      <!-- :rules 表单验证; :model 数据绑定 -->
      <a-form-model
        ref="loginFormRef"
        :rules="rules"
        :model="formdata"
        class="loginForm"
      >
        <!-- prop 表单验证绑定（对应的rule） -->
        <a-form-model-item prop="username">
          <!-- "v-model 具体的数据绑定-->
          <a-input v-model="formdata.username" placeholder="请输入用户名">
            <a-icon
              slot="prefix"
              type="user"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>

        <a-form-model-item prop="password">
          <a-input
            v-model="formdata.password"
            placeholder="请输入密码"
            type="password"
            v-on:keyup.enter="login"
          >
            <a-icon
              slot="prefix"
              type="lock"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>

        <a-form-model-item class="loginBtn">
          <a-button type="primary" style="margin: 10px" @click="login"
            >登录</a-button
          >
          <a-button type="info" style="margin: 10px" @click="resetForm"
            >取消</a-button
          >
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      formdata: {
        username: '',
        password: '',
      },

      rules: {
        username: [
          { required: true, message: '请输入用户名', tigger: 'blur' },
          { min: 4, max: 12, message: '用户名必须在4到12个字符之间' },
        ],
        password: [
          { required: true, message: '请输入密码', tigger: 'blur' },
          { min: 6, max: 20, message: '密码必须在6到20个字符之间' },
        ],
      },
    }
  },
  methods: {
    // 使用 message 前要先在 ant-ui.js 中设置好 Vue.prototype.$message
    resetForm() {
      this.$refs.loginFormRef.resetFields()
    },
    login() {
      this.$refs.loginFormRef.validate(async (valid) => {
        if (!valid) return this.$message.error('输入不合法，请重新输入')
        const { data: res } = await this.$http.post('login', this.formdata)

        if (res.status !== 200) {
          return this.$message.error(res.message)
        }
        window.sessionStorage.setItem('token', res.token)
        this.$router.push('/index')
      })
    },
  },
}
</script>

<style scoped>
.container {
  height: 100%;
  background-color: #202c34;
}
.loginBox {
  width: 450px;
  height: 300px;
  background-color: #fff;
  position: absolute;
  left: 70%;
  top: 50%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}

.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;
}

.loginBtn {
  display: flex;
  justify-content: flex-end;
}
</style>
