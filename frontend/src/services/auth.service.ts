import api from './api'
import type { SignupCreds, LoginCreds, MeResponse } from '@/types/auth'

export async function signup(cred: SignupCreds): Promise<MeResponse> {
  const res = await api.post('/api/auth/signup', cred)
  return res.data
}


export async function login(cred: LoginCreds): Promise<MeResponse> {
  const res = await api.post('/api/auth/signup', cred)
  return res.data
}

export async function getMe(): Promise<MeResponse> {
    const res = await api.get('/api/auth/me')
    return res.data
}