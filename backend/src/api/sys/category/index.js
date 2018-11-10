import request from '@/plugin/axios'

export function getCategories (data) {
  return request({
    url: '/categories',
    method: 'get',
    data
  })
}
