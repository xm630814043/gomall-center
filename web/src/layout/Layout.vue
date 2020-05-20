<template>
  <div>
    <el-container>
      <el-aside :width="asideWidth">
        <el-menu
          class="el-menu-vertical-demo"
          :collapse-transition="false"
          :default-active="activeIndex"
          :collapse="isCollapse"
          :router="true"
        >
          <el-menu-item disabled>
            <i class="el-icon-upload"></i>
            <span slot="title">Gomall中台</span>
          </el-menu-item>
          <el-menu-item @click="collapseHandle" style="text-align:center">
            <i :class="collapseIcon"></i>
          </el-menu-item>
          <el-menu-item index="/">
            <i class="el-icon-menu"></i>
            <span slot="title">桌面</span>
          </el-menu-item>
          <el-menu-item index="/cms">
            <i class="el-icon-document"></i>
            <span slot="title">栏目管理</span>
          </el-menu-item>
          <el-menu-item index="/company">
            <i class="el-icon-star-off"></i>
            <span slot="title">企业管理</span>
          </el-menu-item>
          <el-menu-item index="/goods">
            <i class="el-icon-goods"></i>
            <span slot="title">产品审核</span>
          </el-menu-item>
          <el-menu-item index="/goods/property">
             <i class="el-icon-goods"></i>
             <span slot="title">产品属性管理</span>
             </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-setting"></i>
            <span slot="title">系统设置</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-main style="padding:0px">
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data() {
    return {
      activeIndex: '/',
      isCollapse: false,
      collapseIcon: 'el-icon-d-arrow-left',
      asideWidth: '200px'
    }
  },
  methods: {
    collapseHandle() {
      this.isCollapse = !this.isCollapse
      this.asideWidth = this.isCollapse ? '65px' : '200px'
      this.collapseIcon = this.isCollapse
        ? 'el-icon-d-arrow-right'
        : 'el-icon-d-arrow-left'
    },

  }
}
</script>
<style>
.el-menu {
  min-height: 400px;
  height: 100vh;
}
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 199px;
}
</style>
