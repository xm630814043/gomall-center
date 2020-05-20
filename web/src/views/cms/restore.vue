<template>
    <!-- 优选商家开始 -->
    <div class="goodsShop">
       <div class="good-goodsShop-title">
                    优选商家
          <span>100+大牌聚集 源头直供</span>
           <el-button slot="reference" class="upload-demo" @click="handleCreate()">添加</el-button>
       </div>
       <div class="good-goodsShop" v-model="activeIndex">
          <div class="good-goodsShop-item" v-for="(item,index) in goodShopList" :key="index" >
             <img :src="item.logo" @click="clickImg(item.id)" class="goodsimg"/>
                <hr style="height:1px;border-width:0;color:gray;background-color:gray">
<!--                   <el-button @click.native="handleDelete(item.id)" >删除</el-button>-->
              <el-popconfirm
                      title="这家推荐店铺确定删除吗？"
              >
                  <el-button slot="reference" @click="handleDelete(item.id)">删除</el-button>
              </el-popconfirm>
          </div>
        </div>
    </div>
</template>
<script>
    import {FindBySubjectList,RemoveSubject,FindBySubject}from '../../api/company'
    export default {
        name: "restore",
        data() {
            return {
                id:0,
                contentType:'',
                limit:0,
                activeIndex: '1',
                goodShopList: [],
            }
        },
        mounted() {
            this.id = this.$route.query.sid
            this.contentType = this.$route.query.sname
            this.limit = this.$route.query.slimit
            FindBySubjectList(this.id,this.contentType,this.limit).then(res => {
                // console.log(res.data)
                this.goodShopList = res.data
            })
        },
        methods: {
            handleDelete:function(index) {
                // console.log(index,this.id)
                RemoveSubject(index,this.id).then(res => {
                    console.log(res.data)
                })
            },
            handleCreate:function () {
                this.$router.push({path: '/cms/create', query: {cid:this.id}})
            },
            clickImg:function (id) {
                // console.log(id)
                FindBySubject(id,this.contentType).then(res=>{
                    console.log(res.data)
                })
            }
        }
    };
</script>
<style scoped>
    @import url("../../../public/static/css/index/index.css");
    .upload-demo{
        float:right;
    }
</style>