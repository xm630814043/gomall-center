<template>
        <el-form ref="form"  v-model="form" label-width="80px" size="mini">
            <el-form-item label="名称">
                <el-input v-model="form.product_name"></el-input>
            </el-form-item>
            <el-button type="primary" @click="submitForm()">确 定</el-button>
            <el-button type="primary" @click="submitForms()">修改</el-button>
            <el-button @click="dialogFormVisible = false">取 消</el-button>
        </el-form>
</template>

<script>
    import {AddSubject,AddControlSell,ModifyControlSell}from '../../api/company'
    export default {
        name: "create",
        data() {
            return {
                form:{
                    subject_id:0,
                    object_id:0,
                },
                data:{
                    company_id:1,
                    company_name:'众泰',
                    control_sell_name:'偏远地区',
                    control_sell_type:2,
                    area_option:[{
                        "id" : 7707,
                        "area_name" : "市辖区",
                    }, {
                        "id" : 7708,
                        "area_name" : "王益区",
                    }, {
                        "id" : 7705,
                        "area_name" : "印台区",
                    }, {
                        "id" : 7706,
                        "area_name" : "耀州区",
                    }, {
                        "id" : 7704,
                        "area_name" : "宜君县",
                    }],
                    sys_config_id:[20,21,23],
                    ordinary_scope:["惠民诊所","张辉卫生所","五星大药房"]
                }
            };
        },
        methods: {
            // submitForm(){
            //     this.id = this.$route.query.cid
            //     this.form.subject_id = this.id
            //     this.form.object_id = 90
            //     AddSubject(this.form).then(res => {
            //             console.log(res.data);
            //         })
            //     },
            submitForm() {
                AddControlSell(this.data).then(res => {
                    console.log(res.data);
                })
            },
            submitForms() {
                console.log("进入修改")
                ModifyControlSell(2, this.data).then(res => {
                    console.log(res.data);
                })
            },
        }
    }
</script>

<style>

</style>