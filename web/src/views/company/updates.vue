<template >
    <div class="title">
        <div class="container_table">
            <el-table
                    :data="tableData"
                    stripe
                    style="width: 100%"
                    :default-sort = "{prop: 'date', order: 'descending'}"
            >
                <el-table-column
                        prop="Id"
                        sortable
                        label="待审核企业ID"
                        min-width="8.19%">
                </el-table-column>
                <el-table-column prop="file_url" label="经营许可" min-width="8.19%">
                    <template  slot-scope="scope">
                        <el-popover
                                placement="right"
                                title=""
                                trigger="click">
                            <el-image slot="reference" :src="scope.row.file_url" :alt="scope.row.file_name" style="max-height: 100px;max-width: 100px"></el-image>
                            <el-image :src="scope.row.file_url"></el-image>
                        </el-popover>
                        <!--                <img :src="scope.row.file_name" width="40" height="40"/>-->
                    </template>
                </el-table-column>
                <el-table-column  prop="file_url" label="企业资质" min-width="8.19%">
                    <template slot-scope="scope">
                        <!--                <a :href="scope.row.file_url" target="_blank">"企业资质"</a>-->
                        <router-link :to="{path:scope.row.file_url}">企业资质</router-link>
                    </template>
                </el-table-column>
                <el-table-column label="操作"  min-width="8.19%">
                    <template slot-scope="scope">
                        <el-button type="text" @click="companystatus(scope.row.company_id)">通过</el-button>
                        &nbsp;&nbsp;&nbsp;
                        <el-button type="text" @click="companystatu(scope.row.company_id)">驳回</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>
<script>
    import {ModifyCompany,FindByCompany}from '../../api/company'
    export default {
        name: 'updates',
        data() {
            return {
                form:{
                    company_status:0,
                },
                tableData: [],
            }
        },
        mounted() {
            this.id = this.$route.query.cid
            FindByCompany(this.id).then(res => {
                // console.log(res.data)
                this.tableData = res.data['company_file']
            })
        },
        methods: {
            companystatus: function (id) {
                this.form.company_status = 1
                // console.log(this.form.company_status)
                ModifyCompany(id,this.form).then(res =>
                {
                    console.log(res.data)
                    console.log("审核批准已通过")
                })

            },
            companystatu: function (id) {
                this.form.company_status = 2
                // console.log(this.form.company_status)
                ModifyCompany(id,this.form).then(res =>
                {
                    console.log(res.data)
                    console.log("审核批准驳回")
                })

            }
        }
    }
</script>
<style>
    .checklists{
        background-color: #ffffff;
        text-align: center;
    }
    .listsblock{
        margin-top:10px;

    }
    .listsspan{
        margin-left:10px;
        font-size:28px;
    }
    .buttons{
        margin-top: 20px;
        margin-left: 40px;
        background-color:#ffffff;
    }
    .right{
        margin-left: 70px;
    }
</style>