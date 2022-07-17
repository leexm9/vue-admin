<template>
  <div class="login_container">
    <div class="login_box">
      <!-- 头像区域 -->
      <div class="avatar_box">
        <img src="../assets/logo.png" alt="">
      </div>
      <!-- 登录表单区域 -->
      <el-form ref="loginFormRef" :model="loginForm" :rules="loginFormRules" label-width="0px" class="login_form">
        <!-- 用户名 -->
        <el-form-item prop="username">
          <el-input v-model="loginForm.username" prefix-icon="iconfont icon-user" />
        </el-form-item>
        <!-- 密码 -->
        <el-form-item prop="password">
          <el-input v-model="loginForm.password" type="password" prefix-icon="iconfont icon-3702mima" />
        </el-form-item>
        <!-- 按钮区域 -->
        <el-row justify="end">
          <el-form-item class="login_btn">
            <el-button type="primary" @click="Login">登录</el-button>
            <el-button type="info" @click="resetLoginForm">重置</el-button>
          </el-form-item>
        </el-row>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      // 登录表单的数据绑定对象
      loginForm: {
        username: 'admin',
        password: '123'
      },
      // 表单的验证规则对象
      loginFormRules: {
        // 用户名是否合法
        username: [
          { required: true, message: '请输入登录名称', tigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 10 个字符', tigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入登录密码', tigger: 'blur' },
          { min: 6, max: 15, message: '长度在 6 到 15 个字符', tigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    resetLoginForm () {
      this.$refs.loginFormRef.resetFields()
    },
    Login () {
      this.$refs.loginFormRef.validate(async valid => {
        if (!valid) {
          return
        }
        const { data: result } = await this.$http.post('/login', this.loginForm)
        if (!result.success) {
          return this.$message.error('登录失败')
        }
        this.$message.success('登录成功')
        window.sessionStorage.setItem('token', result.data.token)
        this.$router.push('/home')
      })
    }
  }
}
</script>

<style lang="less" scoped>
.login_container {
  background-color: #2b4b6b;
  height: 100vh;
}
.login_box {
  // 宽450像素
  width: 450px;
  // 高300像素
  height: 300px;
  // 背景颜色
  background-color: #fff;
  // 圆角边框3像素
  border-radius: 3px;
  // 绝对定位
  position: absolute;
  // 距离左则50%
  left: 50%;
  // 上面距离50%
  top: 50%;
  // 进行位移，并且在横轴上位移-50%，纵轴也位移-50%
  transform: translate(-50%, -50%);
  .avatar_box {
    // 盒子高度130像素
    height: 130px;
    // 宽度130像素
    width: 130px;
    // 边框线1像素 实线
    border: 1px solid #eee;
    // 圆角
    border-radius: 50%;
    // 内padding
    padding: 10px;
    // 添加阴影 向外扩散10像素 颜色ddd
    box-shadow: 0 0 10px #ddd;
    // 绝对定位
    position: absolute;
    // 距离左则
    left: 50%;
    // 位移
    transform: translate(-50%, -50%);
    // 背景
    background-color: #fff;
    img {
      width: 100%;
      height: 100%;
      border-radius: 50%;
      background-color: #eee;
    }
  }
}
.login_form {
  position: absolute;
  bottom: 0;
  width: 100%;
  padding: 0 20px;
  box-sizing: border-box;
}
.login_btn {
  // 设置弹性布局
  display: flex;
  // 横轴上尾部对齐
  justify-content: flex-end;
}
</style>
