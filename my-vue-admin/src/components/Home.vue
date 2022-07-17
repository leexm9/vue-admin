<template>
  <el-container class="home-container">
    <!-- 头部区域 -->
    <el-header>
      <div>
        <img src="../assets/heima.png" alt="" />
        <span>电商后台管理系统</span>
      </div>
      <el-button type="info" @click="logout">退出</el-button>
    </el-header>
    <!-- 页面主体区域 -->
    <el-container>
      <!-- 侧边栏 -->
      <el-aside :width="isCollapse ? '64px':'200px'" class="el_aside">
        <div class="toggle-button" @click="toggleCollapse">|||</div>
        <el-menu background-color="#333744" text-color="#fff" active-text-color="#409Eff" :unique-opened="true"
          :collapse="isCollapse" :collapse-transition="false" :router="true" :default-active="activePath">
          <!-- 一级菜单 -->
          <el-submenu :index="item.id+''" v-for="item in menulist" :key="item.id">
            <!-- 一级菜单模板区域 -->
            <template v-slot:title>
              <el-icon :class="iconObj[item.id]"></el-icon>
              <span>{{item.authName}}</span>
            </template>
            <!-- 二级菜单 -->
            <el-menu-item :index="'/'+subItem.path" v-for="subItem in item.children" :key="subItem.id"
              @click="saveNovState('/'+subItem.path)">
              <template v-slot:title>
                <i class="el-icon-menu"></i>
                <span>{{subItem.authName}}</span>
              </template>
            </el-menu-item>
          </el-submenu>
        </el-menu>
      </el-aside>
      <!-- 右侧内容主体区域 -->
      <el-main class="el_main">
        <!-- 路由占位符 -->
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  data () {
    return {
      // 左侧菜单数据对象
      menulist: [],
      // 字体图标库对象
      iconObj: {
        101: 'iconfont icon-users',
        102: 'iconfont icon-tijikongjian',
        103: 'iconfont icon-shangpin',
        104: 'iconfont icon-danju',
        105: 'iconfont icon-baobiao'
      },
      // 是否折叠
      isCollapse: false,
      // 被激活高亮的链接地址
      activePath: ''
    }
  },
  created () {
    this.getMenuList()
    this.activePath = window.sessionStorage.getItem('activePath')
  },
  methods: {
    logout () {
      window.sessionStorage.clear()
      this.$router.push('/login')
    },
    async getMenuList () {
      const { data: res } = await this.$http.get('/menus')
      if (!res.success) {
        return this.$message.error(res.msg)
      }
      this.menulist = res.data
    },
    toggleCollapse () {
      this.isCollapse = !this.isCollapse
    },
    saveNovState (activePath) {
      window.sessionStorage.setItem('activePath', activePath)
      this.activePath = window.sessionStorage.getItem('activePath')
    }
  }
}
</script>

<style lang="less" scoped>
.home-container {
  height: 100%;
}
.el-header {
  background-color: #363d40;
  // 给头部设置一下弹性布局
  display: flex;
  // 让它贴标左右对齐
  justify-content: space-between;
  // 清空图片左侧padding
  padding-left: 0;
  // 按钮居中
  align-items: center;
  // 文本颜色
  color: #fff;
  font-size: 20px;
  > div {
    display: flex;
    align-items: center;
    // 给文本和图片添加间距，使用类选择器
    span {
      margin-left: 15px;
    }
  }
}

.el-aside {
  background-color: #313743;
  .el-menu {
    border-right: none;
  }
}
.el-main {
  background-color: #e9edf1;
}
.iconfont {
  margin-right: 8px;
}
.toggle-button {
  // 添加背景颜色
  background-color: #4a5064;
  // 设置文本大小
  font-size: 10px;
  // 设置文本行高
  line-height: 24px;
  // 设置文本颜色
  color: #fff;
  // 设置文本居中
  text-align: center;
  // 设置文本间距
  letter-spacing: 0.2em;
  // 设置鼠标悬浮变小手效果
  cursor: pointer;
}
</style>
