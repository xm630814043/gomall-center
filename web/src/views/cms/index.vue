<template>
  <div>
    <Split style="height: 100vh">
      <SplitArea :size="25"
                 class="to-do-list">
        <div class="list-title"></div>
        <div>
          <el-collapse v-model="activeIndex">
            <el-collapse-item v-for="c in channels"
                              :name="c.id"
                              :key="c.id">
              <span slot="title"
                    style="margin-left:20px">{{ c.name }}</span>
              <div v-for="s in c.children"
                   :key="s.id"
                   style="margin:20px">
                {{ s.name }}
              </div>
            </el-collapse-item>
          </el-collapse>
        </div>
      </SplitArea>
      <SplitArea :size="75">
        panel right
      </SplitArea>
    </Split>
  </div>
</template>
<script>
export default {
  name: 'cms',
  data () {
    return {
      activeIndex: '1',
      channels: [
        {
          id: 1,
          name: '首页',
          children: [
            {
              id: 112,
              type: 'product',
              name: '推荐产品'
            },
            {
              id: 110,
              type: 'product',
              name: '推荐店铺'
            }
          ]
        },
        {
          id: 2,
          name: '优选店铺',
          children: [
            {
              id: 113,
              type: 'product',
              name: '店铺列表'
            }
          ]
        }
      ]
    }
  },
  methods: {
    resize () {
      console.log('resize')
    }
  }
}
</script>
<style scoped>
.to-do-list {
  background-color: #fff;
}
.list-title {
  line-height: 56px;
  height: 56px;
  background-color: white;
  padding-left: 20px;
}
</style>
