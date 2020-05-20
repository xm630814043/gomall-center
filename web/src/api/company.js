import request from '../ajax'

/**
 * 测试
 */
export function AddControlSell(data) {
    return request({
        url: '/api/v1/cms/subjects/create',
        method: 'post',
        data:data
    })
}
export function FindControlSellByID(id) {
    return request({
        url: '/api/v1/cms/channels/list/' + id,
        method: 'get',
    })
}
export function FindControlSellList(id) {
    return request({
        url: '/api/v1/cms/subjects/list/' + id,
        method: 'get',
    })
}
export function RemoveControlSell(id) {
    return request({
        url: '/api/v1/cms/subjects/' + id ,
        method: 'delete',
    })
}
export function ModifyControlSell(id,data) {
    return request({
        url: '/api/v1/cms/subjects/edit/' + id,
        method: 'put',
        data:data
    })
}
































/**
 * 获取企业列表
 * @param {*} Radio：企业状态
 * @param {*} PageSize：页面大小
 * @param {*} PageNum：页面页数
 * @param {*} KeyWord：关键字
 */
export function FindCompanyList(PageSize,PageNum,KeyWord,Radio) {
  return request({
      url: '/api/v1/company/list?radio=' + Radio,
      method: 'get',
      params:{
          PageSize: PageSize,
          PageNum: PageNum,
          KeyWord: KeyWord
      }
  })
}
/**
 * 获取企业详情
 * @param {*} id：企业id
 */
export function FindByCompany(id) {
    return request({
        url: '/api/v1/company/detail/' + id ,
        method: 'get',
    })
}
/**
 * 修改企业状态
 * @param {*} id：企业id
 * @param {*} data
 */
export function ModifyCompany(id,data) {
    return request({
        url: '/api/v1/company/status/' + id,
        method: 'put',
        data:data
    })
}
/**
 * 获取用户列表
 * @param {*} PageSize：页面大小
 * @param {*} PageNum：页面页数
 * @param {*} KeyWord：关键字
 */
export function FindAccountList(PageSize, PageNum, KeyWord, Radio) {
    return request({
        url: '/api/v1/account/list?radio=' + Radio,
        method: 'get',
        params:{
            PageSize: PageSize,
            PageNum: PageNum,
            KeyWord: KeyWord
        }
    })
}
/**
 * 获取用户详情
 * @param {*} id：用户id
 */
export function FindByAccount(id) {
    return request({
        url: '/api/v1/account/detail/' + id ,
        method: 'get',
    })
}
/**
 * 获取频道专题列表
 */
export function FindChannelList() {
    return request({
        url: '/api/v1/cms/channel/list',
        method: 'get',
    })
}

/**
 * 获取主题内容列表
 * @param {*} id：专题id
 * @param {*} contentType：专题类型
 * @param {*} limit：取值范围
 */
export function FindBySubjectList(id,contentType,limit) {
    return request({
        url: '/api/v1/cms/subject/list',
        method: 'get',
        params:{
            id:id,
            contentType:contentType,
            limit:limit
        }
    })
}
/**
 * 删除专题内容
 * @param {*} id：专题内容id
 * @param {*} subject_id：专题id
 */
export function RemoveSubject(id, subject_id) {
    return request({
        url: '/api/v1/cms/subject/' + id ,
        method: 'delete',
        params:{
            SubjectID:subject_id
        }
    })
}
/**
 * 添加专题内容
 * @param {*}data
 */
export function AddSubject(data) {
    return request({
        url: '/api/v1/cms/subject/create',
        method: 'post',
        data:data
    })
}
/**
 * 获取专题内容详情
 * @param {*} id：专题id
 * @param {*} contentType：专题类型
 */
export function FindBySubject(id,contentType) {
    return request({
        url: '/api/v1/cms/subject/detail/' + id ,
        method: 'get',
        params:{
            contentType:contentType,
        }
    })
}


//种类树category
export function CategoryNode() {
  return request({
    url: '/api/v1/categorys',
    method: 'get',
  })
}

//根据种类获取属性prop
export function GetPropByCategory(categoryId,data) {
  return request({
    url: '/api/v1/category/propertys/'+categoryId,
    method: 'get',
    params:data,
  })
}
//没用
export function GetAllGoodsPropValue() {
  return request({
    url: '/api/v1/goods/prop/value',
    method: 'get',
  })
}
//没用
export function GetGoodsPropValueByPname(name) {
  return request({
    url: '/api/v1/goods/prop/valuebypname',
    method: 'get',
    data: {
      name: name
    }
  })
}
//获取商品基本信息
export function GetProductBasic(PageNum,PageSize,KeyWord,radio) {
  return request({
    url: '/api/v1/product/basic/status?radio='+radio,
    method: 'get',
    params:{
      PageNum:PageNum,
      PageSize:PageSize,
      KeyWord:KeyWord
    }
  })
}
//根据商品id获取商品所要审核信息
export function GetProductAuditInfo(id) {
  return request({
    url: '/api/v1/product/info/'+id,
    method: 'get',
  })
}
//修改商品发布状态
export function UpdateProductStatus(id,status) {
  return request({
    url: '/api/v1/product/status/'+id,
    method: 'put',
    data:{
      publish_status:status
    }

  })
}
//根据prop_id获取prop-value
export function GetPropValueByPropId(id,data) {
  return request({
    url:'/api/v1/category/property/value/'+id,
    method:'get',
    params:data,
  })
}
export function AddValueByPropId(id,value) {
  return request({
    url:'/api/v1/category/property/value',
    method:'post',
    data:{
      prop_id:id,
      prop_value:value
    }
  })
}
//删除商品属性值
export function DelPropValById(id) {
  return request({
    url:'/api/v1/category/property/value/'+id,
    method:'delete',
  })
}
//添加一个商品属性
export function AddProductProp(categoryId,parentId,formData) {
  return request({
    url:'/api/v1/category/property?categoryId='+categoryId+'&parentId='+parentId,
    method:'post',
    data:formData,
  })
}
//修改商品属性值
export function EditValueByPropId(id,value) {
  return request({
    url:'/api/v1/category/property/value/'+id,
    method:'put',
    data:{
      prop_value:value
    }
  })
}
//删除商品属性
export function DelPropNameById(id,categoryId,parentId) {
  return request({
    url:'/api/v1/category/property/remove/'+id+'?categoryId='+categoryId+'&parentId='+parentId,
    method:'delete',
  })
}
export function EditProperty(formdata,id) {
  return request({
    url:'/api/v1/category/property/edit/'+id,
    method:'put',
    data:formdata
  })
}