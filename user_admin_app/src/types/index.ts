export interface User {
  id: number
  username: string
  firstname: string
  lastname: string
  email: string
}

export interface JWTPayload {
  MapClaims: {
    eat: number
    iat: number
  }
  session: string
}
