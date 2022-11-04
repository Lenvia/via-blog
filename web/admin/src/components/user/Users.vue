<template>
  <div>
    <!-- a-card 表框 -->
    <a-card>
      <!-- :gutter 是组件间距 -->
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.username"
            placeholder="输入用户名查找"
            enter-button
            allowClear
            @search="searchUser"
          />
        </a-col>
        <a-col :span="4">
          <!-- addUserVisible 变量控制添加窗口的显示 -->
          <a-button type="primary" @click="addUserVisible = true"
            >新增</a-button
          >
        </a-col>
      </a-row>

      <!--
        dataSource、pagination、columns等号后的的对象对应 script 里data里的对象
        @change 对应 method里的方法。其实替代了 页码改变和页大小改变
       -->
      <a-table
        rowKey="ID"
        :columns="columns"
        :pagination="pagination"
        :dataSource="userlist"
        bordered
        @change="handleTableChange"
      >
        <!--
          我的理解是，span是个名字叫role的插槽，然后其获取的元素用data表示
          在script的data的column中，scopedSlots: { customRender: 'role' }, 将这个插槽插入了role这一列
          所以column中的role那一列每个格子里都是 <span>管理员/订阅者  </span>
         -->
        <span slot="role" slot-scope="data">{{
          data === 1 ? '管理员' : '订阅者'
        }}</span>
        <!--
          下面这个template仍然是个插槽，插入到了column中的action那一列
          由于column的action并没有绑定属性，所以data接收的值就是完整的结构体
         -->
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <!-- 点击后显示编辑框，并请求已有数据 -->
            <a-button
              type="primary"
              icon="edit"
              style="margin-right: 15px"
              @click="editUser(data.ID)"
              >编辑</a-button
            >
            <a-button
              type="danger"
              icon="delete"
              style="margin-right: 15px"
              @click="deleteUser(data.ID)"
              >删除</a-button
            >
            <a-button type="info" icon="info" @click="ChangePassword(data.ID)"
              >修改密码</a-button
            >
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 新增用户区域 -->
    <a-modal
      closable
      title="新增用户"
      :visible="addUserVisible"
      width="60%"
      @ok="addUserOk"
      @cancel="addUserCancel"
      destroyOnClose
    >
      <a-form-model :model="newUser" :rules="addUserRules" ref="addUserRef">
        <a-form-model-item label="用户名" prop="username">
          <a-input v-model="newUser.username"></a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback label="密码" prop="password">
          <a-input-password v-model="newUser.password"></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkpass">
          <a-input-password v-model="newUser.checkpass"></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑用户区域 -->
    <a-modal
      closable
      destroyOnClose
      title="编辑用户"
      :visible="editUserVisible"
      width="60%"
      @ok="editUserOk"
      @cancel="editUserCancel"
    >
      <a-form-model :model="userInfo" :rules="userRules" ref="addUserRef">
        <a-form-model-item label="用户名" prop="username">
          <a-input v-model="userInfo.username"></a-input>
        </a-form-model-item>
        <a-form-model-item label="是否为管理员">
          <a-switch
            :checked="IsAdmin"
            checked-children="是"
            un-checked-children="否"
            @change="adminChange"
          />
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 修改密码 -->
    <a-modal
      closable
      title="修改密码"
      :visible="changePasswordVisible"
      width="60%"
      @ok="changePasswordOk"
      @cancel="changePasswordCancel"
      destroyOnClose
    >
      <a-form-model
        :model="changePassword"
        :rules="changePasswordRules"
        ref="changePasswordRef"
      >
        <a-form-model-item has-feedback label="密码" prop="password">
          <a-input-password
            v-model="changePassword.password"
          ></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkpass">
          <a-input-password
            v-model="changePassword.checkpass"
          ></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
import day from 'dayjs'

const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center',
  },
  {
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center',
  },
  {
    title: '注册时间',
    dataIndex: 'CreatedAt',
    width: '20%',
    key: 'CreatedAt',
    align: 'center',
    customRender: (val) => {
      return val ? day(val).format('YYYY年MM月DD日 HH:mm') : '暂无'
    },
  },
  {
    title: '角色',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    align: 'center',
    scopedSlots: { customRender: 'role' },
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' },
  },
]

