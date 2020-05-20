<template >
    <div class="checklists">
        <div class="listsblock">
            <span class="listsspan">供应商-已审核</span>
        </div>
        <hr>
        <div class="buttons">
            <el-form ref="form" :model="form" label-width="auto">
                <div>
                    <el-row>
                        <el-col :span="5">
                            <el-form-item label="企业ID">
                                <el-input v-model="form.id"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="5" class="right">
                            <el-form-item label="企业名称">
                                <el-input v-model="form.company_name"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
                <div>
                    <el-row>
                        <el-col :span="5">
                            <el-form-item label="法人姓名">
                                <el-input v-model="form.legal_person_name"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="5" class="right">
                            <el-form-item label="法人身份证">
                                <el-input v-model="form.legal_person_id"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
                <div>
                    <el-row>
                        <el-col :span="5">
                            <el-form-item label="营业执照号">
                                <el-input v-model="form.business_license"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="5" class="right">
                            <el-form-item label="纳税识别号">
                                <el-input v-model="form.taxpayer"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
                <div>
                    <el-row>
                        <el-col :span="5">
                            <el-form-item label="联系人">
                                <el-input v-model="form.contacts"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="5" class="right">
                            <el-form-item label="联系电话">
                                <el-input v-model="form.tel"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
                <div>
                    <el-row>
                        <el-col :span="5">
                            <el-form-item label="地址">
                                <el-input v-model="form.address"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="5" class="right">
                            <el-form-item label="省份">
                                <el-input v-model="form.province"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
                <div>
                    <el-row>
                        <el-col :span="5">
                            <el-form-item label="城市">
                                <el-input v-model="form.city"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="5" class="right">
                            <el-form-item label="区域">
                                <el-input v-model="form.area"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
                <div>
                    <el-row>
                        <el-col :span="5">
                            <el-form-item label="经度">
                                <el-input v-model="form.lat"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="5" class="right">
                            <el-form-item label="维度">
                                <el-input v-model="form.lag"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
                <div>
                    <el-row>
                        <el-col :span="5">
                            <el-form-item label="图标logo">
                                <el-input v-model="form.logo"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="5" class="right">
                            <el-form-item label="企业类型">
                                <el-input v-model="form.company_type"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
                <div>
                    <el-row>
                        <el-col :span="5">
                            <el-form-item label="企业现状">
                                <el-input v-model="form.company_status"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
            </el-form>
        </div>
    </div>
</template>
<script>
    import {FindByCompany}from '../../api/company'
    export default {
        name:'checklist',
        data() {
            return {
                id:'',
                form:[]
            }
        },
        mounted() {
            this.id = this.$route.query.cid
            FindByCompany(this.id).then(res => {
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