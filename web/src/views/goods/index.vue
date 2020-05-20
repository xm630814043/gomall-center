<template>
    <div>

        <Split style="height: 100vh">
            <SplitArea :size="40"
                       class="to-do-list">
                <div class="list-title">
                    <el-input v-model="input" placeholder="请输入产品名称" @change="changeRadio"></el-input>
                    <el-radio v-model="radio" label="1" @change="changeRadio">待审核</el-radio>
                    <el-radio v-model="radio" label="4" @change="changeRadio">未通过</el-radio>
                    <hr/>
                    <!--表格-->
                    <el-table
                            :row-style="{height:'75px',textAlign:'center'}"

                            :data="tableData"
                            style="width: 90%"
                            @row-click="getRow"
                            height=450
                            :highlight-current-row="true">
                        <el-table-column
                                prop="id"
                                label="#ID"
                                width=120>
                        </el-table-column>
                        <el-table-column
                                prop="product_name"
                                label="产品名称"
                                width=120>
                        </el-table-column>
                        <el-table-column
                                prop="producer"
                                label="生产商"
                                width=150>
                        </el-table-column>
                    </el-table>
                    <!--分页-->
                    <el-pagination style="margin-left: 80px"
                                   layout="prev, pager, next"
                                   @current-change="handleCurrentChange"
                                   :current-page="page.currentPage"
                                   :page-size="page.pageSize"
                                   :total="page.pageTotal">
                    </el-pagination>
                </div>
            </SplitArea>
            <SplitArea :size="60" class="to-do-list1">
                <div style="margin-top: 5%;margin-left: 9%;height: auto">
                    <div v-if="status==1 ? true : false "
                         style="margin-top: 5%;margin-right: 5%;border: 0px solid black">
                        <div style="float: right;margin-right: 20%">
                            <el-button v-model="status1" label="4" @click="noPassStatus" size="small">不通过</el-button>
                        </div>
                        <div style="float: right">
                            <el-button v-model="status1" label="2" @click="passStatus" size="small">通过</el-button>
                        </div>
                    </div>
                    <div v-else></div>
                    <br/>
                    <div style="margin-top: 10%">
                        <b>商品ID:</b> {{ data.id}}<br/><br/>
                        <b>商品图片:</b> <br/> <img :src="productImg"
                                                style="width:100px; height:130px; margin-left:11%"/><br/><br/>
                        <b>商品描述:</b>{{ data.description }}<br/><br/>
                        <b>发布状态:</b>{{ data.publish_status}}<br/><br/>
                    </div>
                </div>
            </SplitArea>
        </Split>
    </div>
</template>
<script>
    import {GetProductAuditInfo, GetProductBasic, UpdateProductStatus} from "../../api/company"

    export default {
        name: 'goods',
        data() {
            return {
                status: 1,
                status1: 1,
                productImg: '',
                data: [],
                page: {
                    pageNum: 1,
                    pageSize: 5,
                    pageTotal: 0,
                },
                radio: '1',
                activeIndex: '1',
                channels: [],
                input: '',
                tableData: [],
            };
        },
        mounted() {
            GetProductBasic(this.page.currentPage, this.page.pageSize, this.input, this.radio).then(res => {
                this.tableData = res.data.data
                this.page.pageTotal = res.data.count
            })
        },
        methods: {
            noPassStatus() {
                if (this.data.ID == undefined) {
                    alert("请选择商品")
                } else {
                    UpdateProductStatus(this.data.ID, 4).then(res => {
                        alert(res.msg)
                        this.data = []
                        this.productImg = ''
                    })
                }
            },
            passStatus() {
                if (this.data.ID == undefined) {
                    alert("请选择商品")
                    return
                } else {
                    UpdateProductStatus(this.data.ID, 2).then(res => {
                        alert(res.msg)
                        this.data = []
                        this.productImg = ''
                    })
                }
                GetProductBasic(this.page.currentPage, this.page.pageSize, this.input, this.radio).then(res => {
                    this.tableData = res.data.data
                    this.page.pageTotal = res.data.count
                })
            },
            handleCurrentChange(currentPage) {
                this.page.currentPage = currentPage
                GetProductBasic(this.page.currentPage, this.page.pageSize, this.input, this.radio).then(res => {
                    this.tableData = res.data.data
                })
            },
            resize() {
                console.log('resize')
            },
            getRow(row, column, event) {
                GetProductAuditInfo(row.id).then(res => {
                    this.data = res.data
                    this.productImg = res.data.pic
                })
            },
            changeRadio(label) {
                this.data = []
                this.productImg = ''
                this.page.currentPage = 1
                this.status = this.radio
                GetProductBasic(this.page.currentPage, this.page.pageSize, this.input, this.radio).then(res => {
                    this.tableData = res.data.data
                    this.page.pageTotal = res.data.count
                })
            },
        }
    }
</script>
<style scoped>
    .to-do-list {
        background-color: #fff;
    }

    .to-do-list1 {
        background-color: #fefff9
    }

    .list-title {
        line-height: 50px;
        height: 50px;
        background-color: white;
        padding-left: 20px;
    }
</style>
