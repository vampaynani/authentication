const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:3333'

interface Options {
  method?: string
  data?: any
  headers?: any
}
type ApiCallOptions = Options | undefined

export async function apiCall(
  path: string,
  { method = 'GET', data, headers }: ApiCallOptions = {}
) {
  const response = await fetch(BASE_URL + path, {
    credentials: 'include',
    method,
    headers: {
      'Content-Type': 'application/json',
      ...headers
    },
    body: JSON.stringify(data)
  })

  if (method === 'DELETE') {
    return response.ok
  }

  const jsonObj = await response.json()
  return jsonObj
}
