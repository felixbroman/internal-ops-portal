export type User = {
  id: string,
  name: string,
  email: string,
  role: string,
  manager_id?: string | null,
  created_at: string,
}

export type MeResponse = {
  user: User
}

export type SignupCreds = {
    name: string,
    email: string,
    password: string,
}

export type LoginCreds = {
    emal: string,
    password: string,
}