<template>
    <div class="goodshops">
        <div class="good-goodshops-title">
            店铺列表
            <span>大牌聚集 源头直供</span>
            <el-button slot="reference" class="upload-demo" @click="handleCreate()">添加</el-button>
        </div>
        <div class="good-goodshops" v-model="activeIndex">
            <div class="good-goodshops-item" v-for="(item,index) in goodShopsList" :key="index">
                <img :src="item.image_url" @click="clickImg(item.ID)" class="goodsimg" />
                <p style="color: black">{{ item.title }}</p>
                <hr style="height:1px;border-width:0;color:gray;background-color:gray">
                <el-popconfirm
                        title="这家店铺确定删除吗？" class="buttons"
                >
                    <el-button slot="reference" @click="handleDelete(item.ID)">删除</el-button>
                </el-popconfirm>
            </div>
        </div>
    </div>
</template>

<script>
    import {FindBySubjectList,RemoveSubject,FindBySubject}from '../../api/company'
    export default {
        name:"stores",
        data(){
               return{
                   id:0,
                   contentType:'',
                   limit:0,
                   activeIndex: '1',
                   goodShopsList:[]
               }
        },
        mounted() {
            this.id = this.$route.query.sid
            this.contentType = this.$route.query.sname
            this.limit = this.$route.query.slimit
            FindBySubjectList(this.id,this.contentType,this.limit).then(res => {
                // console.log(res.data)
                this.goodShopsList = res.data
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
    }
</script>

<style scoped>
    .buttons{
        margin-top: 1px ;
    }
    .upload-demo{
        float:right;
    }
</style>