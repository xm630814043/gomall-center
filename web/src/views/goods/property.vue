<template>
    <div>
        <Split style="height: 100vh">
            <SplitArea :size="25"
                       class="to-do-list">
                <div class="list-title"></div>
                <div>
                    <el-tree
                            :data="rootdata"
                            ref="tree"
                            node-key="Id"
                            :props="defaultProps"
                            :accordion="true"
                            @node-click="getInfo"
                    >
                    </el-tree>
                </div>
            </SplitArea>
            <SplitArea :size="75">
                <div style="height: 40%">
                    <template>
                        <el-table
                                :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
                                style="width: 100%"
                                @row-click="getRow"
                                :highlight-current-row="true">
                            <el-table-column
                                    label="#prop_id"
                                    prop="id">
                            </el-table-column>
                            <el-table-column
                                    label="prop_name"
                                    prop="prop_name">
                            </el-table-column>
                            <el-table-column
                                    label="Prop_type"
                                    prop="prop_type">
                            </el-table-column>
                            <el-table-column
                                    label="input_type"
                                    prop="input_type">
                            </el-table-column>

                            <el-table-column
                                    align="right">
                                <template slot="header" slot-scope="scope">
                                   <!-- <el-input
                                            v-model="search"
                                            size="mini"
                                            placeholder="输入关键字搜索"/>-->
                                    <el-button type="text" @click="dialogFormVisible = true">添加新的属性</el-button>
                                </template>
                                <template slot-scope="scope">
                                    <el-button
                                            size="mini"
                                            @click="handleEdit(scope.$index, scope.row)">Edit
                                    </el-button>
                                    <el-button
                                            size="mini"
                                            type="danger"
                                            @click="handleDelete(scope.$index, scope.row)">Delete
                                    </el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                    </template>
                    <div style="float: left">
                        <el-pagination
                                layout="prev, pager, next"
                                @current-change="handleCurrentChange"
                                :current-page="page.currentPage"
                                :page-size="page.pageSize"
                                :total="page.pageTotal">
                        </el-pagination>
                    </div>
                    <div style="float: left;margin-left: 5%;">
                        <!-- Form -->


                        <el-dialog :title="this.title" :visible.sync="dialogFormVisible">
                            <el-form :model="form">
                                <el-form-item label="产品属性" :label-width="formLabelWidth">
                                    <el-input v-model="form.prop_name" autocomplete="off"></el-input>
                                </el-form-item>
                                <el-form-item label="属性类型" :label-width="formLabelWidth">
                                    <el-select v-model="form.prop_type" placeholder="请选择属性类型">
                                        <el-option label="规格" value='1'></el-option>
                                        <el-option label="参数" value='2'></el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item label="输入类型" :label-width="formLabelWidth">
                                    <el-select v-model="form.input_type" placeholder="请选择输入类型">
                                        <el-option label="选择" value='0'></el-option>
                                        <el-option label="输入" value='1'></el-option>
                                    </el-select>
                                </el-form-item>
                            </el-form>
                            <div slot="footer" class="dialog-footer">
                                <el-button @click="quxiaoAdd">取 消</el-button>
                                <el-button type="primary" @click="addPropName">确 定</el-button>
                            </div>
                        </el-dialog>
                    </div>
                </div>
                <div style="height: 40%;margin-top: 10%">
                    <el-table
                            :data="butTableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
                            style="width: 100%"
                            :highlight-current-row="true">
                        <el-table-column
                                label="prop_valueId"
                                prop="id">
                        </el-table-column>
                        <el-table-column
                                label="prop_id"
                                prop="prop_id">
                        </el-table-column>
                        <el-table-column
                                label="prop_value"
                                prop="prop_value">
                        </el-table-column>
                        <el-table-column
                                align="right">
                            <template slot="header" slot-scope="scope">
                                <el-button type="text" style="margin-right: 5%" @click="addPropVal">新增</el-button>
                            </template>
                            <template slot-scope="scope">
                                <el-button
                                        size="mini"
                                        @click="buthandleEdit(scope.$index, scope.row)">Edit
                                </el-button>
                                <el-button
                                        size="mini"
                                        type="danger"
                                        @click="buthandleDelete(scope.$index, scope.row)">Delete
                                </el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                    <el-pagination
                            layout="prev, pager, next"
                            @current-change="handleCurrentChange1"
                            :current-page="page1.currentPage"
                            :page-size="page1.pageSize"
                            :total="page1.pageTotal">
                    </el-pagination>
                </div>
            </SplitArea>
        </Split>
    </div>
