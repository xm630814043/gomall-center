<template >
    <div class="checklists">
        <div class="listsblock">
            <span class="listsspan">采购商-信息</span>
        </div>
        <hr>
        <div class="buttons">
            <el-form ref="form" :model="form" label-width="auto">
                <div>
                    <el-row>
                        <el-col :span="5" >
                            <el-form-item label="用户名称">
                                <el-input v-model="form.ID"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="5" class="right">
                            <el-form-item label="性别">
                                <el-input v-model="form.sex"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
                <div>
                    <el-row>
                        <el-col :span="5" >
                            <el-form-item label="邮箱">
                                <el-input v-model="form.email"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="5" class="right">
                            <el-form-item label="密码">
                                <el-input v-model="form.password"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
                <div>
                    <el-row>
                        <el-col :span="5" >
                            <el-form-item label="用户类型">
                                <el-input v-model="form.user_type"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="5" class="right">
                            <el-form-item label="手机">
                                <el-input v-model="form.mobile"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
                <div>
                    <el-row>
                        <el-col :span="5">
                            <el-form-item label="企业ID">
                                <el-input v-model="form.company_id"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="5" class="right">
                            <el-form-item label="登录密码">
                                <el-input v-model="form.pay_password"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
            </el-form>
        </div>
    </div>
</template>
<script>
    import {FindByAccount}from '../../api/company'
    export default {
        name:'lists',
        data() {
            return {
                form:[],
            }
        },
        mounted() {
            this.id = this.$route.query.cid
            FindByAccount(this.id).then(res => {
                console.log(res.data)
                this.form = res.data
            })
        },
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