import request from '../ajax'

function GetCompnayList() {
  request({
    url: '/api/v1/company',
    method: 'get',
  })
}