</template>
<script>
    import {
        AddProductProp,
        AddValueByPropId,
        CategoryNode,
        DelPropValById,
        EditValueByPropId,
        GetProductBasic,
        GetPropByCategory,
        GetPropValueByPropId,
        DelPropNameById,
        EditProperty
    } from "../../api/company"
    export default {
        name: 'property',
        data() {
            return {
                butvalue: '',
                dialogTableVisible: false,
                dialogFormVisible: false,
                edtpropId:0,
                form: {
                    prop_name: '',
                    prop_type:'2',
                    input_type:'1',
                },
                formLabelWidth: '120px',
                ButId: '',
                addValue: '',
                inputtype: '',
                propId: '',
                categoryId: -1,
                parentId: '',
                title: '',
                page: {
                    pageNum: 1,
                    pageSize: 5,
                    pageTotal:0,
                },
                page1: {
                    pageNum: 1,
                    pageSize: 4,
                    pageTotal: 0,
                },
                activeIndex: '1',
                activeName: 'first',
                tableData: [],
                butTableData: [],
                search: '',
                rootdata: [],
                defaultProps: {
                    children: 'children',
                    label: 'category_name'
                }
            }
        },
        mounted() {
            CategoryNode().then(res => {
                this.rootdata = [res.data]
            })
            GetProductBasic().then(res => {
                console.log(res)
            })
        },
        methods: {
            resize() {
                console.log('resize')
            },
            quxiaoAdd(){
                this.dialogFormVisible=false
                this.form.prop_name=''
            },
            addPropName() {
                if(this.categoryId==-1){
                    alert("请输入要添加的种类")
                    this.dialogFormVisible = false;
                    return
                }
                if(this.edtpropId!=0){
                    EditProperty(this.form,this.edtpropId).then(res=>{
                        alert(res.msg)
                        this.form.prop_name=''
                    })
                    this.edtpropId=0
                    this.dialogFormVisible = false;
                    return
                }
                console.log(this.form)
                if(this.form.prop_name!=''){
                    AddProductProp(this.categoryId, this.parentId, this.form)
                    GetPropByCategory(this.categoryId, {
                        PageNum: this.page.currentPage,
                        PageSize: this.page.pageSize,
                        KeyWord: this.search
                    }).then(res => {
                        this.tableData = res.data.data
                        this.page.pageTotal=res.data.count
                    })
                    this.dialogFormVisible = false;
                }else {
                    this.dialogFormVisible = false;
                    this.form.prop_name=''
                    alert("内容不能为空")
                }
            },
            addPropVal() {
                if(this.propId==''){
                    alert('请选择商品属性')
                    return
                }
                if (this.inputtype === 1) {
                    this.$prompt('请输入value', '提示', {
                        confirmButtonText: '确定',
                        cancelButtonText: '取消',
                    }).then(({value}) => {
                        this.$message({
                            message: value
                        })
                        this.addValue = value
                        if (value != '') {
                            AddValueByPropId(this.propId, value)
                            GetPropValueByPropId(this.propId, {
                                PageNum: this.page1.currentPage,
                                PageSize: this.page1.pageSize
                            }).then(res => {
                                this.butTableData = res.data.data
                            })
                        }else{
                            alert("请输入产品属性值")
                        }
                    }).catch(() => {
                            this.$message({
                                type: 'info',
                                message: '取消输入'
                            });
                        }
                    )
                }
            },
            getRow(row, column, event) {
                this.page1.currentPage=1
                GetPropValueByPropId(row.id, {
                    PageNum: this.page1.currentPage,
                    PageSize: this.page1.pageSize
                }).then(res => {
                    this.page1.pageTotal=res.data.count
                    this.butTableData = res.data.data
                    this.inputtype = row.input_type
                })
                this.propId = row.id
            },
            handleCurrentChange(currentPage) {
                this.page.currentPage = currentPage
                //this.getInfo({id: this.categoryId, category_name: this.title, parent_id: this.parentId}, "", "")
                GetPropByCategory(this.categoryId, {
                    PageNum: this.page.currentPage,
                    PageSize: this.page.pageSize,
                }).then(res => {
                    this.tableData = res.data.data
                    this.page.pageTotal=res.data.count
                })
            },
            handleCurrentChange1(currentPage) {
                this.page1.currentPage = currentPage
                GetPropValueByPropId(this.propId, {
                    PageNum: this.page1.currentPage,
                    PageSize: this.page1.pageSize
                }).then(res => {
                    this.butTableData = res.data.data
                })
            },
            getInfo(obj, node, data) {
                this.propId=''
                this.butTableData=[]
                this.page1.pageNum=1
                this.page.currentPage=1
                this.page1.pageTotal=0
                this.categoryId = obj.id
                this.title = obj.category_name
                this.parentId = obj.parent_id
                console.log(obj)
                GetPropByCategory(this.categoryId, {
                    PageNum: this.page.currentPage,
                    PageSize: this.page.pageSize,
                }).then(res => {
                    this.tableData = res.data.data
                    this.page.pageTotal=res.data.count
                })
            },
            handleEdit(index, row) {
                this.edtpropId=row.id
                this.form.prop_name=row.prop_name
                this.form.prop_type=row.prop_type+''
                this.form.input_type=row.input_type+''
                this.dialogFormVisible = true
                console.log(index, row);
            },
            handleDelete(index, row) {
                this.$confirm('此操作将永久删除该内容, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.$message({
                        type: 'success',
                        message: '删除成功!'
                    });
                         DelPropNameById(row.id,this.categoryId,this.parentId)
                         this.butTableData=[]

                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消删除'
                    });
                });
                console.log(index, row);
            },
            buthandleDelete(index, row) {
                this.$confirm('此操作将永久删除该内容, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.$message({
                        type: 'success',
                        message: '删除成功!'
                    });
                    DelPropValById(row.id)
                    GetPropValueByPropId(this.propId, {
                        PageNum: this.page1.currentPage,
                        PageSize: this.page1.pageSize
                    }).then(res => {
                        this.butTableData = res.data
                    })
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消删除'
                    });
                });
            },
            buthandleEdit(index, row) {
                if (this.inputtype === 1) {
                    this.$prompt('请修改value', '提示', {
                        confirmButtonText: '确定',
                        cancelButtonText: '取消',
                    }).then(({value}) => {
                        this.$message({
                            message: value
                        })
                        this.butvalue = value
                        if (value != '') {
                                EditValueByPropId(row.id, this.butvalue)
                                GetPropValueByPropId(this.propId, {
                                    PageNum: this.page1.currentPage,
                                    PageSize: this.page1.pageSize
                                }).then(res => {
                                    this.butTableData = res.data.data
                                })
                        }
                    }).catch(() => {
                            this.$message({
                                type: 'info',
                                message: '取消输入'
                            });
                        }
                    )
                }
            }
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
</style>