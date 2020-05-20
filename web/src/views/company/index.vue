<template>
    <div>
        <Split style="height: 100vh">
            <SplitArea :size="25" class="to-do-list">
                <div class="block">
                    <span class="demonstration">企业管理</span>
                    <hr>
                    <div class="buttons">
                        <el-cascader
                                v-model="activeIndex"
                                placeholder="搜索："
                                :options="options"
                                filterable
                                class="elserach"
                                @change="handlechange"
                        >
                        </el-cascader>
                        <el-button type="primary" icon="el-icon-search" >搜索</el-button>
                    </div>
                </div>
                <div class="container_table"  v-show="tableshow">
                    <el-table
                            :data="tableData.slice((PageNum-1)*pageSize,PageNum*pageSize)"
                            stripe
                            style="width: 100%"
                            :default-sort = "{prop: 'date', order: 'descending'}"
                            @row-click="handdle"
                        >
                        <el-table-column
                                prop="id"
                                label="ID"
                                sortable
                                min-width="8.19%">
                        </el-table-column>
                        <el-table-column
                                prop="company_name"
                                label="公司名称"
                                sortable
                                min-width="8.19%">
                        </el-table-column>
                        <el-table-column
                                prop="legal_person_name"
                                label="法人姓名"
                                min-width="8.19%">
                        </el-table-column>
                        <el-table-column
                                prop="business_license"
                                label="营业执照号"
                                min-width="8.19%">
                        </el-table-column>
                        <el-table-column
                                prop="legal_person_id"
                                label="法人身份证"
                                min-width="8.19%">
                        </el-table-column>
                        <el-table-column
                                prop="company_status"
                                label="审核状态"
                                min-width="8.19%">
                        </el-table-column>
                    </el-table>
                    <div class="eltable" >
                        <span class="demonstration">共{{total}}条</span>
                        <el-pagination class="fy"
                                       layout="prev, pager, next"
                                       @current-change="current_change"
                                       :total="total"
                                       :page-size="pageSize"
                                       background
                        >
                        </el-pagination>
                    </div>
                </div>

                <div class="container_table"  v-show="tableshows">
                    <el-table
                            :data="tableDatas.slice((PageNum-1)*pageSize,PageNum*pageSize)"
                            stripe
                            style="width: 100%"
                            :default-sort = "{prop: 'date', order: 'descending'}"
                            @row-click="handdles"
                    >
                        <el-table-column
                                prop="ID"
                                label="ID"
                                sortable
                                min-width="8.19%">
                        </el-table-column>
                        <el-table-column
                                prop="username"
                                label="用户名称"
                                sortable
                                min-width="8.19%">
                        </el-table-column>
                        <el-table-column
                                prop="password"
                                label="用户密码"
                                min-width="8.19%">
                        </el-table-column>
                        <el-table-column
                                prop="email"
                                label="邮箱"
                                min-width="8.19%">
                        </el-table-column>
                        <el-table-column
                                prop="mobile"
                                label="手机"
                                min-width="8.19%">
                        </el-table-column>
                        <el-table-column
                                prop="user_type"
                                label="用户类型"
                                min-width="8.19%">
                        </el-table-column>
                    </el-table>
                    <div class="eltable" >
                        <span class="demonstration">共{{total}}条</span>
                        <el-pagination class="fy"
                                       layout="prev, pager, next"
                                       @current-change="current_changes"
                                       :total="total"
                                       :page-size="pageSize"
                                       background
                        >
                        </el-pagination>
                    </div>
                </div>
            </SplitArea>
            <SplitArea :size="75" class="splitArea" v-show="errors">
                <router-view/>
            </SplitArea>
        </Split>
    </div>
</template>

<script>
    import {FindCompanyList,FindAccountList}from '../../api/company'
    export default {
        name:'company',
        inject:['reload'],
        data:function() {
            return {
                statusId:'',
                ids:'',
                activeIndex: '1',
                PageNum: 1,//默认开始页面
                pageSize:5,//每页的数据条数
                total:100,//默认数据总数
                tableshow:false,
                errors:false,
                options: [ {
                    value: 'check',
                    label: '供应商',
                    children: [{
                        value: '0',
                        label: '未审核',
                    },
                        {
                            value: '1',
                            label: '已审核',
                        }]
                }],
                tableData: [],
                tableDatas: [],
                filterParams: {
                    PageNum: 0,
                    PageSize: 5,
                    KeyWord:""
                },
            };
        },
        methods: {
            handlechange: function (value) {
                this.errors = false
                this.tableshow = true
                this.tableshows = false
                this.statusId = value[1]
                this.filterParams.PageNum = 1
                if (this.statusId == '0') {
                    FindCompanyList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, this.statusId).then(res => {
                        this.tableData = res.data['data']
                        this.total = res.data.count
                    })

                }
                if (this.statusId == '1') {

                    FindCompanyList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, this.statusId).then(res => {
                        this.tableData = res.data['data']
                        this.total = res.data.count
                    })
                }
            },

            handdle: function (row) {
                this.errors = true
                this.tableshows = true
                if (this.statusId == '0') {
                    FindAccountList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, row.id).then(res => {
                        this.tableDatas = res.data['data']
                        this.total = res.data.count
                    })
                    this.$router.push({path: '/company/updates', query: {cid: row.id}})
                }
                if (this.statusId == '1') {
                    // console.log("获取到的企业ID",row.id)
                    FindAccountList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, row.id).then(res => {
                        this.tableDatas = res.data['data']
                        this.total = res.data.count
                    })
                    this.$router.push({path: '/company/checklist', query: {cid: row.id}})
                }

            },
            handdles:function(row){
                    this.$router.push({path: '/company/lists', query: {cid: row.ID}})
            },
            current_change: function (currentPage) {
                this.filterParams.PageNum = currentPage;
                if (this.statusId == '0') {
                    FindCompanyList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, this.statusId).then(res => {
                        this.tableData = res.data['data']
                    })
                }
                if (this.statusId == '1') {
                    FindCompanyList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, this.statusId).then(res => {
                        this.tableData = res.data['data']
                    })
                }
            },
            current_changes: function (currentPage) {
                this.filterParams.PageNum = currentPage;
                FindAccountList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, this.ids).then(res => {
                        this.tableDatas = res.data['data']
                })
            },
        }
    }
</script>
<style scoped>
    .to-do-list {
        background-color: #fff;
        text-align: center;
    }
    .block{
        margin-top:10px;
    }
    .demonstration{
        margin-left:20px;
        font-size:28px;
    }
    .buttons{
        margin-top:10px;
        float:right;
    }
    .elserach{
        width: 170px;
    }
    .eltable{
        margin-top: 20px;
        text-align: center;
    }
    .splitArea{
        background-color: #ffffff;
    }
</style>
