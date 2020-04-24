<template>
   <div>
      <el-table
              :data="tableData"
              style="width: 100%">
         <el-table-column
                 prop="id"
                 label="编号"
                 width="180">
         </el-table-column>
         <el-table-column
                 prop="company_name"
                 label="企业名称"
                 width="180">
         </el-table-column>
      </el-table>
   </div>
</template>

<script>
   import { GetCompnayList }  from '../../api/company'
    export default {
       name: "company",
       data () {
          return {
             tableData: []
          }
       },
       created() {
          initTable()
       },
       methods: {
         initTable() {
            GetCompnayList().then(res=>{
               if (res.code == 200){
                  this.tableData = res.data
               }
            })
         }
       }
    }
</script>

<style scoped>

</style>