<template>
    <!-- 热销商品开始 -->
    <div class="hot-goods-container">
       <div class="hot-goods-container-title">
           热销商品
           <el-button slot="reference" class="upload-demo" @click="handleCreate()">添加</el-button>
       </div>
          <div class="hot-goods-container-goodscontainer" v-model="activeIndex">
             <div class="goods-item" v-for="item in hotGoodsList">
                <div>
                   <img :src="item.Pic" @click="clickImg(item.ID)" class="goodsimg" />
                </div>
                <p class="goods-item-title">{{item.ProductName}}</p>
                     <hr style="height:1px;border-width:0;color:gray;background-color:gray">
                 <el-popconfirm title="这个推荐产品确定删除吗？">
                    <el-button slot="reference" @click="handleDelete(item.ID)">删除</el-button>
                 </el-popconfirm>
             </div>
          </div>
    </div>
</template>
<script>
    import {FindBySubjectList,RemoveSubject,FindBySubject,RemoveControlSell,FindControlSellByID,FindControlSellList}from '../../api/company'
    export default {
        name: "reproduces",
        data() {
            return {
                id :0,
                contentType:'',
                limit:0,
                activeIndex: '1',
                hotGoodsList: [],
            }
        },
        mounted() {
            this.id = this.$route.query.sid
            this.contentType = this.$route.query.sname
            this.limit = this.$route.query.slimit
            // console.log(this.id,this.contentType,this.limit)
            FindControlSellByID(1).then(res=>{
              console.log(this.tabledata)
              this.tabledata = res.data
            }),
            FindControlSellList(2).then(res=>{
              console.log(this.tabledata)
              this.tabledata = res.data
            })
            FindBySubjectList(this.id,this.contentType,this.limit).then(res => {
                console.log(res.data)
                this.hotGoodsList = res.data
            })
        },
        methods: {
            handleDelete:function(index) {
                // console.log(index,this.id)
                // RemoveSubject(index,this.id).then(res => {
                //     console.log(res.data)
                // }),
                    RemoveControlSell(4).then(res => {
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