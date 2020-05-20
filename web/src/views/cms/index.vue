<template>
  <div>
    <Split style="height: 100vh">
      <SplitArea :size="25"
                 class="to-do-list">
        <div class="list-title"></div>
        <div>
          <el-collapse v-model="activeIndex">
            <el-collapse-item v-for="c in tabledata"
                              :name="c.ID"
                              :key="c.ID">
              <span slot="title" style="margin-left:20px">{{ c.channel_name }}</span>
              <div v-for="s in c.subjects"
                   :key="s.ID"
                   style="margin:20px">
                <router-link :to="{path:s.url}"></router-link>
                <el-button type="text" @click="subjects(s.ID)" style="height: 1px">{{ s.title }}</el-button>
              </div>
            </el-collapse-item>
          </el-collapse>
        </div>
      </SplitArea>
      <SplitArea :size="75" class="splitArea">
        <router-view></router-view>
      </SplitArea>
    </Split>
  </div>
</template>
<script>
  import {FindChannelList} from '../../api/company'
  export default {
    name: 'cms',
    data () {
      return {
        activeIndex: '1',
        tabledata:[],
      }
    },
    mounted(){
      FindChannelList().then(res=>{
        // console.log(this.tabledata)
        this.tabledata = res.data
      })
    },
    methods: {
      subjects(id){
        console.log("点击返回的栏目ID",id)
        if (id == 1){
          this.$router.push({path:'/cms/reproduces',query:{sid:id,sname:'product',slimit:5}})
        }
        if (id == 2){
          this.$router.push({path:'/cms/restore',query:{sid:id,sname:'shop',slimit:5}})
        }
        if (id == 3){
          this.$router.push({path:'/cms/reproduces',query:{sid:id,sname:'product',slimit:5}})
        }
        if (id == 4){
          this.$router.push({path:'/cms/stores',query:{sid:id,sname:'article',slimit:5}})
        }
      },
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
  .splitArea{
    background-color: #e6e6e6;
  }
</style>