export default {
  data() {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true, // 可以改变页码
        showTotal: (total) => `共${total}条`,
      },
      userlist: [],
      // 用于编辑用户
      userInfo: {
        username: '',
        password: '',
        role: 2,
        checkPass: '',
      },
      // 用于添加用户
      newUser: {
        username: '',
        password: '',
        role: 2,
        checkPass: '',
      },
      // 用于修改密码
      changePassword: {
        id: 0,
        password: '',
        checkPass: '',
      },
      columns,
      // 用于获取用户列表/查询
      queryParam: {
        // 请求的参数
        username: '',
        pagesize: 5,
        pagenum: 1,
      },
      editVisible: false,
      // 编辑用户的输入验证
      userRules: {
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.username === '') {
                callback(new Error('请输入用户名'))
              }
              if (
                [...this.userInfo.username].length < 4 ||
                [...this.userInfo.username].length > 12
              ) {
                callback(new Error('用户名应当在4到12个字符之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.password === '') {
                callback(new Error('请输入密码'))
              }
              if (
                [...this.userInfo.password].length < 6 ||
                [...this.userInfo.password].length > 20
              ) {
                callback(new Error('密码应当在6到20位之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.checkpass === '') {
                callback(new Error('请输入密码'))
              }
              if (this.userInfo.password !== this.userInfo.checkpass) {
                callback(new Error('密码不一致，请重新输入'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
      },
      // 添加用户的输入验证
      addUserRules: {
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.username === '') {
                callback(new Error('请输入用户名'))
              }
              if (
                [...this.newUser.username].length < 4 ||
                [...this.newUser.username].length > 12
              ) {
                callback(new Error('用户名应当在4到12个字符之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.password === '') {
                callback(new Error('请输入密码'))
              }
              if (
                [...this.newUser.password].length < 6 ||
                [...this.newUser.password].length > 20
              ) {
                callback(new Error('密码应当在6到20位之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.checkpass === '') {
                callback(new Error('请输入密码'))
              }
              if (this.newUser.password !== this.newUser.checkpass) {
                callback(new Error('密码不一致，请重新输入'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
      },
      changePasswordRules: {
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.changePassword.password === '') {
                callback(new Error('请输入密码'))
              }
              if (
                [...this.changePassword.password].length < 6 ||
                [...this.changePassword.password].length > 20
              ) {
                callback(new Error('密码应当在6到20位之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.changePassword.checkpass === '') {
                callback(new Error('请输入密码'))
              }
              if (
                this.changePassword.password !== this.changePassword.checkpass
              ) {
                callback(new Error('密码不一致，请重新输入'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
      },
      editUserVisible: false,
      addUserVisible: false,
      changePasswordVisible: false,
    }
  },
  created() {
    // 生命周期函数
    this.getUserList()
  },
  computed: {
    IsAdmin: function () {
      if (this.userInfo.role === 1) {
        return true
      } else {
        return false
      }
    },
  },
  methods: {
    // 获取用户列表
    async getUserList() {
      const { data: res } = await this.$http.get('admin/user', {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        },
      })
      if (res.status !== 200) {
        // if (res.status === 1004 || 1005 || 1006 || 1007) {
        // token相关错误
        if ([1004, 1005, 1006, 1007].includes(res.status)) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.userlist = res.data
      this.pagination.total = res.total

      // console.log(res)
    },

    // 搜索用户
    async searchUser() {
      this.queryParam.pagenum = 1
      this.pagination.current = 1
      const { data: res } = await this.$http.get('admin/users', {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        },
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.userlist = res.data
      this.pagination.total = res.total
    },

    // 更改分页 a-table @change 调用
    handleTableChange(pagination, filters, sorter) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getUserList()
    },
    // 删除用户
    deleteUser(id) {
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定要删除该用户吗？一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`user/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getUserList()
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        },
      })
    },
    // 新增用户
    addUserOk() {
      // 把输入的数据绑定到newUser对象，再赋值到请求参数内，向后端发送post请求
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('user/add', {
          username: this.newUser.username,
          password: this.newUser.password,
          role: this.newUser.role,
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$refs.addUserRef.resetFields()
        this.addUserVisible = false
        this.$message.success('添加用户成功')
        this.getUserList()
      })
    },
    addUserCancel() {
      this.$refs.addUserRef.resetFields()
      this.addUserVisible = false
      this.$message.info('新增用户已取消')
    },
    adminChange(checked) {
      if (checked) {
        this.userInfo.role = 1
      } else {
        this.userInfo.role = 2
      }
    },
    // 编辑用户
    async editUser(id) {
      this.editUserVisible = true
      const { data: res } = await this.$http.get(`user/${id}`)
      this.userInfo = res.data
      this.userInfo.id = id
    },
    editUserOk() {
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`user/${this.userInfo.id}`, {
          username: this.userInfo.username,
          role: this.userInfo.role,
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.editUserVisible = false
        this.$message.success('更新用户信息成功')
        this.getUserList()
      })
    },
    editUserCancel() {
      this.$refs.addUserRef.resetFields()
      this.editUserVisible = false
      this.$message.info('编辑已取消')
    },

    // 修改密码
    async ChangePassword(id) {
      this.changePasswordVisible = true
      const { data: res } = await this.$http.get(`user/${id}`)
      this.changePassword.id = id
    },
    changePasswordOk() {
      this.$refs.changePasswordRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(
          `admin/changepwd/${this.changePassword.id}`,
          {
            password: this.changePassword.password,
          }
        )
        if (res.status !== 200) return this.$message.error(res.message)
        this.changePasswordVisible = false
        this.$message.success('修改密码成功')
        this.$refs.changePasswordRef.resetFields()
        this.getUserList()
      })
    },
    changePasswordCancel() {
      this.$refs.changePasswordRef.resetFields()
      this.changePasswordVisible = false
      this.$message.info('已取消')
    },
  },
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
