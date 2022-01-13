let API_ROUTE

process.env.NODE_ENV === 'development'
  ? API_ROUTE = 'http://127.0.0.1:8080/api/v1'
  : API_ROUTE = 'https://forum-prod-api.com/api/v1'

export default API_ROUTE