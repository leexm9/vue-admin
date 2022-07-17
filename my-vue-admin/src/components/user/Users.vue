<template>
  <div>
    <!--面包屑导航-->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>用户管理</el-breadcrumb-item>
      <el-breadcrumb-item>用户列表</el-breadcrumb-item>
    </el-breadcrumb>

    <!--卡片视图区-->
    <el-card>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入内容" v-model="queryInfo.keyword" clearable>
            <template #append>
              <el-button icon="el-icon-search" @click="getUserList"></el-button>
            </template>
          </el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="dialogFormVisible = true">添加用户</el-button>
        </el-col>
      </el-row>
      <!--用户列表区域-->
      <el-table :data="userlist" border stripe>
        <el-table-column type="index"></el-table-column>
        <el-table-column label="姓名" prop="username"></el-table-column>
        <el-table-column label="邮箱" prop="email"></el-table-column>
        <el-table-column label="电话" prop="mobile"></el-table-column>
        <el-table-column label="角色" prop="roleName"></el-table-column>
        <el-table-column label="状态">
          <template v-slot="scope">
            <el-switch v-model="scope.row.mgState" @change="userStateChanged(scope.row)" />
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template v-slot="scope">
            <!--修改按钮-->
            <el-button type="primary" v-model="scope.row.id" size="mini" class="el-icon-edit"
              @click="showEditDialog(scope.row.id)" />
            <!--删除按钮-->
            <el-button type="danger" size="mini" class="el-icon-delete" @click="deleteUser(scope.row.id)" />
            <!--分配角色-->
            <el-tooltip effect="dark" content="分配角色" placement="top" :enterable="false">
              <el-button type="warning" size="mini" class="el-icon-setting" />
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination v-model:currentPage="queryInfo.pageIndex" v-model:page-size="queryInfo.pageSize"
        :page-sizes="[1, 2, 5, 10]" :small="small" :disabled="disabled" :background="background"
        layout="total, sizes, prev, pager, next, jumper" :total="total" @size-change="handleSizeChange"
        @current-change="handleCurrentChange" />
    </el-card>

    <!--添加用户-->
    <el-dialog v-model="dialogFormVisible" title="添加用户" width="50%" @close="addDialogClose">
      <el-form :model="addUserForm" :rules="addUserRules" ref="addUserFormRef" label-width="70px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="addUserForm.username" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="addUserForm.email" />
        </el-form-item>
        <el-form-item label="手机" prop="mobile">
          <el-input v-model="addUserForm.mobile" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogFormVisible = false" size="small">取消</el-button>
          <el-button type="primary" @click="addUser" size="small">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <!--编辑用户-->
    <el-dialog v-model="editUserDialogVisible" title="编辑用户" width=" 50%">
      <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="70px"
        @close="editUserDialogClosed">
        <el-form-item label="用户名">
          <el-input v-model="editForm.username" disabled />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="editForm.email" />
        </el-form-item>
        <el-form-item label="手机号" prop="mobile">
          <el-input v-model="editForm.mobile" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="editUserDialogVisible = false" size="small">取消</el-button>
          <el-button type="primary" @click="editUser" size="small">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data () {
    // 验证手机的规则
    const checkMobile = (rule, value, callback) => {
      const regMobile = /^(13[0-9]|14[5|7]|15[0|1|2|3|4|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$/
      if (regMobile.test(value)) {
        return callback()
      }
      callback(new Error('请输入合法的手机号'))
    }
    // 验证邮箱的规则
    const checkEmail = (rule, value, callback) => {
      const regEmail = /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/
      if (regEmail.test(value)) {
        return callback()
      }
      callback(new Error('请输入合法的邮箱'))
    }
    return {
      queryInfo: {
        keyword: '',
        pageIndex: 1,
        pageSize: 2
      },
      userlist: [],
      total: 0,
      dialogFormVisible: false,
      // 添加用户表单数据
      addUserForm: {
        username: '',
        email: '',
        mobile: ''
      },
      // 添加用户表单数据验证规则
      addUserRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 10, message: '用户名的长度在 3 ~ 10 个字符之间', trigger: 'blur' }
        ],
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { validator: checkEmail, trigger: 'blur' }
        ],
        mobile: [
          { required: true, message: '请输入手机号', trigger: 'blur' },
          { validator: checkMobile, trigger: 'blur' }
        ]
      },
      // 编辑用户对话框
      editUserDialogVisible: false,
      // 用户编辑对象
      editForm: {
        username: '',
        email: '',
        mobile: ''
      },
      // 编辑验证表单规则
      editFormRules: {
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { validator: checkEmail, trigger: 'blur' }
        ],
        mobile: [
          { required: true, message: '请输入手机号', trigger: 'blur' },
          { validator: checkMobile, trigger: 'blur' }
        ]
      }
    }
  },
  created () {
    this.getUserList()
  },
  methods: {
    async getUserList () {
      const { data: res } = await this.$http.get('/user/page', { params: this.queryInfo })
      if (!res.success) {
        return this.$message.error(res.msg)
      }
      this.userlist = res.data.list
      this.total = res.data.total
    },
    // 监听 page size 改变的事件
    handleSizeChange (newSize) {
      this.queryInfo.pageSize = newSize
      this.getUserList()
    },
    // 监听 页码值 改变的事件
    handleCurrentChange (newPage) {
      this.queryInfo.pageIndex = newPage
      console.log(newPage)
      this.getUserList()
    },
    // 监听 Switch 开关状态的变更
    async userStateChanged (userInfo) {
      const { data: res } = await this.$http.post('/user/changeState', { id: userInfo.id, mgState: '' + userInfo.mgState })
      if (!res.success) {
        // 修改失败将 Switch 状态重置为原来的状态
        userInfo.mgState = !userInfo.mgState
        return this.$message.error(res.msg)
      }
      this.$message.success('更新用户状态成功')
    },
    // 监听添加用户对话框的关闭事件
    addDialogClose () {
      this.$refs.addUserFormRef.resetFields()
    },
    // 点击按钮，添加新用户
    addUser () {
      this.$refs.addUserFormRef.validate(async valid => {
        // 校验不通过返回
        if (!valid) return
        // 发起网络请求
        const { data: res } = await this.$http.post('/user/add', this.addUserForm)
        if (!res.success) {
          return this.$message.error(res.msg)
        }
        this.$message.success('添加成功')
        // 隐藏对话框
        this.dialogFormVisible = false
        // 重新获取用户列表
        this.getUserList()
      })
    },
    // 编辑用户对话框
    async showEditDialog (id) {
      const { data: res } = await this.$http.get('/user/get', { params: { id: id } })
      if (!res.success) {
        return this.$message.error(res.msg)
      }
      this.editForm = res.data
      this.editUserDialogVisible = true
    },
    // 监听编辑用户关闭事件
    editUserDialogClosed () {
      this.$refs.editFormRef.resetFields()
    },
    // 编辑用户
    editUser () {
      this.$refs.editFormRef.validate(async vaild => {
        if (!vaild) return
        const { data: res } = await this.$http.post('/user/edit', this.editForm)
        if (!(res.success && res.data)) {
          return this.$message.error(res.msg)
        }
        this.$message.success('修改成功')
        this.editUserDialogVisible = false
        this.getUserList()
      })
    },
    // 删除用户
    deleteUser (id) {
      this.$confirm('确定要删除该用户?', '提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }
      ).then(async () => {
        const { data: res } = await this.$http.get('/user/delete', { params: { id: id } })
        if (res.success && res.data) {
          this.$message.success('删除成功')
          this.getUserList()
        } else {
          this.$message.error(res.msg)
        }
      }).catch(() => {
        this.$message.info('取消删除')
      })
    }
  }
}
</script>

<style lang="less" scoped>
</style>
